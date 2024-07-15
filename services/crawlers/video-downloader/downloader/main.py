import asyncio
import os
import time
from typing import List, Union
from pylog.log import setup_logging
from pyargparse import argsparser
from pydebug import debug
from pysd.sd import new_from_env, ServiceDiscovery
from config_loader.loader import fetch_configs
from pyrabbitmq.consumer import RabbitMQConsumer
from amqp_listener.consumer import EventConsumer
from event_controller.controller import EventController
from job_handler.core import JobHandler


logger = setup_logging(__name__, log_level="DEBUG")

SERVICE_NAME = os.getenv("SERVICE_NAME")
PROVIDER = os.getenv("PROVIDER")
QUEUE_ACTIVE_JOBS = asyncio.Queue()


def parse_args(service_name: str):
    parser = argsparser.new(f"{service_name} args parser")
    return parser.parse_args()


def setup_debug_storage(args):
    if args.enable_debug_storage:
        os.environ["DEBUG_STORAGE_ENABLED"] = "True"
        os.environ["DEBUG_STORAGE_DIR"] = args.debug_storage_dir


def cast_string_to_boolean(input_string):
    if input_string.lower() == "true":
        return True
    else:
        return False


async def create_listeners(
    sd: ServiceDiscovery,
    service_name: str,
    provider: str,
    dbg: Union[debug.DisabledDebug, debug.EnabledDebug]
) -> List[asyncio.Task]:
    configs = await fetch_configs(service_name, provider)
    logger.info(f"Configs: {configs}")
    rmq_consumer = RabbitMQConsumer()
    await rmq_consumer.connect()

    tasks = list()
    for config_id, config in configs.items():
        logger.info(f"Creating listener for config: {config_id}")
        tasks.append(
            asyncio.create_task(
                EventConsumer(sd, rmq_consumer, config, QUEUE_ACTIVE_JOBS, dbg).run(EventController, JobHandler)
            )
        )
    return tasks


async def main():
    service_name = SERVICE_NAME
    provider = PROVIDER
    logger.info(f"Starting {service_name} service.")

    args = parse_args(service_name)
    setup_debug_storage(args)

    dbg = debug.new(cast_string_to_boolean(os.getenv("DEBUG_STORAGE_ENABLED")), os.getenv("DEBUG_STORAGE_DIR"))

    sd = new_from_env()

    tasks = await create_listeners(sd, service_name, provider, dbg)

    await asyncio.gather(*tasks)


if __name__ == "__main__":
    try:
        asyncio.run(main())
    except Exception as e:
        logger.error(f"Error: {e}")
        time.sleep(30)
        asyncio.run(main())
