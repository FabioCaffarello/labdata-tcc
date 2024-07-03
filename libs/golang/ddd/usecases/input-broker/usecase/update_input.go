package usecase

import (
	"libs/golang/ddd/domain/entities/input-broker/entity"
	inputdto "libs/golang/ddd/dtos/input-broker/input"
	outputdto "libs/golang/ddd/dtos/input-broker/output"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/input-broker/converter"
)

// UpdateInputUseCase is the use case for updating an existing input.
type UpdateInputUseCase struct {
	InputRepository entity.InputRepositoryInterface
}

// NewUpdateInputUseCase initializes a new instance of UpdateInputUseCase with the provided InputRepositoryInterface.
//
// Parameters:
//
//	inputRepository: The repository interface for managing Input entities.
//
// Returns:
//
//	A pointer to an instance of UpdateInputUseCase.
func NewUpdateInputUseCase(
	inputRepository entity.InputRepositoryInterface,
) *UpdateInputUseCase {
	return &UpdateInputUseCase{
		InputRepository: inputRepository,
	}
}

// Execute updates an existing input entity based on the provided input DTO and saves it using the repository.
// It then converts the updated entity to an output DTO and returns it.
//
// Parameters:
//
//	input: The input DTO containing the input data.
//
// Returns:
//
//	An output DTO containing the updated input data, and an error if any occurred during the process.
func (uc *UpdateInputUseCase) Execute(input inputdto.InputDTO) (outputdto.InputDTO, error) {
	inputProps := entity.InputProps{
		Provider: input.Provider,
		Service:  input.Service,
		Source:   input.Source,
		Data:     input.Data,
	}

	entityInput, err := entity.NewInput(inputProps)
	if err != nil {
		return outputdto.InputDTO{}, err
	}

	err = uc.InputRepository.Update(entityInput)
	if err != nil {
		return outputdto.InputDTO{}, err
	}

	return outputdto.InputDTO{
		ID:        string(entityInput.ID),
		Data:      input.Data,
		Metadata:  converter.ConvertMetadataEntityToDTO(entityInput.Metadata),
		Status:    converter.ConvertStatusEntityToDTO(entityInput.Status),
		CreatedAt: entityInput.CreatedAt,
		UpdatedAt: entityInput.UpdatedAt,
	}, nil
}
