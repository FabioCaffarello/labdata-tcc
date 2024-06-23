package converter

import (
	"libs/golang/ddd/domain/entities/schema-vault/entity"
	shareddto "libs/golang/ddd/dtos/schema-vault/shared"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SchemaConverterDTOToEntitySuite struct {
	suite.Suite
}

func TestConfigConverterDTOToEntitySuite(t *testing.T) {
	suite.Run(t, new(SchemaConverterDTOToEntitySuite))
}

func (suite *SchemaConverterDTOToEntitySuite) TestConvertJsonSchemaDTOToEntity() {
	jsonSchemaDTO := shareddto.JsonSchema{
		Required:   []string{"field1", "field2"},
		Properties: map[string]interface{}{"field1": "value1", "field2": "value2"},
		JsonType:   "object",
	}

	expected := entity.JsonSchema{
		Required:   jsonSchemaDTO.Required,
		Properties: jsonSchemaDTO.Properties,
		JsonType:   jsonSchemaDTO.JsonType,
	}

	entityJsonSchema := ConvertJsonSchemaDTOToEntity(jsonSchemaDTO)

	assert.Equal(suite.T(), expected, entityJsonSchema)
	assert.Equal(suite.T(), expected.Required, entityJsonSchema.Required)
	assert.Equal(suite.T(), expected.Properties, entityJsonSchema.Properties)
}

func (s *SchemaConverterDTOToEntitySuite) TestConvertJsonSchemaDTOToMap() {
	jsonSchemaDTO := shareddto.JsonSchema{
		Required:   []string{"field1", "field2"},
		Properties: map[string]interface{}{"field1": "value1", "field2": "value2"},
		JsonType:   "object",
	}

	required := make([]interface{}, len(jsonSchemaDTO.Required))
	for i, v := range jsonSchemaDTO.Required {
		required[i] = v
	}

	expected := map[string]interface{}{
		"required":   required,
		"properties": jsonSchemaDTO.Properties,
		"type":  jsonSchemaDTO.JsonType,
	}

	jsonSchemaMap := ConvertJsonSchemaDTOToMap(jsonSchemaDTO)

	assert.Equal(s.T(), expected, jsonSchemaMap)
	assert.Equal(s.T(), expected["required"], jsonSchemaMap["required"])
	assert.Equal(s.T(), expected["properties"], jsonSchemaMap["properties"])
}
