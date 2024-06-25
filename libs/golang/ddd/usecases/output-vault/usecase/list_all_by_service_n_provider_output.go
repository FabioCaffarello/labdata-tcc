package usecase

import (
	"libs/golang/ddd/domain/entities/output-vault/entity"
	outputdto "libs/golang/ddd/dtos/output-vault/output"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/output-vault/converter"
)

// ListAllByServiceAndProviderOutputUseCase is the use case for listing all outputs by service and provider.
type ListAllByServiceAndProviderOutputUseCase struct {
	OutputRepository entity.OutputRepositoryInterface
}

// NewListAllByServiceAndProviderOutputUseCase initializes a new instance of ListAllByServiceAndProviderOutputUseCase with the provided OutputRepositoryInterface.
//
// Parameters:
//
//	outputRepository: The repository interface for managing Output entities.
//
// Returns:
//
//	A pointer to an instance of ListAllByServiceAndProviderOutputUseCase.
func NewListAllByServiceAndProviderOutputUseCase(
	outputRepository entity.OutputRepositoryInterface,
) *ListAllByServiceAndProviderOutputUseCase {
	return &ListAllByServiceAndProviderOutputUseCase{
		OutputRepository: outputRepository,
	}
}

// Execute retrieves all outputs by service and provider from the repository.
//
// Parameters:
//
//	provider: The provider name to filter outputs by.
//	service: The service name to filter outputs by.
//
// Returns:
//
//	A slice of output DTOs containing the output data, and an error if any occurred during the process.
func (uc *ListAllByServiceAndProviderOutputUseCase) Execute(provider, service string) ([]outputdto.OutputDTO, error) {
	outputs, err := uc.OutputRepository.FindAllByServiceAndProvider(provider, service)
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
