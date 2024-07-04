package usecase

import (
	"fmt"
	"libs/golang/ddd/domain/entities/schema-vault/entity"
	mockrepository "libs/golang/ddd/domain/repositories/database/mock/schema-vault/repository"
	inputdto "libs/golang/ddd/dtos/schema-vault/input"
	outputdto "libs/golang/ddd/dtos/schema-vault/output"
	shareddto "libs/golang/ddd/dtos/schema-vault/shared"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/schema-vault/converter"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UpdateSchemaUseCaseSuite struct {
	suite.Suite
	repoMock    *mockrepository.SchemaRepositoryMock
	useCase     *UpdateSchemaUseCase
	inputDTO    inputdto.SchemaDTO
	schemaProps entity.SchemaProps
}

func TestUpdateSchemaUseCaseSuite(t *testing.T) {
	suite.Run(t, new(UpdateSchemaUseCaseSuite))
}

func (suite *UpdateSchemaUseCaseSuite) SetupTest() {
	suite.repoMock = new(mockrepository.SchemaRepositoryMock)
	suite.useCase = NewUpdateSchemaUseCase(suite.repoMock)
	suite.inputDTO = inputdto.SchemaDTO{
		Service:    "test_service",
		Source:     "test_source",
		Provider:   "test_provider",
		SchemaType: "test_schema_type",
		JsonSchema: shareddto.JsonSchemaDTO{
			JsonType: "object",
			Properties: map[string]interface{}{
				"field1": map[string]interface{}{
					"type": "string",
				},
				"field2": map[string]interface{}{
					"type": "string",
				},
			},
			Required: []string{
				"field1",
			},
		},
	}
	suite.schemaProps = entity.SchemaProps{
		Service:    suite.inputDTO.Service,
		Source:     suite.inputDTO.Source,
		Provider:   suite.inputDTO.Provider,
		SchemaType: suite.inputDTO.SchemaType,
		JsonSchema: converter.ConvertJsonSchemaDTOToMap(suite.inputDTO.JsonSchema),
	}
}

func (suite *UpdateSchemaUseCaseSuite) TestExecuteWhenSuccess() {
	expectedSchema, _ := entity.NewSchema(suite.schemaProps)
	suite.repoMock.On("Update", expectedSchema).Return(nil)

	output, err := suite.useCase.Execute(suite.inputDTO)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), suite.inputDTO.Service, output.Service)
	assert.Equal(suite.T(), suite.inputDTO.Source, output.Source)
	assert.Equal(suite.T(), suite.inputDTO.Provider, output.Provider)
	assert.Equal(suite.T(), suite.inputDTO.SchemaType, output.SchemaType)
	assert.Equal(suite.T(), suite.inputDTO.JsonSchema.JsonType, output.JsonSchema.JsonType)
	assert.Equal(suite.T(), suite.inputDTO.JsonSchema.Properties, output.JsonSchema.Properties)
	assert.Equal(suite.T(), suite.inputDTO.JsonSchema.Required, output.JsonSchema.Required)

	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *UpdateSchemaUseCaseSuite) TestExecuteWhenError() {
	expectedSchema, _ := entity.NewSchema(suite.schemaProps)
	suite.repoMock.On("Update", expectedSchema).Return(fmt.Errorf("No schemas found for provider: %s and service: %s", suite.inputDTO.Provider, suite.inputDTO.Service))

	output, err := suite.useCase.Execute(suite.inputDTO)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), outputdto.SchemaDTO{}, output)
	suite.repoMock.AssertExpectations(suite.T())
}
