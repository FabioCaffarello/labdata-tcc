# output-vault/dto

`output-vault/dto` is a Go library that provides data transfer objects (DTOs) for managing output data within a system. This library includes structures for input and output DTOs, facilitating the transfer and validation of output data.

## Features

- Define DTOs for output input and output.
- Facilitate data transfer between different components of the system.
- Ensure consistency and validation of output data.

## Usage

### Defining Output DTOs

The `OutputDTO` struct represents an output data transfer object with attributes such as service, source, provider, and metadata.

#### Output DTO

The `OutputDTO` struct is used to represent output data for output purposes.

```go
package main

import (
    "fmt"
    outputdto "libs/golang/ddd/dtos/output-vault/output"
    shareddto "libs/golang/ddd/dtos/output-vault/shared"
)

func main() {
    output := outputdto.OutputDTO{
        ID:        "12345",
        Data:      map[string]interface{}{"field1": "value1", "field2": "value2"},
        Service:   "exampleService",
        Source:    "exampleSource",
        Provider:  "exampleProvider",
        Metadata:  shareddto.MetadataDTO{
            InputID: "input_id",
            Input: shareddto.InputDTO{
                Data:                map[string]interface{}{"input1": "value1"},
                ProcessingID:        "processing_id",
                ProcessingTimestamp: "2023-06-01 00:00:00",
            },
        },
        CreatedAt: "2023-06-08 12:00:00",
        UpdatedAt: "2023-06-09 12:00:00",
    }

    fmt.Printf("Output OutputDTO: %+v\n", output)
}
```

#### Input DTO

The `OutputDTO` struct is used to represent output data for input purposes.

```go
package main

import (
    "fmt"
    inputdto "libs/golang/ddd/dtos/output-vault/input"
    shareddto "libs/golang/ddd/dtos/output-vault/shared"
)

func main() {
    output := inputdto.OutputDTO{
        Data:      map[string]interface{}{"field1": "value1", "field2": "value2"},
        Service:   "exampleService",
        Source:    "exampleSource",
        Provider:  "exampleProvider",
        Metadata:  shareddto.MetadataDTO{
            InputID: "input_id",
            Input: shareddto.InputDTO{
                Data:                map[string]interface{}{"input1": "value1"},
                ProcessingID:        "processing_id",
                ProcessingTimestamp: "2023-06-01 00:00:00",
            },
        },
    }

    fmt.Printf("Input OutputDTO: %+v\n", output)
}
```

### Shared DTO

The `MetadataDTO` struct is used to represent metadata for both input and output DTOs.

```go
package main

import (
    "fmt"
    shareddto "libs/golang/ddd/dtos/output-vault/shared"
)

func main() {
    metadata := shareddto.MetadataDTO{
        InputID: "input_id",
        Input: shareddto.InputDTO{
            Data:                map[string]interface{}{"input1": "value1"},
            ProcessingID:        "processing_id",
            ProcessingTimestamp: "2023-06-01 00:00:00",
        },
    }

    fmt.Printf("MetadataDTO: %+v\n", metadata)
}
```
