package usecase

import (
	"libs/golang/ddd/domain/entities/schema-vault/entity"
)

// DeleteSchemaUseCase is a use case for deleting a schema.
type DeleteSchemaUseCase struct {
	SchemaRepository entity.SchemaRepositoryInterface
}

// NewDeleteSchemaUseCase initializes a new instance of DeleteSchemaUseCase with the provided SchemaRepositoryInterface.
//
// Parameters:
//
//	schemaRepository: The repository interface for managing Schema entities.
//
// Returns:
//
//	A pointer to an instance of DeleteSchemaUseCase.
func NewDeleteSchemaUseCase(
	schemaRepository entity.SchemaRepositoryInterface,
) *DeleteSchemaUseCase {
	return &DeleteSchemaUseCase{
		SchemaRepository: schemaRepository,
	}
}

// Execute deletes a schema entity based on the provided ID.
//
// Parameters:
//
//	id: The ID of the schema to be deleted.
//
// Returns:
//
//	An error if any occurred during the process.
func (uc *DeleteSchemaUseCase) Execute(id string) error {
	err := uc.SchemaRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
