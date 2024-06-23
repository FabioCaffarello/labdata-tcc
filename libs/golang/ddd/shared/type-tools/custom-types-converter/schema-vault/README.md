Here is the `README.md` file for the `schema-vault/converter` library:

```markdown
# schema-vault/converter

`schema-vault/converter` is a Go library that provides utility functions to convert between data transfer objects (DTOs) and entities within the schema vault domain. This library facilitates the transformation of JSON schema data structures between different layers of the application.

## Features

- Convert JSON schema from DTOs to entities.
- Convert JSON schema from entities to DTOs.
- Convert JSON schema from DTOs to a map.

## Usage

### Converting JSON Schema DTOs to Entities

The `ConvertJsonSchemaDTOToEntity` function converts a `JsonSchema` DTO to a `JsonSchema` entity.

#### Example

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/entities/schema-vault/entity"
    "libs/golang/ddd/dtos/schema-vault/shared"
    "libs/golang/ddd/domain/converter"
)

func main() {
    dtoSchema := shareddto.JsonSchema{
        Required: []string{"field1"},
        Properties: map[string]interface{}{
            "field1": map[string]interface{}{
                "type": "string",
            },
            "field2": map[string]interface{}{
                "type": "string",
            },
        },
        JsonType: "object",
    }

    entitySchema := converter.ConvertJsonSchemaDTOToEntity(dtoSchema)
    fmt.Printf("Converted entity: %+v\n", entitySchema)
}
```

### Converting JSON Schema Entities to DTOs

The `ConvertJsonSchemaEntityToDTO` function converts a `JsonSchema` entity to a `JsonSchema` DTO.

#### Example

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/entities/schema-vault/entity"
    "libs/golang/ddd/dtos/schema-vault/shared"
    "libs/golang/ddd/domain/converter"
)

func main() {
    entitySchema := entity.JsonSchema{
        Required: []string{"field1"},
        Properties: map[string]interface{}{
            "field1": map[string]interface{}{
                "type": "string",
            },
            "field2": map[string]interface{}{
                "type": "string",
            },
        },
        JsonType: "object",
    }

    dtoSchema := converter.ConvertJsonSchemaEntityToDTO(entitySchema)
    fmt.Printf("Converted DTO: %+v\n", dtoSchema)
}
```

### Converting JSON Schema DTOs to a Map

The `ConvertJsonSchemaDTOToMap` function converts a `JsonSchema` DTO to a map.

#### Example

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/entities/schema-vault/entity"
    "libs/golang/ddd/dtos/schema-vault/shared"
    "libs/golang/ddd/domain/converter"
)

func main() {
    dtoSchema := shareddto.JsonSchema{
        Required: []string{"field1"},
        Properties: map[string]interface{}{
            "field1": map[string]interface{}{
                "type": "string",
            },
            "field2": map[string]interface{}{
                "type": "string",
            },
        },
        JsonType: "object",
    }

    schemaMap := converter.ConvertJsonSchemaDTOToMap(dtoSchema)
    fmt.Printf("Converted map: %+v\n", schemaMap)
}
```

## Testing

To run the tests for the `converter` package, use the following command:

```sh
npx nx test libs-golang-ddd-shared-type-tools-custom-types-converter-schema-vault
```
