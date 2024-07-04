package usecase

import (
	"libs/golang/ddd/domain/entities/schema-vault/entity"
	mockrepository "libs/golang/ddd/domain/repositories/database/mock/schema-vault/repository"
	outputdto "libs/golang/ddd/dtos/schema-vault/output"
	shareddto "libs/golang/ddd/dtos/schema-vault/shared"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ListOneByIDSchemaUseCaseSuite struct {
	suite.Suite
	repoMock *mockrepository.SchemaRepositoryMock
	useCase  *ListOneByIDSchemaUseCase
}

func TestListOneByIDSchemaUseCaseSuite(t *testing.T) {
	suite.Run(t, new(ListOneByIDSchemaUseCaseSuite))
}

func (suite *ListOneByIDSchemaUseCaseSuite) SetupTest() {
	suite.repoMock = new(mockrepository.SchemaRepositoryMock)
	suite.useCase = NewListOneByIDSchemaUseCase(suite.repoMock)
}

func (suite *ListOneByIDSchemaUseCaseSuite) TestExecuteWhenSuccess() {
	entitySchemas := []*entity.Schema{
		{
			ID:         "1",
			Service:    "service1",
			Source:     "source1",
			Provider:   "provider",
			SchemaType: "type1",
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
		},
		{
			ID:         "2",
			Service:    "service2",
			Source:     "source2",
			Provider:   "provider",
			SchemaType: "type2",
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
		},
	}

	suite.repoMock.On("FindByID", "1").Return(entitySchemas[0], nil)

	expectedOutput := outputdto.SchemaDTO{
		ID:         "1",
		Service:    "service1",
		Source:     "source1",
		Provider:   "provider",
		SchemaType: "type1",
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
			Required: []string{"field1"},
		},
		CreatedAt: "2023-06-01T00:00:00Z",
		UpdatedAt: "2023-06-01T00:00:00Z",
	}

	output, err := suite.useCase.Execute("1")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, output)
	suite.repoMock.AssertExpectations(suite.T())
}
