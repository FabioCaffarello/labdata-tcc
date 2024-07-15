import os
import unittest
from unittest.mock import AsyncMock, patch

from pyrabbitmq.producer import RabbitMQProducer


class TestRabbitMQProducer(unittest.IsolatedAsyncioTestCase):

    @patch.dict(os.environ, {"RABBITMQ_PORT_6572_TCP": "tcp://localhost:5672"})
    def setUp(self):
        self.producer = RabbitMQProducer()
        self.producer.url = "amqp://user:password@localhost:5672/"

    @patch("aio_pika.connect_robust", new_callable=AsyncMock)
    async def test_send_message(self, mock_connect):
        self.producer.connection = AsyncMock()
        self.producer.exchange = AsyncMock()

        # Mock the publish_message method to ensure it is called correctly
        with patch.object(self.producer, 'publish_message', new_callable=AsyncMock) as mock_publish_message:
            await self.producer.send_message("test_exchange", "test_key", "test_message")
            mock_publish_message.assert_called_once_with("test_exchange", "test_key", "test_message")

    @patch("aio_pika.connect_robust", new_callable=AsyncMock)
    async def test_send_message_no_exchange(self, mock_connect):
        self.producer.connection = AsyncMock()
        self.producer.exchange = None

        # Mock the publish_message method to raise an exception if exchange is None
        with patch.object(self.producer, 'publish_message', new_callable=AsyncMock) as mock_publish_message:
            mock_publish_message.side_effect = RuntimeError("Exchange 'test_exchange' not declared")
            with self.assertRaises(RuntimeError):
                await self.producer.send_message("test_exchange", "test_key", "test_message")
            mock_publish_message.assert_called_once_with("test_exchange", "test_key", "test_message")


if __name__ == "__main__":
    unittest.main()
