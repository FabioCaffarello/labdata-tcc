# go-uuid

`go-uuid` is a Go library that provides utilities for generating UUIDs from various data types. It includes functions for calculating SHA-256 hashes and generating UUIDs based on those hashes.

## Features

- Generate UUIDs from various data types including:
  - `map[string]interface{}`
  - Other data types using serialization to `[]byte`
- Calculate SHA-256 hashes of data

## Usage

### Generating a UUID from a Map

```go
package main

import (
    "fmt"
    "libs/golang/shared/id/go-uuid"
)

func main() {
    data := map[string]interface{}{
        "key1": "value1",
        "key2": 123,
    }
    id, err := gouuid.GenerateUUIDFromMap(data)
    if err != nil {
        fmt.Println("Error generating UUID:", err)
        return
    }
    fmt.Println("Generated UUID:", id)
}
```

### Generating a UUID from Other Data Types

```go
package main

import (
    "fmt"
    "libs/golang/shared/id/go-uuid"
)

func main() {
    data := "test"
    id, err := gouuid.GenerateUUID([]byte(data))
    if err != nil {
        fmt.Println("Error generating UUID:", err)
        return
    }
    fmt.Println("Generated UUID:", id)
}
```

## Testing

To run the tests for the `gouuid` package, use the following command:

```sh
npx nx test libs-golang-shared-id-go-uuid
```
