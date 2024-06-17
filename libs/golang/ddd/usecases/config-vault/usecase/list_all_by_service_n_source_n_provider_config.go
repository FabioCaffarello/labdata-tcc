package usecase

import (
	"libs/golang/ddd/domain/entities/config-vault/entity"
	outputdto "libs/golang/ddd/dtos/config-vault/output"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/config-vault/converter"
)

// ListAllByServiceAndSourceAndProviderConfigUseCase is the use case for listing all configurations by service, source, and provider.
type ListAllByServiceAndSourceAndProviderConfigUseCase struct {
	ConfigRepository entity.ConfigRepositoryInterface
}

// NewListAllByServiceAndSourceAndProviderConfigUseCase initializes a new instance of ListAllByServiceAndSourceAndProviderConfigUseCase with the provided ConfigRepositoryInterface.
//
// Parameters:
//
//	configRepository: The repository interface for managing Config entities.
//
// Returns:
//
//	A pointer to an instance of ListAllByServiceAndSourceAndProviderConfigUseCase.
func NewListAllByServiceAndSourceAndProviderConfigUseCase(
	configRepository entity.ConfigRepositoryInterface,
) *ListAllByServiceAndSourceAndProviderConfigUseCase {
	return &ListAllByServiceAndSourceAndProviderConfigUseCase{
		ConfigRepository: configRepository,
	}
}

// Execute retrieves all configurations by service, source, and provider from the repository.
//
// Parameters:
//
//	service: The service name to filter configurations by.
//	source: The source name to filter configurations by.
//	provider: The provider name to filter configurations by.
//
// Returns:
//
//	A slice of output DTOs containing the configuration data, and an error if any occurred during the process.
func (uc *ListAllByServiceAndSourceAndProviderConfigUseCase) Execute(service, source, provider string) ([]outputdto.ConfigDTO, error) {
	configs, err := uc.ConfigRepository.FindAllByServiceAndSourceAndProvider(service, source, provider)
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
