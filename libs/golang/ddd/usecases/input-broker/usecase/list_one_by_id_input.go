package usecase

import (
	"libs/golang/ddd/domain/entities/input-broker/entity"
	outputdto "libs/golang/ddd/dtos/input-broker/output"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/input-broker/converter"
)

// ListOneByIDInputUseCase is the use case for listing a single input by its ID.
type ListOneByIDInputUseCase struct {
	InputRepository entity.InputRepositoryInterface
}

// NewListOneByIDInputUseCase initializes a new instance of ListOneByIDInputUseCase with the provided InputRepositoryInterface.
//
// Parameters:
//
//	inputRepository: The repository interface for managing Input entities.
//
// Returns:
//
//	A pointer to an instance of ListOneByIDInputUseCase.
func NewListOneByIDInputUseCase(
	inputRepository entity.InputRepositoryInterface,
) *ListOneByIDInputUseCase {
	return &ListOneByIDInputUseCase{
		InputRepository: inputRepository,
	}
}

// Execute retrieves a input by its ID from the repository and converts it to an output DTO.
//
// Parameters:
//
//	id: The ID of the inputuration to retrieve.
//
// Returns:
//
//	An output DTO containing the input data, and an error if any occurred during the process.
func (uc *ListOneByIDInputUseCase) Execute(id string) (outputdto.InputDTO, error) {
	input, err := uc.InputRepository.FindByID(id)
	if err != nil {
		return outputdto.InputDTO{}, err
	}

	dto := outputdto.InputDTO{
		ID:        string(input.ID),
		Metadata:  converter.ConvertMetadataEntityToDTO(input.Metadata),
		Status:    converter.ConvertStatusEntityToDTO(input.Status),
		Data:      input.Data,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	}

	return dto, nil
}
