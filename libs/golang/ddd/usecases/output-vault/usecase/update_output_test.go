package usecase

import (
	"fmt"
	"libs/golang/ddd/domain/entities/output-vault/entity"
	mockrepository "libs/golang/ddd/domain/repositories/database/mock/output-vault/repository"
	inputdto "libs/golang/ddd/dtos/output-vault/input"
	outputdto "libs/golang/ddd/dtos/output-vault/output"
	shareddto "libs/golang/ddd/dtos/output-vault/shared"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/output-vault/converter"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UpdateOutputUseCaseSuite struct {
	suite.Suite
	repoMock    *mockrepository.OutputRepositoryMock
	useCase     *UpdateOutputUseCase
	inputDTO    inputdto.OutputDTO
	outputProps entity.OutputProps
}

func TestUpdateOutputUseCaseSuite(t *testing.T) {
	suite.Run(t, new(UpdateOutputUseCaseSuite))
}

func (suite *UpdateOutputUseCaseSuite) SetupTest() {
	suite.repoMock = new(mockrepository.OutputRepositoryMock)
	suite.useCase = NewUpdateOutputUseCase(suite.repoMock)
	suite.inputDTO = inputdto.OutputDTO{
		Service:  "test_service",
		Source:   "test_source",
		Provider: "test_provider",
		Data: map[string]interface{}{
			"field1": "value1",
			"field2": "value2",
		},
		Metadata: shareddto.MetadataDTO{
			InputID: "input-id",
			Input: shareddto.InputDTO{
				Data:                map[string]interface{}{"key": "value"},
				ProcessingID:        "processing-id",
				ProcessingTimestamp: "2021-06-01 00:00:00",
			},
		},
	}
	suite.outputProps = entity.OutputProps{
		Service:  suite.inputDTO.Service,
		Source:   suite.inputDTO.Source,
		Provider: suite.inputDTO.Provider,
		Data:     suite.inputDTO.Data,
		Metadata: converter.ConvertMetadataDTOToMap(suite.inputDTO.Metadata),
	}
}

func (suite *UpdateOutputUseCaseSuite) TestExecuteWhenSuccess() {
	expectedOutput, _ := entity.NewOutput(suite.outputProps)
	suite.repoMock.On("Update", expectedOutput).Return(nil)

	output, err := suite.useCase.Execute(suite.inputDTO)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), suite.inputDTO.Service, output.Service)
	assert.Equal(suite.T(), suite.inputDTO.Source, output.Source)
	assert.Equal(suite.T(), suite.inputDTO.Provider, output.Provider)
	assert.Equal(suite.T(), suite.inputDTO.Data, output.Data)

	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *UpdateOutputUseCaseSuite) TestExecuteWhenError() {
	expectedOutput, _ := entity.NewOutput(suite.outputProps)
	suite.repoMock.On("Update", expectedOutput).Return(fmt.Errorf("No outputs found for provider: %s and service: %s", suite.inputDTO.Provider, suite.inputDTO.Service))

	output, err := suite.useCase.Execute(suite.inputDTO)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), outputdto.OutputDTO{}, output)
	suite.repoMock.AssertExpectations(suite.T())
}
