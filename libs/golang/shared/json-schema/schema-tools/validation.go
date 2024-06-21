package schematools

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

const metaschemaURL = "http://json-schema.org/draft-07/schema#" // URL of the JSON Schema metaschema

// ValidateJSONSchema validates a JSON Schema to ensure it adheres to the JSON Schema Draft-07 specification.
// It takes a map representation of the JSON Schema as input and returns an error if the schema is invalid.
// If the schema is valid, it returns nil.
//
// Parameters:
// - jsonSchema: map[string]interface{}: The JSON Schema to be validated.
//
// Returns:
// - error: An error object if the JSON Schema is invalid, otherwise nil.
//
// Example:
//  schema := map[string]interface{}{
//      "$schema": "http://json-schema.org/draft-07/schema#",
//      "type": "object",
//      "properties": map[string]interface{}{
//          "name": map[string]interface{}{
//              "type": "string",
//          },
//      },
//  }
//  err := ValidateJSONSchema(schema)
//  if err != nil {
//      // Handle error
//  }
func ValidateJSONSchema(jsonSchema map[string]interface{}) error {
	// Convert the JSON schema map to a JSON string
	jsonSchemaBytes, err := json.Marshal(jsonSchema)
	if err != nil {
		return err
	}

	// Create a JSONLoader for the JSON schema
	schemaLoader := gojsonschema.NewStringLoader(string(jsonSchemaBytes))

	// Validate the JSON Schema structure using the gojsonschema library
	metaschemaLoader := gojsonschema.NewReferenceLoader(metaschemaURL)
	compileResult, err := gojsonschema.Validate(metaschemaLoader, schemaLoader)
	if err != nil {
		return err
	}

	if !compileResult.Valid() {
		validationErrors := compileResult.Errors()
		errorMessages := make([]string, len(validationErrors))
		for i, err := range validationErrors {
			errorMessages[i] = err.String()
		}
		return errors.New("jsonSchema validation failed: " + strings.Join(errorMessages, ", "))
	}

	return nil
}
