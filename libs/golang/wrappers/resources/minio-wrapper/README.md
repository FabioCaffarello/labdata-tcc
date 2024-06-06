# minio-wrapper

`minio-wrapper` is a Go library that provides a wrapper around a Minio client with an interface for creating and managing Minio connections. This library simplifies the initialization of Minio clients using environment variables and provides methods to retrieve the client.

## Features

- Initialize a Minio client using environment variables
- Retrieve the Minio client instance

## Usage

### Initializing a Minio Client

The `MinioWrapper` struct provides an `Init` method to initialize a Minio client using configuration parameters from environment variables.

```go
package main

import (
	"fmt"
	"log"
	"libs/golang/wrappers/resources/minio-wrapper/miniowrapper"
)

func main() {
	wrapper := miniowrapper.NewMinioWrapper()

	err := wrapper.Init()
	if err != nil {
		log.Fatalf("Failed to initialize Minio client: %v", err)
	}

	client := wrapper.GetClient()
	if client == nil {
		log.Fatalf("Failed to retrieve Minio client")
	}

	fmt.Println("Minio client initialized and retrieved successfully")
}
```

### Retrieving the Minio Client

The `MinioWrapper` provides a method to retrieve the initialized Minio client.

```go
package main

import (
	"fmt"
	"log"
	"libs/golang/wrappers/resources/minio-wrapper/miniowrapper"
)

func main() {
	wrapper := miniowrapper.NewMinioWrapper()

	err := wrapper.Init()
	if err != nil {
		log.Fatalf("Failed to initialize Minio client: %v", err)
	}

	client := wrapper.GetClient()
	if client == nil {
		log.Fatalf("Failed to retrieve Minio client")
	}

	fmt.Println("Minio client retrieved successfully")
}
```

## Environment Variables

The `MinioWrapper` uses the following environment variables to configure the Minio client:

- `MINIO_ENDPOINT`: The endpoint URL for the Minio server
- `MINIO_ACCESS_KEY`: The access key for authentication
- `MINIO_SECRET_KEY`: The secret key for authentication
- `MINIO_USE_SSL`: Set to `true` to use SSL/TLS, `false` otherwise

Ensure these variables are set in your environment before initializing the `MinioWrapper`.

## Testing

To run the tests for the `miniowrapper` package, use the following command:

```sh
npx nx test libs-golang-wrappers-resources-minio-wrapper
```

## Example

Here's an example of how to use the `minio-wrapper` library:

```go
package main

import (
	"fmt"
	"log"
	"os"
	"libs/golang/wrappers/resources/minio-wrapper/miniowrapper"
)

func main() {
	os.Setenv("MINIO_ENDPOINT", "localhost:9000")
	os.Setenv("MINIO_ACCESS_KEY", "minioaccesskey")
	os.Setenv("MINIO_SECRET_KEY", "miniosecretkey")
	os.Setenv("MINIO_USE_SSL", "false")

	wrapper := miniowrapper.NewMinioWrapper()

	err := wrapper.Init()
	if err != nil {
		log.Fatalf("Failed to initialize Minio client: %v", err)
	}

	client := wrapper.GetClient()
	if client == nil {
		log.Fatalf("Failed to retrieve Minio client")
	}

	fmt.Println("Minio client initialized and retrieved successfully")
}
```