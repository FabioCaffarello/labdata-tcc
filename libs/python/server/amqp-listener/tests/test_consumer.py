import unittest
from unittest.mock import AsyncMock, MagicMock, patch
import asyncio

from dto_config_vault.output import ConfigDTO
from pydebug import debug
from amqp_listener.consumer import Consumer, EventConsumer


class TestConsumer(unittest.TestCase):

    @patch('amqp_listener.consumer.RabbitMQConsumer')
    @patch('amqp_listener.consumer.ServiceDiscovery')
    def setUp(self, mock_service_discovery, mock_rmq_consumer):
        self.sd = mock_service_discovery()
        self.rmq_consumer = mock_rmq_consumer()
        self.config = MagicMock(spec=ConfigDTO)
        self.config.provider = "test_provider"
        self.config.service = "test_service"
        self.config.source = "test_source"
        self.queue_active_jobs = asyncio.Queue()
        self.dbg = MagicMock(spec=debug.EnabledDebug)
        self.consumer = Consumer(self.sd, self.rmq_consumer, self.config, self.queue_active_jobs, self.dbg)

    def test_exchange_name(self):
        self.sd.services_rabbitmq_exchange = "test_exchange"
        self.assertEqual(self.consumer._Consumer__exchange_name, "test_exchange")

    def test_queue_name(self):
        expected_queue_name = f"input-queue.{self.config.provider}.{self.config.service}.{self.config.source}"
        self.assertEqual(self.consumer._Consumer__queue_name, expected_queue_name)

    def test_routing_key(self):
        expected_routing_key = (
            "input.ready-to-process."
            f"{self.config.provider}."
            f"{self.config.service}."
            f"{self.config.source}")
        self.assertEqual(self.consumer._Consumer__routing_key, expected_routing_key)

    @patch('amqp_listener.consumer.Consumer._Consumer__instantiate_controller')
    def test_instantiate_controller(self, mock_instantiate_controller):
        event_controller = MagicMock()
        job_handler = MagicMock()
        self.consumer._Consumer__instantiate_controller(event_controller, job_handler)
        mock_instantiate_controller.assert_called_once_with(event_controller, job_handler)

    @patch('amqp_listener.consumer.Consumer._Consumer__instantiate_controller')
    @patch('amqp_listener.consumer.RabbitMQConsumer.listen')
    @patch('amqp_listener.consumer.RabbitMQConsumer.create_queue')
    @patch('amqp_listener.consumer.RabbitMQConsumer.create_channel')
    async def test_start(self, mock_create_channel, mock_create_queue, mock_listen, mock_instantiate_controller):
        mock_channel = AsyncMock()
        mock_create_channel.return_value = mock_channel
        mock_queue = AsyncMock()
        mock_create_queue.return_value = mock_queue
        event_controller = AsyncMock()
        job_handler = AsyncMock()

        await self.consumer.start(event_controller, job_handler)

        mock_create_channel.assert_called_once()
        mock_create_queue.assert_called_once_with(
            mock_channel,
            self.consumer._Consumer__queue_name,
            self.consumer._Consumer__exchange_name,
            self.consumer._Consumer__routing_key
        )
        mock_listen.assert_called_once_with(mock_queue, mock_instantiate_controller.return_value.run)


class TestEventConsumer(unittest.TestCase):

    @patch('amqp_listener.consumer.RabbitMQConsumer')
    @patch('amqp_listener.consumer.ServiceDiscovery')
    def setUp(self, mock_service_discovery, mock_rmq_consumer):
        self.sd = mock_service_discovery()
        self.rmq_consumer = mock_rmq_consumer()
        self.config = MagicMock(spec=ConfigDTO)
        self.config.provider = "test_provider"
        self.config.service = "test_service"
        self.config.source = "test_source"
        self.queue_active_jobs = asyncio.Queue()
        self.dbg = MagicMock(spec=debug.EnabledDebug)
        self.event_consumer = EventConsumer(self.sd, self.rmq_consumer, self.config, self.queue_active_jobs, self.dbg)

    @patch('amqp_listener.consumer.Consumer.start')
    async def test_run(self, mock_start):
        event_controller = AsyncMock()
        job_handler = AsyncMock()

        await self.event_consumer.run(event_controller, job_handler)

        mock_start.assert_called_once_with(event_controller, job_handler)


if __name__ == '__main__':
    unittest.main()
