package usecase

import (
	"errors"
	"libs/golang/ddd/domain/entities/schema-vault/entity"
	mockrepository "libs/golang/ddd/domain/repositories/database/mock/schema-vault/repository"
	inputdto "libs/golang/ddd/dtos/schema-vault/input"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ValidateSchemaUseCaseSuite struct {
	suite.Suite
	repoMock *mockrepository.SchemaRepositoryMock
	useCase  *ValidateSchemaUseCase
}

func TestValidateSchemaUseCaseSuite(t *testing.T) {
	suite.Run(t, new(ValidateSchemaUseCaseSuite))
}

func (suite *ValidateSchemaUseCaseSuite) SetupTest() {
	suite.repoMock = new(mockrepository.SchemaRepositoryMock)
	suite.useCase = NewValidateSchemaUseCase(suite.repoMock)
}

func (suite *ValidateSchemaUseCaseSuite) TestExecuteWhenSuccess() {
	schema := &entity.Schema{
		ID:         "1",
		Service:    "service1",
		Source:     "source1",
		Provider:   "provider",
		SchemaType: "input",
		JsonSchema: entity.JsonSchema{
			Required: []string{"field1"},
			Properties: map[string]interface{}{
				"field1": map[string]interface{}{
					"type": "string",
				},
				"field2": map[string]interface{}{
					"type": "string",
				},
			},
			JsonType: "object",
		},
		CreatedAt: "2023-06-01T00:00:00Z",
		UpdatedAt: "2023-06-01T00:00:00Z",
	}

	suite.repoMock.On("FindOneByServiceAndSourceAndProviderAndSchemaType", "service1", "source1", "provider", "input").Return(schema, nil)

	dto := inputdto.SchemaDataDTO{
		Service:    "service1",
		Source:     "source1",
		Provider:   "provider",
		SchemaType: "input",
		Data: map[string]interface{}{
			"field1": "value1",
			"field2": "value2",
		},
	}

	valid, err := suite.useCase.Execute(dto)
	assert.NoError(suite.T(), err)
	assert.True(suite.T(), valid.Valid)
}

func (suite *ValidateSchemaUseCaseSuite) TestExecuteWhenSchemaNotFound() {
	suite.repoMock.On("FindOneByServiceAndSourceAndProviderAndSchemaType", "service1", "source1", "provider", "input").Return(nil, errors.New("schema not found"))

	dto := inputdto.SchemaDataDTO{
		Service:    "service1",
		Source:     "source1",
		Provider:   "provider",
		SchemaType: "input",
		Data: map[string]interface{}{
			"field1": "value1",
			"field2": "value2",
		},
	}

	_, err := suite.useCase.Execute(dto)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "failed to find schema: schema not found", err.Error())

}

func (suite *ValidateSchemaUseCaseSuite) TestExecuteWhenValidationFails() {
	schema := &entity.Schema{
		ID:         "1",
		Service:    "service1",
		Source:     "source1",
		Provider:   "provider",
		SchemaType: "input",
		JsonSchema: entity.JsonSchema{
			Required: []string{"field1"},
			Properties: map[string]interface{}{
				"field1": map[string]interface{}{
					"type": "string",
				},
			},
			JsonType: "object",
		},
		CreatedAt: "2023-06-01T00:00:00Z",
		UpdatedAt: "2023-06-01T00:00:00Z",
	}

	suite.repoMock.On("FindOneByServiceAndSourceAndProviderAndSchemaType", "service1", "source1", "provider", "input").Return(schema, nil)

	dto := inputdto.SchemaDataDTO{
		Service:    "service1",
		Source:     "source1",
		Provider:   "provider",
		SchemaType: "input",
		Data: map[string]interface{}{
			"field1": 123, // Invalid type, should be a string
		},
	}

	_, err := suite.useCase.Execute(dto)
	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "failed to validate JSON data:")
}
