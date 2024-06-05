package resourcemapping

import (
	"context"
	gominio "libs/golang/clients/resources/go-minio/client"
	gomongodb "libs/golang/clients/resources/go-mongo/client"
	gorabbitmq "libs/golang/clients/resources/go-rabbitmq/client"
	"os"
	"sync"
)

type Resources struct {
	MongoDBClient *gomongodb.Client
	// RabbitMQClient   *gorabbitmq.Client
	RabbitMQNotifier *gorabbitmq.RabbitMQNotifier
	RabbitMQConsumer *gorabbitmq.RabbitMQConsumer
	MinioClient      *gominio.Client
}

var instance *Resources
var once sync.Once

func NewResourceMapping(ctx context.Context) *Resources {
	once.Do(func() {
		instance = &Resources{}

		// MongoDB
		mongoDBClient, err := setMongoDBResource()
		if err != nil {
			panic(err)
		}

	})

	return instance
}

func NewMongoDBConfig() gomongodb.Config {
	return gomongodb.Config{
		User:     os.Getenv("MONGODB_USER"),
		Password: os.Getenv("MONGODB_PASSWORD"),
		Host:     os.Getenv("MONGODB_HOST"),
		Port:     os.Getenv("MONGODB_PORT"),
		DBName:   os.Getenv("MONGODB_DBNAME"),
	}
}

func setMongoDBResource() (*gomongodb.Client, error) {
	config := gomongodb.Config{
		User:     os.Getenv("MONGODB_USER"),
		Password: os.Getenv("MONGODB_PASSWORD"),
		Host:     os.Getenv("MONGODB_HOST"),
		Port:     os.Getenv("MONGODB_PORT"),
		DBName:   os.Getenv("MONGODB_DBNAME"),
	}
	client, err := gomongodb.NewClient(config)
	if err != nil {
		return nil, err
	}
	return client, nil
}
