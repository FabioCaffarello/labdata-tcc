package usecase

import (
	"libs/golang/ddd/domain/entities/schema-vault/entity"
	outputdto "libs/golang/ddd/dtos/schema-vault/output"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/schema-vault/converter"
)

// ListOneByServiceAndSourceAndProviderAndSchemaTypeSchemaUseCase is the use case for listing a schema by service, source, provider and schema type.
type ListOneByServiceAndSourceAndProviderAndSchemaTypeSchemaUseCase struct {
	SchemaRepository entity.SchemaRepositoryInterface
}

// NewListAllBySourceAndProviderSchemaUseCase initializes a new instance of ListOneByServiceAndSourceAndProviderAndSchemaTypeSchemaUseCase with the provided SchemaRepositoryInterface.
//
// Parameters:
//
//	schemaRepository: The repository interface for managing Schema entities.
//
// Returns:
//
//	A pointer to an instance of ListOneByServiceAndSourceAndProviderAndSchemaTypeSchemaUseCase.
func NewListOneByServiceAndSourceAndProviderAndSchemaTypeSchemaUseCase(
	schemaRepository entity.SchemaRepositoryInterface,
) *ListOneByServiceAndSourceAndProviderAndSchemaTypeSchemaUseCase {
	return &ListOneByServiceAndSourceAndProviderAndSchemaTypeSchemaUseCase{
		SchemaRepository: schemaRepository,
	}
}

// Execute retrieves a schema by its provider, service, source and schema type from the repository and converts it to an output DTO.
//
// Parameters:
//
//	provider: The provider name to filter schemas by.
//	service: The service name to filter schemas by.
//	source: The source name to filter schemas by.
//	schemaType: The schema type to filter schemas by.
//
// Returns:
//
//	An output DTO containing the schema data, and an error if any occurred during the process.
func (uc *ListOneByServiceAndSourceAndProviderAndSchemaTypeSchemaUseCase) Execute(provider, service, source, schemaType string) (outputdto.SchemaDTO, error) {
	schema, err := uc.SchemaRepository.FindOneByServiceAndSourceAndProviderAndSchemaType(provider, service, source, schemaType)
	if err != nil {
		return outputdto.SchemaDTO{}, err
	}

	dto := outputdto.SchemaDTO{
		ID:              string(schema.ID),
		Service:         schema.Service,
		Source:          schema.Source,
		Provider:        schema.Provider,
		SchemaType:      schema.SchemaType,
		JsonSchema:      converter.ConvertJsonSchemaEntityToDTO(schema.JsonSchema),
		SchemaVersionID: string(schema.SchemaVersionID),
		CreatedAt:       schema.CreatedAt,
		UpdatedAt:       schema.UpdatedAt,
	}

	return dto, nil
}
