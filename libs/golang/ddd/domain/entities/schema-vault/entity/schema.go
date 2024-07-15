package entity

import (
	"errors"
	"reflect"
	"time"

	regularTypesConversion "libs/golang/ddd/shared/type-tools/regular-types-converter/conversion"
	md5id "libs/golang/shared/id/go-md5"
	uuid "libs/golang/shared/id/go-uuid"
	schematools "libs/golang/shared/json-schema/schema-tools"
)

var (
	// ErrMissingID is returned when the ID of a Schema is missing.
	ErrMissingID = errors.New("invalid ID")

	// ErrMissingService is returned when the service of a Schema is missing.
	ErrMissingService = errors.New("invalid service")

	// ErrMissingSource is returned when the source of a Schema is missing.
	ErrMissingSource = errors.New("invalid source")

	// ErrMissingProvider is returned when the provider of a Schema is missing.
	ErrMissingProvider = errors.New("invalid provider")

	// ErrMissingSchemaType is returned when the schema type of a Schema is missing.
	ErrMissingSchemaType = errors.New("invalid schema type")

	// ErrJsonSchemaInvalid is returned when the JSON schema of a Schema is invalid.
	ErrJsonSchemaInvalid = errors.New("invalid JSON schema")

	// dateLayout defines the layout for parsing and formatting dates.
	dateLayout = "2006-01-02 15:04:05"
)

type JsonSchema struct {
	Required   []string               `bson:"required"`   // Required lists the required fields in the JSON schema.
	Properties map[string]interface{} `bson:"properties"` // Properties lists the properties in the JSON schema.
	JsonType   string                 `bson:"type"`       // JsonType specifies the type of JSON schema.
}

// Schema represents a schema entity with various attributes such as service, source, provider, and schema type.
type Schema struct {
	ID              md5id.ID   `bson:"_id"`               // ID is the unique identifier of the Schema entity.
	Service         string     `bson:"service"`           // Service is the service name of the Schema entity.
	Source          string     `bson:"source"`            // Source is the source name of the Schema entity.
	Provider        string     `bson:"provider"`          // Provider is the provider name of the Schema entity.
	SchemaType      string     `bson:"schema_type"`       // SchemaType is the type of the schema entity.
	JsonSchema      JsonSchema `bson:"json_schema"`       // JsonSchema is the JSON schema of the Schema entity.
	SchemaVersionID uuid.ID    `bson:"schema_version_id"` // SchemaVersionID is the unique identifier of the schema version.
	CreatedAt       string     `bson:"created_at"`        // CreatedAt is the timestamp when the Schema entity was created.
	UpdatedAt       string     `bson:"updated_at"`        // UpdatedAt is the timestamp when the Schema entity was last updated.
}

// SchemaProps represents the properties needed to create a new Schema entity.
type SchemaProps struct {
	Service    string
	Source     string
	Provider   string
	SchemaType string
	JsonSchema map[string]interface{}
}

// getIDData constructs a map with the service, source, and provider information.
func getIDData(service, source, provider, schemaType string) map[string]string {
	return map[string]string{
		"service":     service,
		"source":      source,
		"provider":    provider,
		"schema_type": schemaType,
	}
}

// transformJsonSchema transforms the JSON schema to a JsonSchema entity.
func transformJsonSchema(jsonSchema map[string]interface{}) JsonSchema {
	var required []string
	var properties map[string]interface{}
	var jsonType string

	if req, ok := jsonSchema["required"].([]interface{}); ok {
		for _, v := range req {
			if str, ok := v.(string); ok {
				required = append(required, str)
			}
		}
	}

	if props, ok := jsonSchema["properties"].(map[string]interface{}); ok {
		properties = props
	}

	if jt, ok := jsonSchema["type"].(string); ok {
		jsonType = jt
	}

	return JsonSchema{
		Required:   required,
		Properties: properties,
		JsonType:   jsonType,
	}
}

// normalizeJsonSchema normalizes the JSON schema to a map representation.
func normalizeJsonSchema(jsonSchema JsonSchema) map[string]interface{} {
	required := make([]interface{}, len(jsonSchema.Required))
	for i, v := range jsonSchema.Required {
		required[i] = v
	}

	return map[string]interface{}{
		"required":   required,
		"properties": jsonSchema.Properties,
		"type":       jsonSchema.JsonType,
	}
}

