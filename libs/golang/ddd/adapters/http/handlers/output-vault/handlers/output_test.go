package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"libs/golang/ddd/domain/entities/output-vault/entity"
	mockrepository "libs/golang/ddd/domain/repositories/database/mock/output-vault/repository"
	inputdto "libs/golang/ddd/dtos/output-vault/input"
	outputdto "libs/golang/ddd/dtos/output-vault/output"
	shareddto "libs/golang/ddd/dtos/output-vault/shared"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type WebOutputHandlerSuite struct {
	suite.Suite
	handler  *WebOutputHandler
	repoMock *mockrepository.OutputRepositoryMock
}

func TestWebOutputHandlerSuite(t *testing.T) {
	suite.Run(t, new(WebOutputHandlerSuite))
}

func (suite *WebOutputHandlerSuite) SetupTest() {
	suite.repoMock = new(mockrepository.OutputRepositoryMock)
	suite.handler = NewWebOutputHandler(suite.repoMock)
}

// Tests for CreateOutput handler
func (suite *WebOutputHandlerSuite) TestCreateOutputWhenSuccess() {
	inputDTO := inputdto.OutputDTO{
		Service:  "test_service",
		Source:   "test_source",
		Provider: "test_provider",
		Data: map[string]interface{}{
			"field1": "value1",
			"field2": "value2",
		},
		Metadata: shareddto.MetadataDTO{
			InputID: "input1",
			Input: shareddto.InputDTO{
				Data:                map[string]interface{}{"key": "value"},
				ProcessingID:        "processing1",
				ProcessingTimestamp: "2021-06-01 00:00:00",
			},
		},
	}

	expectedOutput := outputdto.OutputDTO{
		ID:       "1",
		Service:  "test_service",
		Source:   "test_source",
		Provider: "test_provider",
		Data: map[string]interface{}{
			"field1": "value1",
			"field2": "value2",
		},
		Metadata: shareddto.MetadataDTO{
			InputID: "input1",
			Input: shareddto.InputDTO{
				Data:                map[string]interface{}{"key": "value"},
				ProcessingID:        "processing1",
				ProcessingTimestamp: "2021-06-01 00:00:00",
			},
		},
		CreatedAt: "2023-06-01 00:00:00",
		UpdatedAt: "2023-06-01 00:00:00",
	}

	suite.repoMock.On("Create", mock.AnythingOfType("*entity.Output")).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*entity.Output)
		arg.ID = "1"
		arg.Service = "test_service"
		arg.Source = "test_source"
		arg.Provider = "test_provider"
		arg.Data = map[string]interface{}{
			"field1": "value1",
			"field2": "value2",
		}
		arg.Metadata = entity.Metadata{
			InputID: "input1",
			Input: entity.Input{
				Data:                map[string]interface{}{"key": "value"},
				ProcessingID:        "processing1",
				ProcessingTimestamp: "2021-06-01 00:00:00",
			},
		}
		arg.CreatedAt = "2023-06-01 00:00:00"
		arg.UpdatedAt = "2023-06-01 00:00:00"
	})

	jsonBody, _ := json.Marshal(inputDTO)
	req := httptest.NewRequest(http.MethodPost, "/outputs", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	suite.handler.CreateOutput(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutput outputdto.OutputDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutput)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedOutput, actualOutput)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebOutputHandlerSuite) TestCreateOutputWhenDecodingFails() {
	req := httptest.NewRequest(http.MethodPost, "/outputs", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	suite.handler.CreateOutput(rr, req)

	assert.Equal(suite.T(), http.StatusBadRequest, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "invalid character")
}

func (suite *WebOutputHandlerSuite) TestCreateOutputWhenRepositoryFails() {
	inputDTO := inputdto.OutputDTO{
		Service:  "test_service",
		Source:   "test_source",
		Provider: "test_provider",
		Data: map[string]interface{}{
			"field1": "value1",
			"field2": "value2",
		},
		Metadata: shareddto.MetadataDTO{
			InputID: "input1",
			Input: shareddto.InputDTO{
				Data:                map[string]interface{}{"key": "value"},
				ProcessingID:        "processing1",
				ProcessingTimestamp: "2021-06-01 00:00:00",
			},
		},
	}

	suite.repoMock.On("Create", mock.AnythingOfType("*entity.Output")).Return(errors.New("repository error"))

	jsonBody, _ := json.Marshal(inputDTO)
	req := httptest.NewRequest(http.MethodPost, "/outputs", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	suite.handler.CreateOutput(rr, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "repository error")
	suite.repoMock.AssertExpectations(suite.T())
}

// Tests for UpdateOutput handler
func (suite *WebOutputHandlerSuite) TestUpdateOutputWhenSuccess() {
	inputDTO := inputdto.OutputDTO{
		Service:  "test_service",
		Source:   "test_source",
		Provider: "test_provider",
		Data: map[string]interface{}{
			"field1": "value1",
			"field2": "value2",
		},
		Metadata: shareddto.MetadataDTO{
			InputID: "input1",
			Input: shareddto.InputDTO{
				Data:                map[string]interface{}{"key": "value"},
				ProcessingID:        "processing1",
				ProcessingTimestamp: "2021-06-01 00:00:00",
			},
		},
	}

	expectedOutput := outputdto.OutputDTO{
		ID:       "1",
		Service:  "test_service",
		Source:   "test_source",
		Provider: "test_provider",
		Data: map[string]interface{}{
			"field1": "value1",
			"field2": "value2",
		},
		Metadata: shareddto.MetadataDTO{
			InputID: "input1",
			Input: shareddto.InputDTO{
				Data:                map[string]interface{}{"key": "value"},
				ProcessingID:        "processing1",
				ProcessingTimestamp: "2021-06-01 00:00:00",
			},
		},
		CreatedAt: "2023-06-01 00:00:00",
		UpdatedAt: "2023-06-01 00:00:00",
	}

	suite.repoMock.On("Update", mock.AnythingOfType("*entity.Output")).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*entity.Output)
		arg.ID = "1"
		arg.Service = "test_service"
		arg.Source = "test_source"
		arg.Provider = "test_provider"
		arg.Data = map[string]interface{}{
			"field1": "value1",
			"field2": "value2",
		}
		arg.Metadata = entity.Metadata{
			InputID: "input1",
			Input: entity.Input{
				Data:                map[string]interface{}{"key": "value"},
				ProcessingID:        "processing1",
				ProcessingTimestamp: "2021-06-01 00:00:00",
			},
		}
		arg.CreatedAt = "2023-06-01 00:00:00"
		arg.UpdatedAt = "2023-06-01 00:00:00"
	})

	jsonBody, _ := json.Marshal(inputDTO)
	req := httptest.NewRequest(http.MethodPut, "/outputs/1", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	suite.handler.UpdateOutput(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutput outputdto.OutputDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutput)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedOutput, actualOutput)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebOutputHandlerSuite) TestUpdateOutputWhenDecodingFails() {
	req := httptest.NewRequest(http.MethodPut, "/outputs/1", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	suite.handler.UpdateOutput(rr, req)

	assert.Equal(suite.T(), http.StatusBadRequest, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "invalid character")
}

func (suite *WebOutputHandlerSuite) TestUpdateOutputWhenRepositoryFails() {
	inputDTO := inputdto.OutputDTO{
		Service:  "test_service",
		Source:   "test_source",
		Provider: "test_provider",
		Data: map[string]interface{}{
			"field1": "value1",
			"field2": "value2",
		},
		Metadata: shareddto.MetadataDTO{
			InputID: "input1",
			Input: shareddto.InputDTO{
				Data:                map[string]interface{}{"key": "value"},
				ProcessingID:        "processing1",
				ProcessingTimestamp: "2021-06-01 00:00:00",
			},
		},
	}

	suite.repoMock.On("Update", mock.AnythingOfType("*entity.Output")).Return(errors.New("repository error"))

	jsonBody, _ := json.Marshal(inputDTO)
	req := httptest.NewRequest(http.MethodPut, "/outputs/1", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	suite.handler.UpdateOutput(rr, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "repository error")
	suite.repoMock.AssertExpectations(suite.T())
}

// Tests for DeleteOutput handler
func (suite *WebOutputHandlerSuite) TestDeleteOutputWhenSuccess() {
	suite.repoMock.On("Delete", "1").Return(nil)

	req := httptest.NewRequest("DELETE", "/outputs/1", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.DeleteOutput(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)
	assert.Equal(suite.T(), "Output deleted successfully", rr.Body.String())
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebOutputHandlerSuite) TestDeleteOutputWhenIDNotProvided() {
	req := httptest.NewRequest("DELETE", "/output/", nil)
	rctx := chi.NewRouteContext()
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.DeleteOutput(rr, req)

	assert.Equal(suite.T(), http.StatusBadRequest, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "ID is required")
}

func (suite *WebOutputHandlerSuite) TestDeleteOutputWhenRepositoryFails() {
	suite.repoMock.On("Delete", "1").Return(errors.New("repository error"))

	req := httptest.NewRequest("DELETE", "/outputs/1", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.DeleteOutput(rr, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "repository error")
	suite.repoMock.AssertExpectations(suite.T())
}

// Tests for ListAllOutputs handler
func (suite *WebOutputHandlerSuite) TestListAllOutputsWhenSuccess() {
	entityOutputs := []*entity.Output{
		{
			ID:       "1",
			Service:  "service1",
			Source:   "source1",
			Provider: "provider",
			Data: map[string]interface{}{
				"field1": "value1",
				"field2": "value2",
			},
			Metadata: entity.Metadata{
				InputID: "input1",
				Input: entity.Input{
					Data:                map[string]interface{}{"key": "value"},
					ProcessingID:        "processing1",
					ProcessingTimestamp: "2021-06-01 00:00:00",
				},
			},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
		{
			ID:       "2",
			Service:  "service2",
			Source:   "source2",
			Provider: "provider",
			Data: map[string]interface{}{
				"field1": "value1",
				"field2": "value2",
			},
			Metadata: entity.Metadata{
				InputID: "input1",
				Input: entity.Input{
					Data:                map[string]interface{}{"key": "value"},
					ProcessingID:        "processing1",
					ProcessingTimestamp: "2021-06-01 00:00:00",
				},
			},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	suite.repoMock.On("FindAll").Return(entityOutputs, nil)

	expectedOutput := []outputdto.OutputDTO{
		{
			ID:       "1",
			Service:  "service1",
			Source:   "source1",
			Provider: "provider",
			Data: map[string]interface{}{
				"field1": "value1",
				"field2": "value2",
			},
			Metadata: shareddto.MetadataDTO{
				InputID: "input1",
				Input: shareddto.InputDTO{
					Data:                map[string]interface{}{"key": "value"},
					ProcessingID:        "processing1",
					ProcessingTimestamp: "2021-06-01 00:00:00",
				},
			},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
		{
			ID:       "2",
			Service:  "service2",
			Source:   "source2",
			Provider: "provider",
			Data: map[string]interface{}{
				"field1": "value1",
				"field2": "value2",
			},
			Metadata: shareddto.MetadataDTO{
				InputID: "input1",
				Input: shareddto.InputDTO{
					Data:                map[string]interface{}{"key": "value"},
					ProcessingID:        "processing1",
					ProcessingTimestamp: "2021-06-01 00:00:00",
				},
			},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/outputs", nil)
	rr := httptest.NewRecorder()

	suite.handler.ListAllOutputs(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutput []outputdto.OutputDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutput)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedOutput, actualOutput)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebOutputHandlerSuite) TestListAllOutputsWhenRepositoryFails() {
	suite.repoMock.On("FindAll").Return(nil, errors.New("repository error"))

	req := httptest.NewRequest(http.MethodGet, "/outputs", nil)
	rr := httptest.NewRecorder()

	suite.handler.ListAllOutputs(rr, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "repository error")
	suite.repoMock.AssertExpectations(suite.T())
}

// Tests for ListOutputByID handler
func (suite *WebOutputHandlerSuite) TestListOutputByIDWhenSuccess() {
	entityOutputs := []*entity.Output{
		{
			ID:       "1",
			Service:  "service1",
			Source:   "source1",
			Provider: "provider",
			Data: map[string]interface{}{
				"field1": "value1",
				"field2": "value2",
			},
			Metadata: entity.Metadata{
				InputID: "input1",
				Input: entity.Input{
					Data:                map[string]interface{}{"key": "value"},
					ProcessingID:        "processing1",
					ProcessingTimestamp: "2021-06-01 00:00:00",
				},
			},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
		{
			ID:       "2",
			Service:  "service2",
			Source:   "source2",
			Provider: "provider",
			Data: map[string]interface{}{
				"field1": "value1",
				"field2": "value2",
			},
			Metadata: entity.Metadata{
				InputID: "input1",
				Input: entity.Input{
					Data:                map[string]interface{}{"key": "value"},
					ProcessingID:        "processing1",
					ProcessingTimestamp: "2021-06-01 00:00:00",
				},
			},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	suite.repoMock.On("FindByID", "1").Return(entityOutputs[0], nil)

	expectedOutput := outputdto.OutputDTO{
		ID:       "1",
		Service:  "service1",
		Source:   "source1",
		Provider: "provider",
		Data: map[string]interface{}{
			"field1": "value1",
			"field2": "value2",
		},
		Metadata: shareddto.MetadataDTO{
			InputID: "input1",
			Input: shareddto.InputDTO{
				Data:                map[string]interface{}{"key": "value"},
				ProcessingID:        "processing1",
				ProcessingTimestamp: "2021-06-01 00:00:00",
			},
		},
		CreatedAt: "2023-06-01T00:00:00Z",
		UpdatedAt: "2023-06-01T00:00:00Z",
	}

	req := httptest.NewRequest(http.MethodGet, "/outputs/1", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListOutputByID(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutput outputdto.OutputDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutput)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedOutput, actualOutput)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebOutputHandlerSuite) TestListOutputByIDWhenIDNotProvided() {
	req := httptest.NewRequest(http.MethodGet, "/output/", nil)
	rctx := chi.NewRouteContext()
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListOutputByID(rr, req)

	assert.Equal(suite.T(), http.StatusBadRequest, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "ID is required")
}

func (suite *WebOutputHandlerSuite) TestListOutputByIDWhenRepositoryFails() {
	suite.repoMock.On("FindByID", "1").Return(nil, errors.New("repository error"))

	req := httptest.NewRequest(http.MethodGet, "/outputs/1", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListOutputByID(rr, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "repository error")
	suite.repoMock.AssertExpectations(suite.T())
}

// Tests for ListOutputsByServiceAndProvider handler
func (suite *WebOutputHandlerSuite) TestListOutputsByServiceAndProviderWhenSuccess() {
	entityOutputs := []*entity.Output{
		{
			ID:       "1",
			Service:  "service1",
			Source:   "source1",
			Provider: "provider",
			Data: map[string]interface{}{
				"field1": "value1",
				"field2": "value2",
			},
			Metadata: entity.Metadata{
				InputID: "input1",
				Input: entity.Input{
					Data:                map[string]interface{}{"key": "value"},
					ProcessingID:        "processing1",
					ProcessingTimestamp: "2021-06-01 00:00:00",
				},
			},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	suite.repoMock.On("FindAllByServiceAndProvider", "test_provider", "test_service").Return(entityOutputs, nil)

	expectedOutput := []outputdto.OutputDTO{
		{
			ID:       "1",
			Service:  "service1",
			Source:   "source1",
			Provider: "provider",
			Data: map[string]interface{}{
				"field1": "value1",
				"field2": "value2",
			},
			Metadata: shareddto.MetadataDTO{
				InputID: "input1",
				Input: shareddto.InputDTO{
					Data:                map[string]interface{}{"key": "value"},
					ProcessingID:        "processing1",
					ProcessingTimestamp: "2021-06-01 00:00:00",
				},
			},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	req := httptest.NewRequest("GET", "/outputs/service/test_service/provider/test_provider", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("service", "test_service")
	rctx.URLParams.Add("provider", "test_provider")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListOutputsByServiceAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutput []outputdto.OutputDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutput)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedOutput, actualOutput)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebOutputHandlerSuite) TestListOutputsByServiceAndProviderWhenServiceOrProviderNotProvided() {
	req := httptest.NewRequest("GET", "/outputs/service/test_service", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("service", "test_service")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListOutputsByServiceAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusBadRequest, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "Service and provider are required")
}

func (suite *WebOutputHandlerSuite) TestListOutputsByServiceAndProviderWhenRepositoryFails() {
	suite.repoMock.On("FindAllByServiceAndProvider", "test_provider", "test_service").Return(nil, errors.New("repository error"))

	req := httptest.NewRequest("GET", "/outputs/service/test_service/provider/test_provider", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("service", "test_service")
	rctx.URLParams.Add("provider", "test_provider")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListOutputsByServiceAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "repository error")
	suite.repoMock.AssertExpectations(suite.T())
}

// Tests for ListOutputsBySourceAndProvider handler
func (suite *WebOutputHandlerSuite) TestListOutputsBySourceAndProviderWhenSuccess() {
	entityOutputs := []*entity.Output{
		{
			ID:       "1",
			Service:  "service1",
			Source:   "source1",
			Provider: "provider",
			Data: map[string]interface{}{
				"field1": "value1",
				"field2": "value2",
			},
			Metadata: entity.Metadata{
				InputID: "input1",
				Input: entity.Input{
					Data:                map[string]interface{}{"key": "value"},
					ProcessingID:        "processing1",
					ProcessingTimestamp: "2021-06-01 00:00:00",
				},
			},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	suite.repoMock.On("FindAllBySourceAndProvider", "test_provider", "test_source").Return(entityOutputs, nil)

	expectedOutput := []outputdto.OutputDTO{
		{
			ID:       "1",
			Service:  "service1",
			Source:   "source1",
			Provider: "provider",
			Data: map[string]interface{}{
				"field1": "value1",
				"field2": "value2",
			},
			Metadata: shareddto.MetadataDTO{
				InputID: "input1",
				Input: shareddto.InputDTO{
					Data:                map[string]interface{}{"key": "value"},
					ProcessingID:        "processing1",
					ProcessingTimestamp: "2021-06-01 00:00:00",
				},
			},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	req := httptest.NewRequest("GET", "/outputs/source/test_source/provider/test_provider", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("source", "test_source")
	rctx.URLParams.Add("provider", "test_provider")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListOutputsBySourceAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutput []outputdto.OutputDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutput)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedOutput, actualOutput)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebOutputHandlerSuite) TestListOutputsBySourceAndProviderWhenRepositoryFails() {
	suite.repoMock.On("FindAllBySourceAndProvider", "test_provider", "test_source").Return(nil, errors.New("repository error"))

	req := httptest.NewRequest("GET", "/outputs/source/test_source/provider/test_provider", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("source", "test_source")
	rctx.URLParams.Add("provider", "test_provider")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListOutputsBySourceAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "repository error")
	suite.repoMock.AssertExpectations(suite.T())
}

// Tests for ListOutputsByServiceAndSourceAndProvider handler
func (suite *WebOutputHandlerSuite) TestListOutputsByServiceAndSourceAndProviderWhenSuccess() {
	entityOutputs := []*entity.Output{
		{
			ID:       "1",
			Service:  "service1",
			Source:   "source1",
			Provider: "provider",
			Data: map[string]interface{}{
				"field1": "value1",
				"field2": "value2",
			},
			Metadata: entity.Metadata{
				InputID: "input1",
				Input: entity.Input{
					Data:                map[string]interface{}{"key": "value"},
					ProcessingID:        "processing1",
					ProcessingTimestamp: "2021-06-01 00:00:00",
				},
			},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	suite.repoMock.On("FindAllByServiceAndSourceAndProvider", "test_service", "test_source", "test_provider").Return(entityOutputs, nil)

	expectedOutput := []outputdto.OutputDTO{
		{
			ID:       "1",
			Service:  "service1",
			Source:   "source1",
			Provider: "provider",
			Data: map[string]interface{}{
				"field1": "value1",
				"field2": "value2",
			},
			Metadata: shareddto.MetadataDTO{
				InputID: "input1",
				Input: shareddto.InputDTO{
					Data:                map[string]interface{}{"key": "value"},
					ProcessingID:        "processing1",
					ProcessingTimestamp: "2021-06-01 00:00:00",
				},
			},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	req := httptest.NewRequest("GET", "/outputs/service/test_service/source/test_source/provider/test_provider", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("service", "test_service")
	rctx.URLParams.Add("source", "test_source")
	rctx.URLParams.Add("provider", "test_provider")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListOutputsByServiceAndSourceAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutput []outputdto.OutputDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutput)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedOutput, actualOutput)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebOutputHandlerSuite) TestListOutputsByServiceAndSourceAndProviderWhenServiceOrSourceOrProviderNotProvided() {
	req := httptest.NewRequest("GET", "/outputs/service/test_service/source/test_source", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("service", "test_service")
	rctx.URLParams.Add("source", "test_source")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListOutputsByServiceAndSourceAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusBadRequest, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "Service, source and provider are required")
}

func (suite *WebOutputHandlerSuite) TestListOutputsByServiceAndSourceAndProviderWhenRepositoryFails() {
	suite.repoMock.On("FindAllByServiceAndSourceAndProvider", "test_service", "test_source", "test_provider").Return(nil, errors.New("repository error"))

	req := httptest.NewRequest("GET", "/outputs/service/test_service/source/test_source/provider/test_provider", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("provider", "test_provider")
	rctx.URLParams.Add("service", "test_service")
	rctx.URLParams.Add("source", "test_source")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListOutputsByServiceAndSourceAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "repository error")
	suite.repoMock.AssertExpectations(suite.T())
}
