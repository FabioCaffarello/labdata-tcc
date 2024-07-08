package client

import (
	"context"
	"fmt"
	inputdto "libs/golang/ddd/dtos/input-broker/input"
	outputdto "libs/golang/ddd/dtos/input-broker/output"
	shareddto "libs/golang/ddd/dtos/input-broker/shared"
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

// UpdateInput sends a request to update an existing input.
//
// Parameters:
//   - id: The input ID.
//   - inputInput: The input data transfer object.
//
// Returns:
//   - outputdto.InputDTO: The updated input data transfer object.
//   - error: An error if the request fails.
func (c *Client) UpdateInput(id string, inputInput inputdto.InputDTO) (outputdto.InputDTO, error) {
	pathParams := []string{"input", id}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, inputInput, defaultHeaders, http.MethodPut)
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

// DeleteInput sends a request to delete an input by ID.
//
// Parameters:
//   - id: The input ID.
//
// Returns:
//   - error: An error if the request fails.
func (c *Client) DeleteInput(id string) error {
	pathParams := []string{"input", id}

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

// ListAllInputs sends a request to retrieve all inputs.
//
// Returns:
//   - []outputdto.InputDTO: A list of input data transfer objects.
//   - error: An error if the request fails.
func (c *Client) ListAllInputs() ([]outputdto.InputDTO, error) {
	pathParams := []string{"input"}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, nil, defaultHeaders, http.MethodGet)
	if err != nil {
		return nil, err
	}

	var inputs []outputdto.InputDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &inputs, c.timeout)
	if err != nil {
		return nil, err
	}

	return inputs, nil
}

// GetInputByID sends a request to retrieve an input by ID.
//
// Parameters:
//   - id: The input ID.
//
// Returns:
//   - outputdto.InputDTO: The input data transfer object.
//   - error: An error if the request fails.
func (c *Client) GetInputByID(id string) (outputdto.InputDTO, error) {
	pathParams := []string{"input", id}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, nil, defaultHeaders, http.MethodGet)
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

// UpdateInputStatus sends a request to update the status of an existing input.
//
// Parameters:
//   - id: The input ID.
//   - status: The status data transfer object.
//
// Returns:
//   - outputdto.InputDTO: The updated input data transfer object.
//   - error: An error if the request fails.
func (c *Client) UpdateInputStatus(id string, status shareddto.StatusDTO) (outputdto.InputDTO, error) {
	pathParams := []string{"input", id, "status"}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, status, defaultHeaders, http.MethodPut)
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

// ListInputsByServiceAndProvider sends a request to retrieve inputs by service and provider.
//
// Parameters:
//   - service: The service name.
//   - provider: The provider name.
//
// Returns:
//   - []outputdto.InputDTO: A list of input data transfer objects.
//   - error: An error if the request fails.
func (c *Client) ListInputsByServiceAndProvider(service, provider string) ([]outputdto.InputDTO, error) {
	pathParams := []string{"input", "provider", provider, "service", service}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, nil, defaultHeaders, http.MethodGet)
	if err != nil {
		return nil, err
	}

	var inputs []outputdto.InputDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &inputs, c.timeout)
	if err != nil {
		return nil, err
	}

	return inputs, nil
}

// ListInputsBySourceAndProvider sends a request to retrieve inputs by source and provider.
//
// Parameters:
//   - source: The source name.
//   - provider: The provider name.
//
// Returns:
//   - []outputdto.InputDTO: A list of input data transfer objects.
//   - error: An error if the request fails.
func (c *Client) ListInputsBySourceAndProvider(source, provider string) ([]outputdto.InputDTO, error) {
	pathParams := []string{"input", "provider", provider, "source", source}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, nil, defaultHeaders, http.MethodGet)
	if err != nil {
		return nil, err
	}

	var inputs []outputdto.InputDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &inputs, c.timeout)
	if err != nil {
		return nil, err
	}

	return inputs, nil
}

