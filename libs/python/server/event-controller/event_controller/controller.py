import asyncio
import json
import time
from typing import Any, Dict, Tuple, Union

import warlock
from cli_schema_vault.client import async_py_schema_vault_client
from dto_config_vault.output import ConfigDTO
from dto_events_router.input import (InputMetadataDTO, MetadataDTO,
                                     OutputMetadataDTO, ServiceFeedBackDTO)
from dto_events_router.output import ProcessOrderDTO
from dto_schema_vault.input import SchemaDataDTO
from dto_schema_vault.output import SchemaDTO
from pydebug import debug
from pylog.log import setup_logging
from pyrabbitmq.producer import RabbitMQProducer
from pysd.sd import ServiceDiscovery
from pyserializer.serializer import (serialize_to_dataclass, serialize_to_dict,
                                     serialize_to_json)
from pywarlock import serializer as warlock_serializer

logger = setup_logging(__name__)

SCHEMA_TYPE_MESSAGE = "input"
SCHEMA_TYPE_RESPONSE = "output"


class Controller:
    """
    Base controller class for handling event processing.

    Attributes:
        config (ConfigDTO): Configuration data transfer object.
        queue_active_jobs (asyncio.Queue): Queue for managing active jobs.
        dbg (Union[debug.EnabledDebug, debug.DisabledDebug]): Debug instance.
    """

    def __init__(
        self,
        config: ConfigDTO,
        queue_active_jobs: asyncio.Queue,
        dbg: Union[debug.EnabledDebug, debug.DisabledDebug]
    ):
        """
        Initializes the Controller instance.

        Args:
            config (ConfigDTO): Configuration data transfer object.
            queue_active_jobs (asyncio.Queue): Queue for managing active jobs.
            dbg (Union[debug.EnabledDebug, debug.DisabledDebug]): Debug instance.
        """
        self.dbg = dbg
        self.config = config
        self.queue_active_jobs = queue_active_jobs
        self.__schema_vault_client = async_py_schema_vault_client()

    @property
    def is_controller_active(self) -> bool:
        """
        Checks if the controller is active.

        Returns:
            bool: True if the controller is active, False otherwise.
        """
        return self.config.active

    @property
    def config_id(self):
        """
        Gets the configuration ID.

        Returns:
            str: Configuration ID.
        """
        return self.config.config_id

    @property
    def provider(self):
        """
        Gets the provider.

        Returns:
            str: Provider.
        """
        return self.config.provider

    @property
    def service(self):
        """
        Gets the service.

        Returns:
            str: Service.
        """
        return self.config.service

    @property
    def source(self):
        """
        Gets the source.

        Returns:
            str: Source.
        """
        return self.config.source

    @property
    def config_version_id(self):
        """
        Gets the configuration version ID.

        Returns:
            str: Configuration version ID.
        """
        return self.config.config_version_id

    def get_json_schema(self, schema: SchemaDTO):
        """
        Gets the JSON schema from the schema DTO.

        Args:
            schema (SchemaDTO): Schema data transfer object.

        Returns:
            dict: JSON schema.
        """
        return schema.json_schema

    def get_schema_version_id(self, schema: SchemaDTO):
        """
        Gets the schema version ID from the schema DTO.

        Args:
            schema (SchemaDTO): Schema data transfer object.

        Returns:
            str: Schema version ID.
        """
        return schema.schema_version_id

    async def get_schema(self, schema_type: str):
        """
        Fetches the schema from the schema vault client.

        Args:
            schema_type (str): Type of schema to fetch.

        Returns:
            Tuple[dict, str]: JSON schema and schema version ID.
        """
        message_schema = await self.__schema_vault_client.list_schemas_by_service_source_provider_and_schema_type(
            provider=self.provider,
            service=self.service,
            source=self.source,
            schema_type=schema_type
        )
        return self.get_json_schema(message_schema), self.get_schema_version_id(message_schema)

    async def validate_schema(self, job_data: Dict[str, Any]) -> bool:
        """
        Validates the job data against the schema.

        Args:
            job_data (Dict[str, Any]): Job data to validate.

        Returns:
            bool: True if the schema is valid, False otherwise.
        """
        return await self.__schema_vault_client.validate_schema(self.get_output_schema_to_validate(job_data))

    def _read_message(self, message) -> Dict[str, Any]:
        """
        Reads and parses the message body.

        Args:
            message: Message to read.

        Returns:
            Dict[str, Any]: Parsed message body.

        Raises:
            ValueError: If the message body is invalid.
        """
        try:
            message_body = message.body.decode()
            return json.loads(message_body)
        except json.JSONDecodeError as e:
            logger.error(f"Failed to parse message body: {e}")
            raise ValueError("Invalid message body")

    def _get_message_data(self, input_dto: ProcessOrderDTO) -> Dict[str, Any]:
        """
        Extracts the data from the process order DTO.

        Args:
            input_dto (ProcessOrderDTO): Process order DTO.

        Returns:
            Dict[str, Any]: Data extracted from the DTO.
        """
        return input_dto.data

    def get_event_order_dto(self, message) -> ProcessOrderDTO:
        """
        Converts the message to a process order DTO.

        Args:
            message: Message to convert.

        Returns:
            ProcessOrderDTO: Process order DTO.
        """
        message_body = self._read_message(message)
        input_dto = serialize_to_dataclass(message_body, ProcessOrderDTO)
        return input_dto

    async def parse_message(self, event_order: ProcessOrderDTO) -> Tuple[type[warlock.model.Model], str]:
        """
        Parses the event order message.

        Args:
            event_order (ProcessOrderDTO): Event order DTO.

        Returns:
            Tuple[type[warlock.model.Model], str]: Parsed message and schema version ID.
        """
        message_schema, schema_version = await self.get_schema(SCHEMA_TYPE_MESSAGE)
        input_data = self._get_message_data(event_order)
        json_schema = serialize_to_dict(message_schema)
        return warlock_serializer.serialize_to_dataclass(json_schema, input_data), schema_version

    async def job_dispatcher(self, event_input, job_handler: callable, metadata: MetadataDTO) -> ServiceFeedBackDTO:
        """
        Dispatches the job for processing.

        Args:
            event_input: Event input data.
            job_handler (callable): Job handler function.
            metadata (MetadataDTO): Metadata DTO.

        Returns:
            ServiceFeedBackDTO: Service feedback DTO.
        """
        await self.queue_active_jobs.put(1)
        job_data = await job_handler(self.config, metadata, self.dbg).execute(event_input)
        return job_data

    async def get_metadata(
        self,
        event_order: ProcessOrderDTO,
        schema_version_input: str,
        schema_version_output: str
    ) -> MetadataDTO:
        """
        Gets the metadata for the event order.

        Args:
            event_order (ProcessOrderDTO): Event order DTO.
            schema_version_input (str): Input schema version ID.
            schema_version_output (str): Output schema version ID.

        Returns:
            MetadataDTO: Metadata DTO.
        """
        return MetadataDTO(
            config_id=self.config_id,
            provider=self.provider,
            service=self.service,
            source=self.source,
            processing_id=event_order.processing_id,
            input_metadata=InputMetadataDTO(
                input_id=event_order.input_id,
                schema_version_id=schema_version_input,
                processing_order_id=event_order.order_id
            ),
            config_version_id=self.config_version_id,
            output_metadata=OutputMetadataDTO(
                schema_version_id=schema_version_output,
            )
        )

    def get_output_schema_to_validate(self, data: Dict[str, Any]) -> SchemaDataDTO:
        """
        Prepares the output schema data for validation.

        Args:
            data (Dict[str, Any]): Data to be validated.

        Returns:
            SchemaDataDTO: Schema data DTO.
        """
        return SchemaDataDTO(
            self.service,
            self.source,
            self.provider,
            SCHEMA_TYPE_RESPONSE,
            data
        )


