package inputdto

import (
	shareddto "libs/golang/ddd/dtos/schema-vault/shared"
)

// SchemaDTO represents the data transfer object for schema input.
// It includes the necessary details required for creating or updating
// a schema, such as service details, source, provider, and JSON schema.
type SchemaDTO struct {
	Service    string                  `json:"service"`     // Service represents the name of the service for which the configuration is created.
	Source     string                  `json:"source"`      // Source indicates the origin or source of the configuration.
	Provider   string                  `json:"provider"`    // Provider specifies the provider of the configuration.
	SchemaType string                  `json:"schema_type"` // SchemaType specifies the type of schema.
	JsonSchema shareddto.JsonSchemaDTO `json:"json_schema"` // JsonSchemaDTO represents the JSON schema of the configuration.
}
