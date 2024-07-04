package shareddto

// MetadataDTO represents the metadata data transfer object.
type MetadataDTO struct {
	Provider            string `json:"provider"`             // Provider represents the provider of the input data.
	Service             string `json:"service"`              // Service represents the service of the input data.
	Source              string `json:"source"`               // Source represents the source of the input data.
	ProcessingID        string `json:"processing_id"`        // ProcessingID represents the unique identifier of the processing job.
	ProcessingTimestamp string `json:"processing_timestamp"` // ProcessingTimestamp represents the timestamp when the processing job was executed.
}

// StatusDTO represents the status data transfer object.
type StatusDTO struct {
	Code   int    `json:"code"`   // Code represents the status code.
	Detail string `json:"detail"` // Detail represents the status detail.
}
