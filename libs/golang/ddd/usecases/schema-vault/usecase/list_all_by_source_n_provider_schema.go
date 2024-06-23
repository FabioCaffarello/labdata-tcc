package usecase

import (
	"libs/golang/ddd/domain/entities/schema-vault/entity"
	outputdto "libs/golang/ddd/dtos/schema-vault/output"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/schema-vault/converter"
)

// ListAllBySourceAndProviderSchemaUseCase is the use case for listing all schemas by source.
type ListAllBySourceAndProviderSchemaUseCase struct {
	SchemaRepository entity.SchemaRepositoryInterface
}

// NewListAllBySourceAndProviderSchemaUseCase initializes a new instance of ListAllBySourceAndProviderSchemaUseCase with the provided SchemaRepositoryInterface.
//
// Parameters:
//
//	schemaRepository: The repository interface for managing Schema entities.
//
// Returns:
//
//	A pointer to an instance of ListAllBySourceAndProviderSchemaUseCase.
func NewListAllBySourceAndProviderSchemaUseCase(
	schemaRepository entity.SchemaRepositoryInterface,
) *ListAllBySourceAndProviderSchemaUseCase {
	return &ListAllBySourceAndProviderSchemaUseCase{
		SchemaRepository: schemaRepository,
	}
}

// Execute retrieves all schemas by source from the repository.
//
// Parameters:
//
//	provider: The provider name to filter schemas by.
//	source: The source name to filter schemas by.
//
// Returns:
//
//	A slice of output DTOs containing the schema data, and an error if any occurred during the process.
func (uc *ListAllBySourceAndProviderSchemaUseCase) Execute(provider, source string) ([]outputdto.SchemaDTO, error) {
	schemas, err := uc.SchemaRepository.FindAllBySourceAndProvider(provider, source)
	if err != nil {
		return []outputdto.SchemaDTO{}, err
	}

	schemaDTOs := make([]outputdto.SchemaDTO, 0, len(schemas))
	for _, schema := range schemas {
		schemaDTOs = append(schemaDTOs, outputdto.SchemaDTO{
			ID:              string(schema.ID),
			Service:         schema.Service,
			Source:          schema.Source,
			Provider:        schema.Provider,
			SchemaType:      schema.SchemaType,
			SchemaVersionID: string(schema.SchemaVersionID),
			JsonSchema:      converter.ConvertJsonSchemaEntityToDTO(schema.JsonSchema),
			CreatedAt:       schema.CreatedAt,
			UpdatedAt:       schema.UpdatedAt,
		})
	}

	return schemaDTOs, nil
}
