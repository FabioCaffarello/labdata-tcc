# mongo-wrapper

`mongo-wrapper` is a Go library that provides a wrapper around a MongoDB client with an interface for creating and managing MongoDB connections. This library simplifies the initialization of MongoDB clients using environment variables and provides methods to retrieve the client.

## Features

- Initialize a MongoDB client using environment variables
- Retrieve the MongoDB client instance

## Usage

### Initializing a MongoDB Client

The `MongoDBWrapper` struct provides an `Init` method to initialize a MongoDB client using configuration parameters from environment variables.

```go
package main

import (
	"fmt"
	"log"
	"libs/golang/wrappers/resources/mongo-wrapper/wrapper"
)

func main() {
	wrapper := mongowrapper.NewMongoDBWrapper()

	err := wrapper.Init()
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB client: %v", err)
	}

	client := wrapper.GetClient()
	if client == nil {
		log.Fatalf("Failed to retrieve MongoDB client")
	}

	fmt.Println("MongoDB client initialized and retrieved successfully")
}
```

### Retrieving the MongoDB Client

The `MongoDBWrapper` provides a method to retrieve the initialized MongoDB client.

```go
package main

import (
	"fmt"
	"log"
	"libs/golang/wrappers/resources/mongo-wrapper/wrapper"
)

func main() {
	wrapper := mongowrapper.NewMongoDBWrapper()

	err := wrapper.Init()
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB client: %v", err)
	}

	client := wrapper.GetClient()
	if client == nil {
		log.Fatalf("Failed to retrieve MongoDB client")
	}

	fmt.Println("MongoDB client retrieved successfully")
}
```

## Environment Variables

The `MongoDBWrapper` uses the following environment variables to configure the MongoDB client:

- `MONGODB_USER`: The username for authentication
- `MONGODB_PASSWORD`: The password for authentication
- `MONGODB_HOST`: The host of the MongoDB instance
- `MONGODB_PORT`: The port of the MongoDB instance
- `MONGODB_DBNAME`: The name of the database to connect to

Ensure these variables are set in your environment before initializing the `MongoDBWrapper`.

## Testing

To run the tests for the `mongowrapper` package, use the following command:

```sh
npx nx test libs-golang-wrappers-resources-mongo-wrapper
```

## Example

Here's an example of how to use the `mongowrapper` library:

```go
package main

import (
	"fmt"
	"log"
	"os"
	"libs/golang/wrappers/resources/mongo-wrapper/wrapper"
)

func main() {
	os.Setenv("MONGODB_USER", "testuser")
	os.Setenv("MONGODB_PASSWORD", "testpassword")
	os.Setenv("MONGODB_HOST", "localhost")
	os.Setenv("MONGODB_PORT", "27017")
	os.Setenv("MONGODB_DBNAME", "testdb")

	wrapper := mongowrapper.NewMongoDBWrapper()

	err := wrapper.Init()
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB client: %v", err)
	}

	client := wrapper.GetClient()
	if client == nil {
		log.Fatalf("Failed to retrieve MongoDB client")
	}

	fmt.Println("MongoDB client initialized and retrieved successfully")
}
```
