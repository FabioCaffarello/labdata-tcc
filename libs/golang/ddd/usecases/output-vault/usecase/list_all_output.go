package usecase

import (
	"libs/golang/ddd/domain/entities/output-vault/entity"
	outputdto "libs/golang/ddd/dtos/output-vault/output"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/output-vault/converter"
)

// ListAllOutputUseCase is the use case for listing all outputs.
type ListAllOutputUseCase struct {
	OutputRepository entity.OutputRepositoryInterface
}

// NewListAllOutputUseCase initializes a new instance of ListAllOutputUseCase with the provided OutputRepositoryInterface.
//
// Parameters:
//
//	outputRepository: The repository interface for managing Output entities.
//
// Returns:
//
//	A pointer to an instance of ListAllOutputUseCase.
func NewListAllOutputUseCase(
	outputRepository entity.OutputRepositoryInterface,
) *ListAllOutputUseCase {
	return &ListAllOutputUseCase{
		OutputRepository: outputRepository,
	}
}

// Execute retrieves all outputs from the repository and converts them to output DTOs.
//
// Returns:
//
//	A slice of output DTOs containing the output data, and an error if any occurred during the process.
func (uc *ListAllOutputUseCase) Execute() ([]outputdto.OutputDTO, error) {
	outputs, err := uc.OutputRepository.FindAll()
	if err != nil {
		return []outputdto.OutputDTO{}, err
	}

	outputDTOs := make([]outputdto.OutputDTO, 0, len(outputs))
	for _, output := range outputs {
		outputDTOs = append(outputDTOs, outputdto.OutputDTO{
			ID:        string(output.ID),
			Service:   output.Service,
			Source:    output.Source,
			Provider:  output.Provider,
			Data:      output.Data,
			Metadata:  converter.ConvertMetadataEntityToDTO(output.Metadata),
			CreatedAt: output.CreatedAt,
			UpdatedAt: output.UpdatedAt,
		})
	}

	return outputDTOs, nil
}
