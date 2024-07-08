# Events Router Entity

The `events-router/entity` package is a Go library that provides structures and functions to manage and manipulate event order entities within a system. This library includes utilities for converting data between different formats, validating event data, and generating necessary identifiers.

## Features

- Define and manage event order entities.
- Convert between `map[string]interface{}` and entity structs.
- Validate event data.
- Generate and handle MD5 and UUID identifiers.

## Usage

### Defining Event Order Entities

The `EventOrder` struct represents an event order entity with attributes such as service, source, provider, and data.

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/entities/events-router/entity"
)

func main() {
    eventOrderProps := entity.EventOrderProps{
        Data: map[string]interface{}{
            "key": "value",
        },
        Service:      "test_service",
        Source:       "test_source",
        Provider:     "test_provider",
		InputID:      "input-id",
        ProcessingID: "xyz789",
    }

    eventOrder, err := entity.NewEventOrder(eventOrderProps)
    if err != nil {
        fmt.Println("Error creating event order:", err)
        return
    }

    fmt.Printf("EventOrder: %+v\n", eventOrder)
}
```

### Converting Event Order Entities to Maps

The `ToMap` method converts an `EventOrder` entity to a `map[string]interface{}` representation.

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/entities/events-router/entity"
)

func main() {
    eventOrderProps := entity.EventOrderProps{
        Data: map[string]interface{}{
            "key": "value",
        },
        Service:      "test_service",
        Source:       "test_source",
        Provider:     "test_provider",
		InputID:      "input-id",
        ProcessingID: "xyz789",
    }

    eventOrder, err := entity.NewEventOrder(eventOrderProps)
    if err != nil {
        fmt.Println("Error creating event order:", err)
        return
    }

    eventOrderMap, err := eventOrder.ToMap()
    if err != nil {
        fmt.Println("Error converting event order to map:", err)
        return
    }

    fmt.Printf("EventOrder as map: %+v\n", eventOrderMap)
}
```

### Validating Event Order Entities

The `isValid` method ensures that all required fields of an `EventOrder` entity are set.

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/entities/events-router/entity"
)

func main() {
    eventOrderProps := entity.EventOrderProps{
        Data: map[string]interface{}{
            "key": "value",
        },
        Service:      "test_service",
        Source:       "test_source",
        Provider:     "test_provider",
		InputID:      "input-id",
        ProcessingID: "xyz789",
    }

    eventOrder, err := entity.NewEventOrder(eventOrderProps)
    if err != nil {
        fmt.Println("Error creating event order:", err)
        return
    }

    err = eventOrder.isValid()
    if err != nil {
        fmt.Println("Event order is invalid:", err)
        return
    }

    fmt.Println("Event order is valid")
}
```

## Testing

To run the tests for the `entity` package, use the following command:

```sh
npx nx test libs-golang-ddd-domain-entities-events-router
```

## Errors

- `ErrInvalidID`: Returned when the ID of an `EventOrder` is invalid.
- `ErrInvalidService`: Returned when the service of an `EventOrder` is invalid.
- `ErrInvalidSource`: Returned when the source of an `EventOrder` is invalid.
- `ErrInvalidProvider`: Returned when the provider of an `EventOrder` is invalid.
- `ErrInvalidProcessingID`: Returned when the processing ID of an `EventOrder` is invalid.