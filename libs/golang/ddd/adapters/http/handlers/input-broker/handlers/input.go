package handlers

import (
	"encoding/json"
	"libs/golang/ddd/domain/entities/input-broker/entity"
	inputdto "libs/golang/ddd/dtos/input-broker/input"
	shareddto "libs/golang/ddd/dtos/input-broker/shared"
	"libs/golang/ddd/usecases/input-broker/usecase"
	events "libs/golang/shared/go-events/amqp_events"
	typetools "libs/golang/shared/type-tools"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// WebInputHandler handles HTTP requests for input-related operations.
type WebInputHandler struct {
	InputRepository   entity.InputRepositoryInterface // Interface for input repository operations.
	EventDispatcher   events.EventDispatcherInterface // Interface for event dispatching.
	InputCreatedEvent events.EventInterface           // Event interface for input creation event.
}

// NewWebInputHandler creates a new instance of WebInputHandler with the provided dependencies.
//
// Parameters:
//   - inputRepository: Interface for input repository operations.
//   - eventDispatcher: Interface for event dispatching.
//   - inputCreatedEvent: Event interface for input creation event.
//
// Returns:
//   - A new instance of WebInputHandler.
func NewWebInputHandler(
	inputRepository entity.InputRepositoryInterface,
	eventDispatcher events.EventDispatcherInterface,
	inputCreatedEvent events.EventInterface,
) *WebInputHandler {
	return &WebInputHandler{
		InputRepository:   inputRepository,
		EventDispatcher:   eventDispatcher,
		InputCreatedEvent: inputCreatedEvent,
	}
}

// CreateInput handles the creation of a new input entity.
//
// This function decodes the request body into an InputDTO, validates it,
// and then creates a new input entity using the use case. If successful,
// it responds with the created input entity as JSON. If there are errors,
// appropriate HTTP error responses are returned.
//
// Parameters:
//   - w: HTTP Response Writer to write the response.
//   - r: HTTP Request containing the input data.
//
// Responses:
//   - 200 OK: If the input entity is created successfully, the response contains the created input entity as JSON.
//   - 400 Bad Request: If there is an error decoding the request body.
//   - 500 Internal Server Error: If there is an error creating the input entity or encoding the response.
func (h *WebInputHandler) CreateInput(w http.ResponseWriter, r *http.Request) {
	var dto inputdto.InputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createInputUseCase := usecase.NewCreateInputUseCase(h.InputRepository, h.InputCreatedEvent, h.EventDispatcher)
	inputCreated, err := createInputUseCase.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(inputCreated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// UpdateInput handles the update of an existing input entity.
//
// This function decodes the request body into an InputDTO, validates it,
// and then updates the input entity using the use case. If successful,
// it responds with the updated input entity as JSON. If there are errors,
// appropriate HTTP error responses are returned.
//
// Parameters:
//   - w: HTTP Response Writer to write the response.
//   - r: HTTP Request containing the input ID and data.
//
// Responses:
//   - 200 OK: If the input entity is updated successfully, the response contains the updated input entity as JSON.
//   - 400 Bad Request: If there is an error decoding the request body.
//   - 500 Internal Server Error: If there is an error updating the input entity or encoding the response.
func (h *WebInputHandler) UpdateInput(w http.ResponseWriter, r *http.Request) {
	var dto inputdto.InputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updateInputUseCase := usecase.NewUpdateInputUseCase(h.InputRepository)
	inputUpdated, err := updateInputUseCase.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(inputUpdated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// DeleteInput handles the deletion of an existing input entity.
//
// This function extracts the input ID from the request URL,
// and then deletes the input entity using the use case. If successful,
// it responds with a success message. If there are errors,
// appropriate HTTP error responses are returned.
//
// Parameters:
//   - w: HTTP Response Writer to write the response.
//   - r: HTTP Request containing the input ID.
//
// Responses:
//   - 200 OK: If the input entity is deleted successfully, the response contains a success message.
//   - 500 Internal Server Error: If there is an error deleting the input entity.
func (h *WebInputHandler) DeleteInput(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	deleteInputUseCase := usecase.NewDeleteInputUseCase(h.InputRepository)
	err := deleteInputUseCase.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Input deleted successfully"))
}

// ListAllInputs handles the retrieval of all input entities.
//
// This function retrieves all input entities using the use case,
// and then responds with the list of input entities as JSON. If there are errors,
// appropriate HTTP error responses are returned.
//
// Parameters:
//   - w: HTTP Response Writer to write the response.
//   - r: HTTP Request.
//
// Responses:
//   - 200 OK: If the input entities are retrieved successfully, the response contains the list of input entities as JSON.
//   - 500 Internal Server Error: If there is an error retrieving the input entities or encoding the response.
func (h *WebInputHandler) ListAllInputs(w http.ResponseWriter, r *http.Request) {
	listInputsUseCase := usecase.NewListAllInputUseCase(h.InputRepository)
	inputs, err := listInputsUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(inputs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ListInputByID handles the retrieval of an input entity by ID.
//
// This function extracts the input ID from the request URL,
// and then retrieves the input entity using the use case. If successful,
// it responds with the input entity as JSON. If there are errors,
// appropriate HTTP error responses are returned.
//
// Parameters:
//   - w: HTTP Response Writer to write the response.
//   - r: HTTP Request containing the input ID.
//
// Responses:
//   - 200 OK: If the input entity is retrieved successfully, the response contains the input entity as JSON.
//   - 500 Internal Server Error: If there is an error retrieving the input entity or encoding the response.
func (h *WebInputHandler) ListInputByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	getInputUseCase := usecase.NewListOneByIDInputUseCase(h.InputRepository)
	input, err := getInputUseCase.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ListInputsByServiceAndProvider handles the retrieval of input entities by service and provider.
//
// This function extracts the service and provider from the request URL,
// and then retrieves the input entities using the use case. If successful,
// it responds with the list of input entities as JSON. If there are errors,
// appropriate HTTP error responses are returned.
//
// Parameters:
//   - w: HTTP Response Writer to write the response.
//   - r: HTTP Request containing the service and provider.
//
// Responses:
//   - 200 OK: If the input entities are retrieved successfully, the response contains the list of input entities as JSON.
//   - 400 Bad Request: If the service or provider is missing.
func (h *WebInputHandler) ListInputsByServiceAndProvider(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	service := chi.URLParam(r, "service")
	if service == "" || provider == "" {
		http.Error(w, "Service and provider are required", http.StatusBadRequest)
		return
	}

	listInputsUseCase := usecase.NewListAllByServiceAndProviderInputUseCase(h.InputRepository)
	inputs, err := listInputsUseCase.Execute(provider, service)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(inputs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ListInputsBySourceAndProvider handles the retrieval of input entities by source and provider.
//
// This function extracts the source and provider from the request URL,
// and then retrieves the input entities using the use case. If successful,
// it responds with the list of input entities as JSON. If there are errors,
// appropriate HTTP error responses are returned.
//
// Parameters:
//   - w: HTTP Response Writer to write the response.
//   - r: HTTP Request containing the source and provider.
//
// Responses:
//   - 200 OK: If the input entities are retrieved successfully, the response contains the list of input entities as JSON.
//   - 400 Bad Request: If the source or provider is missing.
//   - 500 Internal Server Error: If there is an error retrieving the input entities or encoding the response.
func (h *WebInputHandler) ListInputsBySourceAndProvider(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	source := chi.URLParam(r, "source")
	if source == "" || provider == "" {
		http.Error(w, "Source and provider are required", http.StatusBadRequest)
		return
	}

	listInputsUseCase := usecase.NewListAllBySourceAndProviderInputUseCase(h.InputRepository)
	inputs, err := listInputsUseCase.Execute(provider, source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(inputs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ListInputsByServiceAndSourceAndProvider handles the retrieval of input entities by service, source, and provider.
//
// This function extracts the service, source, and provider from the request URL,
// and then retrieves the input entities using the use case. If successful,
// it responds with the list of input entities as JSON. If there are errors,
// appropriate HTTP error responses are returned.
//
// Parameters:
//   - w: HTTP Response Writer to write the response.
//   - r: HTTP Request containing the service, source, and provider.
//
// Responses:
//   - 200 OK: If the input entities are retrieved successfully, the response contains the list of input entities as JSON.
//   - 400 Bad Request: If the service, source, or provider is missing.
//   - 500 Internal Server Error: If there is an error retrieving the input entities or encoding the response.
func (h *WebInputHandler) ListInputsByServiceAndSourceAndProvider(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	service := chi.URLParam(r, "service")
	source := chi.URLParam(r, "source")

	if service == "" || source == "" || provider == "" {
		http.Error(w, "Service, source, and provider are required", http.StatusBadRequest)
		return
	}

	listInputsUseCase := usecase.NewListAllByServiceAndSourceAndProviderInputUseCase(h.InputRepository)
	inputs, err := listInputsUseCase.Execute(provider, service, source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(inputs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ListInputsByStatusAndProvider handles the retrieval of input entities by status and provider.
//
// This function extracts the status and provider from the request URL,
// and then retrieves the input entities using the use case. If successful,
// it responds with the list of input entities as JSON. If there are errors,
// appropriate HTTP error responses are returned.
//
// Parameters:
//   - w: HTTP Response Writer to write the response.
//   - r: HTTP Request containing the status and provider.
//
// Responses:
//   - 200 OK: If the input entities are retrieved successfully, the response contains the list of input entities as JSON.
//   - 400 Bad Request: If the status or provider is missing.
//   - 500 Internal Server Error: If there is an error retrieving the input entities or encoding the response.
func (h *WebInputHandler) ListInputsByStatusAndProvider(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	statusStr := chi.URLParam(r, "status")
	if statusStr == "" || provider == "" {
		http.Error(w, "Status and provider are required", http.StatusBadRequest)
		return
	}
	status, err := typetools.ParseInt(statusStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	listInputsUseCase := usecase.NewListAllByStatusAndProviderInputUseCase(h.InputRepository)
	inputs, err := listInputsUseCase.Execute(provider, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(inputs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ListInputsByStatusAndServiceAndProvider handles the retrieval of input entities by status, service, and provider.
//
// This function extracts the status, service, and provider from the request URL,
// and then retrieves the input entities using the use case. If successful,
// it responds with the list of input entities as JSON. If there are errors,
// appropriate HTTP error responses are returned.
//
// Parameters:
//   - w: HTTP Response Writer to write the response.
//   - r: HTTP Request containing the status, service, and provider.
//
// Responses:
//   - 200 OK: If the input entities are retrieved successfully, the response contains the list of input entities as JSON.
//   - 400 Bad Request: If the status, service, or provider is missing.
//   - 500 Internal Server Error: If there is an error retrieving the input entities or encoding the response.
func (h *WebInputHandler) ListInputsByStatusAndServiceAndProvider(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	service := chi.URLParam(r, "service")
	statusStr := chi.URLParam(r, "status")
	if statusStr == "" || provider == "" || service == "" {
		http.Error(w, "Status, provider and service are required", http.StatusBadRequest)
		return
	}
	status, err := typetools.ParseInt(statusStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	listInputsUseCase := usecase.NewListAllByStatusAndServiceAndProviderInputUseCase(h.InputRepository)
	inputs, err := listInputsUseCase.Execute(provider, service, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(inputs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ListInputsByStatusAndSourceAndProvider handles the retrieval of input entities by status, source, and provider.
//
// This function extracts the status, source, and provider from the request URL,
// and then retrieves the input entities using the use case. If successful,
// it responds with the list of input entities as JSON. If there are errors,
// appropriate HTTP error responses are returned.
//
// Parameters:
//   - w: HTTP Response Writer to write the response.
//   - r: HTTP Request containing the status, source, and provider.
//
// Responses:
//   - 200 OK: If the input entities are retrieved successfully, the response contains the list of input entities as JSON.
//   - 400 Bad Request: If the status, source, or provider is missing.
//   - 500 Internal Server Error: If there is an error retrieving the input entities or encoding the response.
func (h *WebInputHandler) ListInputsByStatusAndSourceAndProvider(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	source := chi.URLParam(r, "source")
	statusStr := chi.URLParam(r, "status")
	if statusStr == "" || provider == "" || source == "" {
		http.Error(w, "Status, provider and source are required", http.StatusBadRequest)
		return
	}
	status, err := typetools.ParseInt(statusStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	listInputsUseCase := usecase.NewListAllByStatusAndSourceAndProviderInputUseCase(h.InputRepository)
	inputs, err := listInputsUseCase.Execute(provider, source, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(inputs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ListInputsByStatusAndServiceAndSourceAndProvider handles the retrieval of input entities by status, service, source, and provider.
//
// This function extracts the status, service, source, and provider from the request URL,
// and then retrieves the input entities using the use case. If successful,
// it responds with the list of input entities as JSON. If there are errors,
// appropriate HTTP error responses are returned.
//
// Parameters:
//   - w: HTTP Response Writer to write the response.
//   - r: HTTP Request containing the status, service, source, and provider.
//
// Responses:
//   - 200 OK: If the input entities are retrieved successfully, the response contains the list of input entities as JSON.
//   - 400 Bad Request: If the status, service, source, or provider is missing.
//   - 500 Internal Server Error: If there is an error retrieving the input entities or encoding the response.
func (h *WebInputHandler) ListInputsByStatusAndServiceAndSourceAndProvider(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	service := chi.URLParam(r, "service")
	source := chi.URLParam(r, "source")
	statusStr := chi.URLParam(r, "status")
	if statusStr == "" || provider == "" || service == "" || source == "" {
		http.Error(w, "Status, provider, service and source are required", http.StatusBadRequest)
		return
	}
	status, err := typetools.ParseInt(statusStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	listInputsUseCase := usecase.NewListAllByStatusAndServiceAndSourceAndProviderInputUseCaseSuite(h.InputRepository)
	inputs, err := listInputsUseCase.Execute(provider, service, source, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(inputs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// UpdateInputStatus handles the update of an existing input entity's status.
//
// This function decodes the request body into a StatusDTO, validates it,
// and then updates the status of the input entity using the use case. If successful,
// it responds with the updated input entity as JSON. If there are errors,
// appropriate HTTP error responses are returned.
//
// Parameters:
//   - w: HTTP Response Writer to write the response.
//   - r: HTTP Request containing the input ID and status data.
//
// Responses:
//   - 200 OK: If the input entity status is updated successfully, the response contains the updated input entity as JSON.
//   - 400 Bad Request: If there is an error decoding the request body.
//   - 500 Internal Server Error: If there is an error updating the input entity status or encoding the response.
func (h *WebInputHandler) UpdateInputStatus(w http.ResponseWriter, r *http.Request) {
	var dto shareddto.StatusDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "missing input ID", http.StatusBadRequest)
		return
	}
	updateStatusInputUseCase := usecase.NewUpdateStatusInputUseCase(h.InputRepository)
	inputUpdated, err := updateStatusInputUseCase.Execute(id, dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(inputUpdated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
