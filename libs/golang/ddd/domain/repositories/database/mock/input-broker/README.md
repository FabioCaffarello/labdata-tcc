# input-broker/mockrepository

`input-broker/mockrepository` is a Go library that provides mock implementations of the repository interfaces used in the `input-broker` domain. This library is primarily used for testing purposes, allowing you to simulate repository behaviors without needing a real database.

## Features

- Mock implementation of `InputRepositoryInterface`.
- Support for creating, finding, updating, and deleting input entities.
- Support for querying inputs based on various attributes.

## Usage

### Mocking an Input Repository

The `InputRepositoryMock` struct provides methods to simulate interactions with input entities.

#### Example

```go
package main

import (
    "fmt"
    "log"
    "testing"

    "libs/golang/ddd/domain/entities/input-broker/entity"
    "libs/golang/ddd/domain/repositories/mock/input-broker/repository"
    "github.com/stretchr/testify/mock"
)

func main() {
    repoMock := new(mockrepository.InputRepositoryMock)

    inputProps := entity.InputProps{
        Data: map[string]interface{}{
            "field1": "value1",
            "field2": "value2",
        },
        Service:  "test_service",
        Source:   "test_source",
        Provider: "test_provider",
    }

    input, err := entity.NewInput(inputProps)
    if err != nil {
        log.Fatal(err)
    }

    repoMock.On("Create", input).Return(nil)

    err = repoMock.Create(input)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Mocked input creation: %+v\n", input)
    repoMock.AssertExpectations(&testing.T{})
}
```

## Mock Implementation Details

The `InputRepositoryMock` struct provides mock implementations for the following methods:

- `Create`: Simulates the creation of an input entity.
- `FindByID`: Simulates finding an input entity by its ID.
- `FindAll`: Simulates finding all input entities.
- `Update`: Simulates updating an input entity.
- `Delete`: Simulates deleting an input entity.
- `FindAllByServiceAndProvider`: Simulates finding all input entities by service and provider.
- `FindAllBySourceAndProvider`: Simulates finding all input entities by source and provider.
- `FindAllByServiceAndSourceAndProvider`: Simulates finding all input entities by service, source, and provider.
- `FindAllByStatusAndProvider`: Simulates finding all input entities by status and provider.
- `FindAllByStatusAndServiceAndProvider`: Simulates finding all input entities by status, service, and provider.
- `FindAllByStatusAndSourceAndProvider`: Simulates finding all input entities by status, source, and provider.
- `FindAllByStatusAndServiceAndSourceAndProvider`: Simulates finding all input entities by status, service, source, and provider.

### Example Test Using the Mock

Here is an example of how to use the `InputRepositoryMock` in a test:

```go
package main

import (
    "testing"

    "libs/golang/ddd/domain/entities/input-broker/entity"
    "libs/golang/ddd/domain/repositories/mock/input-broker/repository"
    "github.com/stretchr/testify/assert"
)

func TestInputRepositoryMock(t *testing.T) {
    repoMock := new(mockrepository.InputRepositoryMock)

    inputProps := entity.InputProps{
        Data: map[string]interface{}{
            "field1": "value1",
            "field2": "value2",
        },
        Service:  "test_service",
        Source:   "test_source",
        Provider: "test_provider",
    }

    input, err := entity.NewInput(inputProps)
    assert.NoError(t, err)

    repoMock.On("Create", input).Return(nil)

    err = repoMock.Create(input)
    assert.NoError(t, err)

    repoMock.AssertExpectations(t)
}
```