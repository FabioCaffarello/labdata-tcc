package inputdto

import (
	shareddto "libs/golang/ddd/dtos/config-vault/shared"
)

// ConfigDTO represents the data transfer object for configuration input.
// It includes the necessary details required for creating or updating
// a configuration, such as service details, source, provider, and dependencies.
type ConfigDTO struct {
	Active    bool                           `json:"active"`     // Active indicates whether the configuration should be activated.
	Service   string                         `json:"service"`    // Service represents the name of the service for which the configuration is created.
	Source    string                         `json:"source"`     // Source indicates the origin or source of the configuration.
	Provider  string                         `json:"provider"`   // Provider specifies the provider of the configuration.
	DependsOn []shareddto.JobDependenciesDTO `json:"depends_on"` // DependsOn lists the dependencies required for the configuration, represented by JobDependenciesDTO.
}
