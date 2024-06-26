package entity

import (
	"errors"
	regularTypesConversion "libs/golang/ddd/shared/type-tools/regular-types-converter/conversion"
	md5id "libs/golang/shared/id/go-md5"
	uuid "libs/golang/shared/id/go-uuid"
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

	// ErrInvalidStatusCode is returned when the status code of an Output is invalid.
	ErrInvalidStatusCode = errors.New("invalid status code")

	// ErrInvalidStatusDetail is returned when the status detail of an Output is invalid.
	ErrInvalidStatusDetail = errors.New("invalid status detail")

	// StatusCodeIdle represents the idle status code.
	StatusCodeIdle = 0

	// StatusDetailIdle represents the idle status detail.
	StatusDetailIdle = "Idle"

	// dateLayout defines the layout for parsing and formatting dates.
	dateLayout = "2006-01-02 15:04:05"
)

type Status struct {
	Code   int    `bson:"code"`   // Code represents the status code.
	Detail string `bson:"detail"` // Detail represents the status detail.
}

type Metadata struct {
	ProcessingID        uuid.ID `bson:"processing_id"`        // ProcessingID is the unique identifier of the processing job.
	ProcessingTimestamp string  `bson:"processing_timestamp"` // ProcessingTimestamp is the timestamp when the processing job was executed.
	Service             string  `bson:"service"`              // Service represents the service of the input data.
	Source              string  `bson:"source"`               // Source represents the source of the input data.
	Provider            string  `bson:"provider"`             // Provider represents the provider of the input data.
}

type Input struct {
	ID        md5id.ID               `bson:"_id"`        // ID is the unique identifier of the Input entity.
	Data      map[string]interface{} `bson:"data"`       // Data represents the input data.
	Metadata  Metadata               `bson:"metadata"`   // Metadata represents the metadata of the input data.
	Status    Status                 `bson:"status"`     // Status represents the status of the input data.
	CreatedAt string                 `bson:"created_at"` // CreatedAt is the timestamp when the Input entity was created.
	UpdatedAt string                 `bson:"updated_at"` // UpdatedAt is the timestamp when the Input entity was last updated.
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

type InputProps struct {
	Data     map[string]interface{}
	Service  string
	Source   string
	Provider string
}

// NewInput creates a new Input entity based on the provided InputProps. It validates the
// properties and returns an error if any of them are invalid.
func NewInput(props InputProps) (*Input, error) {
	idData := getIDData(props.Service, props.Source, props.Provider, props.Data)
	processingID, err := uuid.GenerateUUIDFromMap(idData)
	if err != nil {
		return nil, err
	}

	input := &Input{
		ID:   md5id.NewID(idData),
		Data: props.Data,
		Metadata: Metadata{
			ProcessingID: processingID,
			Service:      props.Service,
			Source:       props.Source,
			Provider:     props.Provider,
		},
		Status: Status{
			Code:   StatusCodeIdle,
			Detail: StatusDetailIdle,
		},
		CreatedAt: time.Now().Format(dateLayout),
		UpdatedAt: time.Now().Format(dateLayout),
	}

	if err := input.isValid(); err != nil {
		return nil, err
	}

	return input, nil
}

// GetEntityID returns the unique identifier of the Input entity.
func (i *Input) GetEntityID() string {
	return string(i.ID)
}

// SetStatus sets the status of the Input entity.
func (i *Input) SetStatus(code int, detail string) {
	i.Status.Code = code
	i.Status.Detail = detail
}

// SetProcessingTimestamp sets the processing timestamp of the Input entity.
func (i *Input) SetProcessingTimestamp(timestamp time.Time) {
	i.Metadata.ProcessingTimestamp = timestamp.Format(dateLayout)
}

// SetCreatedAt sets the created at timestamp of the Input entity.
func (i *Input) SetCreatedAt(createdAt string) {
	i.CreatedAt = createdAt
}

// ToMap converts the Input entity to a map.
func (i *Input) ToMap() (map[string]interface{}, error) {
	doc, err := regularTypesConversion.ConvertFromEntityToMapString(i)
	if err != nil {
		return nil, err
	}

	metadata, err := regularTypesConversion.ConvertFromEntityToMapString(i.Metadata)
	if err != nil {
		return nil, err
	}

	doc["_id"] = string(doc["_id"].(md5id.ID))
	doc["metadata"] = metadata

	return doc, nil
}

// MapToEntity converts a map representation to an Input entity.
func (i *Input) MapToEntity(doc map[string]interface{}) (*Input, error) {
	if id, ok := doc["_id"].(string); ok {
		doc["_id"] = md5id.ID(id)
	} else {
		return nil, errors.New("field _id has invalid type")
	}

	metadata, ok := doc["metadata"].(map[string]interface{})
	if !ok {
		return nil, errors.New("field metadata has invalid type")
	}
	doc["metadata"].(map[string]interface{})["processing_id"] = uuid.ID(metadata["processing_id"].(string))

	inputEntity, err := regularTypesConversion.ConvertFromMapStringToEntity(reflect.TypeOf(Input{}), doc)
	if err != nil {
		return nil, err
	}

	input := inputEntity.(*Input)

	return input, nil
}

// isValid checks if the Input entity has valid properties.
func (i *Input) isValid() error {
	if i.ID == "" {
		return ErrInvalidID
	}

	if i.Metadata.Service == "" {
		return ErrInvalidService
	}

	if i.Metadata.Source == "" {
		return ErrInvalidSource
	}

	if i.Metadata.Provider == "" {
		return ErrInvalidProvider
	}

	if i.Metadata.ProcessingID == "" {
		return ErrInvalidProcessingID
	}

	if i.Data == nil {
		return ErrInvalidData
	}

	return nil
}
