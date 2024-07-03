package usecase

import (
	"libs/golang/ddd/domain/entities/input-broker/entity"
	outputdto "libs/golang/ddd/dtos/input-broker/output"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/input-broker/converter"
)

// ListAllByServiceAndSourceAndProviderInputUseCase is the use case for listing all inputs by service, source, and provider.
type ListAllByServiceAndSourceAndProviderInputUseCase struct {
	InputRepository entity.InputRepositoryInterface
}

// NewListAllByServiceAndSourceAndProviderInputUseCase initializes a new instance of ListAllByServiceAndSourceAndProviderInputUseCase with the provided InputRepositoryInterface.
//
// Parameters:
//
//	inputRepository: The repository interface for managing Input entities.
//
// Returns:
//
//	A pointer to an instance of ListAllByServiceAndSourceAndProviderInputUseCase.
func NewListAllByServiceAndSourceAndProviderInputUseCase(
	inputRepository entity.InputRepositoryInterface,
) *ListAllByServiceAndSourceAndProviderInputUseCase {
	return &ListAllByServiceAndSourceAndProviderInputUseCase{
		InputRepository: inputRepository,
	}
}

// Execute retrieves all inputs by service, source, and provider from the repository.
//
// Parameters:
//
//	service: The service name to filter inputs by.
//	source: The source name to filter inputs by.
//	provider: The provider name to filter inputs by.
//
// Returns:
//
//	A slice of output DTOs containing the input data, and an error if any occurred during the process.
func (uc *ListAllByServiceAndSourceAndProviderInputUseCase) Execute(provider, service, source string) ([]outputdto.InputDTO, error) {
	inputs, err := uc.InputRepository.FindAllByServiceAndSourceAndProvider(provider, service, source)
	if err != nil {
		return []outputdto.InputDTO{}, err
	}

	inputDTOs := make([]outputdto.InputDTO, 0, len(inputs))
	for _, input := range inputs {
		inputDTOs = append(inputDTOs, outputdto.InputDTO{
			ID:        string(input.ID),
			Metadata:  converter.ConvertMetadataEntityToDTO(input.Metadata),
			Status:    converter.ConvertStatusEntityToDTO(input.Status),
			Data:      input.Data,
			CreatedAt: input.CreatedAt,
			UpdatedAt: input.UpdatedAt,
		})
	}

	return inputDTOs, nil
}
