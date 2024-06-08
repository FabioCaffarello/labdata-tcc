# regular-types

`regular-types` is a Go library that provides utilities for converting built-in types to various entity types. This library is designed to be flexible and easy to use, allowing you to convert data into entity types without directly importing those types.

## Features

- Convert `map[string]interface{}` to any struct type.
- Convert an array of `map[string]interface{}` to an array of the specified entity type.
- convert an entity type to a `map[string]interface{}`
- Decouples conversion logic from entity types.
- Makes code more modular and maintainable.

## Usage

### Convert `map[string]interface{}` to Entity

```go
package main

import (
	"fmt"
	"reflect"

	"libs/golang/ddd/entities/config-vault/entity"
	"libs/golang/ddd/entities/shared/type-tools/regular-types/conversion"
)

func main() {
    data := map[string]interface{}{
        "service": "example-service",
        "source":  "example-source",
    }

    entity, err := regulartypetool.ConvertFromMapStringToEntity(reflect.TypeOf(entity.JobDependencies{}), data)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    jobDependencies := entity.(entity.JobDependencies)
    fmt.Println("JobDependencies:", jobDependencies)
}
```

### Convert Array of `map[string]interface{}` to Array of Entities

```go
package main

import (
	"fmt"
	"reflect"

	"libs/golang/ddd/entities/config-vault/entity"
	"libs/golang/ddd/entities/shared/type-tools/regular-types/conversion"
)

func main() {
    dataArray := []map[string]interface{}{
        {
            "service": "example-service-1",
            "source":  "example-source-1",
        },
        {
            "service": "example-service-2",
            "source":  "example-source-2",
        },
    }

    entities, err := regulartypetool.ConvertFromArrayMapStringToEntities(reflect.TypeOf(entity.JobDependencies{}), dataArray)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    for _, e := range entities {
        entityType := e.(entity.JobDependencies)
        fmt.Println("JobDependencies:", entityType)
    }
}
```

### Convert Entity to `map[string]interface{}`

```go
package main

import (
	"fmt"

	"github.com/yourusername/yourrepository/regular-types"
)

type JobDependencies struct {
	Service string `bson:"service"`
	Source  string `bson:"source"`
}

func main() {
    jobDependencies := JobDependencies{
        Service: "example-service",
        Source:  "example-source",
    }

    data, err := regulartypetool.ConvertFromEntityToMapString(jobDependencies)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Data:", data)
}
```

## Testing

To run the tests for the `regular-types` package, use the following command:

```sh
npx nx test libs-golang-ddd-entities-shared/type-tools-regular-types
```
