package usecase

import (
	"libs/golang/ddd/domain/entities/input-broker/entity"
)

// DeleteInputUseCase is the use case for deleting an existing input.
type DeleteInputUseCase struct {
	InputRepository entity.InputRepositoryInterface
}

// NewDeleteInputUseCase initializes a new instance of DeleteInputUseCase with the provided InputRepositoryInterface.
//
// Parameters:
//
//	inputRepository: The repository interface for managing Input entities.
//
// Returns:
//
//	A pointer to an instance of DeleteInputUseCase.
func NewDeleteInputUseCase(
	inputRepository entity.InputRepositoryInterface,
) *DeleteInputUseCase {
	return &DeleteInputUseCase{
		InputRepository: inputRepository,
	}
}

// Execute deletes an existing input entity based on the provided ID.
//
// Parameters:
//
//	id: The ID of the input to be deleted.
//
// Returns:
//
//	An error if any occurred during the process.
func (uc *DeleteInputUseCase) Execute(id string) error {
	err := uc.InputRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
