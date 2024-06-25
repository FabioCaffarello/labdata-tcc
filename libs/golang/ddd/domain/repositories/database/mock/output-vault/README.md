# output-vault/mockrepository

`output-vault/mockrepository` is a Go library that provides mock implementations of the repository interfaces used in the `output-vault` domain. This library is primarily used for testing purposes, allowing you to simulate repository behaviors without needing a real database.

## Features

- Mock implementation of `OutputRepositoryInterface`.
- Support for creating, finding, updating, and deleting output entities.
- Support for querying outputs based on various attributes.

## Usage

### Mocking an Output Repository

The `OutputRepositoryMock` struct provides methods to simulate interactions with output entities.

#### Example

```go
package main

import (
    "fmt"
    "log"
    "testing"

    "libs/golang/ddd/domain/entities/output-vault/entity"
    "libs/golang/ddd/domain/repositories/mock/output-vault/repository"
    "github.com/stretchr/testify/mock"
)

func main() {
    repoMock := new(mockrepository.OutputRepositoryMock)

    outputProps := entity.OutputProps{
        Data: map[string]interface{}{
            "field1": "value1",
            "field2": "value2",
        },
        Service:  "test_service",
        Source:   "test_source",
        Provider: "test_provider",
        Metadata: map[string]interface{}{
            "input_id": "input_id",
            "input": map[string]interface{}{
                "data": map[string]interface{}{
                    "input1": "value1",
                },
                "processing_id":        "processing_id",
                "processing_timestamp": "2023-06-01 00:00:00",
            },
        },
    }

    output, err := entity.NewOutput(outputProps)
    if err != nil {
        log.Fatal(err)
    }

    repoMock.On("Create", output).Return(nil)

    err = repoMock.Create(output)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Mocked output creation: %+v\n", output)
    repoMock.AssertExpectations(&testing.T{})
}
```

## Mock Implementation Details

The `OutputRepositoryMock` struct provides mock implementations for the following methods:

- `Create`: Simulates the creation of an output entity.
- `FindByID`: Simulates finding an output entity by its ID.
- `FindAll`: Simulates finding all output entities.
- `Update`: Simulates updating an output entity.
- `Delete`: Simulates deleting an output entity.
- `FindAllByServiceAndProvider`: Simulates finding all output entities by service and provider.
- `FindAllBySourceAndProvider`: Simulates finding all output entities by source and provider.
- `FindAllByServiceAndSourceAndProvider`: Simulates finding all output entities by service, source, and provider.

### Example Test Using the Mock

Here is an example of how to use the `OutputRepositoryMock` in a test:

```go
package main

import (
    "testing"

    "libs/golang/ddd/domain/entities/output-vault/entity"
    "libs/golang/ddd/domain/repositories/mock/output-vault/repository"
    "github.com/stretchr/testify/assert"
)

func TestOutputRepositoryMock(t *testing.T) {
    repoMock := new(mockrepository.OutputRepositoryMock)

    outputProps := entity.OutputProps{
        Data: map[string]interface{}{
            "field1": "value1",
            "field2": "value2",
        },
        Service:  "test_service",
        Source:   "test_source",
        Provider: "test_provider",
        Metadata: map[string]interface{}{
            "input_id": "input_id",
            "input": map[string]interface{}{
                "data": map[string]interface{}{
                    "input1": "value1",
                },
                "processing_id":        "processing_id",
                "processing_timestamp": "2023-06-01 00:00:00",
            },
        },
    }

    output, err := entity.NewOutput(outputProps)
    assert.NoError(t, err)

    repoMock.On("Create", output).Return(nil)

    err = repoMock.Create(output)
    assert.NoError(t, err)

    repoMock.AssertExpectations(t)
}
```
