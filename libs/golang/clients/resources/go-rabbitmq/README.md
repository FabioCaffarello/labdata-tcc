Here's the updated README file with improved docstrings and consistent formatting:

# go-rabbitmq

`go-rabbitmq` is a Go library that provides a RabbitMQ client wrapper with utilities for connecting to and interacting with a RabbitMQ server. This library facilitates creating a connection to a RabbitMQ instance, declaring exchanges and queues, binding queues, publishing messages, and consuming messages.

## Features

- Connect to a RabbitMQ instance using configuration parameters
- Declare exchanges and queues
- Bind queues to exchanges
- Publish messages to exchanges
- Consume messages from queues

## Usage

### Creating a RabbitMQ Client

```go
package main

import (
	"fmt"
	"log"
	"libs/golang/clients/resources/go-rabbitmq"
)

func main() {
	config := gorabbitmq.Config{
		User:         "guest",
		Password:     "guest",
		Host:         "localhost",
		Port:         "5672",
		Protocol:     "amqp",
		ExchangeName: "my_exchange",
		ExchangeType: "direct",
	}

	client, err := gorabbitmq.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ client: %v", err)
	}

	defer client.Close()

	fmt.Println("Connected to RabbitMQ")
}
```

### Publishing a Message

```go
package main

import (
	"context"
	"fmt"
	"log"
	"libs/golang/clients/resources/go-rabbitmq"
)

func main() {
	config := gorabbitmq.Config{
		User:         "guest",
		Password:     "guest",
		Host:         "localhost",
		Port:         "5672",
		Protocol:     "amqp",
		ExchangeName: "my_exchange",
		ExchangeType: "direct",
	}

	client, err := gorabbitmq.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ client: %v", err)
	}

	defer client.Close()

	ctx := context.Background()
	message := []byte("Hello, RabbitMQ!")
	err = client.Publish(ctx, "text/plain", message, "my_routing_key")
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	fmt.Println("Message published")
}
```

### Sending Notifications

```go
package main

import (
	"fmt"
	"log"
	"libs/golang/clients/resources/go-rabbitmq"
)

func main() {
	config := gorabbitmq.Config{
		User:         "guest",
		Password:     "guest",
		Host:         "localhost",
		Port:         "5672",
		Protocol:     "amqp",
		ExchangeName: "my_exchange",
		ExchangeType: "direct",
	}

	client, err := gorabbitmq.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ client: %v", err)
	}

	defer client.Close()

	notifier := gorabbitmq.NewRabbitMQNotifier(client)
	message := []byte(`{"message": "Hello, World!"}`)
	err = notifier.Notify(message, "my_routing_key")
	if err != nil {
		log.Fatalf("Failed to send notification: %v", err)
	}

	fmt.Println("Notification sent")
}
```

### Consuming Messages

```go
package main

import (
	"context"
	"log"
	"time"
	"libs/golang/clients/resources/go-rabbitmq"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	config := gorabbitmq.Config{
		User:         "guest",
		Password:     "guest",
		Host:         "localhost",
		Port:         "5672",
		Protocol:     "amqp",
		ExchangeName: "my_exchange",
		ExchangeType: "direct",
	}

	consumerConfig := gorabbitmq.ConsumerConfig{
		ConsumerName: "consumer_name",
		AutoAck:      false,
		Args:         nil,
	}

	queueName := "test_queue"
	routingKey := "test_key"
	msgCh := make(chan amqp.Delivery)

	client, err := gorabbitmq.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ client: %v", err)
	}

	defer client.Close()

	consumer := gorabbitmq.NewRabbitMQConsumer(client, consumerConfig)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go consumer.Consume(ctx, msgCh, queueName, routingKey)

	select {
	case msg := <-msgCh:
		log.Printf("Received message: %s", string(msg.Body))
	case <-ctx.Done():
		log.Println("Did not receive message in time")
	}
}
```

## Testing

To run the tests for the `go-rabbitmq` package, use the following command:

```sh
npx nx test libs-golang-clients-resources-go-rabbitmq
```
