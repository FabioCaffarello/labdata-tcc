package client

import (
	"context"
	inputdto "libs/golang/ddd/dtos/input-broker/input"
	outputdto "libs/golang/ddd/dtos/input-broker/output"
	"libs/golang/shared/go-request/requests"
	"net/http"
	"time"
)

var (
	apiTimeout     = 100 * time.Millisecond
	defaultHeaders = map[string]string{"Content-Type": "application/json"}
)

// Client represents the input broker client.
type Client struct {
	ctx     context.Context
	baseURL string
	timeout time.Duration
}

// NewClient initializes a new input broker client.
func NewClient() *Client {
	return &Client{
		ctx:     context.Background(),
		baseURL: "http://input-broker:8000",
		timeout: apiTimeout,
	}
}

// CreateInput sends a request to create a new input.
//
// Parameters:
//   - inputInput: The input data transfer object.
//
// Returns:
//   - outputdto.InputDTO: The created input data transfer object.
//   - error: An error if the request fails.
func (c *Client) CreateInput(inputInput inputdto.InputDTO) (outputdto.InputDTO, error) {
	pathParams := []string{"input"}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, inputInput, defaultHeaders, http.MethodPost)
	if err != nil {
		return outputdto.InputDTO{}, err
	}

	var inputOutput outputdto.InputDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &inputOutput, c.timeout)
	if err != nil {
		return outputdto.InputDTO{}, err
	}

	return inputOutput, nil
}
