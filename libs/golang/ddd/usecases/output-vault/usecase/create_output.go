package usecase

import (
	"libs/golang/ddd/domain/entities/output-vault/entity"
	inputdto "libs/golang/ddd/dtos/output-vault/input"
	outputdto "libs/golang/ddd/dtos/output-vault/output"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/output-vault/converter"
)

// CreateOutputUseCase is the use case for creating a new output.
type CreateOutputUseCase struct {
	OutputRepository entity.OutputRepositoryInterface
}

func NewCreateOutputUseCase(
	outputRepository entity.OutputRepositoryInterface,
) *CreateOutputUseCase {
	return &CreateOutputUseCase{
		OutputRepository: outputRepository,
	}
}

// Execute creates a new output entity based on the provided input DTO and saves it using the repository.
// It then converts the created entity to an output DTO and returns it.
//
// Parameters:
//
//	input: The input DTO containing the output data.
//
// Returns:
//
//	An output DTO containing the created output data, and an error if any occurred during the process.
func (uc *CreateOutputUseCase) Execute(input inputdto.OutputDTO) (outputdto.OutputDTO, error) {
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

	err = uc.OutputRepository.Create(entityOutput)
	if err != nil {
		return outputdto.OutputDTO{}, err
	}

	return outputdto.OutputDTO{
		ID:        string(entityOutput.ID),
		Service:   entityOutput.Service,
		Source:    entityOutput.Source,
		Provider:  entityOutput.Provider,
		Data:      entityOutput.Data,
		Metadata:  converter.ConvertMetadataEntityToDTO(entityOutput.Metadata),
		CreatedAt: entityOutput.CreatedAt,
		UpdatedAt: entityOutput.UpdatedAt,
	}, nil
}
