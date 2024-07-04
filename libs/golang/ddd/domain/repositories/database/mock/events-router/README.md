# events-router/mockrepository

`events-router/mockrepository` is a Go library that provides mock implementations of the repository interfaces used in the `events-router` domain. This library is primarily used for testing purposes, allowing you to simulate repository behaviors without needing a real database.

## Features

- Mock implementation of `EventOrderRepositoryInterface`.
- Support for creating, finding, updating, and deleting event order entities.
- Support for querying event orders based on various attributes.

## Usage

### Mocking an Event Order Repository

The `EventOrderRepositoryMock` struct provides methods to simulate interactions with event order entities.

#### Example

```go
package main

import (
    "fmt"
    "log"
    "testing"

    "libs/golang/ddd/domain/entities/events-router/entity"
    "libs/golang/ddd/domain/repositories/mock/events-router/repository"
    "github.com/stretchr/testify/mock"
)

func main() {
    repoMock := new(mockrepository.EventOrderRepositoryMock)

    eventOrderProps := entity.EventOrderProps{
        Service:      "test_service",
        Source:       "test_source",
        Provider:     "test_provider",
        ProcessingID: "processing_id",
        Data: map[string]interface{}{
            "field1": "value1",
            "field2": "value2",
        },
    }

    eventOrder, err := entity.NewEventOrder(eventOrderProps)
    if err != nil {
        log.Fatal(err)
    }

    repoMock.On("Create", eventOrder).Return(nil)

    err = repoMock.Create(eventOrder)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Mocked event order creation: %+v\n", eventOrder)
    repoMock.AssertExpectations(&testing.T{})
}
```

## Mock Implementation Details

The `EventOrderRepositoryMock` struct provides mock implementations for the following methods:

- `Create`: Simulates the creation of an event order entity.
- `FindByID`: Simulates finding an event order entity by its ID.
- `FindAll`: Simulates finding all event order entities.
- `Delete`: Simulates deleting an event order entity by its ID.

### Example Test Using the Mock

Here is an example of how to use the `EventOrderRepositoryMock` in a test:

```go
package main

import (
    "testing"

    "libs/golang/ddd/domain/entities/events-router/entity"
    "libs/golang/ddd/domain/repositories/mock/events-router/repository"
    "github.com/stretchr/testify/assert"
)

func TestEventOrderRepositoryMock(t *testing.T) {
    repoMock := new(mockrepository.EventOrderRepositoryMock)

    eventOrderProps := entity.EventOrderProps{
        Service:      "test_service",
        Source:       "test_source",
        Provider:     "test_provider",
        ProcessingID: "processing_id",
        Data: map[string]interface{}{
            "field1": "value1",
            "field2": "value2",
        },
    }

    eventOrder, err := entity.NewEventOrder(eventOrderProps)
    assert.NoError(t, err)

    repoMock.On("Create", eventOrder).Return(nil)

    err = repoMock.Create(eventOrder)
    assert.NoError(t, err)

    repoMock.AssertExpectations(t)
}
```
