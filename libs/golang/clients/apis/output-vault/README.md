# output-vault/client

`output-vault/client` is a Go library that provides a client for interacting with the output vault API. This library includes functionalities for creating, reading, updating, and deleting output entities via HTTP requests.

## Features

- Create, read, update, and delete output entities via HTTP requests.
- List outputs based on various attributes such as service, provider, source, and dependencies.
- Handles request creation, sending, and response processing.

## Usage

### Creating a Client

The `Client` struct provides methods to interact with the output vault API.

```go
package main

import (
    "log"
    "fmt"
    "output-vault/client"
    inputdto "libs/golang/ddd/dtos/output-vault/input"
)

func main() {
    cli := client.NewClient()

    // Create a new output
    outputInput := inputdto.OutputDTO{
        Data:      map[string]interface{}{"key": "value"},
        Service:   "example-service",
        Source:    "example-source",
        Provider:  "example-provider",
        Metadata:  inputdto.MetadataDTO{},
    }

    outputOutput, err := cli.CreateOutput(outputInput)
    if err != nil {
        log.Fatalf("Failed to create output: %v", err)
    }
    fmt.Printf("Created output: %+v\n", outputOutput)

    // Other client methods can be used similarly...
}
```

### Client Methods

#### CreateOutput

Creates a new output.

```go
func (c *Client) CreateOutput(outputInput inputdto.OutputDTO) (outputdto.OutputDTO, error)
```

#### UpdateOutput

Updates an existing output.

```go
func (c *Client) UpdateOutput(outputInput inputdto.OutputDTO) (outputdto.OutputDTO, error)
```

#### ListAllOutputs

Lists all outputs.

```go
func (c *Client) ListAllOutputs() ([]outputdto.OutputDTO, error)
```

#### ListOutputByID

Gets an output by its ID.

```go
func (c *Client) ListOutputByID(id string) (outputdto.OutputDTO, error)
```

#### DeleteOutput

Deletes an output by its ID.

```go
func (c *Client) DeleteOutput(id string) error
```

#### ListOutputsByServiceAndProvider

Lists outputs by service and provider.

```go
func (c *Client) ListOutputsByServiceAndProvider(service, provider string) ([]outputdto.OutputDTO, error)
```

#### ListOutputsBySourceAndProvider

Lists outputs by source and provider.

```go
func (c *Client) ListOutputsBySourceAndProvider(source, provider string) ([]outputdto.OutputDTO, error)
```

#### ListOutputsByServiceAndSourceAndProvider

Lists outputs by service, source, and provider.

```go
func (c *Client) ListOutputsByServiceAndSourceAndProvider(service, source, provider string) ([]outputdto.OutputDTO, error)
```

## Testing

To run the tests for the `client` package, use the following command:

```sh
npx nx test libs-golang-clients-apis-output-vault-client
```

## Error Handling

The client methods include error handling for various scenarios, such as:

- Invalid request body
- Request timeout
- Internal server errors during API interaction

These errors are handled and returned as appropriate Go errors.