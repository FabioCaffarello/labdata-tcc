package handlers

import (
	"encoding/json"
	"libs/golang/ddd/domain/entities/config-vault/entity"
	inputdto "libs/golang/ddd/dtos/config-vault/input"
	"libs/golang/ddd/usecases/config-vault/usecase"
	typetools "libs/golang/shared/type-tools"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// WebConfigHandler handles HTTP requests for configuration operations.
type WebConfigHandler struct {
	ConfigRepository entity.ConfigRepositoryInterface
}

// NewWebConfigHandler creates and returns a new WebConfigHandler instance with the provided ConfigRepository.
//
// Parameters:
//
//	configRepository: The repository interface for managing Config entities.
//
// Returns:
//
//	A new WebConfigHandler instance.
func NewWebConfigHandler(
	ConfigRepository entity.ConfigRepositoryInterface,
) *WebConfigHandler {
	return &WebConfigHandler{
		ConfigRepository: ConfigRepository,
	}
}

// CreateConfig handles HTTP POST requests to create a new configuration. It decodes the request body into a ConfigDTO,
// executes the CreateConfigUseCase, and writes the created configuration as a JSON response.
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
func (h *WebConfigHandler) CreateConfig(w http.ResponseWriter, r *http.Request) {
	var dto inputdto.ConfigDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createConfigUseCase := usecase.NewCreateConfigUseCase(h.ConfigRepository)
	configCreated, err := createConfigUseCase.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(configCreated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// UpdateConfig handles HTTP PUT requests to update an existing configuration. It decodes the request body into a ConfigDTO,
// executes the UpdateConfigUseCase, and writes the updated configuration as a JSON response.
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
func (h *WebConfigHandler) UpdateConfig(w http.ResponseWriter, r *http.Request) {
	var dto inputdto.ConfigDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updateConfigUseCase := usecase.NewUpdateConfigUseCase(h.ConfigRepository)
	configUpdated, err := updateConfigUseCase.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(configUpdated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// DeleteConfig handles HTTP DELETE requests to delete an existing configuration by its ID.
// It extracts the ID from the query parameters, executes the DeleteConfigUseCase, and writes a success message.
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
// If the ID is not provided or an error occurs during the deletion process, it responds with the appropriate HTTP status code.
func (h *WebConfigHandler) DeleteConfig(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	deleteConfigUseCase := usecase.NewDeleteConfigUseCase(h.ConfigRepository)
	err := deleteConfigUseCase.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Config deleted successfully"))
}

// ListAllConfigs handles HTTP GET requests to list all configurations. It executes the ListAllConfigUseCase and writes
// the list of configurations as a JSON response.
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
func (h *WebConfigHandler) ListAllConfigs(w http.ResponseWriter, r *http.Request) {
	listConfigsUseCase := usecase.NewListAllConfigUseCase(h.ConfigRepository)
	configs, err := listConfigsUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(configs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ListConfigByID handles HTTP GET requests to list a configuration by its ID. It extracts the ID from the query parameters,
// executes the ListOneByIDConfigUseCase, and writes the configuration as a JSON response.
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
// If the ID is not provided or an error occurs during the listing process, it responds with the appropriate HTTP status code.
func (h *WebConfigHandler) ListConfigByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	getConfigUseCase := usecase.NewListOneByIDConfigUseCase(h.ConfigRepository)
	config, err := getConfigUseCase.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(config)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ListConfigsByServiceAndProvider handles HTTP GET requests to list configurations by service and provider.
// It extracts the service and provider from the query parameters, executes the ListByServiceAndProviderConfigUseCase,
// and writes the list of configurations as a JSON response.
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
// If the service or provider is not provided or an error occurs during the listing process, it responds with the appropriate HTTP status code.
func (h *WebConfigHandler) ListConfigsByServiceAndProvider(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	service := chi.URLParam(r, "service")
	if service == "" || provider == "" {
		http.Error(w, "Service and provider are required", http.StatusBadRequest)
		return
	}

	listConfigsUseCase := usecase.NewListAllByServiceAndProviderConfigUseCase(h.ConfigRepository)
	configs, err := listConfigsUseCase.Execute(provider, service)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(configs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ListConfigsBySourceAndProvider handles HTTP GET requests to list configurations by source and provider.
// It extracts the source and provider from the query parameters, executes the ListBySourceAndProviderConfigUseCase,
// and writes the list of configurations as a JSON response.
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
// If the source or provider is not provided or an error occurs during the listing process, it responds with the appropriate HTTP status code.
func (h *WebConfigHandler) ListConfigsBySourceAndProvider(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	source := chi.URLParam(r, "source")
	if source == "" || provider == "" {
		http.Error(w, "Source and provider are required", http.StatusBadRequest)
		return
	}

	listConfigsUseCase := usecase.NewListAllBySourceAndProviderConfigUseCase(h.ConfigRepository)
	configs, err := listConfigsUseCase.Execute(provider, source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(configs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ListConfigsByServiceAndSourceAndProvider handles HTTP GET requests to list configurations by service, source, and provider.
// It extracts the service, source, and provider from the query parameters, executes the ListByServiceAndSourceAndProviderConfigUseCase,
// and writes the list of configurations as a JSON response.
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
// If the service, source, or provider is not provided or an error occurs during the listing process, it responds with the appropriate HTTP status code.
func (h *WebConfigHandler) ListConfigsByServiceAndSourceAndProvider(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	service := chi.URLParam(r, "service")
	source := chi.URLParam(r, "source")

	if service == "" || source == "" || provider == "" {
		http.Error(w, "Service, source, and provider are required", http.StatusBadRequest)
		return
	}

	listConfigsUseCase := usecase.NewListAllByServiceAndSourceAndProviderConfigUseCase(h.ConfigRepository)
	configs, err := listConfigsUseCase.Execute(service, source, provider)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(configs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ListConfigsByServiceAndProviderAndActive handles HTTP GET requests to list configurations by service, provider, and active status.
// It extracts the service, provider, and active status from the query parameters, executes the ListByServiceAndProviderAndActiveConfigUseCase,
// and writes the list of configurations as a JSON response.
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
// If the service, provider, or active status is not provided or an error occurs during the listing process, it responds with the appropriate HTTP status code.
func (h *WebConfigHandler) ListConfigsByServiceAndProviderAndActive(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	service := chi.URLParam(r, "service")
	active := chi.URLParam(r, "active")

	if service == "" || provider == "" || active == "" {
		http.Error(w, "Service, provider, and active status are required", http.StatusBadRequest)
		return
	}

	activeBool, err := typetools.ParseBool(active)
	if err != nil {
		http.Error(w, "Invalid active status value", http.StatusBadRequest)
		return
	}

	listConfigsUseCase := usecase.NewListAllByServiceAndProviderAndActiveConfigUseCase(h.ConfigRepository)
	configs, err := listConfigsUseCase.Execute(service, provider, activeBool)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(configs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ListConfigsByProviderAndDependencies handles HTTP GET requests to list configurations by their dependencies.
// It extracts the dependencies from the query parameters, executes the ListAllByProviderAndDependsOnConfigUseCase,
// and writes the list of configurations as a JSON response.
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
// If the dependencies are not provided or an error occurs during the listing process, it responds with the appropriate HTTP status code.
func (h *WebConfigHandler) ListConfigsByProviderAndDependencies(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	service := chi.URLParam(r, "service")
	source := chi.URLParam(r, "source")

	if provider == "" && service == "" && source == "" {
		http.Error(w, "At least one dependency is required", http.StatusBadRequest)
		return
	}

	listConfigsUseCase := usecase.NewListAllByProviderAndDependsOnConfigUseCase(h.ConfigRepository)
	configs, err := listConfigsUseCase.Execute(provider, service, source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(configs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
