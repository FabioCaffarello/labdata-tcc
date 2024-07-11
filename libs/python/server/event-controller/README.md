# event-controller

A Python library for managing event processing and handling using RabbitMQ. This library allows you to set up controllers to handle events, validate schemas, and manage job processing.

## Features

- Asynchronous RabbitMQ message handling.
- Schema validation using schema vault client.
- Configurable event controllers and job handlers.
- Easy integration with service discovery and debugging tools.

## Installation

To install the `event-controller` library, use the following command:

```bash
npx nx run <PROJECT>:add --name python-server-event-controller --local
```

## Usage

### Controller

The `Controller` class provides methods to handle event processing, including reading messages, validating schemas, and dispatching jobs.

#### Example

```python
import asyncio
from dto_config_vault.output import ConfigDTO
from pyrabbitmq.producer import RabbitMQProducer
from pysd.sd import ServiceDiscovery
from pydebug import debug
from event_controller import EventController

async def main():
    sd = ServiceDiscovery()
    config = ConfigDTO(provider="provider", service="service", source="source")
    queue_active_jobs = asyncio.Queue()
    dbg = debug.EnabledDebug()
    job_handler = YourJobHandlerFunction

    event_controller = EventController(sd, config, job_handler, queue_active_jobs, dbg)
    await event_controller.run(message)

asyncio.run(main())
```

### EventController

The `EventController` class extends `Controller` for specific event processing.

#### `EventController.__init__(self, sd: ServiceDiscovery, config: ConfigDTO, job_handler: callable, queue_active_jobs: asyncio.Queue, dbg: Union[debug.EnabledDebug, debug.DisabledDebug])`

Initializes the `EventController` instance.

```python
event_controller = EventController(sd, config, job_handler, queue_active_jobs, dbg)
```

#### `EventController.run(self, message) -> None`

Runs the event controller to process the incoming message.

```python
await event_controller.run(message)
```

## Running Tests

To run the tests, use `pytest`:

```sh
npx nx test python-server-event-controller
```

Make sure you have the development dependencies installed:

```sh
npx nx install python-server-event-controller --with dev
```
