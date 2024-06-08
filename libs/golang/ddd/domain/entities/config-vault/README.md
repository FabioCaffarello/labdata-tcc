# config-vault

`config-vault` is a Go library that provides structures and functions to manage and manipulate configuration entities within a system. This library includes utilities for converting data between different formats, validating configuration data, and generating necessary identifiers.

## Features

- Define and manage configuration entities.
- Convert between `map[string]interface{}` and entity structs.
- Validate configuration data.
- Generate and handle MD5 and UUID identifiers.

## Usage

### Defining Configuration Entities

The `Config` struct represents a configuration entity with attributes such as service, source, provider, and dependencies.

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/entities/config-vault/entity"
)

func main() {
    configProps := entity.ConfigProps{
        Active:   true,
        Service:  "exampleService",
        Source:   "exampleSource",
        Provider: "exampleProvider",
        DependsOn: []map[string]interface{}{
            {"service": "dependencyService", "source": "dependencySource"},
        },
        UpdatedAt: "2023-06-08 12:00:00",
    }

    config, err := entity.NewConfig(configProps)
    if err != nil {
        fmt.Println("Error creating config:", err)
        return
    }

    fmt.Printf("Config: %+v\n", config)
}
```

### Converting Configuration Entities to Maps

The `ToMap` method converts a `Config` entity to a `map[string]interface{}` representation.

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/entities/config-vault/entity"
)

func main() {
    configProps := entity.ConfigProps{
        Active:   true,
        Service:  "exampleService",
        Source:   "exampleSource",
        Provider: "exampleProvider",
        DependsOn: []map[string]interface{}{
            {"service": "dependencyService", "source": "dependencySource"},
        },
        UpdatedAt: "2023-06-08 12:00:00",
    }

    config, err := entity.NewConfig(configProps)
    if err != nil {
        fmt.Println("Error creating config:", err)
        return
    }

    configMap, err := config.ToMap()
    if err != nil {
        fmt.Println("Error converting config to map:", err)
        return
    }

    fmt.Printf("Config as map: %+v\n", configMap)
}
```

### Validating Configuration Entities

The `isValid` method ensures that all required fields of a `Config` entity are set.

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/entities/config-vault/entity"
)

func main() {
    configProps := entity.ConfigProps{
        Active:   true,
        Service:  "exampleService",
        Source:   "exampleSource",
        Provider: "exampleProvider",
        DependsOn: []map[string]interface{}{
            {"service": "dependencyService", "source": "dependencySource"},
        },
        UpdatedAt: "2023-06-08 12:00:00",
    }

    config, err := entity.NewConfig(configProps)
    if err != nil {
        fmt.Println("Error creating config:", err)
        return
    }

    err = config.isValid()
    if err != nil {
        fmt.Println("Config is invalid:", err)
        return
    }

    fmt.Println("Config is valid")
}
```

## Testing

To run the tests for the `entity` package, use the following command:

```sh
npx nx test libs-golang-ddd-domain-entities-config-vault
```

## Errors

- `ErrInvalidID`: Returned when the ID of a `Config` is invalid.
- `ErrInvalidService`: Returned when the service of a `Config` is invalid.
- `ErrInvalidSource`: Returned when the source of a `Config` is invalid.
- `ErrInvalidProvider`: Returned when the provider of a `Config` is invalid.
- `ErrInvalidConfigVersionID`: Returned when the config version ID of a `Config` is invalid.
- `ErrInvalidCreatedAt`: Returned when the created at timestamp of a `Config` is invalid.
