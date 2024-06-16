package miniowrapper

import (
	gominio "libs/golang/clients/resources/go-minio/client"
	"os"
)

// MinioWrapper wraps a Minio client and provides initialization and retrieval methods.
type MinioWrapper struct {
	client  *gominio.Client
	factory ClientFactory
}

// NewMinioWrapper creates a new MinioWrapper with the default client factory.
func NewMinioWrapper() *MinioWrapper {
	return &MinioWrapper{
		factory: &DefaultClientFactory{},
	}
}

// Init initializes the Minio client using environment variables for configuration.
// It returns an error if the client could not be created.
func (m *MinioWrapper) Init() error {
	config := gominio.Config{
		Port:      os.Getenv("MINIO_PORT"),
		Host:      os.Getenv("MINIO_HOST"),
		AccessKey: os.Getenv("MINIO_ACCESS_KEY"),
		SecretKey: os.Getenv("MINIO_SECRET_KEY"),
		UseSSL:    os.Getenv("MINIO_USE_SSL") == "true",
	}
	client, err := m.factory.NewClient(config)
	if err != nil {
		return err
	}
	m.client = client
	return nil
}

// GetClient returns the Minio client.
func (m *MinioWrapper) GetClient() interface{} {
	return m.client
}
