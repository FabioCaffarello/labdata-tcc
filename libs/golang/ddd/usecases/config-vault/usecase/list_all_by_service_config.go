package usecase

import (
	"libs/golang/ddd/domain/entities/config-vault/entity"
	outputdto "libs/golang/ddd/dtos/config-vault/output"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/config-vault/converter"
)

// ListAllByServiceConfigUseCase is the use case for listing all configurations by service.
type ListAllByServiceConfigUseCase struct {
	ConfigRepository entity.ConfigRepositoryInterface
}

// NewListAllByServiceConfigUseCase initializes a new instance of ListAllByServiceConfigUseCase with the provided ConfigRepositoryInterface.
//
// Parameters:
//
//	configRepository: The repository interface for managing Config entities.
//
// Returns:
//
//	A pointer to an instance of ListAllByServiceConfigUseCase.
func NewListAllByServiceConfigUseCase(
	configRepository entity.ConfigRepositoryInterface,
) *ListAllByServiceConfigUseCase {
	return &ListAllByServiceConfigUseCase{
		ConfigRepository: configRepository,
	}
}

// Execute retrieves all configurations by service from the repository.
//
// Parameters:
//
//	service: The service name to filter configurations by.
//
// Returns:
//
//	A slice of output DTOs containing the configuration data, and an error if any occurred during the process.
func (uc *ListAllByServiceConfigUseCase) Execute(service string) ([]outputdto.ConfigDTO, error) {
	configs, err := uc.ConfigRepository.FindAllByService(service)
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
