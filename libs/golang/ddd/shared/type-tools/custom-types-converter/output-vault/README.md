# output-vault/converter

`output-vault/converter` is a Go library that provides utility functions to convert between data transfer objects (DTOs) and entities within the output vault domain. This library facilitates the transformation of metadata data structures between different layers of the application.

## Features

- Convert metadata from DTOs to entities.
- Convert metadata from entities to DTOs.
- Convert metadata from DTOs to a map.

## Usage

### Converting Metadata DTOs to Entities

The `ConvertMetadataDTOToEntity` function converts a `MetadataDTO` to a `Metadata` entity.

#### Example

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/entities/output-vault/entity"
    shareddto "libs/golang/ddd/dtos/output-vault/shared"
    "libs/golang/ddd/domain/converter"
)

func main() {
    dtoMetadata := shareddto.MetadataDTO{
        InputID: "input_id",
        Input: shareddto.InputDTO{
            Data: map[string]interface{}{
                "input1": "value1",
            },
            ProcessingID:        "processing_id",
            ProcessingTimestamp: "2023-06-01 00:00:00",
        },
    }

    entityMetadata := converter.ConvertMetadataDTOToEntity(dtoMetadata)
    fmt.Printf("Converted entity: %+v\n", entityMetadata)
}
```

### Converting Metadata Entities to DTOs

The `ConvertMetadataEntityToDTO` function converts a `Metadata` entity to a `MetadataDTO`.

#### Example

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/entities/output-vault/entity"
    shareddto "libs/golang/ddd/dtos/output-vault/shared"
    "libs/golang/ddd/domain/converter"
)

func main() {
    entityMetadata := entity.Metadata{
        InputID: "input_id",
        Input: entity.Input{
            Data: map[string]interface{}{
                "input1": "value1",
            },
            ProcessingID:        "processing_id",
            ProcessingTimestamp: "2023-06-01 00:00:00",
        },
    }

    dtoMetadata := converter.ConvertMetadataEntityToDTO(entityMetadata)
    fmt.Printf("Converted DTO: %+v\n", dtoMetadata)
}
```

### Converting Metadata DTOs to a Map

The `ConvertMetadataDTOToMap` function converts a `MetadataDTO` to a map.

#### Example

```go
package main

import (
    "fmt"
    shareddto "libs/golang/ddd/dtos/output-vault/shared"
    "libs/golang/ddd/domain/converter"
)

func main() {
    dtoMetadata := shareddto.MetadataDTO{
        InputID: "input_id",
        Input: shareddto.InputDTO{
            Data: map[string]interface{}{
                "input1": "value1",
            },
            ProcessingID:        "processing_id",
            ProcessingTimestamp: "2023-06-01 00:00:00",
        },
    }

    metadataMap := converter.ConvertMetadataDTOToMap(dtoMetadata)
    fmt.Printf("Converted map: %+v\n", metadataMap)
}
```

## Testing

To run the tests for the `converter` package, use the following command:

```sh
npx nx test libs-golang-ddd-shared-type-tools-custom-types-converter-output-vault
```
