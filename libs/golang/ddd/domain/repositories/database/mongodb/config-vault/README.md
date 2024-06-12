# config-vault/repository

`config-vault/repository` is a Go library that provides a repository layer for managing configuration entities stored in MongoDB. This library includes functionalities for creating, reading, updating, and deleting configuration entities, as well as querying configurations based on different attributes.

## Features

- Create, read, update, and delete configuration entities in MongoDB.
- Query configurations by service, source, provider, and other attributes.
- Handle collection and database existence checks.

## Usage

### Creating a ConfigRepository

The `ConfigRepository` struct provides methods to interact with the configuration entities stored in MongoDB.

```go
package main

import (
    "context"
    "fmt"
    "log"

    "libs/golang/ddd/domain/entities/config-vault/entity"
    "libs/golang/ddd/domain/repositories/database/mongodb/config-vault/repository"

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
    configProps := entity.ConfigProps{
        Active:   true,
        Service:  "exampleService",
        Source:   "exampleSource",
        Provider: "exampleProvider",
        DependsOn: []map[string]interface{}{
            {"service": "dependencyService", "source": "dependencySource"},
        },
        UpdatedAt: "2023-06-08 12:00:00",
    }

    config, err := entity.NewConfig(configProps)
    if err != nil {
        log.Fatal(err)
    }

    err = repo.Create(config)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Config created: %+v\n", config)
}
```

### Retrieving a Config by ID

Use the `FindByID` method to retrieve a configuration by its ID.

```go
package main

import (
    "context"
    "fmt"
    "log"

    "libs/golang/ddd/domain/repositories/database/mongodb/config-vault/repository"

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
    config, err := repo.FindByID("60d5ec49e17e8e304c8f5310")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Config retrieved: %+v\n", config)
}
```

### Updating a Config

Use the `Update` method to update an existing configuration.

```go
package main

import (
    "context"
    "fmt"
    "log"

    "libs/golang/ddd/domain/entities/config-vault/entity"
    "libs/golang/ddd/domain/repositories/database/mongodb/config-vault/repository"

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
    config, err := repo.FindByID("60d5ec49e17e8e304c8f5310")
    if err != nil {
        log.Fatal(err)
    }

    config.SetActive(false)
    err = repo.Update(config)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Config updated successfully")
}
```

### Deleting a Config

Use the `Delete` method to remove a configuration by its ID.

```go
package main

import (
    "context"
    "fmt"
    "log"

    "libs/golang/ddd/domain/repositories/database/mongodb/config-vault/repository"

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
    err = repo.Delete("60d5ec49e17e8e304c8f5310")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Config deleted successfully")
}
```

### Querying Configurations

Use the various query methods to retrieve configurations based on different attributes.

```go
package main

import (
    "context"
    "fmt"
    "log"

    "libs/golang/ddd/domain/repositories/database/mongodb/config-vault/repository"

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
    configs, err := repo.FindAllByService("exampleService")
    if err != nil {
        log.Fatal(err)
    }

    for _, config := range configs {
        fmt.Printf("Config: %+v\n", config)
    }
}
```

## Testing

To run the tests for the `repository` package, use the following command:

```sh
npx nx test libs-golang-ddd-domain-repositories-database-mongodb-config-vault-repository
```
