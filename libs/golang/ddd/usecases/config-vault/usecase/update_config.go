package usecase

import (
	"libs/golang/ddd/domain/entities/config-vault/entity"
	inputdto "libs/golang/ddd/dtos/config-vault/input"
	outputdto "libs/golang/ddd/dtos/config-vault/output"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/config-vault/converter"
)

// UpdateConfigUseCase is the use case for updating an existing configuration.
type UpdateConfigUseCase struct {
	ConfigRepository entity.ConfigRepositoryInterface
}

// NewUpdateConfigUseCase initializes a new instance of UpdateConfigUseCase with the provided ConfigRepositoryInterface.
//
// Parameters:
//
//	configRepository: The repository interface for managing Config entities.
//
// Returns:
//
//	A pointer to an instance of UpdateConfigUseCase.
func NewUpdateConfigUseCase(
	configRepository entity.ConfigRepositoryInterface,
) *UpdateConfigUseCase {
	return &UpdateConfigUseCase{
		ConfigRepository: configRepository,
	}
}

// Execute updates an existing configuration entity based on the provided input DTO and saves it using the repository.
// It then converts the updated entity to an output DTO and returns it.
//
// Parameters:
//
//	input: The input DTO containing the configuration data.
//
// Returns:
//
//	An output DTO containing the updated configuration data, and an error if any occurred during the process.
func (uc *UpdateConfigUseCase) Execute(input inputdto.ConfigDTO) (outputdto.ConfigDTO, error) {
	configProps := entity.ConfigProps{
		Active:        input.Active,
		Service:       input.Service,
		Source:        input.Source,
		Provider:      input.Provider,
		JobParameters: converter.ConvertJobParametersDTOToMap(input.JobParameters),
		DependsOn:     converter.ConvertJobDependenciesDTOToMap(input.DependsOn),
	}

	entityConfig, err := entity.NewConfig(configProps)
	if err != nil {
		return outputdto.ConfigDTO{}, err
	}

	err = uc.ConfigRepository.Update(entityConfig)
	if err != nil {
		return outputdto.ConfigDTO{}, err
	}

	return outputdto.ConfigDTO{
		ID:              string(entityConfig.ID),
		Active:          entityConfig.Active,
		Service:         entityConfig.Service,
		Source:          entityConfig.Source,
		Provider:        entityConfig.Provider,
		ConfigVersionID: string(entityConfig.ConfigVersionID),
		JobParameters:   converter.ConvertJobParametersEntityToDTO(entityConfig.JobParameters),
		DependsOn:       converter.ConvertJobDependenciesEntityToDTO(entityConfig.DependsOn),
		CreatedAt:       entityConfig.CreatedAt,
		UpdatedAt:       entityConfig.UpdatedAt,
	}, nil
}
