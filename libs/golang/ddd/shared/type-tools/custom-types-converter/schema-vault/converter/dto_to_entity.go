package converter

import (
	"libs/golang/ddd/domain/entities/schema-vault/entity"
	shareddto "libs/golang/ddd/dtos/schema-vault/shared"
)

// ConvertJsonSchemaDTOToEntity converts a JsonSchema DTO to a JsonSchema entity.
// This function maps the fields of the JsonSchema DTO to the corresponding JsonSchema entity fields.
//
// Parameters:
//
//	jsonSchemaDTO: The shareddto.JsonSchema to be converted.
//
// Returns:
//
//	An entity.JsonSchema containing the converted data.
func ConvertJsonSchemaDTOToEntity(jsonSchemaDTO shareddto.JsonSchema) entity.JsonSchema {
	return entity.JsonSchema{
		Required:   jsonSchemaDTO.Required,
		Properties: jsonSchemaDTO.Properties,
		JsonType:   jsonSchemaDTO.JsonType,
	}
}

// ConvertJsonSchemaDTOToMap converts a JsonSchema DTO to a map.
// This function maps the fields of the JsonSchema DTO to a map.
//
// Parameters:
//
//	jsonSchemaDTO: The shareddto.JsonSchema to be converted.
//
// Returns:
//
//	A map containing the converted data.
func ConvertJsonSchemaDTOToMap(jsonSchemaDTO shareddto.JsonSchema) map[string]interface{} {
	required := make([]interface{}, len(jsonSchemaDTO.Required))
	for i, v := range jsonSchemaDTO.Required {
		required[i] = v
	}
	return map[string]interface{}{
		"required":   required,
		"properties": jsonSchemaDTO.Properties,
		"type":       jsonSchemaDTO.JsonType,
	}
}
