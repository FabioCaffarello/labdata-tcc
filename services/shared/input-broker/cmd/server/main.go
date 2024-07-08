package main

import (
	"context"
	gomongodb "libs/golang/clients/resources/go-mongo/client"
	gorabbitmq "libs/golang/clients/resources/go-rabbitmq/client"
	"libs/golang/ddd/adapters/http/handlers/health-check/healthz"
	webHandler "libs/golang/ddd/adapters/http/handlers/input-broker/handlers"
	eventHandlers "libs/golang/ddd/events/input-broker/handlers"
	webserver "libs/golang/server/http/chi-webserver/server"
	servicediscovery "libs/golang/service-discovery/sd"
	events "libs/golang/shared/go-events/amqp_events"
	"log"
	"os"
	"time"
)

var (
	webServerPort = ":8000"
	databaseName  = os.Getenv("MONGODB_DBNAME")
)

// getMongoResource retrieves the MongoDB wrapper client resource from the service discovery.
//
// Parameters:
//   - sd: The service discovery instance.
//
// Returns:
//   - A pointer to the MongoDB wrapper client.
//
// Panics if the MongoDB resource is not found or if the client wrapper type is invalid.
func getMongoResource(sd *servicediscovery.ServiceDiscovery) *gomongodb.Client {
	mongo, err := sd.GetResource("mongodb")
	if err != nil {
		panic(err)
	}
	client, ok := mongo.GetClient().(*gomongodb.Client)
	if !ok {
		panic("invalid MongoDB client type")
	}
	return client
}

// getRabbitMQResource retrieves the RabbitMQ client resource from the service discovery.
//
// Parameters:
//   - sd: The service discovery instance.
//
// Returns:
//   - A pointer to the RabbitMQ client.
//
// Panics if the RabbitMQ resource is not found or if the client type is invalid.
func getRabbitMQResource(sd *servicediscovery.ServiceDiscovery) *gorabbitmq.Client {
	rabbitmq, err := sd.GetResource("rabbitmq")
	if err != nil {
		panic(err)
	}
	client, ok := rabbitmq.GetClient().(*gorabbitmq.Client)
	if !ok {
		panic("invalid RabbitMQ client type")
	}
	return client
}

// getRabbitMQNotifier initializes and configures the RabbitMQ notifier.
//
// Parameters:
//   - sd: The service discovery instance.
//
// Returns:
//   - A pointer to the configured RabbitMQ notifier.
func getRabbitMQNotifier(sd *servicediscovery.ServiceDiscovery) *gorabbitmq.RabbitMQNotifier {
	rabbitmqClient := getRabbitMQResource(sd)
	return gorabbitmq.NewRabbitMQNotifier(rabbitmqClient)
}

// getHTTPServer initializes and configures the HTTP server.
//
// Returns:
//   - A pointer to the configured web server.
func getHTTPServer() *webserver.Server {
	httpServer := webserver.NewWebServer(webServerPort)
	httpServer.ConfigureDefaults()
	return httpServer
}

// makeHTTPHealthzTransport registers the health check route on the HTTP server.
//
// Parameters:
//   - httpServer: The web server instance.
//   - healthzHandler: The health check handler.
func makeHTTPHealthzTransport(httpServer *webserver.Server, healthzHandler *healthz.WebHealthzHandler) {
	httpServer.RegisterRoute("GET", "/healthz", healthzHandler.Healthz)
}

// makeHTTPConfigTransport registers the configuration routes on the HTTP server.
//
// Parameters:
//   - httpServer: The web server instance.
//   - configHandler: The configuration handler.
func makeHTTPConfigTransport(httpServer *webserver.Server, configHandler *webHandler.WebInputHandler) {
	httpServer.RegisterRoute("POST", "/input", configHandler.CreateInput)
	httpServer.RegisterRoute("GET", "/input", configHandler.ListAllInputs)
	httpServer.RegisterRoute("GET", "/input/{id}", configHandler.ListInputByID)
	httpServer.RegisterRoute("UPDATE", "/input/{id}", configHandler.UpdateInput)
	httpServer.RegisterRoute("DELETE", "/input/{id}", configHandler.DeleteInput)
	httpServer.RegisterRoute("UPDATE", "/input/{id}/status", configHandler.UpdateInputStatus)
	httpServer.RegisterRoute("GET", "/input/provider/{provider}/service/{service}", configHandler.ListInputsByServiceAndProvider)
	httpServer.RegisterRoute("GET", "/input/provider/{provider}/source/{source}", configHandler.ListInputsBySourceAndProvider)
	httpServer.RegisterRoute("GET", "/input/provider/{provider}/service/{service}/source/{source}", configHandler.ListInputsByServiceAndSourceAndProvider)
	httpServer.RegisterRoute("GET", "/input/provider/{provider}/service/{service}/status/{status}", configHandler.ListInputsByStatusAndServiceAndProvider)
	httpServer.RegisterRoute("GET", "/input/provider/{provider}/source/{source}/status/{status}", configHandler.ListInputsByStatusAndSourceAndProvider)
	httpServer.RegisterRoute("GET", "/input/provider/{provider}/service/{service}/source/{source}/status/{status}", configHandler.ListInputsByStatusAndServiceAndSourceAndProvider)
	httpServer.RegisterRoute("GET", "/input/provider/{provider}/status/{status}", configHandler.ListInputsByStatusAndProvider)
}

func main() {
	log.New(os.Stdout, "[INPUT-BROKER] - ", log.LstdFlags)
	sd := servicediscovery.NewServiceDiscovery()
	mongoClient := getMongoResource(sd)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	defer mongoClient.Disconnect(ctx)

	notifier := getRabbitMQNotifier(sd)
	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("InputCreated", &eventHandlers.InputCreatedHandler{
		Notifier: notifier,
	})

	healthzHandler := healthz.NewWebHealthzHandler(&healthz.RealTimeProvider{}, 5*time.Second)
	inputHandler := NewWebServiceInputHandler(mongoClient.Client, eventDispatcher, databaseName)

	httpServer := getHTTPServer()
	makeHTTPHealthzTransport(httpServer, healthzHandler)
	makeHTTPConfigTransport(httpServer, inputHandler)

	if err := httpServer.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
