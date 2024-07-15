package usecase

import (
	"libs/golang/ddd/domain/entities/config-vault/entity"
	outputdto "libs/golang/ddd/dtos/config-vault/output"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/config-vault/converter"
)

// ListOneByIDConfigUseCase is the use case for listing a single configuration by its ID.
type ListOneByIDConfigUseCase struct {
	ConfigRepository entity.ConfigRepositoryInterface
}

// NewListOneByIDConfigUseCase initializes a new instance of ListOneByIDConfigUseCase with the provided ConfigRepositoryInterface.
//
// Parameters:
//
//	configRepository: The repository interface for managing Config entities.
//
// Returns:
//
//	A pointer to an instance of ListOneByIDConfigUseCase.
func NewListOneByIDConfigUseCase(
	configRepository entity.ConfigRepositoryInterface,
) *ListOneByIDConfigUseCase {
	return &ListOneByIDConfigUseCase{
		ConfigRepository: configRepository,
	}
}

// Execute retrieves a configuration by its ID from the repository and converts it to an output DTO.
//
// Parameters:
//
//	id: The ID of the configuration to retrieve.
//
// Returns:
//
//	An output DTO containing the configuration data, and an error if any occurred during the process.
func (uc *ListOneByIDConfigUseCase) Execute(id string) (outputdto.ConfigDTO, error) {
	config, err := uc.ConfigRepository.FindByID(id)
	if err != nil {
		return outputdto.ConfigDTO{}, err
	}

	dto := outputdto.ConfigDTO{
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
	}

	return dto, nil
}
