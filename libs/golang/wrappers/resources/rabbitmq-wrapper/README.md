# rabbitmq-wrapper

`rabbitmq-wrapper` is a Go library that provides a wrapper around a RabbitMQ client with an interface for creating and managing RabbitMQ connections. This library simplifies the initialization of RabbitMQ clients using environment variables and provides methods to retrieve the client.

## Features

- Initialize a RabbitMQ client using environment variables
- Retrieve the RabbitMQ client instance

## Usage

### Initializing a RabbitMQ Client

The `RabbitMQWrapper` struct provides an `Init` method to initialize a RabbitMQ client using configuration parameters from environment variables.

```go
package main

import (
	"fmt"
	"log"
	"libs/golang/wrappers/resources/rabbitmq-wrapper/rabbitmqwrapper"
)

func main() {
	wrapper := rabbitmqwrapper.NewRabbitMQWrapper()

	err := wrapper.Init()
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ client: %v", err)
	}

	client := wrapper.GetClient()
	if client == nil {
		log.Fatalf("Failed to retrieve RabbitMQ client")
	}

	fmt.Println("RabbitMQ client initialized and retrieved successfully")
}
```

### Retrieving the RabbitMQ Client

The `RabbitMQWrapper` provides a method to retrieve the initialized RabbitMQ client.

```go
package main

import (
	"fmt"
	"log"
	"libs/golang/wrappers/resources/rabbitmq-wrapper/rabbitmqwrapper"
)

func main() {
	wrapper := rabbitmqwrapper.NewRabbitMQWrapper()

	err := wrapper.Init()
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ client: %v", err)
	}

	client := wrapper.GetClient()
	if client == nil {
		log.Fatalf("Failed to retrieve RabbitMQ client")
	}

	fmt.Println("RabbitMQ client retrieved successfully")
}
```

## Environment Variables

The `RabbitMQWrapper` uses the following environment variables to configure the RabbitMQ client:

- `RABBITMQ_USER`: The username for authentication
- `RABBITMQ_PASSWORD`: The password for authentication
- `RABBITMQ_HOST`: The host of the RabbitMQ instance
- `RABBITMQ_PORT`: The port of the RabbitMQ instance
- `RABBITMQ_PROTOCOL`: The protocol to use for the connection (e.g., "amqp")
- `RABBITMQ_EXCHANGE_NAME`: The name of the RabbitMQ exchange to use
- `RABBITMQ_EXCHANGE_TYPE`: The type of the RabbitMQ exchange (e.g., "direct", "fanout")

Ensure these variables are set in your environment before initializing the `RabbitMQWrapper`.

## Testing

To run the tests for the `rabbitmqwrapper` package, use the following command:

```sh
npx nx test libs-golang-wrappers-resources-rabbitmq-wrapper
```

## Example

Here's an example of how to use the `rabbitmqwrapper` library:

```go
package main

import (
	"fmt"
	"log"
	"os"
	"libs/golang/wrappers/resources/rabbitmq-wrapper/rabbitmqwrapper"
)

func main() {
	os.Setenv("RABBITMQ_USER", "testuser")
	os.Setenv("RABBITMQ_PASSWORD", "testpassword")
	os.Setenv("RABBITMQ_HOST", "localhost")
	os.Setenv("RABBITMQ_PORT", "5672")
	os.Setenv("RABBITMQ_PROTOCOL", "amqp")
	os.Setenv("RABBITMQ_EXCHANGE_NAME", "testexchange")
	os.Setenv("RABBITMQ_EXCHANGE_TYPE", "direct")

	wrapper := rabbitmqwrapper.NewRabbitMQWrapper()

	err := wrapper.Init()
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ client: %v", err)
	}

	client := wrapper.GetClient()
	if client == nil {
		log.Fatalf("Failed to retrieve RabbitMQ client")
	}

	fmt.Println("RabbitMQ client initialized and retrieved successfully")
}
```
