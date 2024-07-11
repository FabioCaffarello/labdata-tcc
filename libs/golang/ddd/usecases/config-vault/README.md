# config-vault/usecase

`config-vault/usecase` is a Go library that provides various use cases for managing configuration entities within a system. This library includes functionalities for creating, updating, deleting, and querying configurations based on different attributes.

## Features

- Create, update, delete, and list configuration entities.
- Query configurations by service, source, provider, and other attributes.
- Validate and convert configuration data between different formats.

## Usage

### Creating a Configuration

The `CreateConfigUseCase` struct provides methods to create a new configuration entity and save it using the repository.

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/entities/config-vault/entity"
    inputdto "libs/golang/ddd/dtos/config-vault/input"
    outputdto "libs/golang/ddd/dtos/config-vault/output"
    "libs/golang/ddd/domain/repositories/database/mongodb/config-vault/repository"
    "libs/golang/ddd/usecases/config-vault/usecase"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    repo := repository.NewConfigRepository(client, "testdb")
    createUseCase := usecase.NewCreateConfigUseCase(repo)

    input := inputdto.ConfigDTO{
        Active:   true,
        Service:  "exampleService",
        Source:   "exampleSource",
        Provider: "exampleProvider",
        JobParameters: shareddto.JobParametersDTO{
            ParserModule: "parser_module1",
        },
        DependsOn: []shareddto.JobDependenciesDTO{
            {Service: "dependencyService", Source: "dependencySource"},
        },
    }

    output, err := createUseCase.Execute(input)
    if err != nil {
        fmt.Println("Error creating config:", err)
        return
    }

    fmt.Printf("Config created: %+v\n", output)
}
```

### Updating a Configuration

The `UpdateConfigUseCase` struct provides methods to update an existing configuration entity and save it using the repository.

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/entities/config-vault/entity"
    inputdto "libs/golang/ddd/dtos/config-vault/input"
    outputdto "libs/golang/ddd/dtos/config-vault/output"
    "libs/golang/ddd/domain/repositories/database/mongodb/config-vault/repository"
    "libs/golang/ddd/usecases/config-vault/usecase"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    repo := repository.NewConfigRepository(client, "testdb")
    updateUseCase := usecase.NewUpdateConfigUseCase(repo)

    input := inputdto.ConfigDTO{
        Active:   true,
        Service:  "exampleService",
        Source:   "exampleSource",
        Provider: "exampleProvider",
        JobParameters: shareddto.JobParametersDTO{
            ParserModule: "parser_module1",
        },
        DependsOn: []shareddto.JobDependenciesDTO{
            {Service: "dependencyService", Source: "dependencySource"},
        },
    }

    output, err := updateUseCase.Execute(input)
    if err != nil {
        fmt.Println("Error updating config:", err)
        return
    }

    fmt.Printf("Config updated: %+v\n", output)
}
```

### Deleting a Configuration

The `DeleteConfigUseCase` struct provides methods to delete an existing configuration entity by its ID.

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/repositories/database/mongodb/config-vault/repository"
    "libs/golang/ddd/usecases/config-vault/usecase"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    repo := repository.NewConfigRepository(client, "testdb")
    deleteUseCase := usecase.NewDeleteConfigUseCase(repo)

    err = deleteUseCase.Execute("exampleID")
    if err != nil {
        fmt.Println("Error deleting config:", err)
        return
    }

    fmt.Println("Config deleted successfully")
}
```

### Listing Configurations by Service

The `ListAllByServiceConfigUseCase` struct provides methods to list all configurations by a specific service.

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/repositories/database/mongodb/config-vault/repository"
    "libs/golang/ddd/usecases/config-vault/usecase"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    repo := repository.NewConfigRepository(client, "testdb")
    listUseCase := usecase.NewListAllByServiceConfigUseCase(repo)

    configs, err := listUseCase.Execute("exampleService")
    if err != nil {
        fmt.Println("Error listing configs:", err)
        return
    }

    for _, config := range configs {
        fmt.Printf("Config: %+v\n", config)
    }
}
```

### Testing

To run the tests for the `usecase` package, use the following command:

```sh
npx nx test libs-golang-ddd-usecases-config-vault
```

## Use Cases

- **CreateConfigUseCase**: Create a new configuration entity.
- **UpdateConfigUseCase**: Update an existing configuration entity.
- **DeleteConfigUseCase**: Delete a configuration entity by its ID.
- **ListAllByServiceConfigUseCase**: List all configurations by a specific service.
- **ListAllConfigUseCase**: List all configurations.
- **ListOneByIDConfigUseCase**: Retrieve a configuration by its ID.
- **ListAllByDependsOnConfigUseCase**: List all configurations by their dependencies.
- **ListAllByServiceAndSourceConfigUseCase**: List all configurations by service and source.
- **ListAllByServiceAndProviderAndActiveConfigUseCase**: List all configurations by service, provider, and active status.
- **ListAllByServiceAndSourceAndProviderConfigUseCase**: List all configurations by service, source, and provider.
- **ListAllBySourceConfigUseCase**: List all configurations by source.

## Errors

- `ErrInvalidID`: Returned when the ID of a `Config` is invalid.
- `ErrInvalidService`: Returned when the service of a `Config` is invalid.
- `ErrInvalidSource`: Returned when the source of a `Config` is invalid.
- `ErrInvalidProvider`: Returned when the provider of a `Config` is invalid.
- `ErrInvalidConfigVersionID`: Returned when the config version ID of a `Config` is invalid.
- `ErrInvalidCreatedAt`: Returned when the created at timestamp of a `Config` is invalid.
