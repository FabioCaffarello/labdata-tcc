# schema-vault/client

`schema-vault/client` is a Go library that provides a client for interacting with the schema vault API. This library includes functionalities for creating, reading, updating, and deleting schema entities via HTTP requests.

## Features

- Create, read, update, and delete schema entities via HTTP requests.
- List schemas based on various attributes such as service, provider, source, and dependencies.
- Handles request creation, sending, and response processing.

## Usage

### Creating a Client

The `Client` struct provides methods to interact with the schema vault API.

```go
package main

import (
    "log"
    "fmt"
    "schema-vault/client"
    inputdto "libs/golang/ddd/dtos/schema-vault/input"
)

func main() {
    cli := client.NewClient()

    // Create a new schema
    schemaInput := inputdto.SchemaDTO{
        Service:    "example-service",
        Source:     "example-source",
        Provider:   "example-provider",
        SchemaType: "example-type",
        JsonSchema: inputdto.JsonSchemaDTO{},
    }

    schemaOutput, err := cli.CreateSchema(schemaInput)
    if err != nil {
        log.Fatalf("Failed to create schema: %v", err)
    }
    fmt.Printf("Created schema: %+v\n", schemaOutput)

    // Other client methods can be used similarly...
}
```

### Client Methods

#### CreateSchema

Creates a new schema.

```go
func (c *Client) CreateSchema(schemaInput inputdto.SchemaDTO) (outputdto.SchemaDTO, error)
```

#### UpdateSchema

Updates an existing schema.

```go
func (c *Client) UpdateSchema(schemaInput inputdto.SchemaDTO) (outputdto.SchemaDTO, error)
```

#### ListAllSchemas

Lists all schemas.

```go
func (c *Client) ListAllSchemas() ([]outputdto.SchemaDTO, error)
```

#### ListSchemaByID

Gets a schema by its ID.

```go
func (c *Client) ListSchemaByID(id string) (outputdto.SchemaDTO, error)
```

#### DeleteSchema

Deletes a schema by its ID.

```go
func (c *Client) DeleteSchema(id string) error
```

#### ListSchemasByServiceAndProvider

Lists schemas by service and provider.

```go
func (c *Client) ListSchemasByServiceAndProvider(service, provider string) ([]outputdto.SchemaDTO, error)
```

#### ListSchemasBySourceAndProvider

Lists schemas by source and provider.

```go
func (c *Client) ListSchemasBySourceAndProvider(source, provider string) ([]outputdto.SchemaDTO, error)
```

#### ListSchemasByServiceAndSourceAndProvider

Lists schemas by service, source, and provider.

```go
func (c *Client) ListSchemasByServiceAndSourceAndProvider(service, source, provider string) ([]outputdto.SchemaDTO, error)
```

## Testing

To run the tests for the `client` package, use the following command:

```sh
npx nx test libs-golang-clients-apis-schema-vault-client
```

## Error Handling

The client methods include error handling for various scenarios, such as:

- Invalid request body
- Request timeout
- Internal server errors during API interaction

These errors are handled and returned as appropriate Go errors.
