# events-router/repository

`events-router/repository` is a Go library that provides a repository layer for managing `EventOrder` entities stored in a document-based database. This library includes functionalities for creating, reading, updating, and deleting `EventOrder` entities, as well as querying `EventOrder` entities based on different attributes.

## Features

- Create, read, and delete `EventOrder` entities in the document-based database.
- Query `EventOrder` entities by ID.
- Handle collection and database existence checks.

## Usage

### Creating an EventOrderRepository

The `EventOrderRepository` struct provides methods to interact with the `EventOrder` entities stored in the document-based database.

```go
package main

import (
    "fmt"
    "log"

    "libs/golang/clients/resources/go-docdb/client"
    "libs/golang/database/go-docdb/database"
    "libs/golang/ddd/domain/entities/events-router/entity"
    "libs/golang/ddd/domain/repositories/events-router/repository"
)

func main() {
    db := database.NewInMemoryDocBD("test_database")
    client := client.NewClient(db)
    repo := repository.NewModelOrderRepository(client, "test_database")
    
    eventOrderProps := entity.EventOrderProps{
        Service:      "exampleService",
        Source:       "exampleSource",
        Provider:     "exampleProvider",
        ProcessingID: "xyz789",
        Data: map[string]interface{}{
            "field1": "value1",
            "field2": "value2",
        },
    }

    eventOrder, err := entity.NewEventOrder(eventOrderProps)
    if err != nil {
        log.Fatal(err)
    }

    err = repo.Create(eventOrder)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("EventOrder created: %+v\n", eventOrder)
}
```

### Retrieving an EventOrder by ID

Use the `FindByID` method to retrieve an `EventOrder` by its ID.

```go
package main

import (
    "fmt"
    "log"

    "libs/golang/clients/resources/go-docdb/client"
    "libs/golang/database/go-docdb/database"
    "libs/golang/ddd/domain/repositories/events-router/repository"
)

func main() {
    db := database.NewInMemoryDocBD("test_database")
    client := client.NewClient(db)
    repo := repository.NewModelOrderRepository(client, "test_database")
    
    eventOrder, err := repo.FindByID("9b97f68f63f3faa91d2d6558428f1863")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("EventOrder retrieved: %+v\n", eventOrder)
}
```

### Deleting an EventOrder

Use the `Delete` method to remove an `EventOrder` by its ID.

```go
package main

import (
    "fmt"
    "log"

    "libs/golang/clients/resources/go-docdb/client"
    "libs/golang/database/go-docdb/database"
    "libs/golang/ddd/domain/repositories/events-router/repository"
)

func main() {
    db := database.NewInMemoryDocBD("test_database")
    client := client.NewClient(db)
    repo := repository.NewModelOrderRepository(client, "test_database")
    
    err := repo.Delete("9b97f68f63f3faa91d2d6558428f1863")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("EventOrder deleted successfully")
}
```

### Retrieving All EventOrders

Use the `FindAll` method to retrieve all `EventOrder` entities.

```go
package main

import (
    "fmt"
    "log"

    "libs/golang/clients/resources/go-docdb/client"
    "libs/golang/database/go-docdb/database"
    "libs/golang/ddd/domain/repositories/events-router/repository"
)

func main() {
    db := database.NewInMemoryDocBD("test_database")
    client := client.NewClient(db)
    repo := repository.NewModelOrderRepository(client, "test_database")
    
    eventOrders, err := repo.FindAll()
    if err != nil {
        log.Fatal(err)
    }

    for _, eventOrder := range eventOrders {
        fmt.Printf("EventOrder: %+v\n", eventOrder)
    }
}
```

## Testing

To run the tests for the `repository` package, use the following command:

```sh
npx nx test libs-golang-ddd-domain-repositories-database-in-memory-go-docdb-events-router
```