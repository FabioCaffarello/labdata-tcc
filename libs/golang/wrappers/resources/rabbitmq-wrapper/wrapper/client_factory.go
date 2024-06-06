package rabbitmqwrapper

import gorabbitmq "libs/golang/clients/resources/go-rabbitmq/client"

// ClientFactory defines an interface for creating new RabbitMQ clients.
type ClientFactory interface {
	// NewClient creates a new RabbitMQ client with the provided configuration.
	// It returns the client and an error if any occurred during the connection.
	NewClient(config gorabbitmq.Config) (*gorabbitmq.Client, error)
}

// DefaultClientFactory is a default implementation of the ClientFactory interface.
type DefaultClientFactory struct{}

// NewClient creates a new RabbitMQ client with the provided configuration using the default implementation.
func (f *DefaultClientFactory) NewClient(config gorabbitmq.Config) (*gorabbitmq.Client, error) {
	return gorabbitmq.NewClient(config)
}
