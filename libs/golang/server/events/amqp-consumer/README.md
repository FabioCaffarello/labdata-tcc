# amqp-consumer

`amqp-consumer` is a Go library that provides an AMQP consumer for handling messages from a RabbitMQ queue. This library includes functionalities for creating and managing a consumer, consuming messages, and handling message channels.

## Features

- Create and configure an AMQP consumer.
- Consume messages from a RabbitMQ queue.
- Handle message channels for processing incoming messages.
- Gracefully stop the consumer.

## Usage

### Creating and Configuring the AMQP Consumer

The `NewAmqpConsumer` function creates a new `AmqpConsumer` instance with the specified RabbitMQ client, queue name, consumer name, and routing key.

```go
package main

import (
    "log"
    "libs/golang/clients/resources/go-rabbitmq/client"
    "libs/golang/server/events/amqp-consumer/consumer"
)

func main() {
    rabbitMQClient := client.NewClient("amqp://guest:guest@localhost:5672/")
    amqpConsumer := consumer.NewAmqpConsumer(rabbitMQClient, "exampleQueue", "exampleConsumer", "exampleRoutingKey")

    go amqpConsumer.Consume()

    for msg := range amqpConsumer.GetMsgCh() {
        log.Printf("Processed message: %s", msg)
    }

    amqpConsumer.Stop()
}
```

### Consuming Messages

The `Consume` method starts consuming messages from the specified queue and processes them.

```go
func main() {
    rabbitMQClient := client.NewClient("amqp://guest:guest@localhost:5672/")
    amqpConsumer := consumer.NewAmqpConsumer(rabbitMQClient, "exampleQueue", "exampleConsumer", "exampleRoutingKey")

    go amqpConsumer.Consume()

    for msg := range amqpConsumer.GetMsgCh() {
        log.Printf("Processed message: %s", msg)
    }

    amqpConsumer.Stop()
}
```

### Handling Message Channels

The `GetMsgCh` method returns a read-only channel where messages are sent for processing.

```go
func main() {
    rabbitMQClient := client.NewClient("amqp://guest:guest@localhost:5672/")
    amqpConsumer := consumer.NewAmqpConsumer(rabbitMQClient, "exampleQueue", "exampleConsumer", "exampleRoutingKey")

    go amqpConsumer.Consume()

    for msg := range amqpConsumer.GetMsgCh() {
        log.Printf("Processed message: %s", msg)
    }

    amqpConsumer.Stop()
}
```

### Stopping the Consumer

The `Stop` method stops the consumer by closing the quit channel.

```go
func main() {
    rabbitMQClient := client.NewClient("amqp://guest:guest@localhost:5672/")
    amqpConsumer := consumer.NewAmqpConsumer(rabbitMQClient, "exampleQueue", "exampleConsumer", "exampleRoutingKey")

    go amqpConsumer.Consume()

    // Simulate processing for a period of time
    // ...

    amqpConsumer.Stop()
}
```