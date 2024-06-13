package usecase

import (
	"libs/golang/ddd/domain/entities/config-vault/entity"
	outputdto "libs/golang/ddd/dtos/config-vault/output"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/config-vault/converter"
)

// ListAllBySourceConfigUseCase is the use case for listing all configurations by source.
type ListAllBySourceConfigUseCase struct {
	ConfigRepository entity.ConfigRepositoryInterface
}

// NewListAllBySourceConfigUseCase initializes a new instance of ListAllBySourceConfigUseCase with the provided ConfigRepositoryInterface.
//
// Parameters:
//
//	configRepository: The repository interface for managing Config entities.
//
// Returns:
//
//	A pointer to an instance of ListAllBySourceConfigUseCase.
func NewListAllBySourceConfigUseCase(
	configRepository entity.ConfigRepositoryInterface,
) *ListAllBySourceConfigUseCase {
	return &ListAllBySourceConfigUseCase{
		ConfigRepository: configRepository,
	}
}

// Execute retrieves all configurations by source from the repository.
//
// Parameters:
//
//	source: The source name to filter configurations by.
//
// Returns:
//
//	A slice of output DTOs containing the configuration data, and an error if any occurred during the process.
func (uc *ListAllBySourceConfigUseCase) Execute(source string) ([]outputdto.ConfigDTO, error) {
	configs, err := uc.ConfigRepository.FindAllBySource(source)
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
