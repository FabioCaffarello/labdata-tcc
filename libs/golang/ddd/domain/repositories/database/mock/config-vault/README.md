# config-vault/mockrepository

`config-vault/mockrepository` is a Go library that provides mock implementations of the repository interfaces used in the `config-vault` domain. This library is primarily used for testing purposes, allowing you to simulate repository behaviors without needing a real database.

## Features

- Mock implementation of `ConfigRepositoryInterface`.
- Support for creating, finding, updating, and deleting configuration entities.
- Support for querying configurations based on various attributes.

## Usage

### Mocking a Config Repository

The `ConfigRepositoryMock` struct provides methods to simulate interactions with configuration entities.

#### Example

```go
package main

import (
    "fmt"
    "log"
    "testing"

    "libs/golang/ddd/domain/entities/config-vault/entity"
    "libs/golang/ddd/domain/repositories/mock/config-vault/repository"
    "github.com/stretchr/testify/mock"
)

func main() {
    repoMock := new(mockrepository.ConfigRepositoryMock)

    configProps := entity.ConfigProps{
        Active:    true,
        Service:   "exampleService",
        Source:    "exampleSource",
        Provider:  "exampleProvider",
        DependsOn: []map[string]interface{}{
            {"service": "dependencyService", "source": "dependencySource"},
        },
    }

    config, err := entity.NewConfig(configProps)
    if err != nil {
        log.Fatal(err)
    }

    repoMock.On("Create", config).Return(nil)

    err = repoMock.Create(config)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Mocked config creation: %+v\n", config)
    repoMock.AssertExpectations(&testing.T{})
}
```
