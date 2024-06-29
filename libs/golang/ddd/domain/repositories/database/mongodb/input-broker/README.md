# input-broker/repository

`input-broker/repository` is a Go library that provides a repository layer for managing input entities stored in MongoDB. This library includes functionalities for creating, reading, updating, and deleting input entities, as well as querying inputs based on different attributes.

## Features

- Create, read, update, and delete input entities in MongoDB.
- Query inputs by service, source, provider, and other attributes.
- Handle collection and database existence checks.

## Usage

### Creating an InputRepository

The `InputRepository` struct provides methods to interact with the input entities stored in MongoDB.

```go
package main

import (
    "context"
    "fmt"
    "log"

    "libs/golang/ddd/domain/entities/input-broker/entity"
    "libs/golang/ddd/domain/repositories/database/mongodb/input-broker/repository"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    repo := repository.NewInputRepository(client, "testdb")
    inputProps := entity.InputProps{
        Active:   true,
        Service:  "exampleService",
        Source:   "exampleSource",
        Provider: "exampleProvider",
        DependsOn: []map[string]interface{}{
            {"service": "dependencyService", "source": "dependencySource"},
        },
    }

    input, err := entity.NewInput(inputProps)
    if err != nil {
        log.Fatal(err)
    }

    err = repo.Create(input)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Input created: %+v\n", input)
}
```

### Retrieving an Input by ID

Use the `FindByID` method to retrieve an input by its ID.

```go
package main

import (
    "context"
    "fmt"
    "log"

    "libs/golang/ddd/domain/repositories/database/mongodb/input-broker/repository"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    repo := repository.NewInputRepository(client, "testdb")
    input, err := repo.FindByID("60d5ec49e17e8e304c8f5310")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Input retrieved: %+v\n", input)
}
```

### Updating an Input

Use the `Update` method to update an existing input.

```go
package main

import (
    "context"
    "fmt"
    "log"

    "libs/golang/ddd/domain/entities/input-broker/entity"
    "libs/golang/ddd/domain/repositories/database/mongodb/input-broker/repository"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    repo := repository.NewInputRepository(client, "testdb")
    input, err := repo.FindByID("60d5ec49e17e8e304c8f5310")
    if err != nil {
        log.Fatal(err)
    }

    input.SetActive(false)
    err = repo.Update(input)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Input updated successfully")
}
```

### Deleting an Input

Use the `Delete` method to remove an input by its ID.

```go
package main

import (
    "context"
    "fmt"
    "log"

    "libs/golang/ddd/domain/repositories/database/mongodb/input-broker/repository"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    repo := repository.NewInputRepository(client, "testdb")
    err = repo.Delete("60d5ec49e17e8e304c8f5310")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Input deleted successfully")
}
```

### Querying Inputs

Use the various query methods to retrieve inputs based on different attributes.

```go
package main

import (
    "context"
    "fmt"
    "log"

    "libs/golang/ddd/domain/repositories/database/mongodb/input-broker/repository"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    repo := repository.NewInputRepository(client, "testdb")
    inputs, err := repo.FindAllByServiceAndProvider("exampleProvider", "exampleService")
    if err != nil {
        log.Fatal(err)
    }

    for _, input := range inputs {
        fmt.Printf("Input: %+v\n", input)
    }
}
```

## Testing

To run the tests for the `repository` package, use the following command:

```sh
npx nx test libs-golang-ddd-domain-repositories-database-mongodb-input-broker-repository
```