class EventController(Controller):
    """
    Event controller class for handling specific event processing.

    Attributes:
        sd (ServiceDiscovery): Service discovery instance.
        job_handler (callable): Job handler function.
    """

    def __init__(
        self,
        sd: ServiceDiscovery,
        config: ConfigDTO,
        job_handler: callable,
        queue_active_jobs: asyncio.Queue,
        dbg: Union[debug.EnabledDebug, debug.DisabledDebug]
    ):
        """
        Initializes the EventController instance.

        Args:
            sd (ServiceDiscovery): Service discovery instance.
            config (ConfigDTO): Configuration data transfer object.
            job_handler (callable): Job handler function.
            queue_active_jobs (asyncio.Queue): Queue for managing active jobs.
            dbg (Union[debug.EnabledDebug, debug.DisabledDebug]): Debug instance.
        """
        super().__init__(config, queue_active_jobs, dbg)
        self.__sd = sd
        self.job_handler = job_handler
        self.__rmq_producer = RabbitMQProducer()

    @property
    def exchange_name(self) -> str:
        """
        Gets the exchange name for RabbitMQ.

        Returns:
            str: Exchange name.
        """
        return self.__sd.services_rabbitmq_exchange

    @property
    def processing_routing_key(self) -> str:
        """
        Gets the routing key for processing jobs.

        Returns:
            str: Processing routing key.
        """
        return "processing-job"

    @property
    def feedback_routing_key(self) -> str:
        """
        Gets the routing key for feedback messages.

        Returns:
            str: Feedback routing key.
        """
        return "service.feedback"

    async def __declare_exchange(self) -> None:
        """
        Declares the RabbitMQ exchange.
        """
        channel = await self.__rmq_producer.create_channel()
        await self.__rmq_producer.declare_exchange(channel, self.exchange_name)

    async def run(self, message) -> None:
        """
        Runs the event controller to process the incoming message.

        Args:
            message: Incoming message to process.
        """
        if not self.is_controller_active:
            logger.info(f"Controller for config_id {self.config_id} is not active")
            return

        await self.__rmq_producer.connect()
        event_order = self.get_event_order_dto(message)
        event_input, input_schema_version = await self.parse_message(event_order)

        await self.__declare_exchange()
        await self.__rmq_producer.send_message(
            exchange_name=self.__sd.services_rabbitmq_exchange,
            routing_key=self.processing_routing_key,
            message=serialize_to_json(event_order)
        )

        _, output_schema_version = await self.get_schema(SCHEMA_TYPE_MESSAGE)
        output_metadata = await self.get_metadata(event_order, input_schema_version, output_schema_version)
        job_result = await self.job_dispatcher(event_input, self.job_handler, output_metadata)
        await self.queue_active_jobs.get()
        is_valid_output = await self.validate_schema(job_result.data)
        if not is_valid_output:
            logger.error("Output schema is not valid")
            return

        output = serialize_to_json(job_result)
        logger.info("sleeping for 5 seconds...")
        time.sleep(5)
        logger.info(f"Service Output: {output}")
        await self.__rmq_producer.send_message(
            exchange_name=self.exchange_name,
            routing_key=self.feedback_routing_key,
            message=output
        )
        await message.ack()
        logger.info("Published message to service")
