# schema-vault/dtos

`schema-vault/dtos` is a Go library that provides data transfer objects (DTOs) for the schema vault domain. These DTOs facilitate the communication between different layers of the application by providing a structured way to transfer schema-related data.

## Features

- Define DTOs for schema input.
- Define DTOs for schema output.
- Shared DTOs for common JSON schema representation.

## Usage

### Input DTO

The `SchemaDTO` in the `inputdto` package represents the data transfer object for schema input. It includes the necessary details required for creating or updating a schema, such as service details, source, provider, and JSON schema.

#### Example

```go
package main

import (
    "fmt"
    "libs/golang/ddd/dtos/schema-vault/inputdto"
    shareddto "libs/golang/ddd/dtos/schema-vault/shared"
)

func main() {
    schemaInput := inputdto.SchemaDTO{
        Service:    "exampleService",
        Source:     "exampleSource",
        Provider:   "exampleProvider",
        SchemaType: "exampleSchemaType",
        JsonSchema: shareddto.JsonSchema{
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
        },
    }

    fmt.Printf("Schema Input DTO: %+v\n", schemaInput)
}
```

### Output DTO

The `SchemaDTO` in the `outputdto` package represents the data transfer object for schema output. It includes the necessary details required for fetching or displaying a schema, such as service details, source, provider, and JSON schema, along with additional metadata like ID, version ID, and timestamps.

#### Example

```go
package main

import (
    "fmt"
    "libs/golang/ddd/dtos/schema-vault/outputdto"
    shareddto "libs/golang/ddd/dtos/schema-vault/shared"
)

func main() {
    schemaOutput := outputdto.SchemaDTO{
        ID:          "exampleID",
        Service:     "exampleService",
        Source:      "exampleSource",
        Provider:    "exampleProvider",
        SchemaType:  "exampleSchemaType",
        JsonSchema: shareddto.JsonSchema{
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
        },
        SchemaVersionID: "exampleVersionID",
        CreatedAt:       "2023-01-01T00:00:00Z",
        UpdatedAt:       "2023-01-02T00:00:00Z",
    }

    fmt.Printf("Schema Output DTO: %+v\n", schemaOutput)
}
```

### Shared DTO

The `JsonSchema` in the `shareddto` package represents a common JSON schema. It includes the required fields, properties, and type of the JSON schema.

#### Example

```go
package main

import (
    "fmt"
    shareddto "libs/golang/ddd/dtos/schema-vault/shared"
)

func main() {
    jsonSchema := shareddto.JsonSchema{
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

    fmt.Printf("JSON Schema DTO: %+v\n", jsonSchema)
}
```
