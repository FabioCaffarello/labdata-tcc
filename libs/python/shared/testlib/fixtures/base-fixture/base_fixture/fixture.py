import asyncio
import json
import os
import unittest
import uuid
from enum import Enum
from pathlib import Path
from typing import Any, Dict, List, Union

from cli_config_vault.client import (AsyncPyConfigVaultClient,
                                     async_py_config_vault_client)
from cli_schema_vault.client import (AsyncPySchemaVaultClient,
                                     async_py_schema_vault_client)
from dto_config_vault.output import ConfigDTO
from dto_events_router.input import ServiceFeedBackDTO
from dto_events_router.output import ProcessOrderDTO
from dto_schema_vault.output import SchemaDTO
from pylog.log import setup_logging
from pyminio.client import MinioClient, minio_client
from pymongodb.client import drop_database
from pyrabbitmq.consumer import RabbitMQConsumer
from pyrabbitmq.producer import RabbitMQProducer
from pyserializer.serializer import serialize_to_dataclass, serialize_to_json
from subprocessd.async_subprocessd import SubprocessDAsync

logger = setup_logging(__name__)


class MappingAPIClients(Enum):
    ConfigVault = async_py_config_vault_client
    SchemaVault = async_py_schema_vault_client


class BaseTestsFixture(unittest.IsolatedAsyncioTestCase):
    """
    Base class for setting up and tearing down test environments asynchronously,
    pushing jobs to queues, and verifying job results.
    """

    async def asyncSetUp(self):
        """
        Sets up the test environment asynchronously.

        Returns:
            None
        """
        self.__set_static_variables()
        logger.debug("Setting up test.")
        self.cleanup_database()
        self.rmq_consumer = await self.__get_rabbitmq_consumer()
        self.rmq_producer = await self.__get_rabbitmq_producer()
        await self.__set_configs_by_service()
        await self.__purge_all_queues()
        self.__create_buckets()
        args = self._get_service_process_args()
        args.append("--enable-debug-storage")
        args.append("--debug-storage-dir")
        args.append(self.get_debug_storage_dir())
        self._subprocessd = SubprocessDAsync(args, self.get_debug_storage_dir())
        await self._subprocessd.start()

    async def asyncTearDown(self):
        """
        Tears down the test environment asynchronously.

        Returns:
            None
        """
        await self._subprocessd.stop()
        self.cleanup_database()
        await self.__purge_all_queues()
        await self.rmq_producer.close_connection()
        await self.rmq_consumer.close_connection()

    def __set_static_variables(self):
        """
        Sets static variables for the test fixture.

        Returns:
            None
        """
        self.__DB_CONFIG_VAULT_NAME = os.getenv("DB_CONFIG_VAULT_NAME")
        self.__DB_SCHEMA_VAULT_NAME = os.getenv("DB_SCHEMA_VAULT_NAME")
        self.__SERVICE_NAME = os.getenv("SERVICE_NAME")
        self.__TESTS_CONFIGS_DIR = "/app/tests/.configs"
        self.__TESTS_DEBUG_DIR = "/app/tests/debug"
        self.__TEST_CONFIG_PATH = "configs"
        self.__TEST_SCHEMA_PATH = "schemas"
        self.__JSON_EXTENTION_REGEX_PATTERN = "*.json"
        self.__BASE_INPUT_QUEUE = "input-queue"
        self.__BASE_OUTPUT_QUEUE = "output-queue"
        self.__BASE_ROUTING_KEY = "input.ready-to-process"
        self.__EXCHANGE_NAME = "services"
        self.__PROCESS_ORDER_STAGE = "pre-processed"
        self.__FEEDBACK_ROUTING_KEY = "service.feedback"
        self.__mapping_api_clients = MappingAPIClients
        self.__all_configs: List[ConfigDTO] = []
        self.__all_schemas: List[SchemaDTO] = []
        self.queue = asyncio.Queue()

    def __get_suite_name(self):
        """
        Gets the suite name for the test.

        Returns:
            str: The suite name.
        """
        return self.__class__.__name__

    def __get_test_name(self):
        """
        Gets the test name for the test.

        Returns:
            str: The test name.
        """
        return self._testMethodName

    def get_debug_storage_dir(self):
        """
        Gets the debug storage directory.

        Returns:
            str: The debug storage directory path.
        """
        return "{debug_dir}/{suite_name}/{test_name}".format(
            debug_dir=self.__TESTS_DEBUG_DIR,
            suite_name=self.__get_suite_name(),
            test_name=self.__get_test_name()
        )

    def _get_service_process_args(self):
        """
        Gets the service process arguments.

        Raises:
            AssertionError: If not implemented by the subclass.

        Returns:
            list: The service process arguments.
        """
        self.fail(
            "Classes inheriting from BaseTestsFixture must implement _get_service_process_args method"
        )

    def cleanup_database(self) -> None:
        """
        Cleans up the database by dropping the config and schema vault databases.

        Returns:
            None
        """
        self.__drop_database(self.__DB_CONFIG_VAULT_NAME)
        self.__drop_database(self.__DB_SCHEMA_VAULT_NAME)

    def __drop_database(self, database_name: str) -> None:
        """
        Drops a database.

        Args:
            database_name (str): The name of the database to drop.

        Returns:
            None
        """
        drop_database(database_name=database_name)

    async def __get_rabbitmq_consumer(self) -> RabbitMQConsumer:
        """
        Gets a RabbitMQ consumer.

        Returns:
            RabbitMQConsumer: The RabbitMQ consumer.
        """
        rmq_consumer = RabbitMQConsumer()
        await rmq_consumer.connect()
        return rmq_consumer

    async def __get_rabbitmq_producer(self) -> RabbitMQProducer:
        """
        Gets a RabbitMQ producer.

        Returns:
            RabbitMQProducer: The RabbitMQ producer.
        """
        rmq_producer = RabbitMQProducer()
        await rmq_producer.connect()
        return rmq_producer

    def __get_file_posix_path(self, config_type: str) -> Path:
        """
        Gets the POSIX path for a config file.

        Args:
            config_type (str): The type of config file.

        Returns:
            Path: The POSIX path for the config file.
        """
        return Path(self.__TESTS_CONFIGS_DIR).joinpath(config_type)

    def __find_local_config_files(self, config_type: str) -> List[Path]:
        """
        Finds local config files.

        Args:
            config_type (str): The type of config files to find.

        Returns:
            List[Path]: A list of paths to the config files.
        """
        configs_path = self.__get_file_posix_path(config_type)
        logger.info(f"Looking for {config_type} files in {configs_path}")
        return list(configs_path.glob(self.__JSON_EXTENTION_REGEX_PATTERN))

    def __read_config_file(self, config_path: Path) -> Dict[str, Any]:
        """
        Reads a config file.

        Args:
            config_path (Path): The path to the config file.

        Returns:
            Dict[str, Any]: The config file data.
        """
        with open(config_path, "r") as config_file:
            return json.load(config_file)

    async def __load_config_and_post_to_api(
        self,
        config_type: str,
        config_client: Union[AsyncPyConfigVaultClient, AsyncPySchemaVaultClient]
    ) -> Union[ConfigDTO, SchemaDTO]:
        """
        Loads a config file and posts it to the API.

        Args:
            config_type (str): The type of config file.
            config_client (Union[AsyncPyConfigVaultClient, AsyncPySchemaVaultClient]): The config client.

        Returns:
            Union[ConfigDTO, SchemaDTO]: The config or schema DTO.
        """
        all_configs_path = self.__find_local_config_files(config_type)
        all_configs = []
        logger.info(f"Found {len(all_configs_path)} {config_type} files")
        for config_path in all_configs_path:
            config_data = self.__read_config_file(config_path)
            logger.info(f"Creating {config_type} with data: {config_data}")

            config_response = await config_client.create(config_data)
            all_configs.append(config_response)
        return all_configs

    async def __set_configs_by_service(self) -> None:
        """
        Sets configs by service.

        Returns:
            None
        """
        self.__all_configs = await self.__load_config_and_post_to_api(
            self.__TEST_CONFIG_PATH,
            self.__mapping_api_clients.ConfigVault()
        )
        self.__all_schemas = await self.__load_config_and_post_to_api(
            self.__TEST_SCHEMA_PATH,
            self.__mapping_api_clients.SchemaVault()
        )

    def __get_queue_name(self, provider: str, service: str, source: str, output: bool = False) -> str:
        """
        Gets the queue name.

        Args:
            provider (str): The provider.
            service (str): The service.
            source (str): The source.
            output (bool, optional): Whether the queue is an output queue. Defaults to False.

        Returns:
            str: The queue name.
        """
        if output:
            return f"{self.__BASE_OUTPUT_QUEUE}.{provider}.{service}.{source}"
        return f"{self.__BASE_INPUT_QUEUE}.{provider}.{service}.{source}"

    def __get_routing_key(self, provider: str, service: str, source: str) -> str:
        """
        Gets the routing key.

        Args:
            provider (str): The provider.
            service (str): The service.
            source (str): The source.

        Returns:
            str: The routing key.
        """
        return f"{self.__BASE_ROUTING_KEY}.{provider}.{service}.{source}"

    async def __purge_queue(self, provider: str, service: str, source: str) -> None:
        """
        Purges a queue.

        Args:
            provider (str): The provider.
            service (str): The service.
            source (str): The source.

        Returns:
            None
        """
        queue_name = self.__get_queue_name(provider, service, source)
        logger.info(f"Purging queue {queue_name}")
        await self.rmq_consumer.purge_queue(queue_name)

    async def __purge_all_queues(self) -> None:
        """
        Purges all queues.

        Returns:
            None
        """
        for config in self.__all_configs:
            await self.__purge_queue(config.provider, config.service, config.source)

    def __create_bucket(self, minio_client: MinioClient, bucket_name: str) -> None:
        """
        Creates a bucket in Minio.

        Args:
            minio_client (MinioClient): The Minio client.
            bucket_name (str): The name of the bucket to create.

        Returns:
            None
        """
        try:
            minio_client.create_bucket(bucket_name)
        except Exception as e:
            logger.warning(f"Error creating bucket {bucket_name}: {e}")

    def __get_bucket_name(self, provider: str, source: str):
        """
        Gets the bucket name.

        Args:
            provider (str): The provider.
            source (str): The source.

        Returns:
            str: The bucket name.
        """
        return f"{provider}-{source}"

    def __create_buckets(self) -> None:
        """
        Creates buckets in Minio.

        Returns:
            None
        """
        minio = minio_client()
        for config in self.__all_configs:
            bucket_name = self.__get_bucket_name(config.provider, config.source)
            self.__create_bucket(minio, bucket_name)

    def __generate_input(self, input_data: Dict[str, Any], provider: str, service: str, source: str) -> ProcessOrderDTO:
        """
        Generates input data for a job.

        Args:
            input_data (Dict[str, Any]): The input data.
            provider (str): The provider.
            service (str): The service.
            source (str): The source.

        Returns:
            ProcessOrderDTO: The process order DTO.
        """
        fake_id = str(uuid.uuid4().hex)
        return ProcessOrderDTO(
            order_id=fake_id,
            provider=provider,
            service=service,
            source=source,
            data=input_data,
            processing_id=fake_id,
            stage=self.__PROCESS_ORDER_STAGE,
            input_id=fake_id
        )

    async def _callback(self, message):
        """
        Callback function for handling messages from the queue.

        Args:
            message: The message from the queue.

        Returns:
            None
        """
        try:
            output = json.loads(message.body.decode())
            logger.info(f"Received message: {output}")
            output_dto = serialize_to_dataclass(output, ServiceFeedBackDTO)
            logger.info(f"Putting message in queue: {output_dto}")
            await self.queue.put(output)
            await message.ack()
        except Exception as e:
            logger.error(f"Error in callback: {e}")
            await message.nack()

    async def pop_job(self, provider="", source="", timeout=60):
        """
        Pops a job from the queue.

        Args:
            provider (str, optional): The provider. Defaults to "".
            source (str, optional): The source. Defaults to "".
            timeout (int, optional): The timeout for popping the job. Defaults to 60.

        Returns:
            None
        """
        queue_name = self.__get_queue_name(provider, self.__SERVICE_NAME, source, output=True)
        logger.info(f"Listening to queue {queue_name}")
        channel = await self.rmq_consumer.create_channel()
        await self.rmq_consumer.declare_exchange(channel, self.__EXCHANGE_NAME)
        queue = await self.rmq_consumer.create_queue(
            channel=channel,
            queue_name=f"{queue_name}",
            exchange_name=self.__EXCHANGE_NAME,
            routing_key=self.__FEEDBACK_ROUTING_KEY
        )

        logger.info(
            f"Queue {queue_name} created and bound to exchange {self.__EXCHANGE_NAME} "
            f"with routing key {self.__FEEDBACK_ROUTING_KEY}"
        )
        await self.rmq_consumer.listen(queue, self._callback, timeout=timeout)

    async def push_job(self, input_data: Dict[str, Any], provider: str = "", source: str = ""):
        """
        Pushes a job to the queue.

        Args:
            input_data (Dict[str, Any]): The input data.
            provider (str, optional): The provider. Defaults to "".
            source (str, optional): The source. Defaults to "".

        Returns:
            None
        """
        queue_name = self.__get_queue_name(provider, self.__SERVICE_NAME, source)
        routing_key = self.__get_routing_key(provider, self.__SERVICE_NAME, source)
        logger.info(f"Pushing job to queue {queue_name}, with routing key {routing_key}")

        input_dto = self.__generate_input(input_data, provider, self.__SERVICE_NAME, source)
        channel = await self.rmq_producer.create_channel()
        await self.rmq_producer.declare_exchange(channel, self.__EXCHANGE_NAME)
        await self.rmq_producer.publish_message(
            exchange_name=self.__EXCHANGE_NAME,
            routing_key=routing_key,
            message=serialize_to_json(input_dto)
        )
        logger.info(f"Job pushed to queue {queue_name}, with routing key {routing_key}")
