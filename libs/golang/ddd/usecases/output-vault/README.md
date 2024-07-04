# output-vault/usecase

`output-vault/usecase` is a Go library that provides various use cases for managing outputs entities within a system. This library includes functionalities for creating, updating, deleting, and querying configurations based on different attributes.

## Features

- Create, update, delete, and list outputs entities.
- Query outputs by service, source, provider, and other attributes.
- Validate and convert outputs data between different formats.

## Usage

### Creating a Output

The `CreateOutputUseCase` struct provides methods to create a new output entity and save it using the repository.

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/entities/output-vault/entity"
    inputdto "libs/golang/ddd/dtos/output-vault/input"
    outputdto "libs/golang/ddd/dtos/output-vault/output"
    "libs/golang/ddd/domain/repositories/database/mongodb/output-vault/repository"
    "libs/golang/ddd/usecases/output-vault/usecase"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    repo := repository.NewOutputRepository(client, "testdb")
    createUseCase := usecase.NewOutputConfigUseCase(repo)

    input = inputdto.OutputDTO{
		Service:    "test_service",
		Source:     "test_source",
		Provider:   "test_provider",
		OutputType: "test_output_type",
		JsonOutput: shareddto.JsonOutputDTO{
			JsonType: "object",
			Properties: map[string]interface{}{
				"field1": map[string]interface{}{
					"type": "string",
				},
				"field2": map[string]interface{}{
					"type": "string",
				},
			},
			Required: []string{
				"field1",
			},
		},
	}

    output, err := createUseCase.Execute(input)
    if err != nil {
        fmt.Println("Error creating output:", err)
        return
    }

    fmt.Printf("Output created: %+v\n", output)
}
```

### Updating a output

The `UpdateOutputUseCase` struct provides methods to update an existing output entity and save it using the repository.

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/entities/output-vault/entity"
    inputdto "libs/golang/ddd/dtos/output-vault/input"
    outputdto "libs/golang/ddd/dtos/output-vault/output"
    "libs/golang/ddd/domain/repositories/database/mongodb/output-vault/repository"
    "libs/golang/ddd/usecases/output-vault/usecase"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    repo := repository.NewOutputRepository(client, "testdb")
    updateUseCase := usecase.NewUpdateOutputUseCase(repo)

    input = inputdto.OutputDTO{
		Service:    "test_service",
		Source:     "test_source",
		Provider:   "test_provider",
		OutputType: "test_output_type",
		JsonOutput: shareddto.JsonOutputDTO{
			JsonType: "object",
			Properties: map[string]interface{}{
				"field1": map[string]interface{}{
					"type": "string",
				},
				"field2": map[string]interface{}{
					"type": "string",
				},
			},
			Required: []string{
				"field1",
			},
		},
	}

    output, err := updateUseCase.Execute(input)
    if err != nil {
        fmt.Println("Error updating output:", err)
        return
    }

    fmt.Printf("Output updated: %+v\n", output)
}
```

### Deleting a Output

The `DeleteOutputUseCase` struct provides methods to delete an existing output entity by its ID.

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/repositories/database/mongodb/output-vault/repository"
    "libs/golang/ddd/usecases/output-vault/usecase"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    repo := repository.NewOutputRepository(client, "testdb")
    deleteUseCase := usecase.NewDeleteOutputUseCase(repo)

    err = deleteUseCase.Execute("exampleID")
    if err != nil {
        fmt.Println("Error deleting output:", err)
        return
    }

    fmt.Println("Output deleted successfully")
}
```

### Listing Output by Service

The `ListAllByServiceOutputUseCase` struct provides methods to list all output by a specific service.

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/repositories/database/mongodb/output-vault/repository"
    "libs/golang/ddd/usecases/output-vault/usecase"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    repo := repository.NewOutputRepository(client, "testdb")
    listUseCase := usecase.NewListAllByServiceOutputUseCase(repo)

    outputs, err := listUseCase.Execute("exampleService")
    if err != nil {
        fmt.Println("Error listing outputs:", err)
        return
    }

    for _, output := range outputs {
        fmt.Printf("Output: %+v\n", output)
    }
}
```

### Testing

To run the tests for the `usecase` package, use the following command:

```sh
npx nx test libs-golang-ddd-usecases-output-vault
```

## Use Cases

- **CreateOutputUseCase**: Create a new output entity.
- **UpdateOutputUseCase**: Update an existing output entity.
- **DeleteOutputUseCase**: Delete a output entity by its ID.
- **ListAllByServiceOutputUseCase**: List all outputs by a specific service.
- **ListAllOutputUseCase**: List all outputs.
- **ListOneByIDOutputUseCase**: Retrieve a output by its ID.
- **ListAllByServiceAndSourceOutputUseCase**: List all outputs by service and source.
- **ListAllByServiceAndSourceAndProviderOutputUseCase**: List all outputs by service, source, and provider.
- **ListAllBySourceOutputUseCase**: List all outputs by source.

## Errors

- `ErrInvalidID`: Returned when the ID of a `Output` is invalid.
- `ErrInvalidService`: Returned when the service of a `Output` is invalid.
- `ErrInvalidSource`: Returned when the source of a `Output` is invalid.
- `ErrInvalidProvider`: Returned when the provider of a `Output` is invalid.
- `ErrInvalidConfigVersionID`: Returned when the config version ID of a `Output` is invalid.
- `ErrInvalidCreatedAt`: Returned when the created at timestamp of a `Output` is invalid.
