# output-vault/entity

`output-vault/entity` is a Go library that provides structures and functions to manage and manipulate output entities within a system. This library includes utilities for converting data between different formats, validating output data, and generating necessary identifiers.

## Features

- Define and manage output entities.
- Convert between `map[string]interface{}` and entity structs.
- Validate output data.
- Generate and handle MD5 identifiers.

## Usage

### Defining Output Entities

The `Output` struct represents an output entity with attributes such as service, source, provider, and data.

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/entities/output-vault/entity"
)

func main() {
    outputProps := entity.OutputProps{
        Data: map[string]interface{}{
            "field1": "value1",
            "field2": "value2",
        },
        Service:  "test_service",
        Source:   "test_source",
        Provider: "test_provider",
        Metadata: map[string]interface{}{
            "input_id": "input_id",
            "input": map[string]interface{}{
                "data": map[string]interface{}{
                    "input1": "value1",
                },
                "processing_id":        "processing_id",
                "processing_timestamp": "2023-06-01 00:00:00",
            },
        },
    }

    output, err := entity.NewOutput(outputProps)
    if err != nil {
        fmt.Println("Error creating output:", err)
        return
    }

    fmt.Printf("Output: %+v\n", output)
}
```

### Converting Output Entities to Maps

The `ToMap` method converts an `Output` entity to a `map[string]interface{}` representation.

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/entities/output-vault/entity"
)

func main() {
    outputProps := entity.OutputProps{
        Data: map[string]interface{}{
            "field1": "value1",
            "field2": "value2",
        },
        Service:  "test_service",
        Source:   "test_source",
        Provider: "test_provider",
        Metadata: map[string]interface{}{
            "input_id": "input_id",
            "input": map[string]interface{}{
                "data": map[string]interface{}{
                    "input1": "value1",
                },
                "processing_id":        "processing_id",
                "processing_timestamp": "2023-06-01 00:00:00",
            },
        },
    }

    output, err := entity.NewOutput(outputProps)
    if err != nil {
        fmt.Println("Error creating output:", err)
        return
    }

    outputMap, err := output.ToMap()
    if err != nil {
        fmt.Println("Error converting output to map:", err)
        return
    }

    fmt.Printf("Output as map: %+v\n", outputMap)
}
```

### Validating Output Entities

The `isValid` method ensures that all required fields of an `Output` entity are set.

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/entities/output-vault/entity"
)

func main() {
    outputProps := entity.OutputProps{
        Data: map[string]interface{}{
            "field1": "value1",
            "field2": "value2",
        },
        Service:  "test_service",
        Source:   "test_source",
        Provider: "test_provider",
        Metadata: map[string]interface{}{
            "input_id": "input_id",
            "input": map[string]interface{}{
                "data": map[string]interface{}{
                    "input1": "value1",
                },
                "processing_id":        "processing_id",
                "processing_timestamp": "2023-06-01 00:00:00",
            },
        },
    }

    output, err := entity.NewOutput(outputProps)
    if err != nil {
        fmt.Println("Error creating output:", err)
        return
    }

    err = output.isValid()
    if err != nil {
        fmt.Println("Output is invalid:", err)
        return
    }

    fmt.Println("Output is valid")
}
```

## Testing

To run the tests for the `entity` package, use the following command:

```sh
npx nx test libs-golang-ddd-domain-entities-output-vault
```

## Errors

- `ErrInvalidID`: Returned when the ID of an `Output` is invalid.
- `ErrInvalidService`: Returned when the service of an `Output` is invalid.
- `ErrInvalidSource`: Returned when the source of an `Output` is invalid.
- `ErrInvalidProvider`: Returned when the provider of an `Output` is invalid.
- `ErrInvalidInputID`: Returned when the input ID of an `Output` is invalid.
- `ErrInvalidProcessingID`: Returned when the processing ID of an `Output` is invalid.
- `ErrInvalidProcessingTimestamp`: Returned when the processing timestamp of an `Output` is invalid.
- `ErrInvalidData`: Returned when the data of an `Output` is invalid.
- `ErrInvalidInputData`: Returned when the input data of an `Output` is invalid.
