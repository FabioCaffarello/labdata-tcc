package usecase

import (
	"libs/golang/ddd/domain/entities/input-broker/entity"
	outputdto "libs/golang/ddd/dtos/input-broker/output"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/input-broker/converter"
)

// ListAllByStatusAndSourceAndProviderInputUseCase is the use case for listing all inputs by source.
type ListAllByStatusAndSourceAndProviderInputUseCase struct {
	InputRepository entity.InputRepositoryInterface
}

// NewListAllByStatusAndSourceAndProviderInputUseCase initializes a new instance of ListAllByStatusAndSourceAndProviderInputUseCase with the provided InputRepositoryInterface.
//
// Parameters:
//
//	inputRepository: The repository interface for managing Input entities.
//
// Returns:
//
//	A pointer to an instance of ListAllByStatusAndSourceAndProviderInputUseCase.
func NewListAllByStatusAndSourceAndProviderInputUseCase(
	inputRepository entity.InputRepositoryInterface,
) *ListAllByStatusAndSourceAndProviderInputUseCase {
	return &ListAllByStatusAndSourceAndProviderInputUseCase{
		InputRepository: inputRepository,
	}
}

// Execute retrieves all inputs by source from the repository.
//
// Parameters:
//
//	provider: The provider name to filter inputs by.
//	source: The source name to filter inputs by.
//	status: The status code to filter inputs by.
//
// Returns:
//
//	A slice of output DTOs containing the input data, and an error if any occurred during the process.
func (uc *ListAllByStatusAndSourceAndProviderInputUseCase) Execute(provider, source string, status int) ([]outputdto.InputDTO, error) {
	inputs, err := uc.InputRepository.FindAllByStatusAndSourceAndProvider(provider, source, status)
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
