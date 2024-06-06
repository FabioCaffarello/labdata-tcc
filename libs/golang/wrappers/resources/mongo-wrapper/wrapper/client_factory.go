package mongowrapper

import gomongodb "libs/golang/clients/resources/go-mongo/client"

// ClientFactory defines an interface for creating new MongoDB clients.
type ClientFactory interface {
	// NewClient creates a new MongoDB client with the provided configuration.
	// It returns the client and an error if any occurred during the connection.
	NewClient(config gomongodb.Config) (*gomongodb.Client, error)
}

// DefaultClientFactory is a default implementation of the ClientFactory interface.
type DefaultClientFactory struct{}

// NewClient creates a new MongoDB client with the provided configuration using the default implementation.
func (f *DefaultClientFactory) NewClient(config gomongodb.Config) (*gomongodb.Client, error) {
	return gomongodb.NewClient(config)
}
