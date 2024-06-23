package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SchemaVaultConfigSuite struct {
	suite.Suite
}

func TestSchemaVaultConfigSuite(t *testing.T) {
	suite.Run(t, new(SchemaVaultConfigSuite))
}

func (suite *SchemaVaultConfigSuite) TestNewSchema() {
	jsonSchema := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"field1": map[string]interface{}{
				"type": "string",
			},
			"field2": map[string]interface{}{
				"type": "string",
			},
		},
		"required": []interface{}{
			"field1",
		},
	}

	schemaProps := SchemaProps{
		Service:    "test-service",
		Source:     "test-source",
		Provider:   "test-provider",
		SchemaType: "test-schema-type",
		JsonSchema: jsonSchema,
	}

	schema, err := NewSchema(schemaProps)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), schema)

	assert.Equal(suite.T(), schemaProps.Service, schema.Service)
	assert.Equal(suite.T(), schemaProps.Source, schema.Source)
	assert.Equal(suite.T(), schemaProps.Provider, schema.Provider)
	assert.Equal(suite.T(), schemaProps.SchemaType, schema.SchemaType)

	expectedJsonSchema := JsonSchema{
		Required: []string{"field1"},
		Properties: map[string]interface{}{
			"field1": map[string]interface{}{
				"type": "string",
			},
			"field2": map[string]interface{}{
				"type": "string",
			},
		},
		JsonType: "object",
	}

	assert.Equal(suite.T(), expectedJsonSchema, schema.JsonSchema)
	assert.NotEmpty(suite.T(), schema.ID)
	assert.NotEmpty(suite.T(), schema.SchemaVersionID)
	assert.NotEmpty(suite.T(), schema.CreatedAt)
	assert.NotEmpty(suite.T(), schema.UpdatedAt)
}

func (suite *SchemaVaultConfigSuite) TestIsSchemaValidWhenEmptyService() {
	jsonSchema := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"field1": map[string]interface{}{
				"type": "string",
			},
			"field2": map[string]interface{}{
				"type": "string",
			},
		},
		"required": []string{
			"field1",
		},
	}

	schemaProps := SchemaProps{
		Service:    "",
		Source:     "test-source",
		Provider:   "test-provider",
		SchemaType: "test-schema-type",
		JsonSchema: jsonSchema,
	}

	schema, err := NewSchema(schemaProps)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), schema)
	assert.Contains(suite.T(), err.Error(), "invalid service")
}

func (suite *SchemaVaultConfigSuite) TestIsSchemaValidWhenEmptySource() {
	jsonSchema := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"field1": map[string]interface{}{
				"type": "string",
			},
			"field2": map[string]interface{}{
				"type": "string",
			},
		},
		"required": []string{
			"field1",
		},
	}

	schemaProps := SchemaProps{
		Service:    "test-service",
		Source:     "",
		Provider:   "test-provider",
		SchemaType: "test-schema-type",
		JsonSchema: jsonSchema,
	}

	schema, err := NewSchema(schemaProps)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), schema)
	assert.Contains(suite.T(), err.Error(), "invalid source")
}

func (suite *SchemaVaultConfigSuite) TestIsSchemaValidWhenEmptyProvider() {
	jsonSchema := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"field1": map[string]interface{}{
				"type": "string",
			},
			"field2": map[string]interface{}{
				"type": "string",
			},
		},
		"required": []string{
			"field1",
		},
	}

	schemaProps := SchemaProps{
		Service:    "test-service",
		Source:     "test-source",
		Provider:   "",
		SchemaType: "test-schema-type",
		JsonSchema: jsonSchema,
	}

	schema, err := NewSchema(schemaProps)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), schema)
	assert.Contains(suite.T(), err.Error(), "invalid provider")
}

func (suite *SchemaVaultConfigSuite) TestIsSchemaValidWhenEmptySchemaType() {
	jsonSchema := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"field1": map[string]interface{}{
				"type": "string",
			},
			"field2": map[string]interface{}{
				"type": "string",
			},
		},
		"required": []string{
			"field1",
		},
	}

	schemaProps := SchemaProps{
		Service:    "test-service",
		Source:     "test-source",
		Provider:   "test-provider",
		SchemaType: "",
		JsonSchema: jsonSchema,
	}

	schema, err := NewSchema(schemaProps)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), schema)
	assert.Contains(suite.T(), err.Error(), "invalid schema type")
}