// ListInputsByServiceAndSourceAndProvider sends a request to retrieve inputs by service, source, and provider.
//
// Parameters:
//   - service: The service name.
//   - source: The source name.
//   - provider: The provider name.
//
// Returns:
//   - []outputdto.InputDTO: A list of input data transfer objects.
//   - error: An error if the request fails.
func (c *Client) ListInputsByServiceAndSourceAndProvider(service, source, provider string) ([]outputdto.InputDTO, error) {
	pathParams := []string{"input", "provider", provider, "service", service, "source", source}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, nil, defaultHeaders, http.MethodGet)
	if err != nil {
		return nil, err
	}

	var inputs []outputdto.InputDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &inputs, c.timeout)
	if err != nil {
		return nil, err
	}

	return inputs, nil
}

// ListInputsByStatusAndProvider sends a request to retrieve inputs by status and provider.
//
// Parameters:
//   - status: The status code.
//   - provider: The provider name.
//
// Returns:
//   - []outputdto.InputDTO: A list of input data transfer objects.
//   - error: An error if the request fails.
func (c *Client) ListInputsByStatusAndProvider(status int, provider string) ([]outputdto.InputDTO, error) {
	pathParams := []string{"input", "provider", provider, "status", fmt.Sprintf("%d", status)}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, nil, defaultHeaders, http.MethodGet)
	if err != nil {
		return nil, err
	}

	var inputs []outputdto.InputDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &inputs, c.timeout)
	if err != nil {
		return nil, err
	}

	return inputs, nil
}

// ListInputsByStatusAndServiceAndProvider sends a request to retrieve inputs by status, service, and provider.
//
// Parameters:
//   - status: The status code.
//   - service: The service name.
//   - provider: The provider name.
//
// Returns:
//   - []outputdto.InputDTO: A list of input data transfer objects.
//   - error: An error if the request fails.
func (c *Client) ListInputsByStatusAndServiceAndProvider(status int, service, provider string) ([]outputdto.InputDTO, error) {
	pathParams := []string{"input", "provider", provider, "service", service, "status", fmt.Sprintf("%d", status)}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, nil, defaultHeaders, http.MethodGet)
	if err != nil {
		return nil, err
	}

	var inputs []outputdto.InputDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &inputs, c.timeout)
	if err != nil {
		return nil, err
	}

	return inputs, nil
}

// ListInputsByStatusAndSourceAndProvider sends a request to retrieve inputs by status, source, and provider.
//
// Parameters:
//   - status: The status code.
//   - source: The source name.
//   - provider: The provider name.
//
// Returns:
//   - []outputdto.InputDTO: A list of input data transfer objects.
//   - error: An error if the request fails.
func (c *Client) ListInputsByStatusAndSourceAndProvider(status int, source, provider string) ([]outputdto.InputDTO, error) {
	pathParams := []string{"input", "provider", provider, "source", source, "status", fmt.Sprintf("%d", status)}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, nil, defaultHeaders, http.MethodGet)
	if err != nil {
		return nil, err
	}

	var inputs []outputdto.InputDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &inputs, c.timeout)
	if err != nil {
		return nil, err
	}

	return inputs, nil
}

// ListInputsByStatusAndServiceAndSourceAndProvider sends a request to retrieve inputs by status, service, source, and provider.
//
// Parameters:
//   - status: The status code.
//   - service: The service name.
//   - source: The source name.
//   - provider: The provider name.
//
// Returns:
//   - []outputdto.InputDTO: A list of input data transfer objects.
//   - error: An error if the request fails.
func (c *Client) ListInputsByStatusAndServiceAndSourceAndProvider(status int, service, source, provider string) ([]outputdto.InputDTO, error) {
	pathParams := []string{"input", "provider", provider, "service", service, "source", source, "status", fmt.Sprintf("%d", status)}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, nil, defaultHeaders, http.MethodGet)
	if err != nil {
		return nil, err
	}

	var inputs []outputdto.InputDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &inputs, c.timeout)
	if err != nil {
		return nil, err
	}

	return inputs, nil
}
