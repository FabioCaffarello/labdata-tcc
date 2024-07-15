package usecase

import (
	"libs/golang/ddd/domain/entities/config-vault/entity"
	outputdto "libs/golang/ddd/dtos/config-vault/output"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/config-vault/converter"
)

// ListAllByServiceAndProviderAndActiveConfigUseCase is the use case for listing all configurations by service, provider, and active status.
type ListAllByServiceAndProviderAndActiveConfigUseCase struct {
	ConfigRepository entity.ConfigRepositoryInterface
}

// NewListAllByServiceAndProviderAndActiveConfigUseCase initializes a new instance of ListAllByServiceAndProviderAndActiveConfigUseCase with the provided ConfigRepositoryInterface.
//
// Parameters:
//
//	configRepository: The repository interface for managing Config entities.
//
// Returns:
//
//	A pointer to an instance of ListAllByServiceAndProviderAndActiveConfigUseCase.
func NewListAllByServiceAndProviderAndActiveConfigUseCase(
	configRepository entity.ConfigRepositoryInterface,
) *ListAllByServiceAndProviderAndActiveConfigUseCase {
	return &ListAllByServiceAndProviderAndActiveConfigUseCase{
		ConfigRepository: configRepository,
	}
}

// Execute retrieves all configurations by service, provider, and active status from the repository.
//
// Parameters:
//
//	service: The service name to filter configurations by.
//	provider: The provider name to filter configurations by.
//	active: The active status to filter configurations by.
//
// Returns:
//
//	A slice of output DTOs containing the configuration data, and an error if any occurred during the process.
func (uc *ListAllByServiceAndProviderAndActiveConfigUseCase) Execute(service, provider string, active bool) ([]outputdto.ConfigDTO, error) {
	configs, err := uc.ConfigRepository.FindAllByServiceAndProviderAndActive(service, provider, active)
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
			JobParameters:   converter.ConvertJobParametersEntityToDTO(config.JobParameters),
			DependsOn:       converter.ConvertJobDependenciesEntityToDTO(config.DependsOn),
			CreatedAt:       config.CreatedAt,
			UpdatedAt:       config.UpdatedAt,
		})
	}

	return configDTOs, nil
}
