package handlers

import (
	"encoding/json"
	"libs/golang/ddd/domain/entities/input-broker/entity"
	inputdto "libs/golang/ddd/dtos/input-broker/input"
	"libs/golang/ddd/usecases/input-broker/usecase"
	events "libs/golang/shared/go-events/amqp_events"
	"net/http"
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