func (suite *SchemaVaultConfigSuite) TestIsSchemaValidWhenEmptyJsonSchema() {
	schemaProps := SchemaProps{
		Service:    "test-service",
		Source:     "test-source",
		Provider:   "test-provider",
		SchemaType: "test-schema-type",
		JsonSchema: nil,
	}

	schema, err := NewSchema(schemaProps)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), schema)
	assert.Contains(suite.T(), err.Error(), "invalid JSON schema")
}

// Test cases for validating JSON schema
func (suite *SchemaVaultConfigSuite) TestIsSchemaValidWhenInvalidJsonSchemaWithEmptySchema() {
	schemaProps := SchemaProps{
		Service:    "test-service",
		Source:     "test-source",
		Provider:   "test-provider",
		SchemaType: "test-schema-type",
		JsonSchema: map[string]interface{}{},
	}

	schema, err := NewSchema(schemaProps)
	assert.Error(suite.T(), err, "An error is expected for an empty JSON schema")
	assert.Nil(suite.T(), schema, "Schema should be nil when JSON schema is invalid")
	if err != nil {
		assert.Contains(suite.T(), err.Error(), "invalid JSON schema", "Error message should contain validation failure info")
	}
}

func (suite *SchemaVaultConfigSuite) TestIsSchemaValidWhenInvalidJsonSchemaWithInvalidType() {
	jsonSchema := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"field1": map[string]interface{}{
				"type": "invalid_type",
			},
			"field2": map[string]interface{}{
				"type": "string",
			},
		},
		"required": []string{
			"field1",
		},
	}
	schemaProps := SchemaProps{
		Service:    "test-service",
		Source:     "test-source",
		Provider:   "test-provider",
		SchemaType: "test-schema-type",
		JsonSchema: jsonSchema,
	}

	schema, err := NewSchema(schemaProps)
	assert.Error(suite.T(), err, "An error is expected for an invalid JSON schema type")
	assert.Nil(suite.T(), schema, "Schema should be nil when JSON schema is invalid")
	if err != nil {
		assert.Contains(suite.T(), err.Error(), "invalid JSON schema", "Error message should contain validation failure info")
	}
}

func (suite *SchemaVaultConfigSuite) TestIsSchemaValidWhenInvalidJsonSchemaWithInvalidTypeAndMissingRequiredField() {
	jsonSchema := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"field1": map[string]interface{}{
				"type": "invalid_type",
			},
			"field2": map[string]interface{}{
				"type": "string",
			},
		},
		"required": []string{
			"field1",
			"field3",
		},
	}
	schemaProps := SchemaProps{
		Service:    "test-service",
		Source:     "test-source",
		Provider:   "test-provider",
		SchemaType: "test-schema-type",
		JsonSchema: jsonSchema,
	}

	schema, err := NewSchema(schemaProps)
	assert.Error(suite.T(), err, "An error is expected for an invalid JSON schema with invalid type and missing required field")
	assert.Nil(suite.T(), schema, "Schema should be nil when JSON schema is invalid")
	if err != nil {
		assert.Contains(suite.T(), err.Error(), "invalid JSON schema", "Error message should contain validation failure info")
	}
}

func (suite *SchemaVaultConfigSuite) TestToMap() {
	jsonSchema := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"field1": map[string]interface{}{
				"type": "string",
			},
			"field2": map[string]interface{}{
				"type": "string",
			},
		},
		"required": []interface{}{
			"field1",
		},
	}

	schemaProps := SchemaProps{
		Service:    "test-service",
		Source:     "test-source",
		Provider:   "test-provider",
		SchemaType: "test-schema-type",
		JsonSchema: jsonSchema,
	}

	schema, err := NewSchema(schemaProps)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), schema)

	schemaMap, err := schema.ToMap()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), string(schema.ID), schemaMap["_id"])
	assert.Equal(suite.T(), schema.Service, schemaMap["service"])
	assert.Equal(suite.T(), schema.Source, schemaMap["source"])
	assert.Equal(suite.T(), schema.Provider, schemaMap["provider"])
	assert.Equal(suite.T(), schema.SchemaType, schemaMap["schema_type"])

	expectedJsonSchema := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"field1": map[string]interface{}{
				"type": "string",
			},
			"field2": map[string]interface{}{
				"type": "string",
			},
		},
		"required": []string{
			"field1",
		},
	}

	assert.Equal(suite.T(), expectedJsonSchema, schemaMap["json_schema"])
	assert.Equal(suite.T(), string(schema.SchemaVersionID), schemaMap["schema_version_id"])
	assert.Equal(suite.T(), schema.CreatedAt, schemaMap["created_at"])
	assert.Equal(suite.T(), schema.UpdatedAt, schemaMap["updated_at"])
}

