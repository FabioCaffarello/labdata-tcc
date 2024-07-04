package usecase

import (
	"libs/golang/ddd/domain/entities/output-vault/entity"
)

// DeleteOutputUseCase is a use case for deleting a output.
type DeleteOutputUseCase struct {
	OutputRepository entity.OutputRepositoryInterface
}

// NewDeleteOutputUseCase initializes a new instance of DeleteOutputUseCase with the provided OutputRepositoryInterface.
//
// Parameters:
//
//	outputRepository: The repository interface for managing Output entities.
//
// Returns:
//
//	A pointer to an instance of DeleteOutputUseCase.
func NewDeleteOutputUseCase(
	outputRepository entity.OutputRepositoryInterface,
) *DeleteOutputUseCase {
	return &DeleteOutputUseCase{
		OutputRepository: outputRepository,
	}
}

// Execute deletes a output entity based on the provided ID.
//
// Parameters:
//
//	id: The ID of the output to be deleted.
//
// Returns:
//
//	An error if any occurred during the process.
func (uc *DeleteOutputUseCase) Execute(id string) error {
	err := uc.OutputRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
