# go-mongo

`go-mongo` is a Go library that provides a MongoDB client wrapper with utilities for connecting to and interacting with a MongoDB database. This library facilitates creating a connection to a MongoDB instance, pinging the server, and disconnecting from it.

## Features

- Connect to a MongoDB instance using configuration parameters
- Ping the MongoDB server to check the connection
- Disconnect from the MongoDB instance

## Usage

### Creating a MongoDB Client

```go
package main

import (
	"context"
	"fmt"
	"log"
	"libs/golang/clients/resources/go-mongo/client"
)

func main() {
	config := mongo.Config{
		User:     "testuser",
		Password: "testpassword",
		Host:     "localhost",
		Port:     "27017",
		DBName:   "testdb",
	}

	client, err := mongo.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	defer client.Disconnect(context.Background())

	fmt.Println("Connected to MongoDB")
}
```

### Pinging the MongoDB Server

```go
package main

import (
	"context"
	"fmt"
	"log"
	"libs/golang/clients/resources/go-mongo/client"
)

func main() {
	config := mongo.Config{
		User:     "testuser",
		Password: "testpassword",
		Host:     "localhost",
		Port:     "27017",
		DBName:   "testdb",
	}

	client, err := mongo.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	fmt.Println("MongoDB server is reachable")
}
```

### Disconnecting from MongoDB

```go
package main

import (
	"context"
	"fmt"
	"log"
	"libs/golang/clients/resources/go-mongo/client"
)

func main() {
	config := mongo.Config{
		User:     "testuser",
		Password: "testpassword",
		Host:     "localhost",
		Port:     "27017",
		DBName:   "testdb",
	}

	client, err := mongo.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	err = client.Disconnect(context.Background())
	if err != nil {
		log.Fatalf("Failed to disconnect from MongoDB: %v", err)
	}

	fmt.Println("Disconnected from MongoDB")
}
```

## Testing

To run the tests for the `go-mongo` package, use the following command:

```sh
npx nx test libs-golang-clients-resources-go-mongo
```

This README provides an overview of the `go-mongo` library, along with usage examples for creating a MongoDB client, pinging the server, and disconnecting from the MongoDB instance. Additionally, it includes a section on how to run the tests for the package.