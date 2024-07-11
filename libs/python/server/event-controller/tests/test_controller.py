import asyncio
import unittest
from unittest.mock import AsyncMock, MagicMock, patch

from dto_config_vault.output import ConfigDTO
from dto_events_router.input import MetadataDTO, ServiceFeedBackDTO
from dto_events_router.output import ProcessOrderDTO
from dto_schema_vault.input import SchemaDataDTO
from dto_schema_vault.output import SchemaDTO
from event_controller.controller import Controller, EventController
from pydebug import debug


class TestController(unittest.TestCase):

    @patch('event_controller.controller.async_py_schema_vault_client')
    def setUp(self, mock_schema_vault_client):
        self.mock_schema_vault_client = mock_schema_vault_client
        self.mock_schema_vault_client.return_value = MagicMock()
        self.config = MagicMock(spec=ConfigDTO)
        self.config.service = "test_service"
        self.config.source = "test_source"
        self.config.provider = "test_provider"
        self.queue_active_jobs = asyncio.Queue()
        self.dbg = MagicMock(spec=debug.EnabledDebug)
        self.controller = Controller(self.config, self.queue_active_jobs, self.dbg)

    def test_is_controller_active(self):
        self.config.active = True
        self.assertTrue(self.controller.is_controller_active)
        self.config.active = False
        self.assertFalse(self.controller.is_controller_active)

    def test_properties(self):
        self.config.config_id = "test_id"
        self.config.provider = "test_provider"
        self.config.service = "test_service"
        self.config.source = "test_source"

        self.assertEqual(self.controller.config_id, "test_id")
        self.assertEqual(self.controller.provider, "test_provider")
        self.assertEqual(self.controller.service, "test_service")
        self.assertEqual(self.controller.source, "test_source")

    def test_get_json_schema(self):
        schema = MagicMock(spec=SchemaDTO)
        schema.json_schema = {"key": "value"}
        self.assertEqual(self.controller.get_json_schema(schema), {"key": "value"})

    def test_get_schema_version_id(self):
        schema = MagicMock(spec=SchemaDTO)
        schema.schema_version_id = "version_id"
        self.assertEqual(self.controller.get_schema_version_id(schema), "version_id")

    @patch('event_controller.controller.Controller.get_schema_version_id')
    @patch('event_controller.controller.Controller.get_json_schema')
    async def test_get_schema(self, mock_get_json_schema, mock_get_schema_version_id):
        mock_get_json_schema.return_value = {"key": "value"}
        mock_get_schema_version_id.return_value = "version_id"
        schema_type = "input"

        schema, version_id = await self.controller.get_schema(schema_type)

        self.assertEqual(schema, {"key": "value"})
        self.assertEqual(version_id, "version_id")

    @patch('event_controller.controller.Controller.get_output_schema_to_validate')
    async def test_validate_schema(self, mock_get_output_schema_to_validate):
        mock_get_output_schema_to_validate.return_value = MagicMock(spec=SchemaDataDTO)
        mock_client = self.mock_schema_vault_client.return_value
        mock_client.validate_schema.return_value = True
        job_data = {"data": "value"}

        is_valid = await self.controller.validate_schema(job_data)

        self.assertTrue(is_valid)

    def test_read_message(self):
        message = MagicMock()
        message.body.decode.return_value = '{"key": "value"}'
        result = self.controller._read_message(message)
        self.assertEqual(result, {"key": "value"})

    def test_read_message_invalid_json(self):
        message = MagicMock()
        message.body.decode.return_value = '{"key": "value"'
        with self.assertRaises(ValueError):
            self.controller._read_message(message)

    def test_read_message_empty_body(self):
        message = MagicMock()
        message.body.decode.return_value = ''
        with self.assertRaises(ValueError):
            self.controller._read_message(message)

    @patch('event_controller.controller.serialize_to_dataclass')
    def test_get_event_order_dto(self, mock_serialize_to_dataclass):
        message = MagicMock()
        message.body.decode.return_value = '{"key": "value"}'
        mock_serialize_to_dataclass.return_value = ProcessOrderDTO(
            order_id="order_id",
            processing_id="processing_id",
            service="service",
            source="source",
            provider="provider",
            stage="stage",
            input_id="input_id",
            data={"data": "value"}
        )

        result = self.controller.get_event_order_dto(message)
        self.assertIsInstance(result, ProcessOrderDTO)

    def test_get_message_data(self):
        process_order_dto = ProcessOrderDTO(
            order_id="order_id",
            processing_id="processing_id",
            service="service",
            source="source",
            provider="provider",
            stage="stage",
            input_id="input_id",
            data={"data": "value"}
        )
        result = self.controller._get_message_data(process_order_dto)
        self.assertEqual(result, {"data": "value"})

    @patch('event_controller.controller.warlock_serializer.serialize_to_dataclass')
    @patch('event_controller.controller.serialize_to_dict')
    @patch('event_controller.controller.Controller.get_schema')
    async def test_parse_message(self, mock_get_schema, mock_serialize_to_dict, mock_serialize_to_dataclass):
        event_order = MagicMock(spec=ProcessOrderDTO)
        mock_get_schema.return_value = ({"schema": "value"}, "version_id")
        mock_serialize_to_dict.return_value = {"schema": "value"}
        mock_serialize_to_dataclass.return_value = MagicMock()

        result, version_id = await self.controller.parse_message(event_order)
        self.assertEqual(version_id, "version_id")
        mock_serialize_to_dataclass.assert_called_once()

    @patch('event_controller.controller.Controller.get_metadata')
    @patch('event_controller.controller.Controller.job_dispatcher')
    async def test_job_dispatcher(self, mock_job_dispatcher, mock_get_metadata):
        event_input = MagicMock()
        job_handler = AsyncMock()
        metadata = MagicMock(spec=MetadataDTO)

        job_result = MagicMock(spec=ServiceFeedBackDTO)
        mock_job_dispatcher.return_value = job_result

        result = await self.controller.job_dispatcher(event_input, job_handler, metadata)
        self.assertEqual(result, job_result)

    @patch('event_controller.controller.MetadataDTO')
    async def test_get_metadata(self, mock_MetadataDTO):
        event_order = MagicMock(spec=ProcessOrderDTO)
        schema_version_input = "input_version"
        schema_version_output = "output_version"

        metadata = await self.controller.get_metadata(event_order, schema_version_input, schema_version_output)
        self.assertIsInstance(metadata, MetadataDTO)
        mock_MetadataDTO.assert_called_once()

    def test_get_output_schema_to_validate(self):
        data = {"key": "value"}
        schema_data = self.controller.get_output_schema_to_validate(data)
        self.assertIsInstance(schema_data, SchemaDataDTO)
        self.assertEqual(schema_data.data, data)


