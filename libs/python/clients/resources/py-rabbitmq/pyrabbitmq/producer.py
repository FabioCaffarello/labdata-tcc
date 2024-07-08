from pyrabbitmq.base import BaseRabbitMQ


class RabbitMQProducer(BaseRabbitMQ):
    """A RabbitMQ producer class that extends BaseRabbitMQ."""

    def __init__(self):
        """
        Initializes the RabbitMQProducer instance.
        """
        super().__init__()

    async def send_message(self, exchange_name: str, routing_key: str, message: str) -> None:
        """
        Sends a message to a RabbitMQ exchange.

        Args:
            exchange_name (str): The name of the exchange to publish to.
            routing_key (str): The routing key to use for publishing.
            message (str): The message to publish.
        """
        await self.publish_message(exchange_name, routing_key, message)
