package shareddto

// MetadataDTO represents the data transfer object for metadata for both input and output dto.
type MetadataDTO struct {
	InputID string   `json:"input_id"` // InputID is the unique identifier of the input data.
	Input   InputDTO `json:"input"`    // Input represents the input data of the Output entity.
}

// InputDTO represents the data transfer object for input metadata for both input and output dto.
type InputDTO struct {
	Data                map[string]interface{} `json:"data"`                 // Data represents the input data.
	ProcessingID        string                 `json:"processing_id"`        // ProcessingID is the unique identifier of the processing job.
	ProcessingTimestamp string                 `json:"processing_timestamp"` // ProcessingTimestamp is the timestamp when the processing job was executed.
}
