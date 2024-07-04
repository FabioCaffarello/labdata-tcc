# input-broker/client

`input-broker/client` is a Go library that provides a client for interacting with the input broker API. This library includes functionalities for creating inputs via HTTP requests.

## Features

- Create input entities via HTTP requests.
- Handles request creation, sending, and response processing.

## Usage

### Creating a Client

The `Client` struct provides methods to interact with the input broker API.

```go
package main

import (
    "log"
    "fmt"
    "input-broker/client"
    inputdto "libs/golang/ddd/dtos/input-broker/input"
)

func main() {
    cli := client.NewClient()

    // Create a new input
    inputInput := inputdto.InputDTO{
        Provider: "example-provider",
        Service:  "example-service",
        Source:   "example-source",
        Data:     map[string]interface{}{"key": "value"},
    }

    inputOutput, err := cli.CreateInput(inputInput)
    if err != nil {
        log.Fatalf("Failed to create input: %v", err)
    }
    fmt.Printf("Created input: %+v\n", inputOutput)
}
```

### Client Methods

#### CreateInput

Creates a new input.

```go
func (c *Client) CreateInput(inputInput inputdto.InputDTO) (outputdto.InputDTO, error)
```

## Testing

To run the tests for the `client` package, use the following command:

```sh
npx nx test libs-golang-clients-apis-input-broker-client
```

## Error Handling

The client methods include error handling for various scenarios, such as:

- Invalid request body
- Request timeout
- Internal server errors during API interaction

These errors are handled and returned as appropriate Go errors.