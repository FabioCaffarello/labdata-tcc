package outputdto

// ErrMsgDTO represents the error message data transfer object.
type ErrMsgDTO struct {
	Err         error  `json:"error"`
	ListenerTag string `json:"listener_tag"`
	Msg         []byte `json:"msg"`
}

// ProcessOrderDTO represents the data transfer object for processing orders.
// It includes the necessary details required for processing an order,
// such as the processing ID, service, source, provider, stage, and data.
type ProcessOrderDTO struct {
	ID           string                 `json:"_id"`           // ID represents the unique identifier of the order.
	ProcessingID string                 `json:"processing_id"` // ProcessingID represents the unique identifier of the order processing.
	Service      string                 `json:"service"`       // Service represents the name of the service for which the order is processed.
	Source       string                 `json:"source"`        // Source indicates the origin or source of the order.
	Provider     string                 `json:"provider"`      // Provider specifies the provider of the order.
	Stage        string                 `json:"stage"`         // Stage represents the current stage of the order processing.
	InputID      string                 `json:"input_id"`      // InputID represents the unique identifier of the input data.
	Data         map[string]interface{} `json:"data"`          // Data represents the order data.
}
