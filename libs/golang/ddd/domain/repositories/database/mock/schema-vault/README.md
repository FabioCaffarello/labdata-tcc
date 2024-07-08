# schema-vault/mockrepository

`schema-vault/mockrepository` is a Go library that provides mock implementations of the repository interfaces used in the `schema-vault` domain. This library is primarily used for testing purposes, allowing you to simulate repository behaviors without needing a real database.

## Features

- Mock implementation of `SchemaRepositoryInterface`.
- Support for creating, finding, updating, and deleting schema entities.
- Support for querying schemas based on various attributes.

## Usage

### Mocking a Schema Repository

The `SchemaRepositoryMock` struct provides methods to simulate interactions with schema entities.

#### Example

```go
package main

import (
    "fmt"
    "log"
    "testing"

    "libs/golang/ddd/domain/entities/schema-vault/entity"
    "libs/golang/ddd/domain/repositories/mock/schema-vault/repository"
    "github.com/stretchr/testify/mock"
)

func main() {
    repoMock := new(mockrepository.SchemaRepositoryMock)

    schemaProps := entity.SchemaProps{
        Service:    "exampleService",
        Source:     "exampleSource",
        Provider:   "exampleProvider",
        SchemaType: "exampleSchemaType",
        JsonSchema: map[string]interface{}{
            "type": "object",
            "properties": map[string]interface{}{
                "field1": map[string]interface{}{
                    "type": "string",
                },
                "field2": map[string]interface{}{
                    "type": "string",
                },
            },
            "required": []string{
                "field1",
            },
        },
    }

    schema, err := entity.NewSchema(schemaProps)
    if err != nil {
        log.Fatal(err)
    }

    repoMock.On("Create", schema).Return(nil)

    err = repoMock.Create(schema)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Mocked schema creation: %+v\n", schema)
    repoMock.AssertExpectations(&testing.T{})
}
```

## Mock Implementation Details

The `SchemaRepositoryMock` struct provides mock implementations for the following methods:

- `Create`: Simulates the creation of a schema entity.
- `FindByID`: Simulates finding a schema entity by its ID.
- `FindAll`: Simulates finding all schema entities.
- `Update`: Simulates updating a schema entity.
- `Delete`: Simulates deleting a schema entity.
- `FindAllByServiceAndProvider`: Simulates finding all schema entities by service and provider.
- `FindAllBySourceAndProvider`: Simulates finding all schema entities by source and provider.
- `FindAllByServiceAndSourceAndProvider`: Simulates finding all schema entities by service, source, and provider.
- `FindOneByServiceAndSourceAndProviderAndSchemaType`: Simulates finding one schema entitiy by service, source, provider and schema type.

### Example Test Using the Mock

Here is an example of how to use the `SchemaRepositoryMock` in a test:

```go
package main

import (
    "testing"

    "libs/golang/ddd/domain/entities/schema-vault/entity"
    "libs/golang/ddd/domain/repositories/mock/schema-vault/repository"
    "github.com/stretchr/testify/assert"
)

func TestSchemaRepositoryMock(t *testing.T) {
    repoMock := new(mockrepository.SchemaRepositoryMock)

    schemaProps := entity.SchemaProps{
        Service:    "exampleService",
        Source:     "exampleSource",
        Provider:   "exampleProvider",
        SchemaType: "exampleSchemaType",
        JsonSchema: map[string]interface{}{
            "type": "object",
            "properties": map[string]interface{}{
                "field1": map[string]interface{}{
                    "type": "string",
                },
                "field2": map[string]interface{}{
                    "type": "string",
                },
            },
            "required": []string{
                "field1",
            },
        },
    }

    schema, err := entity.NewSchema(schemaProps)
    assert.NoError(t, err)

    repoMock.On("Create", schema).Return(nil)

    err = repoMock.Create(schema)
    assert.NoError(t, err)

    repoMock.AssertExpectations(t)
}
```
