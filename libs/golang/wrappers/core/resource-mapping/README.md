# resource-mapping

`resource-mapping` is a Go library that provides a centralized way to manage and access various resource wrappers. It allows you to register, initialize, and retrieve resource instances through a singleton instance, ensuring thread-safe operations and simplifying resource management in your applications.

## Features

- Centralized management of resource instances.
- Thread-safe registration and retrieval of resources.
- Easy initialization and access to resource clients.

## Usage

### Creating a Resource Mapping

To create a resource mapping, use the `NewResourceMapping` function which returns a singleton instance of the `Resources` struct.

```go
package main

import (
	"context"
	"log"
	"libs/golang/wrappers/core/resource-mapping"
)

func main() {
	ctx := context.Background()
	resources := resourcemapping.NewResourceMapping(ctx)
	
	if resources == nil {
		log.Fatal("Failed to create resource mapping instance")
	}
}
```

### Registering and Retrieving Resources

You can register resources using the `RegisterResource` method and retrieve them using the `GetResource` method. Resources must implement the `Resource` interface defined in the `resource-contract` package.

```go
package main

import (
	"context"
	"fmt"
	"log"
	"libs/golang/wrappers/core/resource-mapping"
	resourceWrapper "libs/golang/wrappers/core/resource-contract"
)

// ExampleResource is a mock implementation of the Resource interface
type ExampleResource struct {
	client string
}

func (r *ExampleResource) Init() error {
	r.client = "example_client"
	return nil
}

func (r *ExampleResource) GetClient() interface{} {
	return r.client
}

func main() {
	ctx := context.Background()
	resources := resourcemapping.NewResourceMapping(ctx)
	exampleResource := &ExampleResource{}

	// Register the resource
	resources.RegisterResource("example", exampleResource)

	// Retrieve the resource
	retrievedResource, err := resources.GetResource("example")
	if err != nil {
		log.Fatalf("Failed to retrieve resource: %v", err)
	}

	// Initialize the resource
	err = retrievedResource.Init()
	if err != nil {
		log.Fatalf("Failed to initialize resource: %v", err)
	}

	client := retrievedResource.GetClient()
	if client == nil {
		log.Fatalf("Failed to retrieve client")
	}

	fmt.Println("Resource client:", client)
}
```

## Interface Definition

The `Resource` interface defines the standard methods that all resource wrappers must implement.

```go
package wrappersresourcecontract

// Resource defines the interface that all resources should implement
type Resource interface {
	Init() error
	GetClient() interface{}
}
```

## Example Implementations

### MongoDB Wrapper

Here is an example of how to implement the `Resource` interface for a MongoDB wrapper.

```go
package main

import (
	"fmt"
	"log"
	"libs/golang/wrappers/resources/mongo-wrapper/wrapper"
	resourceImpl "libs/golang/wrappers/core/resource-contract"
)

func main() {
	var resource resourceImpl.Resource = mongowrapper.NewMongoDBWrapper()

	err := resource.Init()
	if err != nil {
		log.Fatalf("Failed to initialize MongoDBWrapper: %v", err)
	}

	client := resource.GetClient()
	if client == nil {
		log.Fatalf("MongoDB client is not initialized")
	}

	fmt.Println("MongoDB client initialized successfully")
}
```

### Minio Wrapper

Here is an example of how to implement the `Resource` interface for a Minio wrapper.

```go
package main

import (
	"fmt"
	"log"

	"libs/golang/wrappers/resources/minio-wrapper/wrapper"
	resourceImpl "libs/golang/wrappers/core/resource-contract"
)

func main() {
	var resource resourceImpl.Resource = miniowrapper.NewMinioWrapper()

	err := resource.Init()
	if err != nil {
		log.Fatalf("Failed to initialize MinioWrapper: %v", err)
	}

	client := resource.GetClient()
	if client == nil {
		log.Fatalf("Minio client is not initialized")
	}

	fmt.Println("Minio client initialized successfully")
}
```

### RabbitMQ Wrapper

Here is an example of how to implement the `Resource` interface for a RabbitMQ wrapper.

```go
package main

import (
	"fmt"
	"log"
	"libs/golang/wrappers/resources/rabbitmq-wrapper/wrapper"
	resourceImpl "libs/golang/wrappers/core/resource-contract"
)

func main() {
	var resource resourceImpl.Resource = rabbitmqwrapper.NewRabbitMQWrapper()

	err := resource.Init()
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQWrapper: %v", err)
	}

	client := resource.GetClient()
	if client == nil {
		log.Fatalf("RabbitMQ client is not initialized")
	}

	fmt.Println("RabbitMQ client initialized successfully")
}
```

## Running Tests

To run the test suite for the `resource-mapping` library, execute the following command:

```sh
npx nx test libs-golang-wrappers-core-resource-mapping
```
