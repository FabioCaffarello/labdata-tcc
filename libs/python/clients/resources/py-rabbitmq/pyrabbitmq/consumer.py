import asyncio

import aio_pika
from pylog.log import setup_logging
from pyrabbitmq.base import BaseRabbitMQ

logger = setup_logging(__name__)


class RabbitMQConsumer(BaseRabbitMQ):
    """A RabbitMQ consumer class that extends BaseRabbitMQ."""

    def __init__(self):
        """
        Initializes the RabbitMQConsumer instance.
        """
        super().__init__()

    async def listen(self, queue: aio_pika.Queue, callback: callable, timeout: float = None) -> None:
        """
        Listens to a RabbitMQ queue and calls the specified callback on message arrival.

        Args:
            queue (aio_pika.Queue): The RabbitMQ queue to listen to.
            callback (callable): The callback function to call when a message arrives.
            timeout (float, optional): The timeout for listening to the queue. If None, listens indefinitely.

        Raises:
            asyncio.TimeoutError: If the listening operation times out.
        """
        async def process_queue():
            async for message in queue.iterator():
                await callback(message)

        try:
            await asyncio.wait_for(process_queue(), timeout)
        except asyncio.TimeoutError:
            logger.debug("Listening to the queue timed out.")
