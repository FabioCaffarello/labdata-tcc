package entity

import (
	"errors"
	regularTypesConversion "libs/golang/ddd/shared/type-tools/regular-types-converter/conversion"
	md5id "libs/golang/shared/id/go-md5"
	uuid "libs/golang/shared/id/go-uuid"
	"reflect"
)

var (
	// ErrInvalidID is returned when the ID of an EventOrder is invalid.
	ErrInvalidID = errors.New("invalid ID")

	// ErrInvalidService is returned when the service of an EventOrder is invalid.
	ErrInvalidService = errors.New("invalid service")

	// ErrInvalidSource is returned when the source of an EventOrder is invalid.
	ErrInvalidSource = errors.New("invalid source")

	// ErrInvalidProvider is returned when the provider of an EventOrder is invalid.
	ErrInvalidProvider = errors.New("invalid provider")

	// ErrInvalidProcessingID is returned when the processing ID of an EventOrder is invalid.
	ErrInvalidProcessingID = errors.New("invalid processing ID")
)

// EventOrder represents an event order entity with various attributes.
type EventOrder struct {
	ID           md5id.ID               `bson:"_id"`
	Service      string                 `bson:"service"`
	Source       string                 `bson:"source"`
	Provider     string                 `bson:"provider"`
	ProcessingID uuid.ID                `bson:"processing_id"`
	Data         map[string]interface{} `bson:"data"`
}

// EventOrderProps holds the properties required to create a new EventOrder.
type EventOrderProps struct {
	Service      string
	Source       string
	Provider     string
	ProcessingID string
	Data         map[string]interface{}
}

// getIDData constructs a map with the service, source, provider, and data information.
func getIDData(service, source, provider string, data map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"service":  service,
		"source":   source,
		"provider": provider,
		"data":     data,
	}
}

// NewEventOrder creates a new EventOrder with the provided properties.
//
// Parameters:
//   - props: The properties required to create a new EventOrder.
//
// Returns:
//   - A pointer to the created EventOrder.
//   - An error if the validation of the EventOrder fails.
func NewEventOrder(props EventOrderProps) (*EventOrder, error) {
	idData := getIDData(props.Service, props.Source, props.Provider, props.Data)
	eventOrder := &EventOrder{
		ID:           md5id.NewID(idData),
		Service:      props.Service,
		Source:       props.Source,
		Provider:     props.Provider,
		ProcessingID: uuid.ID(props.ProcessingID),
		Data:         props.Data,
	}

	if err := eventOrder.isValid(); err != nil {
		return nil, err
	}

	return eventOrder, nil
}

// GetEntityID returns the unique identifier of the EventOrder entity.
func (i *EventOrder) GetEntityID() string {
	return string(i.ID)
}

// ToMap converts the EventOrder entity to a map.
//
// Returns:
//   - A map representation of the EventOrder entity.
//   - An error if the conversion fails.
func (i *EventOrder) ToMap() (map[string]interface{}, error) {
	doc, err := regularTypesConversion.ConvertFromEntityToMapString(i)
	if err != nil {
		return nil, err
	}

	if id, ok := doc["_id"].(md5id.ID); ok {
		doc["_id"] = string(id)
	}
	if processingID, ok := doc["processing_id"].(uuid.ID); ok {
		doc["processing_id"] = string(processingID)
	}

	return doc, nil
}

// MapToEntity converts a map to an EventOrder entity.
//
// Parameters:
//   - doc: The map representation of an EventOrder.
//
// Returns:
//   - A pointer to the EventOrder entity.
//   - An error if the conversion fails.
func (i *EventOrder) MapToEntity(doc map[string]interface{}) (*EventOrder, error) {
	if id, ok := doc["_id"].(string); ok {
		doc["_id"] = md5id.ID(id)
	} else {
		return nil, errors.New("field _id has invalid type")
	}

	if processingID, ok := doc["processing_id"].(string); ok {
		doc["processing_id"] = uuid.ID(processingID)
	} else {
		return nil, errors.New("field processing_id has invalid type")
	}

	eventOrderEntity, err := regularTypesConversion.ConvertFromMapStringToEntity(reflect.TypeOf(EventOrder{}), doc)
	if err != nil {
		return nil, err
	}

	eventOrder := eventOrderEntity.(*EventOrder)

	return eventOrder, nil
}

// isValid validates the EventOrder entity.
//
// Returns:
//   - An error if any of the required fields are invalid.
func (i *EventOrder) isValid() error {
	if i.ID == "" {
		return ErrInvalidID
	}

	if i.Service == "" {
		return ErrInvalidService
	}

	if i.Source == "" {
		return ErrInvalidSource
	}

	if i.Provider == "" {
		return ErrInvalidProvider
	}

	if i.ProcessingID == "" {
		return ErrInvalidProcessingID
	}

	return nil
}
