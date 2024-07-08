import asyncio
import os
import unittest
from unittest.mock import AsyncMock, MagicMock, patch

from pyrabbitmq.consumer import RabbitMQConsumer


class TestRabbitMQConsumer(unittest.IsolatedAsyncioTestCase):

    @patch.dict(os.environ, {"RABBITMQ_PORT_6572_TCP": "tcp://localhost:5672"})
    def setUp(self):
        self.consumer = RabbitMQConsumer()
        self.consumer.url = "amqp://user:password@localhost:5672/"

    @patch("aio_pika.connect_robust", new_callable=AsyncMock)
    async def test_listen(self, mock_connect):
        queue_mock = AsyncMock()
        message_mock = AsyncMock()

        # Mock the queue iterator to yield a message
        async def queue_iterator():
            yield message_mock

        queue_mock.iterator = MagicMock(return_value=queue_iterator())

        async def callback(message):
            pass  # Mock callback to handle the message

        # Run the listen method
        with patch("asyncio.wait_for", wraps=asyncio.wait_for):
            await self.consumer.listen(queue_mock, callback)

        # Ensure the queue's iterator was used
        queue_mock.iterator.assert_called_once()

    @patch("aio_pika.connect_robust", new_callable=AsyncMock)
    async def test_listen_with_timeout(self, mock_connect):
        queue_mock = AsyncMock()

        async def callback(message):
            pass  # Mock callback to handle the message

        # Mock the queue iterator to simulate delay
        async def queue_iterator():
            await asyncio.sleep(2)  # Simulate delay
            yield None

        queue_mock.iterator = MagicMock(return_value=queue_iterator())

        # Run the listen method with a timeout
        with self.assertRaises(asyncio.TimeoutError):
            await asyncio.wait_for(self.consumer.listen(queue_mock, callback), timeout=1)


if __name__ == "__main__":
    unittest.main()
