import asyncio
import urllib.parse
from typing import Optional

import aio_pika
from pylog.log import setup_logging
from pysd.sd import new_from_env

logger = setup_logging(__name__)


class BaseRabbitMQ:
    """A base class for interacting with RabbitMQ."""

    def __init__(self) -> None:
        """
        Initializes the BaseRabbitMQ instance, setting up the service discovery
        and connection attributes.
        """
        self._sd = new_from_env()
        self.url = self._sd.rabbitmq_endpoint
        self.connection: Optional[aio_pika.RobustConnection] = None
        self.exchange: Optional[aio_pika.Exchange] = None

    async def _connect(self) -> None:
        """
        Connects to RabbitMQ using robust connection settings.

        Raises:
            aio_pika.exceptions.AMQPConnectionError: If there is an error connecting to RabbitMQ.
        """
        parsed_url = urllib.parse.urlparse(self.url)
        self.connection = await aio_pika.connect_robust(
            host=parsed_url.hostname,
            port=parsed_url.port,
            login=parsed_url.username,
            password=parsed_url.password,
            timeout=100,
            heartbeat=60,
        )
        logger.info("Connected to RabbitMQ")

    async def connect(self, max_retries: int = 5) -> None:
        """
        Tries to connect to RabbitMQ, retrying on failure.

        Args:
            max_retries (int): Maximum number of retries before giving up. Default is 5.

        Raises:
            RuntimeError: If unable to connect after max retries.
        """
        retries = 0
        while retries < max_retries:
            try:
                await self._connect()
                break
            except Exception as err:
                logger.error('Could not connect to RabbitMQ, retrying in 2 seconds...')
                self.on_connection_error(err)
                await asyncio.sleep(2)
                retries += 1
        else:
            raise RuntimeError("Failed to connect to RabbitMQ after multiple retries")

    async def purge_queue(self, queue_name: str) -> None:
        """
        Purges all messages from a RabbitMQ queue.

        Args:
            queue_name (str): The name of the queue to purge.

        Raises:
            RuntimeError: If not connected to RabbitMQ.
        """
        channel = await self.create_channel()
        async with channel:
            queue = await channel.declare_queue(queue_name, durable=True)
            await queue.purge()
            logger.info(f"Purged queue '{queue_name}'")

    async def delete_queue(self, queue_name: str) -> None:
        """
        Deletes a RabbitMQ queue.

        Args:
            queue_name (str): The name of the queue to delete.

        Raises:
            RuntimeError: If not connected to RabbitMQ.
        """
        channel = await self.create_channel()
        async with channel:
            queue = await channel.declare_queue(queue_name, durable=True)
            await queue.delete()
            logger.info(f"Deleted queue '{queue_name}'")

    def on_connection_error(self, error: Exception) -> None:
        """
        Handles connection errors by logging the error and connection parameters.

        Args:
            error (Exception): The exception that occurred.
        """
        logger.error(f"Connection error: {error}")
        logger.error(f"Connection parameters: {self.url}")

    async def create_channel(self) -> aio_pika.Channel:
        """
        Creates a RabbitMQ channel with QoS settings.

        Returns:
            aio_pika.Channel: The created channel.

        Raises:
            RuntimeError: If not connected to RabbitMQ.
        """
        if self.connection is None:
            raise RuntimeError("Not connected to RabbitMQ")
        channel = await self.connection.channel()
        await channel.set_qos(prefetch_count=1)
        return channel

    async def declare_exchange(self, channel: aio_pika.Channel, exchange_name: str) -> None:
        """
        Declares a RabbitMQ exchange.

        Args:
            channel (aio_pika.Channel): The channel to use for declaring the exchange.
            exchange_name (str): The name of the exchange to declare.

        Raises:
            aio_pika.exceptions.ChannelError: If there is an error declaring the exchange.
        """
        self.exchange = await channel.declare_exchange(
            exchange_name, aio_pika.ExchangeType.TOPIC, durable=True
        )
        logger.info(f"Declared exchange '{exchange_name}'")

    async def create_queue(
        self,
        channel: aio_pika.Channel,
        queue_name: str,
        exchange_name: str,
        routing_key: str
    ) -> aio_pika.Queue:
        """
        Creates a RabbitMQ queue and binds it to an exchange.

        Args:
            channel (aio_pika.Channel): The channel to use for creating the queue.
            queue_name (str): The name of the queue to create.
            exchange_name (str): The name of the exchange to bind the queue to.
            routing_key (str): The routing key to use for binding.

        Returns:
            aio_pika.Queue: The created and bound queue.

        Raises:
            aio_pika.exceptions.ChannelError: If there is an error creating or binding the queue.
        """
        await self.declare_exchange(channel, exchange_name)
        queue = await channel.declare_queue(queue_name, durable=True)
        await queue.bind(self.exchange, routing_key)
        logger.info(
            f"Created and bound queue '{queue_name}' to exchange '{exchange_name}' with""routing key '{routing_key}'"
        )
        return queue

    async def close_connection(self) -> None:
        """
        Closes the RabbitMQ connection.

        Raises:
            aio_pika.exceptions.AMQPError: If there is an error closing the connection.
        """
        if self.connection:
            await self.connection.close()
            logger.info("Closed RabbitMQ connection")
            self.connection = None

    async def publish_message(self, exchange_name: str, routing_key: str, message: str) -> None:
        """
        Publishes a message to a RabbitMQ exchange.

        Args:
            exchange_name (str): The name of the exchange to publish to.
            routing_key (str): The routing key to use for publishing.
            message (str): The message to publish.

        Raises:
            RuntimeError: If the exchange is not declared.
            aio_pika.exceptions.AMQPError: If there is an error publishing the message.
        """
        if self.exchange is None:
            raise RuntimeError(f"Exchange '{exchange_name}' not declared")
        try:
            await self.exchange.publish(
                aio_pika.Message(
                    body=message.encode(),
                    delivery_mode=aio_pika.DeliveryMode.PERSISTENT,
                ),
                routing_key=routing_key,
            )
            logger.info(f"Published message to exchange '{exchange_name}' with routing key '{routing_key}'")
        except Exception as e:
            logger.error(f"Error while publishing message: {e}")
