package mongowrapper

import (
	gomongodb "libs/golang/clients/resources/go-mongo/client"
	"os"
)

// MongoDBWrapper wraps a MongoDB client and provides initialization and retrieval methods.
type MongoDBWrapper struct {
	client  *gomongodb.Client
	factory ClientFactory
}

// NewMongoDBWrapper creates a new MongoDBWrapper with the default client factory.
func NewMongoDBWrapper() *MongoDBWrapper {
	return &MongoDBWrapper{
		factory: &DefaultClientFactory{},
	}
}

// Init initializes the MongoDB client using environment variables for configuration.
// It returns an error if the client could not be created.
func (m *MongoDBWrapper) Init() error {
	config := gomongodb.Config{
		User:     os.Getenv("MONGODB_USER"),
		Password: os.Getenv("MONGODB_PASSWORD"),
		Host:     os.Getenv("MONGODB_HOST"),
		Port:     os.Getenv("MONGODB_PORT"),
		DBName:   os.Getenv("MONGODB_DBNAME"),
	}
	client, err := m.factory.NewClient(config)
	if err != nil {
		return err
	}
	m.client = client
	return nil
}

// GetClient returns the MongoDB client.
func (m *MongoDBWrapper) GetClient() interface{} {
	return m.client
}
