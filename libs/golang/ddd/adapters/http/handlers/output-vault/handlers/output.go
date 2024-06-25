package handlers

import (
	"encoding/json"
	"libs/golang/ddd/domain/entities/output-vault/entity"
	inputdto "libs/golang/ddd/dtos/output-vault/input"
	"libs/golang/ddd/usecases/output-vault/usecase"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// WebOutputHandler represents the handler for the output vault.
type WebOutputHandler struct {
	OutputRepository entity.OutputRepositoryInterface
}

// NewWebOutputHandler initializes a new instance of WebOutputHandler with the provided OutputRepositoryInterface.
//
// Parameters:
//
//	outputRepository: The repository interface for managing Output entities.
//
// Returns:
//
//	A pointer to an instance of WebOutputHandler.
func NewWebOutputHandler(
	outputRepository entity.OutputRepositoryInterface,
) *WebOutputHandler {
	return &WebOutputHandler{
		OutputRepository: outputRepository,
	}
}

// CreateOutput handles HTTP POST requests to create a new output. It decodes the request body into a OutputDTO,
// executes the CreateOutputUseCase, and writes the created output as a JSON response.
//
// Parameters:
//
//	w: The HTTP response writer.
//	r: The HTTP request.
//
// Returns:
//
//	None.
//
// If the request body cannot be decoded, it responds with HTTP status 400 (Bad Request).
// If an error occurs during the creation process, it responds with HTTP status 500 (Internal Server Error).
func (h *WebOutputHandler) CreateOutput(w http.ResponseWriter, r *http.Request) {
	var dto inputdto.OutputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOutputUseCase := usecase.NewCreateOutputUseCase(h.OutputRepository)
	outputCreated, err := createOutputUseCase.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(outputCreated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// UpdateOutput handles HTTP PUT requests to update an existing output. It decodes the request body into a OutputDTO,
// executes the UpdateOutputUseCase, and writes the updated output as a JSON response.
//
// Parameters:
//
//	w: The HTTP response writer.
//	r: The HTTP request.
//
// Returns:
//
//	None.
//
// If the request body cannot be decoded, it responds with HTTP status 400 (Bad Request).
// If an error occurs during the update process, it responds with HTTP status 500 (Internal Server Error).
func (h *WebOutputHandler) UpdateOutput(w http.ResponseWriter, r *http.Request) {
	var dto inputdto.OutputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updateOutputUseCase := usecase.NewUpdateOutputUseCase(h.OutputRepository)
	outputUpdated, err := updateOutputUseCase.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(outputUpdated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// DeleteOutput handles HTTP DELETE requests to delete an existing output by its ID. It extracts the output ID from the
// request URL, executes the DeleteOutputUseCase, and writes the deleted output as a JSON response.
//
// Parameters:
//
//	w: The HTTP response writer.
//	r: The HTTP request.
//
// Returns:
//
//	None.
//
// If an error occurs during the deletion process, it responds with HTTP status 500 (Internal Server Error).
func (h *WebOutputHandler) DeleteOutput(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	deleteOutputUseCase := usecase.NewDeleteOutputUseCase(h.OutputRepository)
	err := deleteOutputUseCase.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`Output deleted successfully`))
}

// ListAllOutputs handles HTTP GET requests to list all outputs. It executes the ListAllOutputUseCase and writes the
// outputs as a JSON response.
//
// Parameters:
//
//	w: The HTTP response writer.
//	r: The HTTP request.
//
// Returns:
//
//	None.
//
// If an error occurs during the listing process, it responds with HTTP status 500 (Internal Server Error).
func (h *WebOutputHandler) ListAllOutputs(w http.ResponseWriter, r *http.Request) {
	listAllOutputUseCase := usecase.NewListAllOutputUseCase(h.OutputRepository)
	outputs, err := listAllOutputUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(outputs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ListOutputByID handles HTTP GET requests to list a output by its ID. It extracts the output ID from the request URL,
// executes the ListOneByIDOutputUseCase, and writes the output as a JSON response.
//
// Parameters:
//
//	w: The HTTP response writer.
//	r: The HTTP request.
//
// Returns:
//
//	None.
//
// If an error occurs during the listing process, it responds with HTTP status 500 (Internal Server Error).
func (h *WebOutputHandler) ListOutputByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	listOneByIDOutputUseCase := usecase.NewListOneByIDOutputUseCase(h.OutputRepository)
	output, err := listOneByIDOutputUseCase.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ListOutputsByServiceAndProvider handles HTTP GET requests to list all outputs by service and provider. It extracts the
// service and provider names from the request URL, executes the ListAllBySourceAndProviderOutputUseCase, and writes the
// outputs as a JSON response.
//
// Parameters:
//
//	w: The HTTP response writer.
//	r: The HTTP request.
//
// Returns:
//
//	None.
//
// If an error occurs during the listing process, it responds with HTTP status 500 (Internal Server Error).
func (h *WebOutputHandler) ListOutputsByServiceAndProvider(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	service := chi.URLParam(r, "service")
	if service == "" || provider == "" {
		http.Error(w, "Service and provider are required", http.StatusBadRequest)
		return
	}

	listAllByServiceAndSourceOutputUseCase := usecase.NewListAllBySourceAndProviderOutputUseCase(h.OutputRepository)
	outputs, err := listAllByServiceAndSourceOutputUseCase.Execute(provider, service)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(outputs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ListOutputsBySourceAndProvider handles HTTP GET requests to list all outputs by source and provider. It extracts the
// source and provider names from the request URL, executes the ListAllBySourceAndProviderOutputUseCase, and writes the
// outputs as a JSON response.
//
// Parameters:
//
//	w: The HTTP response writer.
//	r: The HTTP request.
//
// Returns:
//
//	None.
//
// If an error occurs during the listing process, it responds with HTTP status 500 (Internal Server Error).
func (h *WebOutputHandler) ListOutputsBySourceAndProvider(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	source := chi.URLParam(r, "source")
	if source == "" || provider == "" {
		http.Error(w, "Source and provider are required", http.StatusBadRequest)
		return
	}

	listAllBySourceAndProviderOutputUseCase := usecase.NewListAllBySourceAndProviderOutputUseCase(h.OutputRepository)
	outputs, err := listAllBySourceAndProviderOutputUseCase.Execute(provider, source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(outputs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ListOutputsByServiceAndSourceAndProvider handles HTTP GET requests to list all outputs by service, source, and provider.
// It extracts the service, source, and provider names from the request URL, executes the ListAllByServiceAndSourceAndProviderOutputUseCase,
// and writes the outputs as a JSON response.
//
// Parameters:
//
//	w: The HTTP response writer.
//	r: The HTTP request.
//
// Returns:
//
//	None.
//
// If an error occurs during the listing process, it responds with HTTP status 500 (Internal Server Error).
func (h *WebOutputHandler) ListOutputsByServiceAndSourceAndProvider(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	service := chi.URLParam(r, "service")
	source := chi.URLParam(r, "source")
	if service == "" || source == "" || provider == "" {
		http.Error(w, "Service, source and provider are required", http.StatusBadRequest)
		return
	}

	listAllByServiceAndSourceAndProviderOutputUseCase := usecase.NewListAllByServiceAndSourceAndProviderOutputUseCase(h.OutputRepository)
	outputs, err := listAllByServiceAndSourceAndProviderOutputUseCase.Execute(provider, service, source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(outputs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
