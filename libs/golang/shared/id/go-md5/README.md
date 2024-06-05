# go-md5

`go-md5` is a Go library that provides utilities for generating MD5-based IDs from various data types. This library leverages the `type-tools` library to convert different types of data into a string representation before generating the MD5 hash.

## Features

- Generate MD5-based IDs from various data types including:
  - `string`
  - `float64`
  - `int` and other integer types
  - `bool`
  - `map[string]interface{}`
  - `map[string]string`

## Usage

### Generating an ID from a String

```go
package main

import (
	"fmt"
	"libs/golang/shared/id/go-md5"
)

func main() {
	id := md5id.NewID("test")
	fmt.Println(id) // Output: 098f6bcd4621d373cade4e832627b4f6
}
```

### Generating an ID from a Float64

```go
package main

import (
	"fmt"
	"libs/golang/shared/id/go-md5"
)

func main() {
	id := md5id.NewID(123.456)
	fmt.Println(id) // Output: f6e809317508ea1fdcb5e6d878e166ef
}
```

### Generating an ID from a Map

#### Map of String to Interface

```go
package main

import (
	"fmt"
	"libs/golang/shared/id/go-md5"
)

func main() {
	data := map[string]interface{}{
		"foo": "bar",
		"baz": 123,
	}
	id := md5id.NewID(data)
	fmt.Println(id) // Output: 7cc94a32929de9da271e6f19ef1392d7
}
```

#### Map of String to String

```go
package main

import (
	"fmt"
	"libs/golang/shared/id/go-md5"
)

func main() {
	data := map[string]string{
		"foo": "bar",
		"baz": "qux",
	}
	id := md5id.NewID(data)
	fmt.Println(id) // Output: 491ebd8bf73d6a9b2fabf44575e98fbe
}
```

### Generating an ID from an Integer

```go
package main

import (
	"fmt"
	"libs/golang/shared/id/go-md5"
)

func main() {
	id := md5id.NewID(123)
	fmt.Println(id) // Output: 202cb962ac59075b964b07152d234b70
}
```

### Generating an ID from a Boolean

```go
package main

import (
	"fmt"
	"libs/golang/shared/id/go-md5"
)

func main() {
	id := md5id.NewID(true)
	fmt.Println(id) // Output: b326b5062b2f0e69046810717534cb09

	id = md5id.NewID(false)
	fmt.Println(id) // Output: 68934a3e9455fa72420237eb05902327
}
```

## Testing

To run the tests for the `md5id` package, use the following command:

```sh
npx nx test libs-golang-shared-id-go-md5
```
