package shareddto

// JsonSchema is a DTO that represents a JSON schema.
// It includes the required fields, properties, and type of the JSON schema.
type JsonSchema struct {
	Required   []string               `json:"required"`   // Required lists the required fields in the JSON schema.
	Properties map[string]interface{} `json:"properties"` // Properties lists the properties in the JSON schema.
	JsonType   string                 `json:"type"`       // JsonType specifies the type of JSON schema.
}
