package usecase

import (
	"libs/golang/ddd/domain/entities/config-vault/entity"
	inputdto "libs/golang/ddd/dtos/config-vault/input"
	outputdto "libs/golang/ddd/dtos/config-vault/output"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/config-vault/converter"
	"log"
)

// CreateConfigUseCase is the use case for creating a new configuration.
type CreateConfigUseCase struct {
	ConfigRepository entity.ConfigRepositoryInterface
}

// NewCreateConfigUseCase initializes a new instance of CreateConfigUseCase with the provided ConfigRepositoryInterface.
//
// Parameters:
//
//	configRepository: The repository interface for managing Config entities.
//
// Returns:
//
//	A pointer to an instance of CreateConfigUseCase.
func NewCreateConfigUseCase(
	configRepository entity.ConfigRepositoryInterface,
) *CreateConfigUseCase {
	return &CreateConfigUseCase{
		ConfigRepository: configRepository,
	}
}

// Execute creates a new configuration entity based on the provided input DTO and saves it using the repository.
// It then converts the created entity to an output DTO and returns it.
//
// Parameters:
//
//	input: The input DTO containing the configuration data.
//
// Returns:
//
//	An output DTO containing the created configuration data, and an error if any occurred during the process.
func (uc *CreateConfigUseCase) Execute(input inputdto.ConfigDTO) (outputdto.ConfigDTO, error) {
	configProps := entity.ConfigProps{
		Active:        input.Active,
		Service:       input.Service,
		Source:        input.Source,
		Provider:      input.Provider,
		JobParameters: converter.ConvertJobParametersDTOToMap(input.JobParameters),
		DependsOn:     converter.ConvertJobDependenciesDTOToMap(input.DependsOn),
	}

	log.Printf("Creating new configuration with properties: %+v", configProps)

	entityConfig, err := entity.NewConfig(configProps)
	if err != nil {
		return outputdto.ConfigDTO{}, err
	}

	err = uc.ConfigRepository.Create(entityConfig)
	if err != nil {
		return outputdto.ConfigDTO{}, err
	}

	dtoDependsOn := converter.ConvertJobDependenciesEntityToDTO(entityConfig.DependsOn)
	jobParams := converter.ConvertJobParametersEntityToDTO(entityConfig.JobParameters)

	log.Printf("Job parameters: %+v", jobParams)

	dto := outputdto.ConfigDTO{
		ID:              string(entityConfig.ID),
		Active:          entityConfig.Active,
		Service:         entityConfig.Service,
		Source:          entityConfig.Source,
		Provider:        entityConfig.Provider,
		DependsOn:       dtoDependsOn,
		JobParameters:   jobParams,
		ConfigVersionID: string(entityConfig.ConfigVersionID),
		CreatedAt:       entityConfig.CreatedAt,
		UpdatedAt:       entityConfig.UpdatedAt,
	}

	log.Printf("Configuration DTO created successfully: %+v", dto)

	return dto, nil
}
