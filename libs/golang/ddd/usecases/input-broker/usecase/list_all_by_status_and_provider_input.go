package usecase

import (
	"libs/golang/ddd/domain/entities/input-broker/entity"
	outputdto "libs/golang/ddd/dtos/input-broker/output"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/input-broker/converter"
)

// ListAllByStatusAndProviderInputUseCase is the use case for listing all inputs.
type ListAllByStatusAndProviderInputUseCase struct {
	InputRepository entity.InputRepositoryInterface
}

// NewListAllByStatusAndProviderInputUseCase initializes a new instance of ListAllByStatusAndProviderInputUseCase with the provided InputRepositoryInterface.
//
// Parameters:
//
//	inputRepository: The repository interface for managing Input entities.
//
// Returns:
//
//	A pointer to an instance of ListAllByStatusAndProviderInputUseCase.
func NewListAllByStatusAndProviderInputUseCase(
	inputRepository entity.InputRepositoryInterface,
) *ListAllByStatusAndProviderInputUseCase {
	return &ListAllByStatusAndProviderInputUseCase{
		InputRepository: inputRepository,
	}
}

// Execute retrieves all inputs from the repository and converts them to output DTOs.
//
// Parameters:
//
//	provider: The provider name to filter inputs by.
//	status: The status to filter inputs by.
//
// Returns:
//
//	A slice of output DTOs containing the input data, and an error if any occurred during the process.
func (uc *ListAllByStatusAndProviderInputUseCase) Execute(provider string, status int) ([]outputdto.InputDTO, error) {
	inputs, err := uc.InputRepository.FindAllByStatusAndProvider(provider, status)
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
