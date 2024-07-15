# py-rabbitmq

A Python library for interacting with RabbitMQ, providing easy-to-use producer and consumer classes with robust connection handling.

## Features

- Easy-to-use producer and consumer classes for RabbitMQ.
- Robust connection handling with automatic retries.
- Support for declaring exchanges and queues.
- Methods for purging and deleting queues.
- Timeout handling for consumer operations.

## Installation

To install the `py-rabbitmq` library, use the following command:

```bash
npx nx run <PROJECT>:add --name python-clients-resources-py-rabbitmq --local
```

## Usage

### RabbitMQProducer

The `RabbitMQProducer` class extends `BaseRabbitMQ` and provides methods to send messages to a RabbitMQ exchange.

#### Example

```python
from pyrabbitmq.producer import RabbitMQProducer
import asyncio

async def main():
    producer = RabbitMQProducer()
    await producer.connect()
    await producer.send_message('my_exchange', 'routing_key', 'Hello, World!')
    await producer.close_connection()

asyncio.run(main())
```

### RabbitMQConsumer

The `RabbitMQConsumer` class extends `BaseRabbitMQ` and provides methods to listen to a RabbitMQ queue and process messages with a callback function.

#### Example

```python
from pyrabbitmq.consumer import RabbitMQConsumer
import asyncio
import aio_pika

async def callback(message: aio_pika.IncomingMessage):
    print("Received message:", message.body.decode())

async def main():
    consumer = RabbitMQConsumer()
    await consumer.connect()
    channel = await consumer.create_channel()
    queue = await channel.declare_queue('my_queue', durable=True)
    await consumer.listen(queue, callback, timeout=10)
    await consumer.close_connection()

asyncio.run(main())
```

## Running Tests

To run the tests, use `pytest`:

```sh
npx nx test python-clients-resources-py-rabbitmq
```

Make sure you have the development dependencies installed:

```sh
npx nx install python-clients-resources-py-rabbitmq --with dev
```
