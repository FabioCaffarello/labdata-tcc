# input-broker/handlers

`input-broker/handlers` is a Go library that provides HTTP handlers for managing input entities within a web application. This library includes functionalities for creating, updating, deleting, and retrieving input entities through HTTP requests and dispatching events.

## Features

- Create input entities via HTTP requests.
- Update existing input entities.
- Delete input entities.
- Retrieve input entities by various criteria.
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

    "github.com/go-chi/chi/v5"
)

func main() {
    inputRepository := entity.NewInputRepository()
    eventDispatcher := events.NewEventDispatcher()
    inputCreatedEvent := events.NewInputCreatedEvent()

    handler := handlers.NewWebInputHandler(inputRepository, eventDispatcher, inputCreatedEvent)

    r := chi.NewRouter()
    r.Post("/inputs", handler.CreateInput)
    r.Put("/inputs/{id}", handler.UpdateInput)
    r.Delete("/inputs/{id}", handler.DeleteInput)
    r.Get("/inputs", handler.ListAllInputs)
    r.Get("/inputs/{id}", handler.ListInputByID)
    r.Get("/inputs/service/{service}/provider/{provider}", handler.ListInputsByServiceAndProvider)
    r.Get("/inputs/source/{source}/provider/{provider}", handler.ListInputsBySourceAndProvider)
    r.Get("/inputs/service/{service}/source/{source}/provider/{provider}", handler.ListInputsByServiceAndSourceAndProvider)
    r.Get("/inputs/status/{status}/provider/{provider}", handler.ListInputsByStatusAndProvider)
    r.Get("/inputs/status/{status}/service/{service}/provider/{provider}", handler.ListInputsByStatusAndServiceAndProvider)
    r.Get("/inputs/status/{status}/source/{source}/provider/{provider}", handler.ListInputsByStatusAndSourceAndProvider)
    r.Get("/inputs/status/{status}/service/{service}/source/{source}/provider/{provider}", handler.ListInputsByStatusAndServiceAndSourceAndProvider)
    r.Put("/inputs/{id}/status", handler.UpdateInputStatus)

    log.Fatal(http.ListenAndServe(":8080", r))
}
```

### HTTP Endpoints

- `POST /inputs` - Create a new input entity.
- `PUT /inputs/{id}` - Update an existing input entity.
- `DELETE /inputs/{id}` - Delete an input entity.
- `GET /inputs` - List all input entities.
- `GET /inputs/{id}` - Retrieve an input entity by ID.
- `GET /inputs/service/{service}/provider/{provider}` - Retrieve input entities by service and provider.
- `GET /inputs/source/{source}/provider/{provider}` - Retrieve input entities by source and provider.
- `GET /inputs/service/{service}/source/{source}/provider/{provider}` - Retrieve input entities by service, source, and provider.
- `GET /inputs/status/{status}/provider/{provider}` - Retrieve input entities by status and provider.
- `GET /inputs/status/{status}/service/{service}/provider/{provider}` - Retrieve input entities by status, service, and provider.
- `GET /inputs/status/{status}/source/{source}/provider/{provider}` - Retrieve input entities by status, source, and provider.
- `GET /inputs/status/{status}/service/{service}/source/{source}/provider/{provider}` - Retrieve input entities by status, service, source, and provider.
- `PUT /inputs/{id}/status` - Update the status of an existing input entity.

## Testing

To run the tests for the `handlers` package, use the following command:

```sh
npx nx test libs-golang-ddd-adapters-http-handlers-input-broker
```

## Error Handling

The handlers include error handling for various scenarios, such as:

- Invalid request body
- Missing required parameters
- Internal server errors during use case execution

These errors are responded to with appropriate HTTP status codes and error messages.

### Example Error Responses

- `400 Bad Request` - Returned when the request body is invalid or required parameters are missing.
- `500 Internal Server Error` - Returned when there is an error during use case execution or encoding the response.
