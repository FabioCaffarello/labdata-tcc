package usecase

import (
	"libs/golang/ddd/domain/entities/config-vault/entity"
	outputdto "libs/golang/ddd/dtos/config-vault/output"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/config-vault/converter"
)

// ListAllConfigUseCase is the use case for listing all configurations.
type ListAllConfigUseCase struct {
	ConfigRepository entity.ConfigRepositoryInterface
}

// NewListAllConfigUseCase initializes a new instance of ListAllConfigUseCase with the provided ConfigRepositoryInterface.
//
// Parameters:
//
//	configRepository: The repository interface for managing Config entities.
//
// Returns:
//
//	A pointer to an instance of ListAllConfigUseCase.
func NewListAllConfigUseCase(
	configRepository entity.ConfigRepositoryInterface,
) *ListAllConfigUseCase {
	return &ListAllConfigUseCase{
		ConfigRepository: configRepository,
	}
}

// Execute retrieves all configurations from the repository and converts them to output DTOs.
//
// Returns:
//
//	A slice of output DTOs containing the configuration data, and an error if any occurred during the process.
func (uc *ListAllConfigUseCase) Execute() ([]outputdto.ConfigDTO, error) {
	configs, err := uc.ConfigRepository.FindAll()
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