class TestEventController(unittest.TestCase):

    @patch('event_controller.controller.async_py_schema_vault_client')
    @patch('event_controller.controller.RabbitMQProducer')
    def setUp(self, mock_rabbitmq_producer, mock_schema_vault_client):
        self.mock_schema_vault_client = mock_schema_vault_client
        self.mock_rabbitmq_producer = mock_rabbitmq_producer
        self.mock_schema_vault_client.return_value = MagicMock()
        self.mock_rabbitmq_producer.return_value = MagicMock()
        self.config = MagicMock(spec=ConfigDTO)
        self.config.service = "test_service"
        self.config.source = "test_source"
        self.config.provider = "test_provider"
        self.queue_active_jobs = asyncio.Queue()
        self.dbg = MagicMock(spec=debug.EnabledDebug)
        self.sd = MagicMock()
        self.sd.services_rabbitmq_exchange = "test_exchange"
        self.job_handler = AsyncMock()
        self.event_controller = EventController(
            sd=self.sd,
            config=self.config,
            job_handler=self.job_handler,
            queue_active_jobs=self.queue_active_jobs,
            dbg=self.dbg
        )

    def test_exchange_name(self):
        self.assertEqual(self.event_controller.exchange_name, "test_exchange")

    def test_processing_routing_key(self):
        self.assertEqual(self.event_controller.processing_routing_key, "processing-job")

    def test_feedback_routing_key(self):
        self.assertEqual(self.event_controller.feedback_routing_key, "service.feedback")

    @patch('event_controller.controller.Controller.is_controller_active', new_callable=MagicMock)
    async def test_run_controller_not_active(self, mock_is_controller_active):
        mock_is_controller_active.return_value = False
        message = MagicMock()

        await self.event_controller.run(message)
        self.event_controller._EventController__rmq_producer.connect.assert_not_called()
        message.ack.assert_not_called()

    @patch('event_controller.controller.Controller.get_event_order_dto')
    @patch('event_controller.controller.Controller.parse_message')
    @patch('event_controller.controller.Controller.get_schema')
    @patch('event_controller.controller.Controller.get_metadata')
    @patch('event_controller.controller.Controller.job_dispatcher')
    @patch('event_controller.controller.Controller.validate_schema')
    async def test_run_controller_active(
        self,
        mock_validate_schema,
        mock_job_dispatcher,
        mock_get_metadata,
        mock_get_schema,
        mock_parse_message,
        mock_get_event_order_dto
    ):
        mock_is_controller_active = patch(
            'event_controller.controller.Controller.is_controller_active',
            new_callable=MagicMock
        )
        mock_is_controller_active.start().return_value = True
        message = MagicMock()
        mock_get_event_order_dto.return_value = MagicMock(spec=ProcessOrderDTO)
        mock_parse_message.return_value = (MagicMock(), "input_version")
        mock_get_schema.return_value = (MagicMock(), "output_version")
        mock_get_metadata.return_value = MagicMock(spec=MetadataDTO)
        mock_job_dispatcher.return_value = MagicMock(spec=ServiceFeedBackDTO)
        mock_validate_schema.return_value = True

        await self.event_controller.run(message)

        self.event_controller._EventController__rmq_producer.connect.assert_called_once()
        self.event_controller._EventController__rmq_producer.send_message.assert_called()
        message.ack.assert_called_once()

    @patch('event_controller.controller.RabbitMQProducer.create_channel')
    async def test_declare_exchange(self, mock_create_channel):
        mock_channel = AsyncMock()
        mock_create_channel.return_value = mock_channel

        await self.event_controller._EventController__declare_exchange()

        self.event_controller._EventController__rmq_producer.declare_exchange.assert_called_once_with(
            mock_channel,
            "test_exchange"
        )

    @patch('event_controller.controller.Controller.validate_schema')
    async def test_run_controller_active_invalid_schema(self, mock_validate_schema):
        mock_is_controller_active = patch(
            'event_controller.controller.Controller.is_controller_active',
            new_callable=MagicMock
        )
        mock_is_controller_active.start().return_value = True
        message = MagicMock()
        mock_validate_schema.return_value = False

        await self.event_controller.run(message)

        self.event_controller._EventController__rmq_producer.connect.assert_called_once()
        self.event_controller._EventController__rmq_producer.send_message.assert_called()
        message.ack.assert_called_once()

        mock_validate_schema.assert_called_once()
        self.event_controller._EventController__rmq_producer.send_message.assert_any_call(
            exchange_name=self.event_controller.exchange_name,
            routing_key=self.event_controller.feedback_routing_key,
            message='{"data": "invalid"}'
        )

    @patch('event_controller.controller.Controller._read_message')
    @patch('event_controller.controller.serialize_to_dataclass')
    def test_get_event_order_dto_invalid_message(self, mock_serialize_to_dataclass, mock_read_message):
        message = MagicMock()
        mock_read_message.side_effect = ValueError("Invalid message body")
        with self.assertRaises(ValueError):
            self.event_controller.get_event_order_dto(message)

    @patch('event_controller.controller.Controller.get_output_schema_to_validate')
    async def test_validate_schema_invalid(self, mock_get_output_schema_to_validate):
        mock_get_output_schema_to_validate.return_value = MagicMock(spec=SchemaDataDTO)
        mock_client = self.mock_schema_vault_client.return_value
        mock_client.validate_schema.return_value = False
        job_data = {"data": "value"}

        is_valid = await self.controller.validate_schema(job_data)

        self.assertFalse(is_valid)

    @patch('event_controller.controller.Controller.get_output_schema_to_validate')
    async def test_validate_schema_exception(self, mock_get_output_schema_to_validate):
        mock_get_output_schema_to_validate.side_effect = Exception("Unexpected Error")
        job_data = {"data": "value"}

        with self.assertRaises(Exception):
            await self.controller.validate_schema(job_data)

    @patch('event_controller.controller.Controller.get_event_order_dto')
    @patch('event_controller.controller.Controller.parse_message')
    @patch('event_controller.controller.Controller.get_schema')
    @patch('event_controller.controller.Controller.get_metadata')
    @patch('event_controller.controller.Controller.job_dispatcher')
    @patch('event_controller.controller.Controller.validate_schema')
    async def test_run_controller_exception_in_declare_exchange(
        self,
        mock_validate_schema,
        mock_job_dispatcher,
        mock_get_metadata,
        mock_get_schema,
        mock_parse_message,
        mock_get_event_order_dto
    ):
        mock_is_controller_active = patch(
            'event_controller.controller.Controller.is_controller_active',
            new_callable=MagicMock
        )
        mock_is_controller_active.start().return_value = True
        message = MagicMock()
        mock_get_event_order_dto.return_value = MagicMock(spec=ProcessOrderDTO)
        mock_parse_message.return_value = (MagicMock(), "input_version")
        mock_get_schema.return_value = (MagicMock(), "output_version")
        mock_get_metadata.return_value = MagicMock(spec=MetadataDTO)
        mock_job_dispatcher.return_value = MagicMock(spec=ServiceFeedBackDTO)
        mock_validate_schema.return_value = True
        self.event_controller._EventController__declare_exchange = AsyncMock(
            side_effect=Exception("Declare Exchange Error")
        )

        await self.event_controller.run(message)

        self.event_controller._EventController__rmq_producer.connect.assert_called_once()
        message.ack.assert_not_called()


if __name__ == '__main__':
    unittest.main()
