package outputdto

import (
	shareddto "libs/golang/ddd/dtos/config-vault/shared"
)

// ConfigDTO represents the data transfer object for configuration output.
// It contains detailed information about the configuration, including its
// status, service details, dependencies, and timestamps for creation and updates.
type ConfigDTO struct {
	ID              string                         `json:"id"`                // ID is the unique identifier of the configuration.
	Active          bool                           `json:"active"`            // Active indicates whether the configuration is currently active.
	Service         string                         `json:"service"`           // Service represents the name of the service associated with the configuration.
	Source          string                         `json:"source"`            // Source indicates the origin or source of the configuration.
	Provider        string                         `json:"provider"`          // Provider specifies the provider of the configuration.
	DependsOn       []shareddto.JobDependenciesDTO `json:"depends_on"`        // DependsOn lists the dependencies of the configuration, represented by JobDependenciesDTO.
	ConfigVersionID string                         `json:"config_version_id"` // ConfigVersionID is the identifier of the configuration version.
	CreatedAt       string                         `json:"created_at"`        // CreatedAt is the timestamp when the configuration was created.
	UpdatedAt       string                         `json:"updated_at"`        // UpdatedAt is the timestamp when the configuration was last updated.
}
