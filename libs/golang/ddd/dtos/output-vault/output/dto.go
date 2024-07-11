package outputdto

import (
	shareddto "libs/golang/ddd/dtos/output-vault/shared"
)

// OutputDTO represents the data transfer object for output output.
// It includes the necessary details required for creating or updating or listing
// an output, such as service details, source, provider and metadata.
type OutputDTO struct {
	ID        string                 `json:"_id"`        // ID is the unique identifier of the output.
	Data      map[string]interface{} `json:"data"`       // Data represents the output data.
	Service   string                 `json:"service"`    // Service represents the name of the service for which the output is created.
	Source    string                 `json:"source"`     // Source indicates the origin or source of the output.
	Provider  string                 `json:"provider"`   // Provider specifies the provider of the output.
	Metadata  shareddto.MetadataDTO  `json:"metadata"`   // Metadata represents the metadata of the output.
	CreatedAt string                 `json:"created_at"` // CreatedAt is the timestamp when the output was created.
	UpdatedAt string                 `json:"updated_at"` // UpdatedAt is the timestamp when the output was last updated.
}
