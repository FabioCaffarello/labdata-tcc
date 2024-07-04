# schema-vault/handlers

`schema-vault/handlers` is a Go library that provides HTTP handlers for managing schemas entities within a web application. This library includes functionalities for creating, reading, updating, and deleting schemas entities through HTTP requests.

## Features

- Create, read, update, and delete schema entities via HTTP requests.
- List schemas based on various attributes such as service, provider, and source.
- Handle input validation and error responses.

## Usage

### Creating a WebSchemaHandler

The `WebSchemaHandler` struct provides methods to handle HTTP requests for schema operations.

```go
package main

import (
    "log"
    "net/http"

    "libs/golang/ddd/adapters/http/handlers/schema-vault/handlers"
    "libs/golang/ddd/domain/entities/schema-vault/entity"
    "libs/golang/ddd/domain/repositories/database/mongodb/schema-vault/repository"

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

    http.HandleFunc("/schemas", handler.CreateConfig)
    http.HandleFunc("/schemas", handler.UpdateConfig)
    http.HandleFunc("/schemas", handler.DeleteConfig)
    http.HandleFunc("/schemas", handler.ListAllConfigs)
    http.HandleFunc("/schemas", handler.ListConfigByID)
    http.HandleFunc("/schemas", handler.ListConfigsByServiceAndProvider)
    http.HandleFunc("/schemas", handler.ListConfigsBySourceAndProvider)
    http.HandleFunc("/schemas", handler.ListConfigsByServiceAndSourceAndProvider)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
```


## Testing

To run the tests for the `handlers` package, use the following command:

```sh
npx nx test libs-golang-ddd-adapters-http-handlers-schema-vault
```

## Error Handling

The handlers include error handling for various scenarios, such as:

- Invalid request body
- Missing required query parameters
- Internal server errors during use case execution

These errors are responded to with appropriate HTTP status codes and error messages.