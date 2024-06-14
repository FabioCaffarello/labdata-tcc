package usecase

import (
	"libs/golang/ddd/domain/entities/config-vault/entity"
	outputdto "libs/golang/ddd/dtos/config-vault/output"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/config-vault/converter"
)

// ListAllByServiceAndProviderConfigUseCase is the use case for listing all configurations by service.
type ListAllByServiceAndProviderConfigUseCase struct {
	ConfigRepository entity.ConfigRepositoryInterface
}

// NewListAllByServiceConfigUseCase initializes a new instance of ListAllByServiceAndProviderConfigUseCase with the provided ConfigRepositoryInterface.
//
// Parameters:
//
//	configRepository: The repository interface for managing Config entities.
//
// Returns:
//
//	A pointer to an instance of ListAllByServiceAndProviderConfigUseCase.
func NewListAllByServiceAndProviderConfigUseCase(
	configRepository entity.ConfigRepositoryInterface,
) *ListAllByServiceAndProviderConfigUseCase {
	return &ListAllByServiceAndProviderConfigUseCase{
		ConfigRepository: configRepository,
	}
}

// Execute retrieves all configurations by service from the repository.
//
// Parameters:
//
//	provider: The provider name to filter configurations by.
//	service: The service name to filter configurations by.
//
// Returns:
//
//	A slice of output DTOs containing the configuration data, and an error if any occurred during the process.
func (uc *ListAllByServiceAndProviderConfigUseCase) Execute(provider, service string) ([]outputdto.ConfigDTO, error) {
	configs, err := uc.ConfigRepository.FindAllByServiceAndProvider(provider, service)
	if err != nil {
		return []outputdto.ConfigDTO{}, err
	}

	configDTOs := make([]outputdto.ConfigDTO, 0, len(configs))
	for _, config := range configs {
		configDTOs = append(configDTOs, outputdto.ConfigDTO{
			ID:              string(config.ID),
			Active:          config.Active,
			Service:         config.Service,
			Source:          config.Source,
			Provider:        config.Provider,
			ConfigVersionID: string(config.ConfigVersionID),
			DependsOn:       converter.ConvertJobDependenciesEntityToDTO(config.DependsOn),
			CreatedAt:       config.CreatedAt,
			UpdatedAt:       config.UpdatedAt,
		})
	}

	return configDTOs, nil
}
