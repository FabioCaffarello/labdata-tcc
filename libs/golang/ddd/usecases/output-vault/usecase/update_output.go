package usecase

import (
	"libs/golang/ddd/domain/entities/output-vault/entity"
	inputdto "libs/golang/ddd/dtos/output-vault/input"
	outputdto "libs/golang/ddd/dtos/output-vault/output"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/output-vault/converter"
)

// UpdateOutputUseCase is the use case for updating an existing output.
type UpdateOutputUseCase struct {
	OutputRepository entity.OutputRepositoryInterface
}

// NewUpdateOutputUseCase initializes a new instance of UpdateOutputUseCase with the provided OutputRepositoryInterface.
//
// Parameters:
//
//	outputRepository: The repository interface for managing Output entities.
//
// Returns:
//
//	A pointer to an instance of UpdateOutputUseCase.
func NewUpdateOutputUseCase(
	outputRepository entity.OutputRepositoryInterface,
) *UpdateOutputUseCase {
	return &UpdateOutputUseCase{
		OutputRepository: outputRepository,
	}
}

// Execute updates an existing output entity based on the provided input DTO and saves it using the repository.
// It then converts the updated entity to an output DTO and returns it.
//
// Parameters:
//
//	input: The input DTO containing the output data.
//
// Returns:
//
//	An output DTO containing the updated output data, and an error if any occurred during the process.
func (uc *UpdateOutputUseCase) Execute(input inputdto.OutputDTO) (outputdto.OutputDTO, error) {
	outputProps := entity.OutputProps{
		Service:  input.Service,
		Source:   input.Source,
		Provider: input.Provider,
		Data:     input.Data,
		Metadata: converter.ConvertMetadataDTOToMap(input.Metadata),
	}

	entityOutput, err := entity.NewOutput(outputProps)
	if err != nil {
		return outputdto.OutputDTO{}, err
	}

	err = uc.OutputRepository.Update(entityOutput)
	if err != nil {
		return outputdto.OutputDTO{}, err
	}

	dtoMetadata := converter.ConvertMetadataEntityToDTO(entityOutput.Metadata)

	return outputdto.OutputDTO{
		ID:        string(entityOutput.ID),
		Service:   entityOutput.Service,
		Source:    entityOutput.Source,
		Provider:  entityOutput.Provider,
		Data:      entityOutput.Data,
		Metadata:  dtoMetadata,
		CreatedAt: entityOutput.CreatedAt,
		UpdatedAt: entityOutput.UpdatedAt,
	}, nil
}
