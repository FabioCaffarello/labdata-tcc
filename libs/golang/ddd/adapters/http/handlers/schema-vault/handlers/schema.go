package handlers

import (
	"encoding/json"
	"libs/golang/ddd/domain/entities/schema-vault/entity"
	inputdto "libs/golang/ddd/dtos/schema-vault/input"
	"libs/golang/ddd/usecases/schema-vault/usecase"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// WebSchemaHandler represents the handler for the schema vault.
type WebSchemaHandler struct {
	SchemaRepository entity.SchemaRepositoryInterface
}

// NewWebSchemaHandler initializes a new instance of WebSchemaHandler with the provided SchemaRepositoryInterface.
//
// Parameters:
//
//	schemaRepository: The repository interface for managing Schema entities.
//
// Returns:
//
//	A pointer to an instance of WebSchemaHandler.
func NewWebSchemaHandler(
	schemaRepository entity.SchemaRepositoryInterface,
) *WebSchemaHandler {
	return &WebSchemaHandler{
		SchemaRepository: schemaRepository,
	}
}

// CreateSchema handles HTTP POST requests to create a new schema. It decodes the request body into a SchemaDTO,
// executes the CreateSchemaUseCase, and writes the created schema as a JSON response.
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
func (h *WebSchemaHandler) CreateSchema(w http.ResponseWriter, r *http.Request) {
	var dto inputdto.SchemaDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createSchemaUseCase := usecase.NewCreateSchemaUseCase(h.SchemaRepository)
	schemaCreated, err := createSchemaUseCase.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(schemaCreated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// UpdateSchema handles HTTP PUT requests to update an existing schema. It decodes the request body into a SchemaDTO,
// executes the UpdateSchemaUseCase, and writes the updated schema as a JSON response.
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
func (h *WebSchemaHandler) UpdateSchema(w http.ResponseWriter, r *http.Request) {
	var dto inputdto.SchemaDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updateSchemaUseCase := usecase.NewUpdateSchemaUseCase(h.SchemaRepository)
	schemaUpdated, err := updateSchemaUseCase.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(schemaUpdated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// DeleteSchema handles HTTP DELETE requests to delete an existing schema by its ID. It extracts the schema ID from the
// request URL, executes the DeleteSchemaUseCase, and writes the deleted schema as a JSON response.
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
func (h *WebSchemaHandler) DeleteSchema(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	deleteSchemaUseCase := usecase.NewDeleteSchemaUseCase(h.SchemaRepository)
	err := deleteSchemaUseCase.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`Schema deleted successfully`))
}

// ListAllSchemas handles HTTP GET requests to list all schemas. It executes the ListAllSchemaUseCase and writes the
// schemas as a JSON response.
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
func (h *WebSchemaHandler) ListAllSchemas(w http.ResponseWriter, r *http.Request) {
	listAllSchemaUseCase := usecase.NewListAllSchemaUseCase(h.SchemaRepository)
	schemas, err := listAllSchemaUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(schemas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ListSchemaByID handles HTTP GET requests to list a schema by its ID. It extracts the schema ID from the request URL,
// executes the ListOneByIDSchemaUseCase, and writes the schema as a JSON response.
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
func (h *WebSchemaHandler) ListSchemaByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	listOneByIDSchemaUseCase := usecase.NewListOneByIDSchemaUseCase(h.SchemaRepository)
	schema, err := listOneByIDSchemaUseCase.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(schema)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ListSchemasByServiceAndProvider handles HTTP GET requests to list all schemas by service and provider. It extracts the
// service and provider names from the request URL, executes the ListAllBySourceAndProviderSchemaUseCase, and writes the
// schemas as a JSON response.
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
func (h *WebSchemaHandler) ListSchemasByServiceAndProvider(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	service := chi.URLParam(r, "service")
	if service == "" || provider == "" {
		http.Error(w, "Service and provider are required", http.StatusBadRequest)
		return
	}

	listAllByServiceAndProviderSchemaUseCase := usecase.NewListAllByServiceAndProviderSchemaUseCase(h.SchemaRepository)
	schemas, err := listAllByServiceAndProviderSchemaUseCase.Execute(provider, service)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(schemas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ListSchemasBySourceAndProvider handles HTTP GET requests to list all schemas by source and provider. It extracts the
// source and provider names from the request URL, executes the ListAllBySourceAndProviderSchemaUseCase, and writes the
// schemas as a JSON response.
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
func (h *WebSchemaHandler) ListSchemasBySourceAndProvider(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	source := chi.URLParam(r, "source")
	if source == "" || provider == "" {
		http.Error(w, "Source and provider are required", http.StatusBadRequest)
		return
	}

	listAllBySourceAndProviderSchemaUseCase := usecase.NewListAllBySourceAndProviderSchemaUseCase(h.SchemaRepository)
	schemas, err := listAllBySourceAndProviderSchemaUseCase.Execute(provider, source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(schemas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ListSchemasByServiceAndSourceAndProvider handles HTTP GET requests to list all schemas by service, source, and provider.
// It extracts the service, source, and provider names from the request URL, executes the ListAllByServiceAndSourceAndProviderSchemaUseCase,
// and writes the schemas as a JSON response.
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
func (h *WebSchemaHandler) ListSchemasByServiceAndSourceAndProvider(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	service := chi.URLParam(r, "service")
	source := chi.URLParam(r, "source")
	if service == "" || source == "" || provider == "" {
		http.Error(w, "Service, source and provider are required", http.StatusBadRequest)
		return
	}

	listAllByServiceAndSourceAndProviderSchemaUseCase := usecase.NewListAllByServiceAndSourceAndProviderSchemaUseCase(h.SchemaRepository)
	schemas, err := listAllByServiceAndSourceAndProviderSchemaUseCase.Execute(provider, service, source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(schemas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *WebSchemaHandler) ListSchemasByServiceAndSourceAndProviderAndSchemaType(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	service := chi.URLParam(r, "service")
	source := chi.URLParam(r, "source")
	schemaType := chi.URLParam(r, "schemaType")
	if service == "" || source == "" || provider == "" || schemaType == "" {
		http.Error(w, "Service, source, provider and schema type are required", http.StatusBadRequest)
		return
	}

	listAllByServiceAndSourceAndProviderAndSchemaTypeSchemaUseCase := usecase.NewListOneByServiceAndSourceAndProviderAndSchemaTypeSchemaUseCase(h.SchemaRepository)
	schemas, err := listAllByServiceAndSourceAndProviderAndSchemaTypeSchemaUseCase.Execute(provider, service, source, schemaType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(schemas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *WebSchemaHandler) ValidateSchema(w http.ResponseWriter, r *http.Request) {
	var dto inputdto.SchemaDataDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	validateSchemaUseCase := usecase.NewValidateSchemaUseCase(h.SchemaRepository)
	valid, err := validateSchemaUseCase.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(valid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
