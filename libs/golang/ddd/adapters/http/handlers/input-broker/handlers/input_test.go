package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"libs/golang/ddd/domain/entities/input-broker/entity"
	mockrepository "libs/golang/ddd/domain/repositories/database/mock/input-broker/repository"
	mockevent "libs/golang/ddd/events/event-mock/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	inputdto "libs/golang/ddd/dtos/input-broker/input"
	outputdto "libs/golang/ddd/dtos/input-broker/output"
	shareddto "libs/golang/ddd/dtos/input-broker/shared"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type WebInputHandlerSuite struct {
	suite.Suite
	handler        *WebInputHandler
	repoMock       *mockrepository.InputRepositoryMock
	eventMock      *mockevent.MockEvent
	dispatcherMock *mockevent.MockEventDispatcher
}

func TestWebInputHandlerSuite(t *testing.T) {
	suite.Run(t, new(WebInputHandlerSuite))
}

func (suite *WebInputHandlerSuite) SetupTest() {
	suite.repoMock = new(mockrepository.InputRepositoryMock)
	suite.eventMock = new(mockevent.MockEvent)
	suite.dispatcherMock = new(mockevent.MockEventDispatcher)
	suite.handler = NewWebInputHandler(suite.repoMock, suite.dispatcherMock, suite.eventMock)
}

func (suite *WebInputHandlerSuite) TestCreateInput() {
	inputDTO := inputdto.InputDTO{
		Provider: "test_provider",
		Service:  "test_service",
		Source:   "test_source",
		Data:     map[string]interface{}{"key": "value"},
	}

	expectedOutput := outputdto.InputDTO{
		ID: "test_id",
		Metadata: shareddto.MetadataDTO{
			Service:             "test_service",
			Source:              "test_source",
			Provider:            "test_provider",
			ProcessingID:        "test_processing_id",
			ProcessingTimestamp: "2023-06-01 00:00:00",
		},
		Status: shareddto.StatusDTO{
			Code:   0,
			Detail: "test_detail",
		},
		Data: map[string]interface{}{
			"key": "value",
		},
		CreatedAt: "2023-06-01 00:00:00",
		UpdatedAt: "2023-06-01 00:00:00",
	}

	suite.repoMock.On("Create", mock.AnythingOfType("*entity.Input")).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*entity.Input)
		arg.ID = "test_id"
		arg.Metadata.ProcessingID = "test_processing_id"
		arg.Metadata.ProcessingTimestamp = "2023-06-01 00:00:00"
		arg.Metadata.Service = inputDTO.Service
		arg.Metadata.Source = inputDTO.Source
		arg.Metadata.Provider = inputDTO.Provider
		arg.Status.Code = 0
		arg.Status.Detail = "test_detail"
		arg.Data = inputDTO.Data
		arg.CreatedAt = "2023-06-01 00:00:00"
		arg.UpdatedAt = "2023-06-01 00:00:00"
	})

	jsonBody, _ := json.Marshal(inputDTO)
	req := httptest.NewRequest(http.MethodPost, "/inputs", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	suite.eventMock.On("SetPayload", mock.AnythingOfType("outputdto.InputDTO")).Return(nil)
	suite.dispatcherMock.On("Dispatch", suite.eventMock, fmt.Sprintf("input.created.%s.%s.%s", inputDTO.Provider, inputDTO.Service, inputDTO.Source)).Return(nil)

	suite.handler.CreateInput(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutput outputdto.InputDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutput)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedOutput, actualOutput)
	suite.repoMock.AssertExpectations(suite.T())
	suite.eventMock.AssertExpectations(suite.T())
	suite.dispatcherMock.AssertExpectations(suite.T())
}
