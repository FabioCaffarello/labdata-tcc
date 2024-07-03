# input-broker/usecase

`input-broker/usecase` is a Go library that provides various use cases for managing input entities within a system. This library includes functionalities for creating, updating, deleting, and querying inputs based on different attributes.

## Features

- Create, update, delete, and list input entities.
- Query inputs by service, source, provider, status, and other attributes.
- Convert input data between different formats.

## Usage

### Creating an Input

The `CreateInputUseCase` struct provides methods to create a new input entity and save it using the repository.

```go
package main

import (
    "fmt"
    "log"
    "libs/golang/ddd/domain/entities/input-broker/entity"
    inputdto "libs/golang/ddd/dtos/input-broker/input"
    outputdto "libs/golang/ddd/dtos/input-broker/output"
    "libs/golang/ddd/domain/repositories/database/mongodb/input-broker/repository"
    "libs/golang/ddd/usecases/input-broker/usecase"
    events "libs/golang/shared/go-events/amqp_events"
    "context"
)

func main() {
    inputRepo := repository.NewInputRepository("mongodb://localhost:27017", "testdb")
    inputCreatedEvent := events.NewEvent("input.created")
    eventDispatcher := events.NewEventDispatcher()

    createUseCase := usecase.NewCreateInputUseCase(inputRepo, inputCreatedEvent, eventDispatcher)

    input := inputdto.InputDTO{
        Provider: "exampleProvider",
        Service:  "exampleService",
        Source:   "exampleSource",
        Data:     "exampleData",
    }

    output, err := createUseCase.Execute(input)
    if err != nil {
        log.Fatalf("Error creating input: %v", err)
    }

    fmt.Printf("Input created: %+v\n", output)
}
```

### Updating an Input

The `UpdateInputUseCase` struct provides methods to update an existing input entity and save it using the repository.

```go
package main

import (
    "fmt"
    "log"
    "libs/golang/ddd/domain/entities/input-broker/entity"
    inputdto "libs/golang/ddd/dtos/input-broker/input"
    outputdto "libs/golang/ddd/dtos/input-broker/output"
    "libs/golang/ddd/domain/repositories/database/mongodb/input-broker/repository"
    "libs/golang/ddd/usecases/input-broker/usecase"
)

func main() {
    inputRepo := repository.NewInputRepository("mongodb://localhost:27017", "testdb")

    updateUseCase := usecase.NewUpdateInputUseCase(inputRepo)

    input := inputdto.InputDTO{
        Provider: "exampleProvider",
        Service:  "exampleService",
        Source:   "exampleSource",
        Data:     "updatedData",
    }

    output, err := updateUseCase.Execute(input)
    if err != nil {
        log.Fatalf("Error updating input: %v", err)
    }

    fmt.Printf("Input updated: %+v\n", output)
}
```

### Deleting an Input

The `DeleteInputUseCase` struct provides methods to delete an existing input entity by its ID.

```go
package main

import (
    "fmt"
    "log"
    "libs/golang/ddd/domain/repositories/database/mongodb/input-broker/repository"
    "libs/golang/ddd/usecases/input-broker/usecase"
)

func main() {
    inputRepo := repository.NewInputRepository("mongodb://localhost:27017", "testdb")

    deleteUseCase := usecase.NewDeleteInputUseCase(inputRepo)

    err := deleteUseCase.Execute("exampleID")
    if err != nil {
        log.Fatalf("Error deleting input: %v", err)
    }

    fmt.Println("Input deleted successfully")
}
```

### Listing Inputs by Service and Provider

The `ListAllByServiceAndProviderInputUseCase` struct provides methods to list all inputs by a specific service and provider.

```go
package main

import (
    "fmt"
    "log"
    "libs/golang/ddd/domain/repositories/database/mongodb/input-broker/repository"
    "libs/golang/ddd/usecases/input-broker/usecase"
)

func main() {
    inputRepo := repository.NewInputRepository("mongodb://localhost:27017", "testdb")

    listUseCase := usecase.NewListAllByServiceAndProviderInputUseCase(inputRepo)

    inputs, err := listUseCase.Execute("exampleProvider", "exampleService")
    if err != nil {
        log.Fatalf("Error listing inputs: %v", err)
    }

    for _, input := range inputs {
        fmt.Printf("Input: %+v\n", input)
    }
}
```

### Testing

To run the tests for the `usecase` package, use the following command:

```sh
npx nx test libs-golang-ddd-usecases-input-broker
```

## Use Cases

- **CreateInputUseCase**: Create a new input entity.
- **UpdateInputUseCase**: Update an existing input entity.
- **DeleteInputUseCase**: Delete an input entity by its ID.
- **ListAllByServiceAndProviderInputUseCase**: List all inputs by service and provider.
- **ListAllInputUseCase**: List all inputs.
- **ListOneByIDInputUseCase**: Retrieve an input by its ID.
- **ListAllByStatusAndProviderInputUseCase**: List all inputs by status and provider.
- **ListAllByServiceAndSourceAndProviderInputUseCase**: List all inputs by service, source, and provider.
- **ListAllByStatusAndSourceAndProviderInputUseCase**: List all inputs by status, source, and provider.

## Errors

- `ErrInvalidID`: Returned when the ID of an `Input` is invalid.
- `ErrInvalidService`: Returned when the service of an `Input` is invalid.
- `ErrInvalidSource`: Returned when the source of an `Input` is invalid.
- `ErrInvalidProvider`: Returned when the provider of an `Input` is invalid.
- `ErrInvalidData`: Returned when the data of an `Input` is invalid.
- `ErrInvalidStatus`: Returned when the status of an `Input` is invalid.