import os
import unittest
from unittest.mock import AsyncMock, patch

import aio_pika
from pylog.log import setup_logging
from pyrabbitmq.base import BaseRabbitMQ

logger = setup_logging("pyrabbitmq.base")


class TestBaseRabbitMQ(unittest.IsolatedAsyncioTestCase):

    @patch.dict(os.environ, {"RABBITMQ_PORT_6572_TCP": "tcp://localhost:5672"})
    def setUp(self):
        self.rabbitmq = BaseRabbitMQ()
        self.rabbitmq.url = "amqp://user:password@localhost:5672/"

    @patch("aio_pika.connect_robust", new_callable=AsyncMock)
    async def test_connect_success(self, mock_connect):
        await self.rabbitmq.connect()
        mock_connect.assert_called_once()
        self.assertIsNotNone(self.rabbitmq.connection)

    @patch("aio_pika.connect_robust", new_callable=AsyncMock, side_effect=Exception("Connection error"))
    async def test_connect_failure(self, mock_connect):
        with self.assertRaises(RuntimeError):
            await self.rabbitmq.connect(max_retries=2)
        self.assertIsNone(self.rabbitmq.connection)

    @patch("aio_pika.connect_robust", new_callable=AsyncMock)
    async def test_purge_queue(self, mock_connect):
        mock_channel = AsyncMock()
        mock_queue = AsyncMock()
        mock_channel.declare_queue.return_value = mock_queue
        self.rabbitmq.connection = AsyncMock()
        self.rabbitmq.connection.channel.return_value = mock_channel

        await self.rabbitmq.purge_queue("test_queue")

        mock_channel.declare_queue.assert_called_once_with("test_queue", durable=True)
        mock_queue.purge.assert_called_once()

    @patch("aio_pika.connect_robust", new_callable=AsyncMock)
    async def test_delete_queue(self, mock_connect):
        mock_channel = AsyncMock()
        mock_queue = AsyncMock()
        mock_channel.declare_queue.return_value = mock_queue
        self.rabbitmq.connection = AsyncMock()
        self.rabbitmq.connection.channel.return_value = mock_channel

        await self.rabbitmq.delete_queue("test_queue")

        mock_channel.declare_queue.assert_called_once_with("test_queue", durable=True)
        mock_queue.delete.assert_called_once()

    @patch("aio_pika.connect_robust", new_callable=AsyncMock)
    async def test_publish_message(self, mock_connect):
        mock_channel = AsyncMock()
        self.rabbitmq.connection = AsyncMock()
        self.rabbitmq.connection.channel.return_value = mock_channel
        self.rabbitmq.exchange = AsyncMock()

        await self.rabbitmq.publish_message("test_exchange", "test_key", "test_message")

        self.rabbitmq.exchange.publish.assert_called_once()
        args, kwargs = self.rabbitmq.exchange.publish.call_args
        self.assertEqual(kwargs['routing_key'], "test_key")
        self.assertEqual(args[0].body, b"test_message")

    @patch.dict(os.environ, {"RABBITMQ_PORT_6572_TCP": "tcp://localhost:5672"})
    async def test_on_connection_error(self):
        amqp_url = "amqp://user:password@localhost:5672/"
        with self.assertLogs("pyrabbitmq.base", level='ERROR') as log:
            self.rabbitmq.on_connection_error(Exception("Test error"))
            error_message = "Connection error: Test error"
            self.assertTrue(any(error_message in message for message in log.output))
            self.assertTrue(any(f"Connection parameters: {amqp_url}" in message for message in log.output))

    @patch("aio_pika.connect_robust", new_callable=AsyncMock)
    async def test_create_channel(self, mock_connect):
        self.rabbitmq.connection = AsyncMock()
        mock_channel = AsyncMock()
        self.rabbitmq.connection.channel.return_value = mock_channel

        channel = await self.rabbitmq.create_channel()

        self.rabbitmq.connection.channel.assert_called_once()
        mock_channel.set_qos.assert_called_once_with(prefetch_count=1)
        self.assertEqual(channel, mock_channel)

    @patch("aio_pika.connect_robust", new_callable=AsyncMock)
    async def test_declare_exchange(self, mock_connect):
        mock_channel = AsyncMock()
        await self.rabbitmq.declare_exchange(mock_channel, "test_exchange")
        mock_channel.declare_exchange.assert_called_once_with(
            "test_exchange",
            aio_pika.ExchangeType.TOPIC,
            durable=True
        )
        self.assertIsNotNone(self.rabbitmq.exchange)

    @patch("aio_pika.connect_robust", new_callable=AsyncMock)
    async def test_create_queue(self, mock_connect):
        mock_channel = AsyncMock()
        mock_queue = AsyncMock()
        mock_channel.declare_queue.return_value = mock_queue
        self.rabbitmq.exchange = AsyncMock()

        queue = await self.rabbitmq.create_queue(mock_channel, "test_queue", "test_exchange", "test_key")

        mock_channel.declare_queue.assert_called_once_with("test_queue", durable=True)
        mock_queue.bind.assert_called_once_with(self.rabbitmq.exchange, "test_key")
        self.assertEqual(queue, mock_queue)

    @patch("aio_pika.connect_robust", new_callable=AsyncMock)
    async def test_close_connection(self, mock_connect):
        self.rabbitmq.connection = AsyncMock()
        mock_connection = self.rabbitmq.connection
        mock_connection.close = AsyncMock()  # Ensure close method is present and mock it
        await self.rabbitmq.close_connection()
        mock_connection.close.assert_called_once()  # Check that close was called once
        self.assertIsNone(self.rabbitmq.connection)  # Check that connection is set to None


if __name__ == "__main__":
    unittest.main()
