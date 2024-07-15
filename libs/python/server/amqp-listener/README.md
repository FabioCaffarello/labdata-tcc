# amqp-listener

A Python library for managing RabbitMQ consumers, allowing you to listen for messages and process them using event controllers and job handlers.

## Features

- Asynchronous RabbitMQ consumer management.
- Configurable event controllers and job handlers.
- Easy integration with service discovery and debugging tools.

## Installation

To install the `-amqp-listener` library, use the following command:

```bash
npx nx run <PROJECT>:add --name python-server-amqp-listener --local
```

## Usage

### Consumer

The `Consumer` class provides methods to manage RabbitMQ consumer operations.

#### Example

```python
import asyncio
from dto_config_vault.output import ConfigDTO
from pyrabbitmq.consumer import RabbitMQConsumer
from pysd.sd import ServiceDiscovery
from pydebug import debug
from amqp_listener import Consumer, EventConsumer

async def main():
    sd = ServiceDiscovery()
    rmq_consumer = RabbitMQConsumer()
    config = ConfigDTO(provider="provider", service="service", source="source")
    queue_active_jobs = asyncio.Queue()
    dbg = debug.EnabledDebug()

    consumer = EventConsumer(sd, rmq_consumer, config, queue_active_jobs, dbg)
    await consumer.run(event_controller=YourEventControllerClass, job_handler=your_job_handler_function)

asyncio.run(main())
```

## Running Tests

To run the tests, use `pytest`:

```sh
npx nx test python-server-amqp-listener
```

Make sure you have the development dependencies installed:

```sh
npx nx install python-server-amqp-listener --with dev
```