// NewSchema creates a new Schema entity with the provided properties.
func NewSchema(schemaProps SchemaProps) (*Schema, error) {
	idData := getIDData(schemaProps.Service, schemaProps.Source, schemaProps.Provider, schemaProps.SchemaType)

	jsonSchema := transformJsonSchema(schemaProps.JsonSchema)

	schema := &Schema{
		ID:         md5id.NewID(idData),
		Service:    schemaProps.Service,
		Source:     schemaProps.Source,
		Provider:   schemaProps.Provider,
		SchemaType: schemaProps.SchemaType,
		JsonSchema: jsonSchema,
		UpdatedAt:  time.Now().Format(dateLayout),
		CreatedAt:  time.Now().Format(dateLayout),
	}

	versionID, err := uuid.GenerateUUIDFromMap(schema.GetVersionIDData())
	if err != nil {
		return nil, err
	}
	schema.SetSchemaVersionID(versionID)
	if err := schema.isValid(); err != nil {
		return nil, err
	}

	return schema, nil
}

// GetVersionIDData constructs a map with the service, source, provider, and schema type information.
func (s *Schema) GetVersionIDData() map[string]interface{} {
	return map[string]interface{}{
		"service":     s.Service,
		"source":      s.Source,
		"provider":    s.Provider,
		"schema_type": s.SchemaType,
		"json_schema": s.JsonSchema,
	}
}

// SetSchemaVersionID sets the schema version ID.
func (s *Schema) SetSchemaVersionID(versionID uuid.ID) {
	s.SchemaVersionID = versionID
}

// SetCreatedAt sets the created at timestamp of the Schema entity.
func (s *Schema) SetCreatedAt(createdAt string) {
	s.CreatedAt = createdAt
}

// SetJsonSchema sets the JSON schema of the Schema entity.
func (s *Schema) SetJsonSchema(jsonSchema map[string]interface{}) {
	s.JsonSchema = transformJsonSchema(jsonSchema)
}

// GetEntityID returns the ID of the Schema entity.
func (s *Schema) GetEntityID() string {
	return string(s.ID)
}

// ToMap converts the Schema entity to a map representation.
func (s *Schema) ToMap() (map[string]interface{}, error) {
	if err := s.isValid(); err != nil {
		return nil, err
	}

	doc, err := regularTypesConversion.ConvertFromEntityToMapString(s)
	if err != nil {
		return nil, err
	}

	doc["_id"] = string(doc["_id"].(md5id.ID))
	doc["schema_version_id"] = string(doc["schema_version_id"].(uuid.ID))

	jsonSchema := doc["json_schema"].(map[string]interface{})
	required := jsonSchema["required"].([]interface{})
	strRequired := make([]string, len(required))
	for i, v := range required {
		strRequired[i] = v.(string)
	}
	jsonSchema["required"] = strRequired
	doc["json_schema"] = jsonSchema

	return doc, nil
}

// ToMap converts the JsonSchema to a map[string]interface{}.
func (js JsonSchema) ToMap() (map[string]interface{}, error) {
	result := make(map[string]interface{})
	result["required"] = js.Required
	result["properties"] = js.Properties
	result["type"] = js.JsonType
	return result, nil
}

// MapToEntity converts a map to a Schema entity.
func (s *Schema) MapToEntity(doc map[string]interface{}) (*Schema, error) {
	if id, ok := doc["_id"].(string); ok {
		doc["_id"] = md5id.ID(id)
	} else {
		return nil, errors.New("field _id has invalid type")
	}

	if schemaVersionID, ok := doc["schema_version_id"].(string); ok {
		doc["schema_version_id"] = uuid.ID(schemaVersionID)
	} else {
		return nil, errors.New("field schema_version_id has invalid type")
	}

	schemaEntity, err := regularTypesConversion.ConvertFromMapStringToEntity(reflect.TypeOf(Schema{}), doc)
	if err != nil {
		return nil, err
	}
	schema := schemaEntity.(*Schema)
	return schema, nil
}

// isValid checks if the Schema entity is valid.
func (s *Schema) isValid() error {
	if s.ID == "" {
		return ErrMissingID
	}
	if s.Service == "" {
		return ErrMissingService
	}
	if s.Source == "" {
		return ErrMissingSource
	}
	if s.Provider == "" {
		return ErrMissingProvider
	}
	if s.SchemaType == "" {
		return ErrMissingSchemaType
	}

	jsonSchema := normalizeJsonSchema(s.JsonSchema)
	if err := schematools.ValidateJSONSchema(jsonSchema); err != nil {
		return ErrJsonSchemaInvalid
	}
	return nil
}
