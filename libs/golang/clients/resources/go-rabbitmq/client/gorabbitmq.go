package gorabbitmq

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Config holds the configuration for connecting to a RabbitMQ instance.
type Config struct {
	User         string // Username for RabbitMQ authentication
	Password     string // Password for RabbitMQ authentication
	Host         string // Host address of the RabbitMQ server
	Port         string // Port number of the RabbitMQ server
	Protocol     string // Protocol to use for the connection (e.g., "amqp")
	ExchangeName string // Name of the RabbitMQ exchange to use
	ExchangeType string // Type of the RabbitMQ exchange (e.g., "direct", "fanout")
}

// Client represents a RabbitMQ client.
type Client struct {
	Dsn           string           // Data Source Name for connecting to RabbitMQ
	Conn          *amqp.Connection // RabbitMQ connection instance
	Channel       *amqp.Channel    // RabbitMQ channel instance
	ExchangeName  string           // Name of the RabbitMQ exchange in use
	ExchangeType  string           // Type of the RabbitMQ exchange in use
	totalAttempts int              // Total number of attempts to connect/reconnect
}

// NewClient creates a new RabbitMQ client with the given configuration.
//
// Parameters:
//   - config: The configuration for connecting to RabbitMQ.
//
// Returns:
//   - A pointer to the newly created Client.
//   - An error if the client could not be created.
func NewClient(config Config) (*Client, error) {
	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/", config.Protocol, config.User, config.Password, config.Host, config.Port)
	log.Printf("Connecting to RabbitMQ with DSN: %s", dsn)
	rabbitClient := &Client{
		Dsn:           dsn,
		ExchangeName:  config.ExchangeName,
		ExchangeType:  config.ExchangeType,
		totalAttempts: 20,
	}

	var err error
	for i := 0; i < rabbitClient.totalAttempts; i++ {
		err = rabbitClient.connect()
		if err == nil {
			err = rabbitClient.channel()
			if err == nil {
				break
			}
		}
		log.Printf("Failed to connect or open channel, retrying... (%d/%d)", i+1, rabbitClient.totalAttempts)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	if err := rabbitClient.declareExchange(); err != nil {
		return nil, fmt.Errorf("failed to declare exchange: %w", err)
	}

	return rabbitClient, nil
}

// connect establishes a connection to the RabbitMQ server.
//
// Returns:
//   - An error if the connection could not be established.
func (c *Client) connect() error {
	log.Println("Connecting to RabbitMQ...")
	var err error
	c.Conn, err = amqp.Dial(c.Dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}
	log.Println("Connected to RabbitMQ")
	return nil
}

// channel opens a channel for the RabbitMQ client.
//
// Returns:
//   - An error if the channel could not be opened.
func (c *Client) channel() error {
	if c.Conn == nil {
		return fmt.Errorf("connection is nil")
	}
	log.Println("Opening channel...")
	var err error
	c.Channel, err = c.Conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %w", err)
	}
	log.Println("Opened channel")
	return nil
}

// declareExchange declares an exchange for the RabbitMQ client.
//
// Returns:
//   - An error if the exchange could not be declared.
func (c *Client) declareExchange() error {
	if c.Channel == nil {
		return fmt.Errorf("channel is nil")
	}
	err := c.Channel.ExchangeDeclare(
		c.ExchangeName,
		c.ExchangeType,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to declare exchange: %w", err)
	}
	log.Println("Declared exchange")
	return nil
}

// declareQueue declares a queue for the RabbitMQ client.
//
// Parameters:
//   - queueName: The name of the queue to declare.
//   - args: Additional arguments for the queue declaration.
//
// Returns:
//   - A pointer to the declared queue.
//   - An error if the queue could not be declared.
func (c *Client) declareQueue(queueName string, args amqp.Table) (*amqp.Queue, error) {
	if c.Channel == nil {
		return nil, fmt.Errorf("channel is nil")
	}
	var err error
	for i := 0; i < c.totalAttempts; i++ {
		q, err := c.Channel.QueueDeclare(
			queueName,
			true,
			false,
			false,
			false,
			args,
		)
		if err == nil {
			log.Printf("Declared queue: %s", queueName)
			return &q, nil
		}
		log.Printf("Failed to declare queue: %s, retrying... (%d/%d)", queueName, i+1, c.totalAttempts)
		time.Sleep(2 * time.Second)
	}
	return nil, fmt.Errorf("failed to declare queue: %w", err)
}

// bindQueue binds a queue to an exchange for the RabbitMQ client.
//
// Parameters:
//   - queueName: The name of the queue to bind.
//   - routingKey: The routing key to use for the binding.
//
// Returns:
//   - An error if the queue could not be bound.
func (c *Client) bindQueue(queueName, routingKey string) error {
	if c.Channel == nil {
		return fmt.Errorf("channel is nil")
	}
	var err error
	for i := 0; i < c.totalAttempts; i++ {
		err = c.Channel.QueueBind(
			queueName,
			routingKey,
			c.ExchangeName,
			false,
			nil,
		)
		if err == nil {
			log.Printf("Bound queue: %s with routing key: %s", queueName, routingKey)
			return nil
		}
		log.Printf("Failed to bind queue: %s, retrying... (%d/%d)", queueName, i+1, c.totalAttempts)
		time.Sleep(2 * time.Second)
	}
	return fmt.Errorf("failed to bind queue: %w", err)
}

// consume starts consuming messages from the specified queue.
//
// Parameters:
//   - msgCh: A channel to receive the messages.
//   - consumerName: The name of the consumer.
//   - queueName: The name of the queue to consume from.
//   - autoAck: Whether to automatically acknowledge messages.
func (c *Client) consume(msgCh chan amqp.Delivery, consumerName string, queueName string, autoAck bool) {
	if c.Channel == nil {
		log.Println("Channel is nil")
		return
	}
	deliveryCh, err := c.Channel.Consume(
		queueName,
		consumerName,
		autoAck,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("Failed to consume messages from queue: %s", queueName)
		return
	}
	log.Printf("Started consuming messages from queue: %s", queueName)
	go func() {
		for message := range deliveryCh {
			log.Printf("Received message %s from queue: %s", string(message.Body), queueName)
			msgCh <- message
		}
		log.Println("RabbitMQ channel closed")
		close(msgCh)
	}()
}

// publish sends a message to the RabbitMQ exchange.
//
// Parameters:
//   - ctx: The context for publishing the message.
//   - contentType: The content type of the message.
//   - message: The message to be sent as a byte slice.
//   - routingKey: The routing key to use for routing the message.
//
// Returns:
//   - An error if the message could not be published.
func (c *Client) publish(ctx context.Context, contentType string, message []byte, routingKey string) error {
	if c.Channel == nil {
		return fmt.Errorf("channel is nil")
	}
	err := c.Channel.PublishWithContext(
		ctx,
		c.ExchangeName,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: contentType,
			Body:        message,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}
	log.Printf("Published message to routing key: %s", routingKey)
	return nil
}

// Close closes the RabbitMQ client's channel and connection.
//
// Returns:
//   - An error if the channel or connection could not be closed.
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
	log.Println("Closed RabbitMQ connection and channel")
	return err
}
