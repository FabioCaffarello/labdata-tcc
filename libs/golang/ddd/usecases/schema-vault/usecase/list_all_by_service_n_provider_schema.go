package usecase

import (
	"libs/golang/ddd/domain/entities/schema-vault/entity"
	outputdto "libs/golang/ddd/dtos/schema-vault/output"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/schema-vault/converter"
)

// ListAllByServiceAndProviderSchemaUseCase is the use case for listing all schemas by service and provider.
type ListAllByServiceAndProviderSchemaUseCase struct {
	SchemaRepository entity.SchemaRepositoryInterface
}

// NewListAllByServiceAndProviderSchemaUseCase initializes a new instance of ListAllByServiceAndProviderSchemaUseCase with the provided SchemaRepositoryInterface.
//
// Parameters:
//
//	schemaRepository: The repository interface for managing Schema entities.
//
// Returns:
//
//	A pointer to an instance of ListAllByServiceAndProviderSchemaUseCase.
func NewListAllByServiceAndProviderSchemaUseCase(
	schemaRepository entity.SchemaRepositoryInterface,
) *ListAllByServiceAndProviderSchemaUseCase {
	return &ListAllByServiceAndProviderSchemaUseCase{
		SchemaRepository: schemaRepository,
	}
}

// Execute retrieves all schemas by service and provider from the repository.
//
// Parameters:
//
//	provider: The provider name to filter schemas by.
//	service: The service name to filter schemas by.
//
// Returns:
//
//	A slice of output DTOs containing the schema data, and an error if any occurred during the process.
func (uc *ListAllByServiceAndProviderSchemaUseCase) Execute(provider, service string) ([]outputdto.SchemaDTO, error) {
	schemas, err := uc.SchemaRepository.FindAllByServiceAndProvider(provider, service)
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
