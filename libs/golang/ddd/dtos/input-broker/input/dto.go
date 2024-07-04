package inputdto

import (
	shareddto "libs/golang/ddd/dtos/input-broker/shared"
)

// InputDTO represents the input data transfer object.
type InputDTO struct {
	Provider string                 `json:"provider"` // Provider represents the provider of the input data.
	Service  string                 `json:"service"`  // Service represents the service of the input data.
	Source   string                 `json:"source"`   // Source represents the source of the input data.
	Data     map[string]interface{} `json:"data"`     // Data represents the input data.
}

// StatusDTO represents the status data transfer object.
type StatusDTO struct {
	shareddto.StatusDTO `json:"status"` // Status represents the status of the input data.
}
