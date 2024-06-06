# resource-contract

`resource-contract` is a Go library that defines a contract for resource wrappers, providing a standardized interface for initializing and retrieving clients. This library is designed to facilitate the integration and management of various resources within a consistent framework.

## Features

- Define a standard interface for initializing resources.
- Retrieve the client instance for the resource.

## Usage

### Defining a Resource Wrapper

To define a resource wrapper, implement the `Resource` interface provided by the `resource-contract` package. This interface requires two methods: `Init` and `GetClient`.

```go
package main

import (
	"fmt"
	"log"
	resourceWrapper "libs/golang/wrappers/core/resource-contract"
)

// ExampleResource is an example implementation of the Resource interface
type ExampleResource struct {
	client string
}

// Init initializes the ExampleResource
func (r *ExampleResource) Init() error {
	r.client = "example_client"
	return nil
}

// GetClient returns the client instance for the ExampleResource
func (r *ExampleResource) GetClient() interface{} {
	return r.client
}

func main() {
	var resource resourceWrapper.Resource = &ExampleResource{}

	err := resource.Init()
	if err != nil {
		log.Fatalf("Failed to initialize resource: %v", err)
	}

	client := resource.GetClient()
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

## Example Implementation: MongoDB Wrapper

Here is an example of how to implement the `Resource` interface for a MongoDB wrapper.

```go
package main

import (
	"fmt"
	"log"
	"libs/golang/wrappers/resources/mongo-wrapper/wrapper"
	resourceWrapper "libs/golang/wrappers/core/resource-contract"
)

func main() {
	var resource resourceWrapper.Resource = mongowrapper.NewMongoDBWrapper()

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

## Example Implementation: Minio Wrapper

Here is an example of how to implement the `Resource` interface for a Minio wrapper.

```go
package main

import (
	"fmt"
	"log"
	"miniowrapper"
	"wrappersresourcecontract"
)

func main() {
	var resource wrappersresourcecontract.Resource = miniowrapper.NewMinioWrapper()

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