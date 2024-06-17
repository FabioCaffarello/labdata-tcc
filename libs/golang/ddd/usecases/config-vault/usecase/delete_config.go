package usecase

import (
	"libs/golang/ddd/domain/entities/config-vault/entity"
)

// DeleteConfigUseCase is the use case for deleting an existing configuration.
type DeleteConfigUseCase struct {
	ConfigRepository entity.ConfigRepositoryInterface
}

// NewDeleteConfigUseCase initializes a new instance of DeleteConfigUseCase with the provided ConfigRepositoryInterface.
//
// Parameters:
//
//	configRepository: The repository interface for managing Config entities.
//
// Returns:
//
//	A pointer to an instance of DeleteConfigUseCase.
func NewDeleteConfigUseCase(
	configRepository entity.ConfigRepositoryInterface,
) *DeleteConfigUseCase {
	return &DeleteConfigUseCase{
		ConfigRepository: configRepository,
	}
}

// Execute deletes an existing configuration entity based on the provided ID.
//
// Parameters:
//
//	id: The ID of the configuration to be deleted.
//
// Returns:
//
//	An error if any occurred during the process.
func (uc *DeleteConfigUseCase) Execute(id string) error {
	err := uc.ConfigRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
