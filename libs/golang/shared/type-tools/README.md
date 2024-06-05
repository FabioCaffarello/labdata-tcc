# type-tools

`type-tools` is a Go library that provides utilities for converting various data types (Go's built-in types). This library is designed to be simple and efficient, making it easy to handle different types of data in a consistent manner.

## Features

- Convert `map[string]interface{}` to a sorted string representation.
- Convert `map[string]string` to a sorted string representation.
- Convert `float64` to a string representation.
- Convert `interface{}` to specific Go built-in types.
- Generic function to convert any type to a string.


## Usage

### Convert `map[string]interface{}` to String

```go
package main

import (
	"fmt"
	"libs/golang/shared/type-tools"
)

func main() {
	data := map[string]interface{}{
		"foo": "bar",
		"baz": 123,
	}
	str := typetools.MapInterfaceToString(data)
	fmt.Println(str) // Output: baz:123;foo:bar;
}
```

### Convert `map[string]string` to String

```go
package main

import (
	"fmt"
	"libs/golang/shared/type-tools"
)

func main() {
	data := map[string]string{
		"foo": "bar",
		"baz": "qux",
	}
	str := typetools.MapStringToString(data)
	fmt.Println(str) // Output: baz:qux;foo:bar;
}
```

### Convert Various Types to String

```go
package main

import (
	"fmt"
	"libs/golang/shared/type-tools"
)

func main() {
	str := typetools.ToString("test")
	fmt.Println(str) // Output: test

	str = typetools.ToString(123.456)
	fmt.Println(str) // Output: 123.456000

	data := map[string]interface{}{
		"foo": "bar",
		"baz": 123,
	}
	str = typetools.ToString(data)
	fmt.Println(str) // Output: baz:123;foo:bar;

	dataString := map[string]string{
		"foo": "bar",
		"baz": "qux",
	}
	str = typetools.ToString(dataString)
	fmt.Println(str) // Output: baz:qux;foo:bar;
}
```

### Convert Interface to Specific Types
```go
package main

import (
	"fmt"
	"libs/golang/shared/type-tools"
)

func main() {
	var intf interface{} = "123"
	str, err := typetools.ToString(intf)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(str) // Output: 123
	}

	var num interface{} = 123.456
	floatStr, err := typetools.ToFloat64(num)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(floatStr) // Output: 123.456
	}
}
```

## Testing

To run the tests for `typetools` package, use the following command:

```sh
npx nx test libs-golang-shared-type-tools
```
