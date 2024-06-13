# config-vault/dto

`config-vault/dto` is a Go library that provides data transfer objects (DTOs) for managing configuration data within a system. This library includes structures for input and output DTOs, facilitating the transfer and validation of configuration data.

## Features

- Define DTOs for configuration input and output.
- Facilitate data transfer between different components of the system.
- Ensure consistency and validation of configuration data.

## Usage

### Defining Configuration DTOs

The `ConfigDTO` struct represents a configuration data transfer object with attributes such as service, source, provider, and dependencies.

#### Output DTO

The `outputdto.ConfigDTO` struct is used to represent configuration data for output purposes.

```go
package main

import (
    "fmt"
    "libs/golang/ddd/dtos/config-vault/outputdto"
    shareddto "libs/golang/ddd/dtos/config-vault/shared"
)

func main() {
    config := outputdto.ConfigDTO{
        ID:              "12345",
        Active:          true,
        Service:         "exampleService",
        Source:          "exampleSource",
        Provider:        "exampleProvider",
        DependsOn:       []shareddto.JobDependenciesDTO{
            {Service: "dependencyService", Source: "dependencySource"},
        },
        ConfigVersionID: "v1",
        CreatedAt:       "2023-06-08 12:00:00",
        UpdatedAt:       "2023-06-09 12:00:00",
    }

    fmt.Printf("Output ConfigDTO: %+v\n", config)
}
```

#### Input DTO

The `inputdto.ConfigDTO` struct is used to represent configuration data for input purposes.

```go
package main

import (
    "fmt"
    "libs/golang/ddd/dtos/config-vault/inputdto"
    shareddto "libs/golang/ddd/dtos/config-vault/shared"
)

func main() {
    config := inputdto.ConfigDTO{
        Active:    true,
        Service:   "exampleService",
        Source:    "exampleSource",
        Provider:  "exampleProvider",
        DependsOn: []shareddto.JobDependenciesDTO{
            {Service: "dependencyService", Source: "dependencySource"},
        },
    }

    fmt.Printf("Input ConfigDTO: %+v\n", config)
}
```

### Shared DTO

The `shareddto.JobDependenciesDTO` struct is used to represent dependencies between services.

```go
package main

import (
    "fmt"
    shareddto "libs/golang/ddd/dtos/config-vault/shared"
)

func main() {
    dependency := shareddto.JobDependenciesDTO{
        Service: "exampleService",
        Source:  "exampleSource",
    }

    fmt.Printf("JobDependenciesDTO: %+v\n", dependency)
}
```
