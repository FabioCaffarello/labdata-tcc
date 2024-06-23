package outputdto

import (
	shareddto "libs/golang/ddd/dtos/schema-vault/shared"
)

// SchemaDTO represents the data transfer object for schema output.
// It includes the necessary details required for fetching or displaying
// a schema, such as service details, source, provider, and JSON schema.
type SchemaDTO struct {
	ID              string               `json:"id"`                // ID is the unique identifier of the Schema entity.
	Service         string               `json:"service"`           // Service represents the name of the service for which the configuration is created.
	Source          string               `json:"source"`            // Source indicates the origin or source of the configuration.
	Provider        string               `json:"provider"`          // Provider specifies the provider of the configuration.
	SchemaType      string               `json:"schema_type"`       // SchemaType specifies the type of schema.
	JsonSchema      shareddto.JsonSchema `json:"json_schema"`       // JsonSchema represents the JSON schema of the configuration.
	SchemaVersionID string               `json:"schema_version_id"` // SchemaVersionID is the unique identifier of the schema version.
	CreatedAt       string               `json:"created_at"`        // CreatedAt is the timestamp when the Schema entity was created.
	UpdatedAt       string               `json:"updated_at"`        // UpdatedAt is the timestamp when the Schema entity was last updated.
}
