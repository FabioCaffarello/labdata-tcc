package gomongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client is a wrapper around the mongo.Client.
type Client struct {
	*mongo.Client
}

// Config holds the configuration for connecting to a MongoDB instance.
type Config struct {
	User     string // Username for authentication
	Password string // Password for authentication
	Host     string // Host of the MongoDB instance
	Port     string // Port of the MongoDB instance
	DBName   string // Name of the database to connect to
}

// NewClient creates a new MongoDB client with the given configuration.
// It returns the Client and an error if any occurred during connection.
//
// Example:
//   config := Config{
//       User:     "testuser",
//       Password: "testpassword",
//       Host:     "localhost",
//       Port:     "27017",
//       DBName:   "testdb",
//   }
//   client, err := NewClient(config)
//   if err != nil {
//       log.Fatal(err)
//   }
//   defer client.Disconnect(context.Background())
func NewClient(config Config) (*Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		config.User, config.Password, config.Host, config.Port, config.DBName)

	clientOptions := options.Client().ApplyURI(uri).SetAuth(options.Credential{
		Username: config.User,
		Password: config.Password,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	return &Client{Client: client}, nil
}

// Disconnect closes the connection to the MongoDB instance.
// It returns an error if any occurred during disconnection.
func (c *Client) Disconnect(ctx context.Context) error {
	if err := c.Client.Disconnect(ctx); err != nil {
		return fmt.Errorf("failed to disconnect from MongoDB: %w", err)
	}
	return nil
}