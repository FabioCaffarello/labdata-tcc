package gorabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQConsumer represents a RabbitMQ consumer.
type RabbitMQConsumer struct {
	rmqClient    *Client
	autoAck      bool
	args         amqp.Table
	consumerName string
}

// ConsumerConfig holds the configuration for the RabbitMQ consumer.
type ConsumerConfig struct {
	ConsumerName string
	AutoAck      bool
	Args         amqp.Table
}

// NewRabbitMQConsumer creates a new RabbitMQ consumer with the given configuration.
//
// Parameters:
//   - rmqClient: A pointer to a Client instance that is already connected to RabbitMQ.
//   - config: The configuration for the consumer.
//
// Returns:
//   - A pointer to the newly created RabbitMQConsumer.
func NewRabbitMQConsumer(rmqClient *Client, config ConsumerConfig) *RabbitMQConsumer {
	return &RabbitMQConsumer{
		rmqClient:    rmqClient,
		autoAck:      config.AutoAck,
		args:         config.Args,
		consumerName: config.ConsumerName,
	}
}

// Consume starts consuming messages from the specified queue and sends them to the provided channel.
//
// Parameters:
//   - msgCh: A channel to send the consumed messages to.
//   - queueName: The name of the queue to consume messages from.
//   - routingKey: The routing key to use for binding the queue.
//
// This method will panic if the queue declaration, binding, or consumption fails.
func (c *RabbitMQConsumer) Consume(msgCh chan amqp.Delivery, queueName string, routingKey string) {
	if err := c.rmqClient.declareQueue(queueName, c.args); err != nil {
		panic(err)
	}

	if err := c.rmqClient.bindQueue(queueName, routingKey); err != nil {
		panic(err)
	}

	deliveriesCh, err := c.rmqClient.consume(c.consumerName, queueName, c.autoAck)
	if err != nil {
		panic(err)
	}

	go func() {
		for message := range deliveriesCh {
			log.Println("Incoming new message")
			msgCh <- message
		}
		log.Println("RabbitMQ channel closed")
		close(msgCh)
	}()
}
