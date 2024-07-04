package repository

import (
	"fmt"
	"libs/golang/clients/resources/go-docdb/client"
	"libs/golang/ddd/domain/entities/events-router/entity"
	"log"
)

var (
	schemaCollection = "events-order"
)

// EventOrderRepository is a repository for EventOrder entities.
type EventOrderRepository struct {
	log            *log.Logger
	client         *client.Client
	database       string
	collectionName string
}

// NewEventOrderRepository creates a new instance of EventOrderRepository.
//
// Parameters:
//   - client: The client instance to interact with the document-based database.
//   - database: The name of the database.
//
// Returns:
//   - A pointer to the newly created EventOrderRepository instance.
func NewEventOrderRepository(
	client *client.Client,
	database string,
) *EventOrderRepository {
	inMemoryRepository := &EventOrderRepository{
		log:            log.New(log.Writer(), "[EVENT-ORDER-REPOSITORY] ", log.LstdFlags),
		client:         client,
		database:       database,
		collectionName: schemaCollection,
	}
	inMemoryRepository.client.CreateCollection(inMemoryRepository.collectionName)
	return inMemoryRepository
}

// getOneByID retrieves a single EventOrder by its ID.
//
// Parameters:
//   - id: The ID of the EventOrder to retrieve.
//
// Returns:
//   - A pointer to the EventOrder if found, otherwise nil.
//   - An error if the document is not found or cannot be mapped to an EventOrder entity.
func (r *EventOrderRepository) getOneByID(id string) (*entity.EventOrder, error) {
	document, err := r.client.FindOne(r.collectionName, id)
	if err != nil {
		return nil, err
	}
	result := &entity.EventOrder{}
	result, err = result.MapToEntity(document)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Create saves a new EventOrder to the repository.
//
// Parameters:
//   - eventOrder: The EventOrder entity to save.
//
// Returns:
//   - An error if the EventOrder already exists or cannot be saved.
func (r *EventOrderRepository) Create(eventOrder *entity.EventOrder) error {
	r.log.Printf("Save event order: %+v to collection %s\n", eventOrder, r.collectionName)
	eventOrderMap, err := eventOrder.ToMap()
	if err != nil {
		return err
	}
	entityID := eventOrder.GetEntityID()
	_, err = r.getOneByID(entityID)
	if err == nil {
		r.log.Printf("Event order with ID %s already exists\n", entityID)
		return fmt.Errorf("Event order with ID %s already exists", entityID)
	}

	err = r.client.InsertOne(r.collectionName, eventOrderMap)
	if err != nil {
		return err
	}

	r.log.Printf("Config saved with ID: %s\n", entityID)
	return nil
}

// FindByID retrieves an EventOrder by its ID.
//
// Parameters:
//   - id: The ID of the EventOrder to retrieve.
//
// Returns:
//   - A pointer to the EventOrder if found, otherwise nil.
//   - An error if the document is not found or cannot be mapped to an EventOrder entity.
func (r *EventOrderRepository) FindByID(id string) (*entity.EventOrder, error) {
	return r.getOneByID(id)
}

// FindAll retrieves all EventOrders from the repository.
//
// Returns:
//   - A slice of pointers to EventOrder entities.
//   - An error if the documents cannot be retrieved or mapped to EventOrder entities.
func (r *EventOrderRepository) FindAll() ([]*entity.EventOrder, error) {
	documents, err := r.client.FindAll(r.collectionName)
	if err != nil {
		return nil, err
	}

	var result []*entity.EventOrder
	for _, document := range documents {
		eventOrder := &entity.EventOrder{}
		eventOrder, err = eventOrder.MapToEntity(document)
		if err != nil {
			return nil, err
		}
		result = append(result, eventOrder)
	}

	return result, nil
}

// Delete removes an EventOrder by its ID.
//
// Parameters:
//   - id: The ID of the EventOrder to delete.
//
// Returns:
//   - An error if the document cannot be found or deleted.
func (r *EventOrderRepository) Delete(id string) error {
	err := r.client.DeleteOne(r.collectionName, id)
	if err != nil {
		return err
	}
	return nil
}
