package entity

import (
	"errors"
	regularTypesConversion "libs/golang/ddd/shared/type-tools/regular-types-converter/conversion"
	md5id "libs/golang/shared/id/go-md5"
	"log"
	"reflect"
	"time"
)

var (
	// ErrInvalidID is returned when the ID of an Output is invalid.
	ErrInvalidID = errors.New("invalid ID")

	// ErrInvalidService is returned when the service of an Output is invalid.
	ErrInvalidService = errors.New("invalid service")

	// ErrInvalidSource is returned when the source of an Output is invalid.
	ErrInvalidSource = errors.New("invalid source")

	// ErrInvalidProvider is returned when the provider of an Output is invalid.
	ErrInvalidProvider = errors.New("invalid provider")

	// ErrInvalidInputID is returned when the input ID of an Output is invalid.
	ErrInvalidInputID = errors.New("invalid input ID")

	// ErrInvalidProcessingID is returned when the processing ID of an Output is invalid.
	ErrInvalidProcessingID = errors.New("invalid processing ID")

	// ErrInvalidProcessingTimestamp is returned when the processing timestamp of an Output is invalid.
	ErrInvalidProcessingTimestamp = errors.New("invalid processing timestamp")

	// ErrInvalidData is returned when the data of an Output is invalid.
	ErrInvalidData = errors.New("invalid data")

	// ErrInvalidInputData is returned when the input data of an Output is invalid.
	ErrInvalidInputData = errors.New("invalid input data")

	// dateLayout defines the layout for parsing and formatting dates.
	dateLayout = "2006-01-02 15:04:05"
)

// Input represents the input data of an Output entity.
type Input struct {
	Data                map[string]interface{} `bson:"data"`                 // Data represents the input data.
	ProcessingID        string                 `bson:"processing_id"`        // ProcessingID is the unique identifier of the processing job.
	ProcessingTimestamp string                 `bson:"processing_timestamp"` // ProcessingTimestamp is the timestamp when the processing job was executed.
}

// Metadata represents the metadata of an Output entity.
type Metadata struct {
	InputID string `bson:"input_id"` // InputID is the unique identifier of the input data.
	Input   Input  `bson:"input"`    // Input represents the input data of the Output entity.
}

// Output represents an output entity with various attributes such as service, source, provider, and data.
type Output struct {
	ID        md5id.ID               `bson:"_id"`        // ID is the unique identifier of the Output entity.
	Data      map[string]interface{} `bson:"data"`       // Data represents the output data.
	Service   string                 `bson:"service"`    // Service represents the name of the service for which the output is created.
	Source    string                 `bson:"source"`     // Source indicates the origin or source of the output.
	Provider  string                 `bson:"provider"`   // Provider specifies the provider of the output.
	Metadata  Metadata               `bson:"metadata"`   // Metadata represents the metadata of the output.
	CreatedAt string                 `bson:"created_at"` // CreatedAt is the timestamp when the Output entity was created.
	UpdatedAt string                 `bson:"updated_at"` // UpdatedAt is the timestamp when the Output entity was last updated.
}

// OutputProps represents the properties needed to create a new Output entity.
type OutputProps struct {
	Data     map[string]interface{}
	Service  string
	Source   string
	Provider string
	Metadata map[string]interface{}
}

// getIDData constructs a map with the service, source, provider and data information.
func getIDData(service, source, provider string, data map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"service":  service,
		"source":   source,
		"provider": provider,
		"data":     data,
	}
}

// transformMetadata converts the raw metadata map to a Metadata struct.
func transformMetadata(metadataRaw map[string]interface{}) (Metadata, error) {
	var metadata Metadata
	if inputID, ok := metadataRaw["input_id"].(string); ok {
		metadata.InputID = inputID
	} else {
		return Metadata{}, ErrInvalidInputID
	}
	metadataInput, ok := metadataRaw["input"].(map[string]interface{})
	if !ok {
		return Metadata{}, ErrInvalidInputData
	}
	if processingID, ok := metadataInput["processing_id"].(string); ok {
		metadata.Input.ProcessingID = processingID
	} else {
		return Metadata{}, ErrInvalidProcessingID
	}
	if processingTimestamp, ok := metadataInput["processing_timestamp"].(string); ok {
		metadata.Input.ProcessingTimestamp = processingTimestamp
	} else {
		return Metadata{}, ErrInvalidProcessingTimestamp
	}
	if data, ok := metadataInput["data"].(map[string]interface{}); ok {
		metadata.Input.Data = data
	} else {
		return Metadata{}, ErrInvalidInputData
	}
	return metadata, nil
}

// NewOutput creates a new Output entity based on the provided OutputProps. It validates the
// properties and returns an error if any of them are invalid.
func NewOutput(props OutputProps) (*Output, error) {
	idData := getIDData(props.Service, props.Source, props.Provider, props.Data)

	metadata, err := transformMetadata(props.Metadata)
	log.Printf("Metadata: %+v", metadata)
	if err != nil {
		return nil, err
	}

	output := &Output{
		ID:        md5id.NewID(idData),
		Data:      props.Data,
		Service:   props.Service,
		Source:    props.Source,
		Provider:  props.Provider,
		Metadata:  metadata,
		UpdatedAt: time.Now().Format(dateLayout),
		CreatedAt: time.Now().Format(dateLayout),
	}
	log.Printf("Output: %+v", output)

	if err := output.isValid(); err != nil {
		return nil, err
	}

	return output, nil
}

// GetEntityID returns the ID of the Output entity.
func (o *Output) GetEntityID() string {
	return string(o.ID)
}

// SetCreatedAt sets the created at timestamp of the Output entity.
func (o *Output) SetCreatedAt(createdAt string) {
	o.CreatedAt = createdAt
}

// ToMap converts the Output entity to a map representation.
func (o *Output) ToMap() (map[string]interface{}, error) {
	doc, err := regularTypesConversion.ConvertFromEntityToMapString(o)
	if err != nil {
		return nil, err
	}

	doc["_id"] = string(doc["_id"].(md5id.ID))

	return doc, nil
}

// MapToEntity converts a map representation to an Output entity.
func (o *Output) MapToEntity(doc map[string]interface{}) (*Output, error) {
	if id, ok := doc["_id"].(string); ok {
		doc["_id"] = md5id.ID(id)
	} else {
		return nil, errors.New("field _id has invalid type")
	}

	outputEntity, err := regularTypesConversion.ConvertFromMapStringToEntity(reflect.TypeOf(Output{}), doc)
	if err != nil {
		return nil, err
	}

	output := outputEntity.(*Output)

	return output, nil
}

// isValid checks if the Output entity has valid properties.
func (o *Output) isValid() error {
	if o.ID == "" {
		return ErrInvalidID
	}
	if o.Service == "" {
		return ErrInvalidService
	}
	if o.Source == "" {
		return ErrInvalidSource
	}
	if o.Provider == "" {
		return ErrInvalidProvider
	}
	if o.Metadata.InputID == "" {
		return ErrInvalidInputID
	}
	if o.Metadata.Input.ProcessingID == "" {
		return ErrInvalidProcessingID
	}
	if o.Metadata.Input.ProcessingTimestamp == "" {
		return ErrInvalidProcessingTimestamp
	}
	if o.Data == nil {
		return ErrInvalidData
	}
	return nil
}
