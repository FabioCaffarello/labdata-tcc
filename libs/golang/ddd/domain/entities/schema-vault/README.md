# schema-vault/entity

`schema-vault/entity` is a Go library that provides structures and functions to manage and manipulate schema entities within a system. This library includes utilities for converting data between different formats, validating schema data, and generating necessary identifiers.

## Features

- Define and manage schema entities.
- Convert between `map[string]interface{}` and entity structs.
- Validate schema data.
- Generate and handle MD5 and UUID identifiers.

## Usage

### Defining Schema Entities

The `Schema` struct represents a schema entity with attributes such as service, source, provider, and schema type.

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/entities/schema-vault/entity"
)

func main() {
    schemaProps := entity.SchemaProps{
        Service:    "exampleService",
        Source:     "exampleSource",
        Provider:   "exampleProvider",
        SchemaType: "exampleSchemaType",
        JsonSchema: map[string]interface{}{
            "type": "object",
            "properties": map[string]interface{}{
                "field1": map[string]interface{}{
                    "type": "string",
                },
                "field2": map[string]interface{}{
                    "type": "string",
                },
            },
            "required": []interface{}{"field1"},
        },
    }

    schema, err := entity.NewSchema(schemaProps)
    if err != nil {
        fmt.Println("Error creating schema:", err)
        return
    }

    fmt.Printf("Schema: %+v\n", schema)
}
```

### Converting Schema Entities to Maps

The `ToMap` method converts a `Schema` entity to a `map[string]interface{}` representation.

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/entities/schema-vault/entity"
)

func main() {
    schemaProps := entity.SchemaProps{
        Service:    "exampleService",
        Source:     "exampleSource",
        Provider:   "exampleProvider",
        SchemaType: "exampleSchemaType",
        JsonSchema: map[string]interface{}{
            "type": "object",
            "properties": map[string]interface{}{
                "field1": map[string]interface{}{
                    "type": "string",
                },
                "field2": map[string]interface{}{
                    "type": "string",
                },
            },
            "required": []interface{}{"field1"},
        },
    }

    schema, err := entity.NewSchema(schemaProps)
    if err != nil {
        fmt.Println("Error creating schema:", err)
        return
    }

    schemaMap, err := schema.ToMap()
    if err != nil {
        fmt.Println("Error converting schema to map:", err)
        return
    }

    fmt.Printf("Schema as map: %+v\n", schemaMap)
}
```

### Validating Schema Entities

The `isValid` method ensures that all required fields of a `Schema` entity are set.

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/entities/schema-vault/entity"
)

func main() {
    schemaProps := entity.SchemaProps{
        Service:    "exampleService",
        Source:     "exampleSource",
        Provider:   "exampleProvider",
        SchemaType: "exampleSchemaType",
        JsonSchema: map[string]interface{}{
            "type": "object",
            "properties": map[string]interface{}{
                "field1": map[string]interface{}{
                    "type": "string",
                },
                "field2": map[string]interface{}{
                    "type": "string",
                },
            },
            "required": []interface{}{"field1"},
        },
    }

    schema, err := entity.NewSchema(schemaProps)
    if err != nil {
        fmt.Println("Error creating schema:", err)
        return
    }

    err = schema.isValid()
    if err != nil {
        fmt.Println("Schema is invalid:", err)
        return
    }

    fmt.Println("Schema is valid")
}
```

## Testing

To run the tests for the `entity` package, use the following command:

```sh
npx nx test libs-golang-ddd-domain-entities-schema-vault
```

## Errors

- `ErrMissingID`: Returned when the ID of a `Schema` is missing.
- `ErrMissingService`: Returned when the service of a `Schema` is missing.
- `ErrMissingSource`: Returned when the source of a `Schema` is missing.
- `ErrMissingProvider`: Returned when the provider of a `Schema` is missing.
- `ErrMissingSchemaType`: Returned when the schema type of a `Schema` is missing.
- `ErrJsonSchemaInvalid`: Returned when the JSON schema of a `Schema` is invalid.
