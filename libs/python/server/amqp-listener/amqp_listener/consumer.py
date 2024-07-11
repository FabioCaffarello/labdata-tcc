import asyncio
from typing import Union

from dto_config_vault.output import ConfigDTO
from pylog.log import setup_logging
from pyrabbitmq.consumer import RabbitMQConsumer
from pysd.sd import ServiceDiscovery
from pydebug import debug

logger = setup_logging(__name__)

BASE_ROUTING_KEY = "input.ready-to-process"
BASE_INPUT_QUEUE = "input-queue"


class Consumer:
    """
    A class to manage RabbitMQ consumer operations.

    Attributes:
        sd (ServiceDiscovery): Service discovery instance.
        rmq_consumer (RabbitMQConsumer): RabbitMQ consumer instance.
        config (ConfigDTO): Configuration data transfer object.
        queue_active_jobs (asyncio.Queue): Queue for managing active jobs.
        dbg (Union[debug.EnabledDebug, debug.DisabledDebug]): Debug instance.
    """

    def __init__(
        self,
        sd: ServiceDiscovery,
        rmq_consumer: RabbitMQConsumer,
        config: ConfigDTO,
        queue_active_jobs: asyncio.Queue,
        dbg: Union[debug.EnabledDebug, debug.DisabledDebug]
    ) -> None:
        """
        Initializes the Consumer instance with necessary attributes.

        Args:
            sd (ServiceDiscovery): Service discovery instance.
            rmq_consumer (RabbitMQConsumer): RabbitMQ consumer instance.
            config (ConfigDTO): Configuration data transfer object.
            queue_active_jobs (asyncio.Queue): Queue for managing active jobs.
            dbg (Union[debug.EnabledDebug, debug.DisabledDebug]): Debug instance.
        """
        self.sd = sd
        self.config = config
        self.dbg = dbg
        self.queue_active_jobs = queue_active_jobs
        self.rmq_consumer = rmq_consumer

    @property
    def __exchange_name(self) -> str:
        """
        Returns the exchange name for RabbitMQ.

        Returns:
            str: Exchange name.
        """
        return self.sd.services_rabbitmq_exchange

    @property
    def __queue_name(self) -> str:
        """
        Returns the queue name for RabbitMQ.

        Returns:
            str: Queue name.
        """
        return f"{BASE_INPUT_QUEUE}.{self.config.provider}.{self.config.service}.{self.config.source}"

    @property
    def __routing_key(self) -> str:
        """
        Returns the routing key for RabbitMQ.

        Returns:
            str: Routing key.
        """
        return f"{BASE_ROUTING_KEY}.{self.config.provider}.{self.config.service}.{self.config.source}"

    def __instantiate_controller(self, event_controller: callable, job_handler: callable):
        """
        Instantiates the event controller with provided parameters.

        Args:
            event_controller (callable): The event controller class.
            job_handler (callable): The job handler function.

        Returns:
            callable: Instantiated event controller.
        """
        return event_controller(self.sd, self.config, job_handler, self.queue_active_jobs, self.dbg)

    async def start(self, event_controller: callable, job_handler: callable):
        """
        Starts the RabbitMQ consumer to listen for messages and process them.

        Args:
            event_controller (callable): The event controller class.
            job_handler (callable): The job handler function.
        """
        channel = await self.rmq_consumer.create_channel()
        queue = await self.rmq_consumer.create_queue(
            channel,
            self.__queue_name,
            self.__exchange_name,
            self.__routing_key
        )
        await self.rmq_consumer.listen(queue, self.__instantiate_controller(event_controller, job_handler).run)


class EventConsumer(Consumer):
    """
    A class to extend Consumer for specific event consumption.

    Methods:
        run(event_controller: callable, job_handler: callable): Runs the consumer.
    """

    def __init__(
        self,
        sd: ServiceDiscovery,
        rmq_consumer: RabbitMQConsumer,
        config: ConfigDTO,
        queue_active_jobs: asyncio.Queue,
        dbg: Union[debug.EnabledDebug, debug.DisabledDebug]
    ) -> None:
        """
        Initializes the EventConsumer instance.

        Args:
            sd (ServiceDiscovery): Service discovery instance.
            rmq_consumer (RabbitMQConsumer): RabbitMQ consumer instance.
            config (ConfigDTO): Configuration data transfer object.
            queue_active_jobs (asyncio.Queue): Queue for managing active jobs.
            dbg (Union[debug.EnabledDebug, debug.DisabledDebug]): Debug instance.
        """
        super().__init__(sd, rmq_consumer, config, queue_active_jobs, dbg)

    async def run(self, event_controller: callable, job_handler: callable):
        """
        Runs the consumer to start listening and processing messages.

        Args:
            event_controller (callable): The event controller class.
            job_handler (callable): The job handler function.
        """
        await self.start(event_controller, job_handler)
