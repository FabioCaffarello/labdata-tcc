package usecase

import (
	"libs/golang/ddd/domain/entities/input-broker/entity"
	outputdto "libs/golang/ddd/dtos/input-broker/output"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/input-broker/converter"
)

// ListAllBySourceAndProviderInputUseCase is the use case for listing all inputs by source.
type ListAllBySourceAndProviderInputUseCase struct {
	InputRepository entity.InputRepositoryInterface
}

// NewListAllBySourceAndProviderInputUseCase initializes a new instance of ListAllBySourceAndProviderInputUseCase with the provided InputRepositoryInterface.
//
// Parameters:
//
//	inputRepository: The repository interface for managing Input entities.
//
// Returns:
//
//	A pointer to an instance of ListAllBySourceAndProviderInputUseCase.
func NewListAllBySourceAndProviderInputUseCase(
	inputRepository entity.InputRepositoryInterface,
) *ListAllBySourceAndProviderInputUseCase {
	return &ListAllBySourceAndProviderInputUseCase{
		InputRepository: inputRepository,
	}
}

// Execute retrieves all inputs by source from the repository.
//
// Parameters:
//
//	provider: The provider name to filter inputs by.
//	source: The source name to filter inputs by.
//
// Returns:
//
//	A slice of output DTOs containing the input data, and an error if any occurred during the process.
func (uc *ListAllBySourceAndProviderInputUseCase) Execute(provider, source string) ([]outputdto.InputDTO, error) {
	inputs, err := uc.InputRepository.FindAllBySourceAndProvider(provider, source)
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
