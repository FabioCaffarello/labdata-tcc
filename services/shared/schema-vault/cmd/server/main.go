package main

import (
	"context"
	gomongodb "libs/golang/clients/resources/go-mongo/client"
	"libs/golang/ddd/adapters/http/handlers/health-check/healthz"
	webHandler "libs/golang/ddd/adapters/http/handlers/schema-vault/handlers"
	webserver "libs/golang/server/http/chi-webserver/server"
	servicediscovery "libs/golang/service-discovery/sd"
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
func makeHTTPSchemaTransport(httpServer *webserver.Server, schemaHandler *webHandler.WebSchemaHandler) {
	httpServer.RegisterRoute("POST", "/schema", schemaHandler.CreateSchema)
	httpServer.RegisterRoute("PUT", "/schema", schemaHandler.UpdateSchema)
	httpServer.RegisterRoute("GET", "/schema", schemaHandler.ListAllSchemas)
	httpServer.RegisterRoute("GET", "/schema/{id}", schemaHandler.ListSchemaByID)
	httpServer.RegisterRoute("DELETE", "/schema/{id}", schemaHandler.DeleteSchema)
	httpServer.RegisterRoute("GET", "/schema/provider/{provider}/service/{service}", schemaHandler.ListSchemasByServiceAndProvider)
	httpServer.RegisterRoute("GET", "/schema/provider/{provider}/source/{source}", schemaHandler.ListSchemasBySourceAndProvider)
	httpServer.RegisterRoute("GET", "/schema/provider/{provider}/service/{service}/source/{source}", schemaHandler.ListSchemasByServiceAndSourceAndProvider)
}

// main is the entry point of the application.
// It initializes the service discovery, MongoDB client, HTTP server, and handlers,
// then starts the HTTP server.
func main() {
	log.New(os.Stdout, "[SCHEMA-VAULT] - ", log.LstdFlags)
	sd := servicediscovery.NewServiceDiscovery()
	mongoClient := getMongoResource(sd)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	defer mongoClient.Disconnect(ctx)

	healthzHandler := healthz.NewWebHealthzHandler(&healthz.RealTimeProvider{}, 5*time.Second)
	schemaHandler := NewWebServiceSchemaHandler(mongoClient.Client, databaseName)

	httpServer := getHTTPServer()
	makeHTTPHealthzTransport(httpServer, healthzHandler)
	makeHTTPSchemaTransport(httpServer, schemaHandler)

	if err := httpServer.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
