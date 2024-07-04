# Input Broker Entity

The `input-broker/entity` package is a Go library that provides structures and functions to manage and manipulate input entities within a system. This library includes utilities for converting data between different formats, validating input data, and generating necessary identifiers.

## Features

- Define and manage input entities.
- Convert between `map[string]interface{}` and entity structs.
- Validate input data.
- Generate and handle MD5 and UUID identifiers.

## Usage

### Defining Input Entities

The `Input` struct represents an input entity with attributes such as service, source, provider, and data.

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/entities/input-broker/entity"
)

func main() {
    inputProps := entity.InputProps{
        Data: map[string]interface{}{
            "field1": "value1",
            "field2": "value2",
        },
        Service:  "test_service",
        Source:   "test_source",
        Provider: "test_provider",
    }

    input, err := entity.NewInput(inputProps)
    if err != nil {
        fmt.Println("Error creating input:", err)
        return
    }

    fmt.Printf("Input: %+v\n", input)
}
```

### Converting Input Entities to Maps

The `ToMap` method converts an `Input` entity to a `map[string]interface{}` representation.

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/entities/input-broker/entity"
)

func main() {
    inputProps := entity.InputProps{
        Data: map[string]interface{}{
            "field1": "value1",
            "field2": "value2",
        },
        Service:  "test_service",
        Source:   "test_source",
        Provider: "test_provider",
    }

    input, err := entity.NewInput(inputProps)
    if err != nil {
        fmt.Println("Error creating input:", err)
        return
    }

    inputMap, err := input.ToMap()
    if err != nil {
        fmt.Println("Error converting input to map:", err)
        return
    }

    fmt.Printf("Input as map: %+v\n", inputMap)
}
```

### Validating Input Entities

The `isValid` method ensures that all required fields of an `Input` entity are set.

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/entities/input-broker/entity"
)

func main() {
    inputProps := entity.InputProps{
        Data: map[string]interface{}{
            "field1": "value1",
            "field2": "value2",
        },
        Service:  "test_service",
        Source:   "test_source",
        Provider: "test_provider",
    }

    input, err := entity.NewInput(inputProps)
    if err != nil {
        fmt.Println("Error creating input:", err)
        return
    }

    err = input.isValid()
    if err != nil {
        fmt.Println("Input is invalid:", err)
        return
    }

    fmt.Println("Input is valid")
}
```

## Testing

To run the tests for the `entity` package, use the following command:

```sh
npx nx test libs-golang-ddd-domain-entities-input-broker
```

## Errors

- `ErrInvalidID`: Returned when the ID of an `Input` is invalid.
- `ErrInvalidService`: Returned when the service of an `Input` is invalid.
- `ErrInvalidSource`: Returned when the source of an `Input` is invalid.
- `ErrInvalidProvider`: Returned when the provider of an `Input` is invalid.
- `ErrInvalidProcessingID`: Returned when the processing ID of an `Input` is invalid.
- `ErrInvalidProcessingTimestamp`: Returned when the processing timestamp of an `Input` is invalid.
- `ErrInvalidData`: Returned when the data of an `Input` is invalid.