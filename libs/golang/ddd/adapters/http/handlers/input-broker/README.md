# input-broker/handlers

`input-broker/handlers` is a Go library that provides HTTP handlers for managing input entities within a web application. This library includes functionalities for creating input entities through HTTP requests and dispatching events.

## Features

- Create input entities via HTTP requests.
- Dispatch events upon successful creation of input entities.
- Handle input validation and error responses.

## Usage

### Creating a WebInputHandler

The `WebInputHandler` struct provides methods to handle HTTP requests for input operations.

```go
package main

import (
    "log"
    "net/http"

    "libs/golang/ddd/domain/entities/input-broker/entity"
    "libs/golang/ddd/usecases/input-broker/usecase"
    "libs/golang/ddd/adapters/http/handlers/input-broker/handlers"
    events "libs/golang/shared/go-events/amqp_events"
)

func main() {
    inputRepository := entity.NewInputRepository()
    eventDispatcher := events.NewEventDispatcher()
    inputCreatedEvent := events.NewInputCreatedEvent()

    handler := handlers.NewWebInputHandler(inputRepository, eventDispatcher, inputCreatedEvent)

    http.HandleFunc("/inputs", handler.CreateInput)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## Testing

To run the tests for the `handlers` package, use the following command:

```sh
npx nx test libs-golang-ddd-adapters-http-handlers-input-broker
```

## Error Handling

The handlers include error handling for various scenarios, such as:

- Invalid request body
- Internal server errors during use case execution

These errors are responded to with appropriate HTTP status codes and error messages.
