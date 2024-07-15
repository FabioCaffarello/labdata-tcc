package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"libs/golang/ddd/domain/entities/input-broker/entity"
	mockrepository "libs/golang/ddd/domain/repositories/database/mock/input-broker/repository"
	mockevent "libs/golang/ddd/events/event-mock/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	inputdto "libs/golang/ddd/dtos/input-broker/input"
	outputdto "libs/golang/ddd/dtos/input-broker/output"
	shareddto "libs/golang/ddd/dtos/input-broker/shared"

	"github.com/go-chi/chi/v5"
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

func (suite *WebInputHandlerSuite) TestUpdateInput() {
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

	suite.repoMock.On("Update", mock.AnythingOfType("*entity.Input")).Return(nil).Run(func(args mock.Arguments) {
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
	req := httptest.NewRequest(http.MethodPut, "/inputs/test_id", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	suite.handler.UpdateInput(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutput outputdto.InputDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutput)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedOutput, actualOutput)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebInputHandlerSuite) TestDeleteInput() {
	suite.repoMock.On("Delete", "1").Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/inputs/1", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.DeleteInput(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)
	assert.Equal(suite.T(), "Input deleted successfully", rr.Body.String())
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebInputHandlerSuite) TestListInputByID() {
	expectedInput := outputdto.InputDTO{
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

	suite.repoMock.On("FindByID", "test_id").Return(&entity.Input{
		ID: "test_id",
		Metadata: entity.Metadata{
			Service:             "test_service",
			Source:              "test_source",
			Provider:            "test_provider",
			ProcessingID:        "test_processing_id",
			ProcessingTimestamp: "2023-06-01 00:00:00",
		},
		Status: entity.Status{
			Code:   0,
			Detail: "test_detail",
		},
		Data: map[string]interface{}{
			"key": "value",
		},
		CreatedAt: "2023-06-01 00:00:00",
		UpdatedAt: "2023-06-01 00:00:00",
	}, nil)

	req := httptest.NewRequest(http.MethodGet, "/inputs/test_id", nil)
	rr := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "test_id")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	suite.handler.ListInputByID(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutput outputdto.InputDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutput)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedInput, actualOutput)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebInputHandlerSuite) TestListAllInputs() {
	expectedInputs := []outputdto.InputDTO{
		{
			ID: "test_id_1",
			Metadata: shareddto.MetadataDTO{
				Service:             "test_service_1",
				Source:              "test_source_1",
				Provider:            "test_provider_1",
				ProcessingID:        "test_processing_id_1",
				ProcessingTimestamp: "2023-06-01 00:00:00",
			},
			Status: shareddto.StatusDTO{
				Code:   0,
				Detail: "test_detail_1",
			},
			Data: map[string]interface{}{
				"key": "value_1",
			},
			CreatedAt: "2023-06-01 00:00:00",
			UpdatedAt: "2023-06-01 00:00:00",
		},
	}

	suite.repoMock.On("FindAll").Return([]*entity.Input{
		{
			ID: "test_id_1",
			Metadata: entity.Metadata{
				Service:             "test_service_1",
				Source:              "test_source_1",
				Provider:            "test_provider_1",
				ProcessingID:        "test_processing_id_1",
				ProcessingTimestamp: "2023-06-01 00:00:00",
			},
			Status: entity.Status{
				Code:   0,
				Detail: "test_detail_1",
			},
			Data: map[string]interface{}{
				"key": "value_1",
			},
			CreatedAt: "2023-06-01 00:00:00",
			UpdatedAt: "2023-06-01 00:00:00",
		},
	}, nil)

	req := httptest.NewRequest(http.MethodGet, "/inputs", nil)
	rr := httptest.NewRecorder()

	suite.handler.ListAllInputs(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutputs []outputdto.InputDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutputs)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedInputs, actualOutputs)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebInputHandlerSuite) TestListInputsByServiceAndProvider() {
	expectedInputs := []outputdto.InputDTO{
		{
			ID: "test_id_1",
			Metadata: shareddto.MetadataDTO{
				Service:             "test_service",
				Source:              "test_source",
				Provider:            "test_provider",
				ProcessingID:        "test_processing_id_1",
				ProcessingTimestamp: "2023-06-01 00:00:00",
			},
			Status: shareddto.StatusDTO{
				Code:   0,
				Detail: "test_detail_1",
			},
			Data: map[string]interface{}{
				"key": "value_1",
			},
			CreatedAt: "2023-06-01 00:00:00",
			UpdatedAt: "2023-06-01 00:00:00",
		},
	}

	suite.repoMock.On("FindAllByServiceAndProvider", "test_provider", "test_service").Return([]*entity.Input{
		{
			ID: "test_id_1",
			Metadata: entity.Metadata{
				Service:             "test_service",
				Source:              "test_source",
				Provider:            "test_provider",
				ProcessingID:        "test_processing_id_1",
				ProcessingTimestamp: "2023-06-01 00:00:00",
			},
			Status: entity.Status{
				Code:   0,
				Detail: "test_detail_1",
			},
			Data: map[string]interface{}{
				"key": "value_1",
			},
			CreatedAt: "2023-06-01 00:00:00",
			UpdatedAt: "2023-06-01 00:00:00",
		},
	}, nil)

	req := httptest.NewRequest(http.MethodGet, "/inputs/provider/test_provider/service/test_service", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("provider", "test_provider")
	rctx.URLParams.Add("service", "test_service")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListInputsByServiceAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutputs []outputdto.InputDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutputs)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedInputs, actualOutputs)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebInputHandlerSuite) TestListInputsBySourceAndProvider() {
	expectedInputs := []outputdto.InputDTO{
		{
			ID: "test_id_1",
			Metadata: shareddto.MetadataDTO{
				Service:             "test_service",
				Source:              "test_source",
				Provider:            "test_provider",
				ProcessingID:        "test_processing_id_1",
				ProcessingTimestamp: "2023-06-01 00:00:00",
			},
			Status: shareddto.StatusDTO{
				Code:   0,
				Detail: "test_detail_1",
			},
			Data: map[string]interface{}{
				"key": "value_1",
			},
			CreatedAt: "2023-06-01 00:00:00",
			UpdatedAt: "2023-06-01 00:00:00",
		},
	}

	suite.repoMock.On("FindAllBySourceAndProvider", "test_provider", "test_source").Return([]*entity.Input{
		{
			ID: "test_id_1",
			Metadata: entity.Metadata{
				Service:             "test_service",
				Source:              "test_source",
				Provider:            "test_provider",
				ProcessingID:        "test_processing_id_1",
				ProcessingTimestamp: "2023-06-01 00:00:00",
			},
			Status: entity.Status{
				Code:   0,
				Detail: "test_detail_1",
			},
			Data: map[string]interface{}{
				"key": "value_1",
			},
			CreatedAt: "2023-06-01 00:00:00",
			UpdatedAt: "2023-06-01 00:00:00",
		},
	}, nil)

	req := httptest.NewRequest(http.MethodGet, "/inputs/provider/test_provider/source/test_source", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("provider", "test_provider")
	rctx.URLParams.Add("source", "test_source")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListInputsBySourceAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutputs []outputdto.InputDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutputs)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedInputs, actualOutputs)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebInputHandlerSuite) TestListInputsByServiceAndSourceAndProvider() {
	expectedInputs := []outputdto.InputDTO{
		{
			ID: "test_id_1",
			Metadata: shareddto.MetadataDTO{
				Service:             "test_service",
				Source:              "test_source",
				Provider:            "test_provider",
				ProcessingID:        "test_processing_id_1",
				ProcessingTimestamp: "2023-06-01 00:00:00",
			},
			Status: shareddto.StatusDTO{
				Code:   0,
				Detail: "test_detail_1",
			},
			Data: map[string]interface{}{
				"key": "value_1",
			},
			CreatedAt: "2023-06-01 00:00:00",
			UpdatedAt: "2023-06-01 00:00:00",
		},
	}

	suite.repoMock.On("FindAllByServiceAndSourceAndProvider", "test_provider", "test_service", "test_source").Return([]*entity.Input{
		{
			ID: "test_id_1",
			Metadata: entity.Metadata{
				Service:             "test_service",
				Source:              "test_source",
				Provider:            "test_provider",
				ProcessingID:        "test_processing_id_1",
				ProcessingTimestamp: "2023-06-01 00:00:00",
			},
			Status: entity.Status{
				Code:   0,
				Detail: "test_detail_1",
			},
			Data: map[string]interface{}{
				"key": "value_1",
			},
			CreatedAt: "2023-06-01 00:00:00",
			UpdatedAt: "2023-06-01 00:00:00",
		},
	}, nil)

	req := httptest.NewRequest(http.MethodGet, "/inputs/provider/test_provider/service/test_service/source/test_source", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("provider", "test_provider")
	rctx.URLParams.Add("service", "test_service")
	rctx.URLParams.Add("source", "test_source")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListInputsByServiceAndSourceAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutputs []outputdto.InputDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutputs)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedInputs, actualOutputs)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebInputHandlerSuite) TestListInputsByStatusAndProvider() {
	expectedInputs := []outputdto.InputDTO{
		{
			ID: "test_id_1",
			Metadata: shareddto.MetadataDTO{
				Service:             "test_service",
				Source:              "test_source",
				Provider:            "test_provider",
				ProcessingID:        "test_processing_id_1",
				ProcessingTimestamp: "2023-06-01 00:00:00",
			},
			Status: shareddto.StatusDTO{
				Code:   0,
				Detail: "test_detail_1",
			},
			Data: map[string]interface{}{
				"key": "value_1",
			},
			CreatedAt: "2023-06-01 00:00:00",
			UpdatedAt: "2023-06-01 00:00:00",
		},
	}

	suite.repoMock.On("FindAllByStatusAndProvider", "test_provider", 0).Return([]*entity.Input{
		{
			ID: "test_id_1",
			Metadata: entity.Metadata{
				Service:             "test_service",
				Source:              "test_source",
				Provider:            "test_provider",
				ProcessingID:        "test_processing_id_1",
				ProcessingTimestamp: "2023-06-01 00:00:00",
			},
			Status: entity.Status{
				Code:   0,
				Detail: "test_detail_1",
			},
			Data: map[string]interface{}{
				"key": "value_1",
			},
			CreatedAt: "2023-06-01 00:00:00",
			UpdatedAt: "2023-06-01 00:00:00",
		},
	}, nil)

	req := httptest.NewRequest(http.MethodGet, "/inputs/provider/test_provider/status/0", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("provider", "test_provider")
	rctx.URLParams.Add("status", "0")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListInputsByStatusAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutputs []outputdto.InputDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutputs)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedInputs, actualOutputs)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebInputHandlerSuite) TestListInputsByStatusAndServiceAndProvider() {
	expectedInputs := []outputdto.InputDTO{
		{
			ID: "test_id_1",
			Metadata: shareddto.MetadataDTO{
				Service:             "test_service",
				Source:              "test_source",
				Provider:            "test_provider",
				ProcessingID:        "test_processing_id_1",
				ProcessingTimestamp: "2023-06-01 00:00:00",
			},
			Status: shareddto.StatusDTO{
				Code:   0,
				Detail: "test_detail_1",
			},
			Data: map[string]interface{}{
				"key": "value_1",
			},
			CreatedAt: "2023-06-01 00:00:00",
			UpdatedAt: "2023-06-01 00:00:00",
		},
	}

	suite.repoMock.On("FindAllByStatusAndServiceAndProvider", "test_service", "test_provider", 0).Return([]*entity.Input{
		{
			ID: "test_id_1",
			Metadata: entity.Metadata{
				Service:             "test_service",
				Source:              "test_source",
				Provider:            "test_provider",
				ProcessingID:        "test_processing_id_1",
				ProcessingTimestamp: "2023-06-01 00:00:00",
			},
			Status: entity.Status{
				Code:   0,
				Detail: "test_detail_1",
			},
			Data: map[string]interface{}{
				"key": "value_1",
			},
			CreatedAt: "2023-06-01 00:00:00",
			UpdatedAt: "2023-06-01 00:00:00",
		},
	}, nil)

	req := httptest.NewRequest(http.MethodGet, "/inputs/provider/test_provider/service/test_service/status/0", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("provider", "test_provider")
	rctx.URLParams.Add("service", "test_service")
	rctx.URLParams.Add("status", "0")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListInputsByStatusAndServiceAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutputs []outputdto.InputDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutputs)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedInputs, actualOutputs)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebInputHandlerSuite) TestListInputsByStatusAndSourceAndProvider() {
	expectedInputs := []outputdto.InputDTO{
		{
			ID: "test_id_1",
			Metadata: shareddto.MetadataDTO{
				Service:             "test_service",
				Source:              "test_source",
				Provider:            "test_provider",
				ProcessingID:        "test_processing_id_1",
				ProcessingTimestamp: "2023-06-01 00:00:00",
			},
			Status: shareddto.StatusDTO{
				Code:   0,
				Detail: "test_detail_1",
			},
			Data: map[string]interface{}{
				"key": "value_1",
			},
			CreatedAt: "2023-06-01 00:00:00",
			UpdatedAt: "2023-06-01 00:00:00",
		},
	}

	suite.repoMock.On("FindAllByStatusAndSourceAndProvider", "test_source", "test_provider", 0).Return([]*entity.Input{
		{
			ID: "test_id_1",
			Metadata: entity.Metadata{
				Service:             "test_service",
				Source:              "test_source",
				Provider:            "test_provider",
				ProcessingID:        "test_processing_id_1",
				ProcessingTimestamp: "2023-06-01 00:00:00",
			},
			Status: entity.Status{
				Code:   0,
				Detail: "test_detail_1",
			},
			Data: map[string]interface{}{
				"key": "value_1",
			},
			CreatedAt: "2023-06-01 00:00:00",
			UpdatedAt: "2023-06-01 00:00:00",
		},
	}, nil)

	req := httptest.NewRequest(http.MethodGet, "/inputs/provider/test_provider/source/test_source/status/0", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("provider", "test_provider")
	rctx.URLParams.Add("source", "test_source")
	rctx.URLParams.Add("status", "0")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListInputsByStatusAndSourceAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutputs []outputdto.InputDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutputs)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedInputs, actualOutputs)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebInputHandlerSuite) TestListInputsByStatusAndServiceAndSourceAndProvider() {
	expectedInputs := []outputdto.InputDTO{
		{
			ID: "test_id_1",
			Metadata: shareddto.MetadataDTO{
				Service:             "test_service",
				Source:              "test_source",
				Provider:            "test_provider",
				ProcessingID:        "test_processing_id_1",
				ProcessingTimestamp: "2023-06-01 00:00:00",
			},
			Status: shareddto.StatusDTO{
				Code:   0,
				Detail: "test_detail_1",
			},
			Data: map[string]interface{}{
				"key": "value_1",
			},
			CreatedAt: "2023-06-01 00:00:00",
			UpdatedAt: "2023-06-01 00:00:00",
		},
	}

	suite.repoMock.On("FindAllByStatusAndServiceAndSourceAndProvider", "test_service", "test_source", "test_provider", 0).Return([]*entity.Input{
		{
			ID: "test_id_1",
			Metadata: entity.Metadata{
				Service:             "test_service",
				Source:              "test_source",
				Provider:            "test_provider",
				ProcessingID:        "test_processing_id_1",
				ProcessingTimestamp: "2023-06-01 00:00:00",
			},
			Status: entity.Status{
				Code:   0,
				Detail: "test_detail_1",
			},
			Data: map[string]interface{}{
				"key": "value_1",
			},
			CreatedAt: "2023-06-01 00:00:00",
			UpdatedAt: "2023-06-01 00:00:00",
		},
	}, nil)

	req := httptest.NewRequest(http.MethodGet, "/inputs/provider/test_provider/service/test_service/source/test_source/status/0", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("provider", "test_provider")
	rctx.URLParams.Add("service", "test_service")
	rctx.URLParams.Add("source", "test_source")
	rctx.URLParams.Add("status", "0")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListInputsByStatusAndServiceAndSourceAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutputs []outputdto.InputDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutputs)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedInputs, actualOutputs)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebInputHandlerSuite) TestUpdateInputStatus() {
	statusDTO := shareddto.StatusDTO{
		Code:   1,
		Detail: "updated_detail",
	}

	// Mock the current time
	mockTime := time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC)
	entity.DateLayout = "2006-01-02 15:04:05"

	expectedOutput := outputdto.InputDTO{
		ID: "test_id",
		Metadata: shareddto.MetadataDTO{
			Service:             "test_service",
			Source:              "test_source",
			Provider:            "test_provider",
			ProcessingID:        "test_processing_id",
			ProcessingTimestamp: mockTime.Format(entity.DateLayout),
		},
		Status: shareddto.StatusDTO{
			Code:   1,
			Detail: "updated_detail",
		},
		Data: map[string]interface{}{
			"key": "value",
		},
		CreatedAt: "2023-06-01 00:00:00",
		UpdatedAt: mockTime.Format(entity.DateLayout),
	}

	suite.repoMock.On("FindByID", "test_id").Return(&entity.Input{
		ID: "test_id",
		Metadata: entity.Metadata{
			Service:             "test_service",
			Source:              "test_source",
			Provider:            "test_provider",
			ProcessingID:        "test_processing_id",
			ProcessingTimestamp: "2023-06-01 00:00:00",
		},
		Status: entity.Status{
			Code:   0,
			Detail: "test_detail",
		},
		Data: map[string]interface{}{
			"key": "value",
		},
		CreatedAt: "2023-06-01 00:00:00",
		UpdatedAt: "2023-06-01 00:00:00",
	}, nil)

	suite.repoMock.On("Update", mock.AnythingOfType("*entity.Input")).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*entity.Input)
		arg.UpdatedAt = mockTime.Format(entity.DateLayout)
		arg.Metadata.ProcessingTimestamp = mockTime.Format(entity.DateLayout)
	})

	reqBody, _ := json.Marshal(statusDTO)
	req := httptest.NewRequest(http.MethodPut, "/inputs/test_id/status", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "test_id")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.UpdateInputStatus(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutput outputdto.InputDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutput)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedOutput, actualOutput)
	suite.repoMock.AssertExpectations(suite.T())
}
