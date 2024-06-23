package usecase

import (
	"libs/golang/ddd/domain/entities/schema-vault/entity"
	inputdto "libs/golang/ddd/dtos/schema-vault/input"
	outputdto "libs/golang/ddd/dtos/schema-vault/output"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/schema-vault/converter"
)

// UpdateSchemaUseCase is the use case for updating an existing schema.
type UpdateSchemaUseCase struct {
	SchemaRepository entity.SchemaRepositoryInterface
}

// NewUpdateSchemaUseCase initializes a new instance of UpdateSchemaUseCase with the provided SchemaRepositoryInterface.
//
// Parameters:
//
//	schemaRepository: The repository interface for managing Schema entities.
//
// Returns:
//
//	A pointer to an instance of UpdateSchemaUseCase.
func NewUpdateSchemaUseCase(
	schemaRepository entity.SchemaRepositoryInterface,
) *UpdateSchemaUseCase {
	return &UpdateSchemaUseCase{
		SchemaRepository: schemaRepository,
	}
}

// Execute updates an existing schema entity based on the provided input DTO and saves it using the repository.
// It then converts the updated entity to an output DTO and returns it.
//
// Parameters:
//
//	input: The input DTO containing the schema data.
//
// Returns:
//
//	An output DTO containing the updated schema data, and an error if any occurred during the process.
func (uc *UpdateSchemaUseCase) Execute(input inputdto.SchemaDTO) (outputdto.SchemaDTO, error) {
	schemaProps := entity.SchemaProps{
		Service:    input.Service,
		Source:     input.Source,
		Provider:   input.Provider,
		SchemaType: input.SchemaType,
		JsonSchema: converter.ConvertJsonSchemaDTOToMap(input.JsonSchema),
	}

	entitySchema, err := entity.NewSchema(schemaProps)
	if err != nil {
		return outputdto.SchemaDTO{}, err
	}

	err = uc.SchemaRepository.Update(entitySchema)
	if err != nil {
		return outputdto.SchemaDTO{}, err
	}

	dtoJsonSchema := converter.ConvertJsonSchemaEntityToDTO(entitySchema.JsonSchema)

	return outputdto.SchemaDTO{
		ID:              string(entitySchema.ID),
		Service:         entitySchema.Service,
		Source:          entitySchema.Source,
		Provider:        entitySchema.Provider,
		SchemaType:      entitySchema.SchemaType,
		JsonSchema:      dtoJsonSchema,
		SchemaVersionID: string(entitySchema.SchemaVersionID),
		CreatedAt:       entitySchema.CreatedAt,
		UpdatedAt:       entitySchema.UpdatedAt,
	}, nil
}
