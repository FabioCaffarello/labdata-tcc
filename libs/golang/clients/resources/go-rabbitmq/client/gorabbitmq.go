package gorabbitmq

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Config holds the configuration for connecting to a RabbitMQ instance.
type Config struct {
	User         string
	Password     string
	Host         string
	Port         string
	Protocol     string
	ExchangeName string
	ExchangeType string
}

// Client represents a RabbitMQ client.
type Client struct {
	Dsn          string
	Conn         *amqp.Connection
	Channel      *amqp.Channel
	ExchangeName string
	ExchangeType string
}

// NewClient creates a new RabbitMQ client with the given configuration.
//
// Parameters:
//   - config: Configuration for connecting to RabbitMQ.
//
// Returns:
//   - A pointer to the newly created Client, or an error if the connection fails.
func NewClient(config Config) (*Client, error) {
	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/", config.Protocol, config.User, config.Password, config.Host, config.Port)
	rabbitClient := &Client{
		Dsn:          dsn,
		ExchangeName: config.ExchangeName,
		ExchangeType: config.ExchangeType,
	}

	if err := rabbitClient.connect(); err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	if err := rabbitClient.channel(); err != nil {
		return nil, fmt.Errorf("failed to open an amqp channel: %w", err)
	}

	if err := rabbitClient.declareExchange(); err != nil {
		return nil, fmt.Errorf("failed to declare exchange: %w", err)
	}

	return rabbitClient, nil
}

// connect establishes a connection to the RabbitMQ server.
//
// Returns:
//   - An error if the connection fails, or nil if the connection is successful.
func (c *Client) connect() error {
	var err error
	c.Conn, err = amqp.Dial(c.Dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}
	return nil
}

// channel opens a channel for the RabbitMQ client.
//
// Returns:
//   - An error if the channel could not be opened, or nil if the channel is successfully opened.
func (c *Client) channel() error {
	var err error
	c.Channel, err = c.Conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %w", err)
	}
	return nil
}

// declareExchange declares an exchange for the RabbitMQ client.
//
// Returns:
//   - An error if the exchange could not be declared, or nil if the exchange is successfully declared.
func (c *Client) declareExchange() error {
	return c.Channel.ExchangeDeclare(
		c.ExchangeName,
		c.ExchangeType,
		true,
		false,
		false,
		false,
		nil,
	)
}

// declareQueue declares a queue for the RabbitMQ client.
//
// Parameters:
//   - queueName: The name of the queue to declare.
//   - args: Additional arguments for the queue declaration.
//
// Returns:
//   - An error if the queue could not be declared, or nil if the queue is successfully declared.
func (c *Client) declareQueue(queueName string, args amqp.Table) error {
	_, err := c.Channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		args,
	)
	return err
}

// bindQueue binds a queue to an exchange for the RabbitMQ client.
//
// Parameters:
//   - queueName: The name of the queue to bind.
//   - routingKey: The routing key to use for the binding.
//
// Returns:
//   - An error if the queue could not be bound, or nil if the queue is successfully bound.
func (c *Client) bindQueue(queueName, routingKey string) error {
	return c.Channel.QueueBind(
		queueName,
		routingKey,
		c.ExchangeName,
		false,
		nil,
	)
}

// consume starts consuming messages from the specified queue.
//
// Parameters:
//   - consumerName: The name of the consumer.
//   - queueName: The name of the queue to consume messages from.
//   - autoAck: Whether to automatically acknowledge messages.
//
// Returns:
//   - A channel to receive messages from, or an error if the consumption fails.
func (c *Client) consume(consumerName string, queueName string, autoAck bool) (<-chan amqp.Delivery, error) {
	msgCh, err := c.Channel.Consume(
		queueName,
		consumerName,
		autoAck,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to consume messages: %w", err)
	}

	return msgCh, nil
}

// Publish sends a message to the RabbitMQ exchange.
//
// Parameters:
//   - ctx: The context for the publish operation.
//   - contentType: The content type of the message.
//   - message: The message to be sent as a byte slice.
//   - routingKey: The routing key to be used for routing the message.
//
// Returns:
//   - An error if the message could not be published, or nil if the message was successfully published.
func (c *Client) publish(ctx context.Context, contentType string, message []byte, routingKey string) error {
	err := c.Channel.PublishWithContext(
		ctx,
		c.ExchangeName, // Use the Exchange property of the struct
		routingKey,     // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: contentType,
			Body:        message,
		})

	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}

// Close closes the RabbitMQ client's channel and connection.
//
// Returns:
//   - An error if the channel or connection could not be closed, or nil if they were successfully closed.
func (c *Client) Close() error {
	var err error
	if c.Channel != nil && !c.Channel.IsClosed() {
		if closeErr := c.Channel.Close(); closeErr != nil {
			err = closeErr
		}
	}
	if c.Conn != nil && !c.Conn.IsClosed() {
		if closeErr := c.Conn.Close(); closeErr != nil {
			err = closeErr
		}
	}
	return err
}
