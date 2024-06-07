package miniowrapper

import gominio "libs/golang/clients/resources/go-minio/client"

// ClientFactory defines an interface for creating new Minio clients.
type ClientFactory interface {
	// NewClient creates a new Minio client with the provided configuration.
	// It returns the client and an error if any occurred during the connection.
	NewClient(config gominio.Config) (*gominio.Client, error)
}

// DefaultClientFactory is a default implementation of the ClientFactory interface.
type DefaultClientFactory struct{}

// NewClient creates a new Minio client with the provided configuration using the default implementation.
func (f *DefaultClientFactory) NewClient(config gominio.Config) (*gominio.Client, error) {
	return gominio.NewClient(config)
}
