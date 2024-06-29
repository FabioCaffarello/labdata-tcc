package outputdto

import (
	shareddto "libs/golang/ddd/dtos/input-broker/shared"
)

// InputDTO represents the input data transfer object.
type InputDTO struct {
	ID        string                 `json:"id"`         // ID represents the unique identifier of the input data.
	Data      map[string]interface{} `json:"data"`       // Data represents the input data.
	Metadata  shareddto.MetadataDTO  `json:"metadata"`   // Metadata represents the metadata of the input data.
	Status    shareddto.StatusDTO    `json:"status"`     // Status represents the status of the input data.
	CreatedAt string                 `json:"created_at"` // CreatedAt represents the timestamp when the input data was created.
	UpdatedAt string                 `json:"updated_at"` // UpdatedAt represents the timestamp when the input data was last updated.
}
