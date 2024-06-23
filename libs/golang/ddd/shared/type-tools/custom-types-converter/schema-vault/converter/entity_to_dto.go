package converter

import (
	"libs/golang/ddd/domain/entities/schema-vault/entity"
	shareddto "libs/golang/ddd/dtos/schema-vault/shared"
)

// ConvertJsonSchemaEntityToDTO converts a JsonSchema entity to a JsonSchema DTO.
// This function maps the fields of the JsonSchema entity to the corresponding JsonSchema DTO fields.
//
// Parameters:
//
//	jsonSchema: The entity.JsonSchema to be converted.
//
// Returns:
//
//	A shareddto.JsonSchema containing the converted data.
func ConvertJsonSchemaEntityToDTO(jsonSchema entity.JsonSchema) shareddto.JsonSchema {
	return shareddto.JsonSchema{
		Required:   jsonSchema.Required,
		Properties: jsonSchema.Properties,
		JsonType:   jsonSchema.JsonType,
	}
}