func (suite *SchemaVaultConfigSuite) TestToMapWhenEmptySchema() {
	schema := &Schema{}

	schemaMap, err := schema.ToMap()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), schemaMap)
}

func (suite *SchemaVaultConfigSuite) TestToMapWhenEmptyID() {
	schema := &Schema{
		Service:    "test-service",
		Source:     "test-source",
		Provider:   "test-provider",
		SchemaType: "test-schema-type",
		JsonSchema: JsonSchema{
			JsonType: "object",
			Properties: map[string]interface{}{
				"field1": map[string]interface{}{
					"type": "string",
				},
				"field2": map[string]interface{}{
					"type": "string",
				},
			},
			Required: []string{
				"field1",
			},
		},
		CreatedAt: "2024-06-22 15:04:05",
		UpdatedAt: "2024-06-22 15:04:05",
	}

	schemaMap, err := schema.ToMap()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), schemaMap)
}

func (suite *SchemaVaultConfigSuite) TestToMapWhenEmptyService() {
	schema := &Schema{
		ID:         "test-id",
		Source:     "test-source",
		Provider:   "test-provider",
		SchemaType: "test-schema-type",
		JsonSchema: JsonSchema{
			JsonType: "object",
			Properties: map[string]interface{}{
				"field1": map[string]interface{}{
					"type": "string",
				},
				"field2": map[string]interface{}{
					"type": "string",
				},
			},
			Required: []string{
				"field1",
			},
		},
		CreatedAt: "2024-06-22 15:04:05",
		UpdatedAt: "2024-06-22 15:04:05",
	}

	schemaMap, err := schema.ToMap()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), schemaMap)
}

func (suite *SchemaVaultConfigSuite) TestToMapWhenEmptySource() {
	schema := &Schema{
		ID:         "test-id",
		Service:    "test-service",
		Provider:   "test-provider",
		SchemaType: "test-schema-type",
		JsonSchema: JsonSchema{
			JsonType: "object",
			Properties: map[string]interface{}{
				"field1": map[string]interface{}{
					"type": "string",
				},
				"field2": map[string]interface{}{
					"type": "string",
				},
			},
			Required: []string{
				"field1",
			},
		},
		CreatedAt: "2024-06-22 15:04:05",
		UpdatedAt: "2024-06-22 15:04:05",
	}

	schemaMap, err := schema.ToMap()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), schemaMap)
}

func (suite *SchemaVaultConfigSuite) TestToMapWhenEmptyProvider() {
	schema := &Schema{
		ID:         "test-id",
		Service:    "test-service",
		Source:     "test-source",
		SchemaType: "test-schema-type",
		JsonSchema: JsonSchema{
			JsonType: "object",
			Properties: map[string]interface{}{
				"field1": map[string]interface{}{
					"type": "string",
				},
				"field2": map[string]interface{}{
					"type": "string",
				},
			},
			Required: []string{
				"field1",
			},
		},
		CreatedAt: "2024-06-22 15:04:05",
		UpdatedAt: "2024-06-22 15:04:05",
	}

	schemaMap, err := schema.ToMap()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), schemaMap)
}

func (suite *SchemaVaultConfigSuite) TestToMapWhenEmptySchemaType() {
	schema := &Schema{
		ID:       "test-id",
		Service:  "test-service",
		Source:   "test-source",
		Provider: "test-provider",
		JsonSchema: JsonSchema{
			JsonType: "object",
			Properties: map[string]interface{}{
				"field1": map[string]interface{}{
					"type": "string",
				},
				"field2": map[string]interface{}{
					"type": "string",
				},
			},
			Required: []string{
				"field1",
			},
		},
		CreatedAt: "2024-06-22 15:04:05",
		UpdatedAt: "2024-06-22 15:04:05",
	}

	schemaMap, err := schema.ToMap()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), schemaMap)
}

func (suite *SchemaVaultConfigSuite) TestToMapWhenEmptyJsonSchema() {
	schema := &Schema{
		ID:         "test-id",
		Service:    "test-service",
		Source:     "test-source",
		Provider:   "test-provider",
		SchemaType: "test-schema-type",
		CreatedAt:  "2024-06-22 15:04:05",
		UpdatedAt:  "2024-06-22 15:04:05",
	}

	schemaMap, err := schema.ToMap()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), schemaMap)
}
