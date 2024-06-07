package rabbitmqwrapper

import (
	gorabbitmq "libs/golang/clients/resources/go-rabbitmq/client"
	"os"
)

// RabbitMQWrapper wraps a RabbitMQ client and provides initialization and retrieval methods.
type RabbitMQWrapper struct {
	client  *gorabbitmq.Client
	factory ClientFactory
}

// NewRabbitMQWrapper creates a new RabbitMQWrapper with the default client factory.
func NewRabbitMQWrapper() *RabbitMQWrapper {
	return &RabbitMQWrapper{
		factory: &DefaultClientFactory{},
	}
}

// Init initializes the RabbitMQ client using environment variables for configuration.
// It returns an error if the client could not be created.
func (r *RabbitMQWrapper) Init() error {
	config := gorabbitmq.Config{
		User:         os.Getenv("RABBITMQ_USER"),
		Password:     os.Getenv("RABBITMQ_PASSWORD"),
		Host:         os.Getenv("RABBITMQ_HOST"),
		Port:         os.Getenv("RABBITMQ_PORT"),
		Protocol:     os.Getenv("RABBITMQ_PROTOCOL"),
		ExchangeName: os.Getenv("RABBITMQ_EXCHANGE_NAME"),
		ExchangeType: os.Getenv("RABBITMQ_EXCHANGE_TYPE"),
	}
	client, err := r.factory.NewClient(config)
	if err != nil {
		return err
	}
	r.client = client
	return nil
}

// GetClient returns the RabbitMQ client.
func (r *RabbitMQWrapper) GetClient() interface{} {
	return r.client
}
