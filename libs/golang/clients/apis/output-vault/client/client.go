package client

import (
	"context"
	inputdto "libs/golang/ddd/dtos/output-vault/input"
	outputdto "libs/golang/ddd/dtos/output-vault/output"
	"libs/golang/shared/go-request/requests"
	"net/http"
	"time"
)

var (
	apiTimeout     = 100 * time.Millisecond
	defaultHeaders = map[string]string{"Content-Type": "application/json"}
)

// Client represents the output vault client.
type Client struct {
	ctx     context.Context
	baseURL string
	timeout time.Duration
}

// NewClient initializes a new output vault client.
func NewClient() *Client {
	return &Client{
		ctx:     context.Background(),
		baseURL: "http://output-handler:8000",
		timeout: apiTimeout,
	}
}

// CreateOutput sends a request to create a new output.
//
// Parameters:
//   - outputInput: The output data transfer object.
//
// Returns:
//   - outputdto.OutputDTO: The created output data transfer object.
//   - error: An error if the request fails.
func (c *Client) CreateOutput(outputInput inputdto.OutputDTO) (outputdto.OutputDTO, error) {
	pathParams := []string{"output"}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, outputInput, defaultHeaders, http.MethodPost)
	if err != nil {
		return outputdto.OutputDTO{}, err
	}

	var outputOutput outputdto.OutputDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &outputOutput, c.timeout)
	if err != nil {
		return outputdto.OutputDTO{}, err
	}

	return outputOutput, nil
}

// UpdateOutput sends a request to update an existing output.
//
// Parameters:
//   - outputInput: The output data transfer object.
//
// Returns:
//   - outputdto.OutputDTO: The updated output data transfer object.
//   - error: An error if the request fails.
func (c *Client) UpdateOutput(outputInput inputdto.OutputDTO) (outputdto.OutputDTO, error) {
	pathParams := []string{"output"}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, outputInput, defaultHeaders, http.MethodPut)
	if err != nil {
		return outputdto.OutputDTO{}, err
	}

	var outputOutput outputdto.OutputDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &outputOutput, c.timeout)
	if err != nil {
		return outputdto.OutputDTO{}, err
	}

	return outputOutput, nil
}

// ListAllOutputs sends a request to retrieve all outputs.
//
// Returns:
//   - []outputdto.OutputDTO: A slice of output data transfer objects.
//   - error: An error if the request fails.
func (c *Client) ListAllOutputs() ([]outputdto.OutputDTO, error) {
	pathParams := []string{"output"}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, nil, defaultHeaders, http.MethodGet)
	if err != nil {
		return nil, err
	}

	var outputList []outputdto.OutputDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &outputList, c.timeout)
	if err != nil {
		return nil, err
	}

	return outputList, nil
}

// ListOutputByID sends a request to retrieve an output by its ID.
//
// Parameters:
//   - id: The ID of the output.
//
// Returns:
//   - outputdto.OutputDTO: The output data transfer object.
//   - error: An error if the request fails.
func (c *Client) ListOutputByID(id string) (outputdto.OutputDTO, error) {
	pathParams := []string{"output", id}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, nil, defaultHeaders, http.MethodGet)
	if err != nil {
		return outputdto.OutputDTO{}, err
	}

	var outputOutput outputdto.OutputDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &outputOutput, c.timeout)
	if err != nil {
		return outputdto.OutputDTO{}, err
	}

	return outputOutput, nil
}

// DeleteOutput sends a request to delete an output by its ID.
//
// Parameters:
//   - id: The ID of the output.
//
// Returns:
//   - error: An error if the request fails.
func (c *Client) DeleteOutput(id string) error {
	pathParams := []string{"output", id}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, nil, defaultHeaders, http.MethodDelete)
	if err != nil {
		return err
	}

	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, nil, c.timeout)
	if err != nil {
		return err
	}

	return nil
}

// ListOutputsByServiceAndProvider sends a request to retrieve outputs by service and provider.
//
// Parameters:
//   - service: The service name.
//   - provider: The provider name.
//
// Returns:
//   - []outputdto.OutputDTO: A slice of output data transfer objects.
//   - error: An error if the request fails.
func (c *Client) ListOutputsByServiceAndProvider(service, provider string) ([]outputdto.OutputDTO, error) {
	pathParams := []string{"output", "provider", provider, "service", service}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, nil, defaultHeaders, http.MethodGet)
	if err != nil {
		return nil, err
	}

	var outputList []outputdto.OutputDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &outputList, c.timeout)
	if err != nil {
		return nil, err
	}

	return outputList, nil
}

// ListOutputsBySourceAndProvider sends a request to retrieve outputs by source and provider.
//
// Parameters:
//   - source: The source name.
//   - provider: The provider name.
//
// Returns:
//   - []outputdto.OutputDTO: A slice of output data transfer objects.
//   - error: An error if the request fails.
func (c *Client) ListOutputsBySourceAndProvider(source, provider string) ([]outputdto.OutputDTO, error) {
	pathParams := []string{"output", "provider", provider, "source", source}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, nil, defaultHeaders, http.MethodGet)
	if err != nil {
		return nil, err
	}

	var outputList []outputdto.OutputDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &outputList, c.timeout)
	if err != nil {
		return nil, err
	}

	return outputList, nil
}

// ListOutputsByServiceAndSourceAndProvider sends a request to retrieve outputs by service, source, and provider.
//
// Parameters:
//   - service: The service name.
//   - source: The source name.
//   - provider: The provider name.
//
// Returns:
//   - []outputdto.OutputDTO: A slice of output data transfer objects.
//   - error: An error if the request fails.
func (c *Client) ListOutputsByServiceAndSourceAndProvider(service, source, provider string) ([]outputdto.OutputDTO, error) {
	pathParams := []string{"output", "provider", provider, "service", service, "source", source}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, nil, defaultHeaders, http.MethodGet)
	if err != nil {
		return nil, err
	}

	var outputList []outputdto.OutputDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &outputList, c.timeout)
	if err != nil {
		return nil, err
	}

	return outputList, nil
}
