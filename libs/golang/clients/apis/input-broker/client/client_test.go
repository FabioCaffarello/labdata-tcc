package client

import (
	"context"
	"encoding/json"
	inputdto "libs/golang/ddd/dtos/input-broker/input"
	outputdto "libs/golang/ddd/dtos/input-broker/output"
	shareddto "libs/golang/ddd/dtos/input-broker/shared"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ClientTestSuite struct {
	suite.Suite
	client     *Client
	mockServer *httptest.Server
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}

func (suite *ClientTestSuite) SetupTest() {
	// Create a mock HTTP server
	suite.mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/input" && r.Method == http.MethodPost:
			var inputInput inputdto.InputDTO
			if err := json.NewDecoder(r.Body).Decode(&inputInput); err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}
			inputOutput := outputdto.InputDTO{
				ID:        "1",
				Data:      inputInput.Data,
				Metadata:  shareddto.MetadataDTO{Provider: inputInput.Provider, Service: inputInput.Service, Source: inputInput.Source},
				Status:    shareddto.StatusDTO{Code: 200, Detail: "Success"},
				CreatedAt: "2023-06-01T00:00:00Z",
				UpdatedAt: "2023-06-01T00:00:00Z",
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(inputOutput)

		default:
			http.NotFound(w, r)
		}
	}))

	// Initialize the client with the mock server's URL
	suite.client = &Client{
		ctx:     context.Background(),
		baseURL: suite.mockServer.URL,
		timeout: apiTimeout,
	}
}

func (suite *ClientTestSuite) TearDownTest() {
	suite.mockServer.Close()
}

func (suite *ClientTestSuite) TestCreateInputWhenSuccess() {
	inputInput := inputdto.InputDTO{
		Provider: "test_provider",
		Service:  "test_service",
		Source:   "test_source",
		Data:     map[string]interface{}{"key": "value"},
	}

	expectedOutput := outputdto.InputDTO{
		ID:        "1",
		Data:      map[string]interface{}{"key": "value"},
		Metadata:  shareddto.MetadataDTO{Provider: "test_provider", Service: "test_service", Source: "test_source"},
		Status:    shareddto.StatusDTO{Code: 200, Detail: "Success"},
		CreatedAt: "2023-06-01T00:00:00Z",
		UpdatedAt: "2023-06-01T00:00:00Z",
	}

	inputOutput, err := suite.client.CreateInput(inputInput)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, inputOutput)
}
