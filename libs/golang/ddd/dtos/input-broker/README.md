# input-broker/dto

`input-broker/dto` is a Go library that provides data transfer objects (DTOs) for managing input and output data within a system. This library includes structures for input, output, and shared DTOs, facilitating the transfer and validation of input data.

## Features

- Define DTOs for input and output data.
- Facilitate data transfer between different components of the system.
- Ensure consistency and validation of input data.

## Usage

### Defining Input DTOs

The `InputDTO` struct represents an input data transfer object with attributes such as provider, service, source, and data.

#### Input DTO

The `inputdto.InputDTO` struct is used to represent input data for various services.

```go
package main

import (
    "fmt"
    "libs/golang/ddd/dtos/input-broker/inputdto"
)

func main() {
    input := inputdto.InputDTO{
        Provider: "exampleProvider",
        Service:  "exampleService",
        Source:   "exampleSource",
        Data: map[string]interface{}{
            "key1": "value1",
            "key2": "value2",
        },
    }

    fmt.Printf("InputDTO: %+v\n", input)
}
```

### Defining Output DTOs

The `OutputDTO` struct represents an output data transfer object with attributes such as ID, data, metadata, status, created at, and updated at.

#### Output DTO

The `outputdto.OutputDTO` struct is used to represent output data, including metadata and status information.

```go
package main

import (
    "fmt"
    "libs/golang/ddd/dtos/input-broker/outputdto"
    shareddto "libs/golang/ddd/dtos/input-broker/shared"
)

func main() {
    output := outputdto.OutputDTO{
        ID:   "12345",
        Data: map[string]interface{}{
            "key1": "value1",
            "key2": "value2",
        },
        Metadata: shareddto.MetadataDTO{
            Provider:            "exampleProvider",
            Service:             "exampleService",
            Source:              "exampleSource",
            ProcessingID:        "process123",
            ProcessingTimestamp: "2023-06-08 12:00:00",
        },
        Status: shareddto.StatusDTO{
            Code:   200,
            Detail: "Success",
        },
        CreatedAt: "2023-06-08 12:00:00",
        UpdatedAt: "2023-06-09 12:00:00",
    }

    fmt.Printf("OutputDTO: %+v\n", output)
}
```

### Shared DTO

The `shareddto` package includes shared DTOs such as `MetadataDTO` and `StatusDTO`, which are used by both input and output DTOs.

```go
package main

import (
    "fmt"
    shareddto "libs/golang/ddd/dtos/input-broker/shared"
)

func main() {
    metadata := shareddto.MetadataDTO{
        Provider:            "exampleProvider",
        Service:             "exampleService",
        Source:              "exampleSource",
        ProcessingID:        "process123",
        ProcessingTimestamp: "2023-06-08 12:00:00",
    }

    status := shareddto.StatusDTO{
        Code:   200,
        Detail: "Success",
    }

    fmt.Printf("MetadataDTO: %+v\n", metadata)
    fmt.Printf("StatusDTO: %+v\n", status)
}
```
