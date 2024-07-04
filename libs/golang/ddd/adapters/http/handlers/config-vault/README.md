# config-vault/handlers

`config-vault/handlers` is a Go library that provides HTTP handlers for managing configuration entities within a web application. This library includes functionalities for creating, reading, updating, and deleting configuration entities through HTTP requests.

## Features

- Create, read, update, and delete configuration entities via HTTP requests.
- List configurations based on various attributes such as service, provider, and source.
- Handle input validation and error responses.

## Usage

### Creating a WebConfigHandler

The `WebConfigHandler` struct provides methods to handle HTTP requests for configuration operations.

```go
package main

import (
    "log"
    "net/http"

    "libs/golang/ddd/adapters/http/handlers/config-vault/handlers"
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
    handler := handlers.NewWebConfigHandler(repo)

    http.HandleFunc("/configs", handler.CreateConfig)
    http.HandleFunc("/configs", handler.UpdateConfig)
    http.HandleFunc("/configs", handler.DeleteConfig)
    http.HandleFunc("/configs", handler.ListAllConfigs)
    http.HandleFunc("/configs", handler.ListConfigByID)
    http.HandleFunc("/configs", handler.ListConfigsByServiceAndProvider)
    http.HandleFunc("/configs", handler.ListConfigsBySourceAndProvider)
    http.HandleFunc("/configs", handler.ListConfigsByServiceAndSourceAndProvider)
    http.HandleFunc("/configs", handler.ListConfigsByServiceAndProviderAndActive)
    http.HandleFunc("/configs/depends-on", handler.ListConfigsByProviderAndDependencies)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
```


## Testing

To run the tests for the `handlers` package, use the following command:

```sh
npx nx test libs-golang-ddd-adapters-http-handlers-config-vault
```

## Error Handling

The handlers include error handling for various scenarios, such as:

- Invalid request body
- Missing required query parameters
- Internal server errors during use case execution

These errors are responded to with appropriate HTTP status codes and error messages.