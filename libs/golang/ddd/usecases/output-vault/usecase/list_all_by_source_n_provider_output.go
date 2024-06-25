package usecase

import (
	"libs/golang/ddd/domain/entities/output-vault/entity"
	outputdto "libs/golang/ddd/dtos/output-vault/output"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/output-vault/converter"
)

// ListAllBySourceAndProviderOutputUseCase is the use case for listing all outputs by source.
type ListAllBySourceAndProviderOutputUseCase struct {
	OutputRepository entity.OutputRepositoryInterface
}

// NewListAllBySourceAndProviderOutputUseCase initializes a new instance of ListAllBySourceAndProviderOutputUseCase with the provided OutputRepositoryInterface.
//
// Parameters:
//
//	outputRepository: The repository interface for managing Output entities.
//
// Returns:
//
//	A pointer to an instance of ListAllBySourceAndProviderOutputUseCase.
func NewListAllBySourceAndProviderOutputUseCase(
	outputRepository entity.OutputRepositoryInterface,
) *ListAllBySourceAndProviderOutputUseCase {
	return &ListAllBySourceAndProviderOutputUseCase{
		OutputRepository: outputRepository,
	}
}

// Execute retrieves all outputs by source from the repository.
//
// Parameters:
//
//	provider: The provider name to filter outputs by.
//	source: The source name to filter outputs by.
//
// Returns:
//
//	A slice of output DTOs containing the output data, and an error if any occurred during the process.
func (uc *ListAllBySourceAndProviderOutputUseCase) Execute(provider, source string) ([]outputdto.OutputDTO, error) {
	outputs, err := uc.OutputRepository.FindAllBySourceAndProvider(provider, source)
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
