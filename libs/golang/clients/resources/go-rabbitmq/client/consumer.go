package gorabbitmq

import (
	"context"
	"log"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQConsumer represents a RabbitMQ consumer.
type RabbitMQConsumer struct {
	rmqClient    *Client         // RabbitMQ client instance
	autoAck      bool            // Automatic acknowledgment flag
	args         amqp.Table      // Additional arguments for the queue declaration
	consumerName string          // Name of the consumer
	wg           *sync.WaitGroup // WaitGroup to manage goroutines
}

// ConsumerConfig holds the configuration for the RabbitMQ consumer.
type ConsumerConfig struct {
	ConsumerName string     // Name of the consumer
	AutoAck      bool       // Automatic acknowledgment flag
	Args         amqp.Table // Additional arguments for the queue declaration
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
		wg:           &sync.WaitGroup{},
	}
}

// Consume starts consuming messages from the specified queue and sends them to the provided channel.
//
// Parameters:
//   - ctx: The context to use for the consumer.
//   - msgCh: A channel to send the consumed messages to.
//   - queueName: The name of the queue to consume messages from.
//   - routingKey: The routing key to use for binding the queue.
//
// This method will panic if the queue declaration, binding, or consumption fails.
func (c *RabbitMQConsumer) Consume(ctx context.Context, msgCh chan amqp.Delivery, queueName string, routingKey string) {
	if c.rmqClient == nil || c.rmqClient.Channel == nil {
		panic("rmqClient or Channel is nil")
	}

	q, err := c.rmqClient.declareQueue(queueName, c.args)
	if err != nil {
		panic(err)
	}
	log.Printf("Declared queue: %s", q.Name)

	if err := c.rmqClient.bindQueue(q.Name, routingKey); err != nil {
		panic(err)
	}
	log.Printf("Bound queue: %s with routing key: %s", q.Name, routingKey)

	deliveryCh := make(chan amqp.Delivery)
	go c.rmqClient.consume(deliveryCh, c.consumerName, q.Name, c.autoAck)
	log.Println("Started internal consume routine")

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		for {
			select {
			case message, ok := <-deliveryCh:
				if !ok {
					log.Println("Deliveries channel closed")
					close(msgCh)
					return
				}
				log.Println("Incoming new message")
				if !c.autoAck {
					message.Ack(false)
				}
				log.Printf("Received message %s from queue: %s", string(message.Body), queueName)
				msgCh <- message
			case <-ctx.Done():
				log.Println("Context done, stopping consumer")
				close(msgCh)
				return
			}
		}
	}()
}

// Wait waits for all goroutines to finish.
func (c *RabbitMQConsumer) Wait() {
	c.wg.Wait()
}
