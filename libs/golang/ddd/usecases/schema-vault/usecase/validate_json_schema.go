package usecase

import (
	"fmt"
	"libs/golang/ddd/domain/entities/schema-vault/entity"
	inputdto "libs/golang/ddd/dtos/schema-vault/input"
	schematools "libs/golang/shared/json-schema/schema-tools"
)

type ValidateSchemaUseCase struct {
	SchemaRepository entity.SchemaRepositoryInterface
}

func NewValidateSchemaUseCase(schemaRepository entity.SchemaRepositoryInterface) *ValidateSchemaUseCase {
	return &ValidateSchemaUseCase{
		SchemaRepository: schemaRepository,
	}
}

func (uc *ValidateSchemaUseCase) Execute(dto inputdto.SchemaDataDTO) error {
	schema, err := uc.SchemaRepository.FindOneByServiceAndSourceAndProviderAndSchemaType(dto.Service, dto.Source, dto.Provider, dto.SchemaType)
	if err != nil {
		return err
	}

	jsonSchema, err := schema.JsonSchema.ToMap()
	if err != nil {
		return fmt.Errorf("failed to convert JSON schema to map: %w", err)
	}
	err = schematools.ValidateJSONData(jsonSchema, dto.Data)
	if err != nil {
		return fmt.Errorf("data validation against schema failed: %w", err)
	}

	return nil
}
