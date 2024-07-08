package inputdto

// InputDTO represents the input data transfer object.
type InputDTO struct {
	Provider string                 `json:"provider"` // Provider represents the provider of the input data.
	Service  string                 `json:"service"`  // Service represents the service of the input data.
	Source   string                 `json:"source"`   // Source represents the source of the input data.
	Data     map[string]interface{} `json:"data"`     // Data represents the input data.
}
