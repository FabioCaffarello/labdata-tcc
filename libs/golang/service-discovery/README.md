# service-discovery

`service-discovery` is a Go library that facilitates the registration and retrieval of various service resources. It follows the singleton pattern and provides methods for registering and accessing resources like MongoDB, Minio, and RabbitMQ.

## Features

- Register and initialize resources such as MongoDB, Minio, and RabbitMQ.
- Retrieve registered resources by key.
- Singleton instance ensuring a single point of resource management.

## Usage

### Initializing Service Discovery

```go
package main

import (
	"fmt"
	"libs/golang/service-discovery/sd"
)

func main() {
	serviceDiscovery := servicediscovery.NewServiceDiscovery()
	fmt.Println("Service Discovery initialized")
}
```

### Registering a Resource

```go
package main

import (
	"fmt"
	"libs/golang/service-discovery/sd"
	resourceImpl "libs/golang/wrappers/core/resource-contract"
)

type MockResource struct{}

func (m *MockResource) Init() error {
	// Initialize the resource
	return nil
}

func (m *MockResource) GetClient() interface{} {
	// Return the resource client
	return nil
}

func main() {
	serviceDiscovery := servicediscovery.NewServiceDiscovery()

	mockResource := &MockResource{}
	serviceDiscovery.RegisterResource("mock", mockResource)

	fmt.Println("Mock resource registered")
}
```

### Retrieving a Registered Resource

```go
package main

import (
	"fmt"
	"libs/golang/service-discovery/sd"
)

func main() {
	serviceDiscovery := servicediscovery.NewServiceDiscovery()

	resource, err := serviceDiscovery.GetResource("mock")
	if err != nil {
		fmt.Println("Error retrieving resource:", err)
		return
	}

	fmt.Println("Resource retrieved:", resource)
}
```

## Testing

To run the tests for the `servicediscovery` package, use the following command:

```sh
npx nx test libs-golang-service-discovery
```
