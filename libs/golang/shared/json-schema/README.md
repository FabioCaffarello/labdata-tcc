# json-schema

`json-schema` is a Go library designed to validate JSON schemas according to the JSON Schema Draft-07 specification. This library leverages the `gojsonschema` package to ensure that your JSON schemas are correctly formatted and adhere to the defined standards.

## Features

- Validate JSON Schema Draft-07 structures.
- Return detailed validation error messages.

## Usage

### Validate JSON Schema

The `ValidateJSONSchema` function validates a JSON schema to ensure it adheres to the JSON Schema Draft-07 specification. It takes a map representation of the JSON schema as input and returns an error if the schema is invalid. If the schema is valid, it returns `nil`.

Here's an example of how to use the `ValidateJSONSchema` function to validate a JSON schema.

```go
package main

import (
	"fmt"
	"github.com/yourusername/yourrepository/schematools"
)

func main() {
	schema := map[string]interface{}{
		"$schema": "http://json-schema.org/draft-07/schema#",
		"type":    "object",
		"properties": map[string]interface{}{
			"age": map[string]interface{}{
				"type": "integer",
			},
			"name": map[string]interface{}{
				"type": "string",
			},
		},
		"required": []string{}{"age", "name"},
	}

	err := schematools.ValidateJSONSchema(schema)
	if err != nil {
		fmt.Println("Validation error:", err)
	} else {
		fmt.Println("Schema is valid.")
	}
}
```

## Testing

To run the tests for the `schematools` package, use the following command:

```sh
npx nx test libs-golang-shared-json-schema 
```
