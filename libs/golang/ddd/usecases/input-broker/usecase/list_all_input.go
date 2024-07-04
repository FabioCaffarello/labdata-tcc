package usecase

import (
	"libs/golang/ddd/domain/entities/input-broker/entity"
	outputdto "libs/golang/ddd/dtos/input-broker/output"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/input-broker/converter"
)

// ListAllInputUseCase is the use case for listing all inputs.
type ListAllInputUseCase struct {
	InputRepository entity.InputRepositoryInterface
}

// NewListAllInputUseCase initializes a new instance of ListAllInputUseCase with the provided InputRepositoryInterface.
//
// Parameters:
//
//	inputRepository: The repository interface for managing Input entities.
//
// Returns:
//
//	A pointer to an instance of ListAllInputUseCase.
func NewListAllInputUseCase(
	inputRepository entity.InputRepositoryInterface,
) *ListAllInputUseCase {
	return &ListAllInputUseCase{
		InputRepository: inputRepository,
	}
}

// Execute retrieves all inputs from the repository and converts them to output DTOs.
//
// Returns:
//
//	A slice of output DTOs containing the input data, and an error if any occurred during the process.
func (uc *ListAllInputUseCase) Execute() ([]outputdto.InputDTO, error) {
	inputs, err := uc.InputRepository.FindAll()
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
