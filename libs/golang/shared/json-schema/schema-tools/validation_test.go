package schematools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateJSONSchema(t *testing.T) {
	validSchema := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"test": map[string]interface{}{
				"type": "string",
			},
		},
		"required": []string{
			"test",
		},
	}

	invalidSchema := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"test": map[string]interface{}{
				"type": "invalid_type",
			},
		},
		"required": []string{
			"test",
		},
	}

	err := ValidateJSONSchema(validSchema)
	assert.NoError(t, err, "Valid schema should not produce an error")

	err = ValidateJSONSchema(invalidSchema)
	assert.Error(t, err, "Invalid schema should produce an error")
	assert.Contains(t, err.Error(), "jsonSchema validation failed", "Error message should contain validation failure info")
}

func TestValidateJSONData(t *testing.T) {
	schema := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"name": map[string]interface{}{
				"type": "string",
			},
			"age": map[string]interface{}{
				"type": "integer",
			},
		},
		"required": []string{
			"name",
		},
	}

	validData := map[string]interface{}{
		"name": "John Doe",
		"age":  30,
	}

	invalidData := map[string]interface{}{
		"name": 123,      // Should be a string
		"age":  "thirty", // Should be an integer
	}

	missingRequiredData := map[string]interface{}{
		"age": 30,
	}

	err := ValidateJSONData(schema, validData)
	assert.NoError(t, err, "Valid data should not produce an error")

	err = ValidateJSONData(schema, invalidData)
	assert.Error(t, err, "Invalid data should produce an error")
	assert.Contains(t, err.Error(), "data validation failed", "Error message should contain validation failure info")

	err = ValidateJSONData(schema, missingRequiredData)
	assert.Error(t, err, "Data missing required fields should produce an error")
	assert.Contains(t, err.Error(), "data validation failed", "Error message should contain validation failure info")
}
