# input-broker/converter

`input-broker/converter` is a Go library that provides utility functions to convert between data transfer objects (DTOs) and entities within the input-broker domain. This library facilitates the transformation of metadata and status data structures between different layers of the application.

## Features

- Convert metadata from DTOs to entities.
- Convert metadata from entities to DTOs.
- Convert status from DTOs to entities.
- Convert status from entities to DTOs.

## Usage

### Converting Metadata DTOs to Entities

The `ConvertMetadataDTOToEntity` function converts a `MetadataDTO` to a `Metadata` entity.

#### Example

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/entities/input-broker/entity"
    shareddto "libs/golang/ddd/dtos/input-broker/shared"
    "libs/golang/ddd/domain/converter"
)

func main() {
    metadataDTO := shareddto.MetadataDTO{
        Provider:            "test_provider",
        Service:             "test_service",
        Source:              "test_source",
        ProcessingID:        "test_processing_id",
        ProcessingTimestamp: "2023-07-02T12:34:56Z",
    }

    metadata := converter.ConvertMetadataDTOToEntity(metadataDTO)
    fmt.Printf("Converted entity: %+v\n", metadata)
}
```

### Converting Metadata Entities to DTOs

The `ConvertMetadataEntityToDTO` function converts a `Metadata` entity to a `MetadataDTO`.

#### Example

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/entities/input-broker/entity"
    shareddto "libs/golang/ddd/dtos/input-broker/shared"
    "libs/golang/ddd/domain/converter"
)

func main() {
    metadata := entity.Metadata{
        Provider:            "test_provider",
        Service:             "test_service",
        Source:              "test_source",
        ProcessingID:        "test_processing_id",
        ProcessingTimestamp: "2023-07-02T12:34:56Z",
    }

    metadataDTO := converter.ConvertMetadataEntityToDTO(metadata)
    fmt.Printf("Converted DTO: %+v\n", metadataDTO)
}
```

### Converting Status DTOs to Entities

The `ConvertStatusDTOToEntity` function converts a `StatusDTO` to a `Status` entity.

#### Example

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/entities/input-broker/entity"
    shareddto "libs/golang/ddd/dtos/input-broker/shared"
    "libs/golang/ddd/domain/converter"
)

func main() {
    statusDTO := shareddto.StatusDTO{
        Code:   "200",
        Detail: "OK",
    }

    status := converter.ConvertStatusDTOToEntity(statusDTO)
    fmt.Printf("Converted entity: %+v\n", status)
}
```

### Converting Status Entities to DTOs

The `ConvertStatusEntityToDTO` function converts a `Status` entity to a `StatusDTO`.

#### Example

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/entities/input-broker/entity"
    shareddto "libs/golang/ddd/dtos/input-broker/shared"
    "libs/golang/ddd/domain/converter"
)

func main() {
    status := entity.Status{
        Code:   "200",
        Detail: "OK",
    }

    statusDTO := converter.ConvertStatusEntityToDTO(status)
    fmt.Printf("Converted DTO: %+v\n", statusDTO)
}
```

## Testing

To run the tests for the `converter` package, use the following command:

```sh
npx nx test libs-golang-ddd-shared-type-tools-custom-types-converter-input-broker
```
