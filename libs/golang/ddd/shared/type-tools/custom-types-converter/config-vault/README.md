# config-vault/converter

`config-vault/converter` is a Go library that provides utility functions to convert between data transfer objects (DTOs) and entities within the configuration vault domain. This library facilitates the transformation of job dependency data structures between different layers of the application.

## Features

- Convert job dependencies from DTOs to entities.
- Convert job dependencies from entities to DTOs.
- Convert job dependencies from DTOs to a map.

## Usage

### Converting Job Dependencies DTOs to Entities

The `ConvertJobDependenciesDTOToEntity` function converts a slice of `JobDependenciesDTO` to a slice of `JobDependencies` entities.

#### Example

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/entities/config-vault/entity"
    "libs/golang/ddd/dtos/config-vault/shared"
    "libs/golang/ddd/domain/converter"
)

func main() {
    dtoDeps := []shareddto.JobDependenciesDTO{
        {
            Service: "service1",
            Source:  "source1",
        },
        {
            Service: "service2",
            Source:  "source2",
        },
    }

    entityDeps := converter.ConvertJobDependenciesDTOToEntity(dtoDeps)
    fmt.Printf("Converted entities: %+v\n", entityDeps)
}
```

### Converting Job Dependencies Entities to DTOs

The `ConvertJobDependenciesEntityToDTO` function converts a slice of `JobDependencies` entities to a slice of `JobDependenciesDTO`.

#### Example

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/entities/config-vault/entity"
    "libs/golang/ddd/dtos/config-vault/shared"
    "libs/golang/ddd/domain/converter"
)

func main() {
    entityDeps := []entity.JobDependencies{
        {
            Service: "service1",
            Source:  "source1",
        },
        {
            Service: "service2",
            Source:  "source2",
        },
    }

    dtoDeps := converter.ConvertJobDependenciesEntityToDTO(entityDeps)
    fmt.Printf("Converted DTOs: %+v\n", dtoDeps)
}
```


### Converting Job Dependencies DTOs to a Map

The `ConvertJobDependenciesDTOToMap` function converts a slice of `JobDependenciesDTO` to a map of `JobDependencies` entities.

#### Example

```go
package main

import (
    "fmt"
    "libs/golang/ddd/domain/entities/config-vault/entity"
    "libs/golang/ddd/dtos/config-vault/shared"
    "libs/golang/ddd/domain/converter"
)

func main() {
    dtoDeps := []shareddto.JobDependenciesDTO{
        {
            Service: "service1",
            Source:  "source1",
        },
        {
            Service: "service2",
            Source:  "source2",
        },
    }

    entityMap := converter.ConvertJobDependenciesDTOToMap(dtoDeps)
    fmt.Printf("Converted entity map: %+v\n", entityMap)
}
```


## Testing

To run the tests for the `converter` package, use the following command:

```sh
npx nx test libs-golang-ddd-shared-type-tools-custom-types-converter-config-vault
```
