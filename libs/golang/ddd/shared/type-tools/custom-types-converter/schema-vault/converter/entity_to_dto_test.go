package converter

import (
	"libs/golang/ddd/domain/entities/schema-vault/entity"
	shareddto "libs/golang/ddd/dtos/schema-vault/shared"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SchemaConverterEntityToDTOSuite struct {
	suite.Suite
}

func TestSchemaConverterEntityToDTOSuite(t *testing.T) {
	suite.Run(t, new(SchemaConverterEntityToDTOSuite))
}

func (suite *SchemaConverterEntityToDTOSuite) TestConvertJsonSchemaEntityToDTO() {
	jsonSchema := entity.JsonSchema{
		Required:   []string{"field1", "field2"},
		Properties: map[string]interface{}{"field1": "value1", "field2": "value2"},
		JsonType:   "object",
	}

	expected := shareddto.JsonSchemaDTO{
		Required:   jsonSchema.Required,
		Properties: jsonSchema.Properties,
		JsonType:   jsonSchema.JsonType,
	}

	dtoJsonSchema := ConvertJsonSchemaEntityToDTO(jsonSchema)

	suite.Equal(expected, dtoJsonSchema)
	suite.Equal(expected.Required, dtoJsonSchema.Required)
	suite.Equal(expected.Properties, dtoJsonSchema.Properties)
}
