package gorabbitmq

import (
	"context"
)

// RabbitMQNotifier is a struct that handles sending notifications through RabbitMQ.
type RabbitMQNotifier struct {
	rmqClient *Client
}

// NewRabbitMQNotifier creates a new RabbitMQNotifier with the given RabbitMQ client.
//
// Parameters:
//   - rmqClient: A pointer to a Client instance that is already connected to RabbitMQ.
//
// Returns:
//   - A pointer to the newly created RabbitMQNotifier.
func NewRabbitMQNotifier(rmqClient *Client) *RabbitMQNotifier {
	return &RabbitMQNotifier{
		rmqClient: rmqClient,
	}
}

// Notify sends a notification message to the RabbitMQ exchange using the specified routing key.
//
// Parameters:
//   - message: The message to be sent as a byte slice.
//   - routingKey: The routing key to be used for routing the message.
//
// Returns:
//   - An error if the message could not be published, or nil if the message was successfully published.
func (n *RabbitMQNotifier) Notify(message []byte, routingKey string) error {
	var (
		contentType = "application/json"
		ctx         = context.Background()
	)
	return n.rmqClient.publish(ctx, contentType, message, routingKey)
}
