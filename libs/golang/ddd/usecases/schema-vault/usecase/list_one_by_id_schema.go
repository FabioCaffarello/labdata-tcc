package usecase

import (
	"libs/golang/ddd/domain/entities/schema-vault/entity"
	outputdto "libs/golang/ddd/dtos/schema-vault/output"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/schema-vault/converter"
)

// ListOneByIDSchemaUseCase is the use case for listing a schema by ID.
type ListOneByIDSchemaUseCase struct {
	SchemaRepository entity.SchemaRepositoryInterface
}

// NewListOneByIDSchemaUseCase initializes a new instance of ListOneByIDSchemaUseCase with the provided SchemaRepositoryInterface.
//
// Parameters:
//
//	schemaRepository: The repository interface for managing Schema entities.
//
// Returns:
//
//	A pointer to an instance of ListOneByIDSchemaUseCase.
func NewListOneByIDSchemaUseCase(
	schemaRepository entity.SchemaRepositoryInterface,
) *ListOneByIDSchemaUseCase {
	return &ListOneByIDSchemaUseCase{
		SchemaRepository: schemaRepository,
	}
}

// Execute retrieves a schema by its ID from the repository and converts it to an output DTO.
//
// Parameters:
//
//	id: The ID of the schema to retrieve.
//
// Returns:
//
//	An output DTO containing the schema data, and an error if any occurred during the process.
func (uc *ListOneByIDSchemaUseCase) Execute(id string) (outputdto.SchemaDTO, error) {
	schema, err := uc.SchemaRepository.FindByID(id)
	if err != nil {
		return outputdto.SchemaDTO{}, err
	}

	dto := outputdto.SchemaDTO{
		ID:         string(schema.ID),
		Service:    schema.Service,
		Source:     schema.Source,
		Provider:   schema.Provider,
		SchemaType: schema.SchemaType,
		JsonSchema: converter.ConvertJsonSchemaEntityToDTO(schema.JsonSchema),
		CreatedAt:  schema.CreatedAt,
		UpdatedAt:  schema.UpdatedAt,
	}

	return dto, nil
}
