package schematools

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type JsonSchemaValidatorSuite struct {
	suite.Suite
}

func TestJsonSchemaValidatorSuite(t *testing.T) {
	suite.Run(t, new(JsonSchemaValidatorSuite))
}

func (suite *JsonSchemaValidatorSuite) TestValidateJSONSchemaValidSchema() {
	jsonSchema := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"name": map[string]interface{}{
				"type": "string",
			},
		},
	}

	err := ValidateJSONSchema(jsonSchema)

	assert.NoError(suite.T(), err)
}

func (suite *JsonSchemaValidatorSuite) TestValidateJSONSchemaInvalidSchema() {
	invalidJsonSchema := map[string]interface{}{
		"type": "invalid_type",
	}

	err := ValidateJSONSchema(invalidJsonSchema)

	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "jsonSchema validation failed")
}
