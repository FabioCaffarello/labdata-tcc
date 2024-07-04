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
