package usecase

import (
	"libs/golang/ddd/domain/entities/schema-vault/entity"
	outputdto "libs/golang/ddd/dtos/schema-vault/output"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/schema-vault/converter"
)

// ListAllSchemaUseCase is the use case for listing all schemas.
type ListAllSchemaUseCase struct {
	SchemaRepository entity.SchemaRepositoryInterface
}

// NewListAllSchemaUseCase initializes a new instance of ListAllSchemaUseCase with the provided SchemaRepositoryInterface.
//
// Parameters:
//
//	schemaRepository: The repository interface for managing Schema entities.
//
// Returns:
//
//	A pointer to an instance of ListAllSchemaUseCase.
func NewListAllSchemaUseCase(
	schemaRepository entity.SchemaRepositoryInterface,
) *ListAllSchemaUseCase {
	return &ListAllSchemaUseCase{
		SchemaRepository: schemaRepository,
	}
}

// Execute retrieves all schemas from the repository and converts them to output DTOs.
//
// Returns:
//
//	A slice of output DTOs containing the schema data, and an error if any occurred during the process.
func (uc *ListAllSchemaUseCase) Execute() ([]outputdto.SchemaDTO, error) {
	schemas, err := uc.SchemaRepository.FindAll()
	if err != nil {
		return []outputdto.SchemaDTO{}, err
	}

	schemaDTOs := make([]outputdto.SchemaDTO, 0, len(schemas))
	for _, schema := range schemas {
		schemaDTOs = append(schemaDTOs, outputdto.SchemaDTO{
			ID:         string(schema.ID),
			Service:    schema.Service,
			Source:     schema.Source,
			Provider:   schema.Provider,
			SchemaType: schema.SchemaType,
			JsonSchema: converter.ConvertJsonSchemaEntityToDTO(schema.JsonSchema),
			CreatedAt:  schema.CreatedAt,
			UpdatedAt:  schema.UpdatedAt,
		})
	}

	return schemaDTOs, nil
}
