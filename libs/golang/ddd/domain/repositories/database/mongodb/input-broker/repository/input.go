package repository

import (
	"context"
	"fmt"
	"libs/golang/ddd/domain/entities/input-broker/entity"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	schemaCollection = "inputs"
)

// InputRepository manages the operations on the inputs collection in MongoDB
type InputRepository struct {
	log        *log.Logger
	client     *mongo.Client
	database   string
	collection *mongo.Collection
}

// NewInputRepository creates a new InputRepository instance.
// It initializes the collection for the specified database.
//
// Parameters:
//   - client: The MongoDB client.
//   - database: The name of the database.
//
// Returns:
//   - A pointer to a InputRepository instance.
//
// Example:
//
//	client := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
//	repository := NewInputRepository(client, "testdb")
func NewInputRepository(client *mongo.Client, database string) *InputRepository {
	return &InputRepository{
		log:        log.New(log.Writer(), "[INPUT-REPOSITORY] ", log.LstdFlags),
		client:     client,
		database:   database,
		collection: client.Database(database).Collection(schemaCollection),
	}
}

// getOneByID retrieves a single Input document by its ID.
//
// Parameters:
//   - id: The ID of the Input document.
//
// Returns:
//   - A pointer to the Input entity.
//   - An error if the document is not found or cannot be decoded.
//
// Example:
//
//	input, err := repository.getOneByID("5f7b3b3b7b3b3b3b3b3b3b3b")
//	if err != nil {
func (r *InputRepository) getOneByID(id string) (*entity.Input, error) {
	filter := bson.M{"_id": id}
	document := r.collection.FindOne(context.Background(), filter)

	if document.Err() != nil {
		return nil, document.Err()
	}

	var input entity.Input
	if err := document.Decode(&input); err != nil {
		return nil, err
	}

	return &input, nil
}

// Create inserts a new Input document into the collection.
//
// Parameters:
//   - input: The Input entity to insert.
//
// Returns:
//   - An error if the document already exists or cannot be inserted.
//
// Example:
//
//	err := repository.Create(newInput)
//	if err != nil {
//		log.Fatal(err)
//	}
func (r *InputRepository) Create(input *entity.Input) error {
	r.log.Printf("Saving input: %+v to collection: %s\n", input, schemaCollection)
	inputMap, err := input.ToMap()
	if err != nil {
		return err
	}
	entityID := input.GetEntityID()
	_, err = r.getOneByID(entityID)
	if err == nil {
		r.log.Printf("Input with ID: %s already exists\n", entityID)
		return fmt.Errorf("input with ID: %s already exists", entityID)
	}

	doc, err := r.collection.InsertOne(context.Background(), inputMap)
	if err != nil {
		return err
	}
	r.log.Printf("Inserted document with ID: %s\n", doc.InsertedID)

	return nil
}

// FindByID retrieves a single Input document by its ID.
//
// Parameters:
//   - id: The ID of the Input document.
//
// Returns:
//   - A pointer to the Input entity.
//   - An error if the document is not found or cannot be decoded.
//
// Example:
//
//	inputs, err := repository.FindByID()
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Input: %+v\n", input)
func (r *InputRepository) FindByID(id string) (*entity.Input, error) {
	return r.getOneByID(id)
}

// FindAll retrieves all Input documents from the collection.
//
// Returns:
//   - A slice of Input entities.
//   - An error if the documents cannot be decoded.
//
// Example:
//
//	inputs, err := repository.FindAll()
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, input := range inputs {
//		fmt.Println(input)
//	}
func (r *InputRepository) FindAll() ([]*entity.Input, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var inputs []*entity.Input
	for cursor.Next(context.Background()) {
		var input entity.Input
		if err := cursor.Decode(&input); err != nil {
			return nil, err
		}
		inputs = append(inputs, &input)
	}

	return inputs, nil
}

// Update modifies an existing Input document in the collection.
//
// Parameters:
//   - input: The Input entity to update.
//
// Returns:
//   - An error if the document does not exist or cannot be updated.
//
// Example:
//
//	err := repository.Update(updatedInput)
//	if err != nil {
//		log.Fatal(err)
//	}
func (r *InputRepository) Update(input *entity.Input) error {
	r.log.Printf("Updating input: %+v in collection: %s\n", input, schemaCollection)
	inputID := input.GetEntityID()
	inputStored, err := r.getOneByID(inputID)
	if err != nil {
		r.log.Printf("Input with ID: %s does not exist\n", inputID)
		return fmt.Errorf("input with ID: %s does not exist", inputID)
	}

	input.SetCreatedAt(inputStored.CreatedAt)

	inputMap, err := input.ToMap()
	if err != nil {
		return err
	}

	filter := bson.M{"_id": inputID}
	update := bson.M{"$set": inputMap}
	_, err = r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	r.log.Printf("Updated input with ID: %s\n", inputID)
	return nil
}

// Delete removes an existing Input document from the collection by its ID.
//
// Parameters:
//   - id: The ID of the Input document.
//
// Returns:
//   - An error if the document does not exist or cannot be deleted.
//
// Example:
//
//	err := repository.Delete("5f7b3b3b7b3b3b3b3b3b3b3b")
//	if err != nil {
//		log.Fatal(err)
//	}
func (r *InputRepository) Delete(id string) error {
	r.log.Printf("Deleting input with ID: %s from collection: %s\n", id, schemaCollection)
	filter := bson.M{"_id": id}
	_, err := r.getOneByID(id)
	if err != nil {
		r.log.Printf("Input with ID: %s not found\n", id)
		return fmt.Errorf("input with ID: %s not found", id)
	}
	_, err = r.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	r.log.Printf("Input deleted with ID: %s\n", id)
	return nil
}

// find executes a query on the collection and returns the matching Inputs documents.
//
// Parameters:
//   - query: The BSON query to execute.
//
// Returns:
//   - A slice of Input entities.
//   - An error if the documents cannot be decoded.
//
// Example:
//
//	inputs, err := repository.find(bson.M{"status.code": 0})
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, input := range inputs {
//		fmt.Println(input)
//	}
func (r *InputRepository) find(query bson.M) ([]*entity.Input, error) {
	cursor, err := r.collection.Find(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var inputs []*entity.Input
	for cursor.Next(context.Background()) {
		var input entity.Input
		if err := cursor.Decode(&input); err != nil {
			return nil, err
		}
		inputs = append(inputs, &input)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(inputs) == 0 {
		return []*entity.Input{}, nil
	}

	return inputs, nil
}

// FindAllByStatus retrieves all Input documents from the collection with the specified status.
//
// Parameters:
//   - status: The status code.
//
// Returns:
//   - A slice of Input entities.
//   - An error if the documents cannot be decoded.
//
// Example:
//
//	inputs, err := repository.FindAllByStatus(0)
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, input := range inputs {
//		fmt.Println(input)
//	}
func (r *InputRepository) FindAllByStatus(status int) ([]*entity.Input, error) {
	query := bson.M{"status.code": status}
	return r.find(query)
}

// FindAllByServiceAndProvider retrieves all Input documents that match the given provider and service.
//
// Parameters:
//   - provider: The provider name to match.
//   - service: The service name to match.
//
// Returns:
//   - A slice of Input entities.
//   - An error if the documents cannot be decoded.
//
// Example:
//
//	inputs, err := repository.FindAllByServiceAndProvider("myprovider", "myservice")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, input := range inputs {
//		fmt.Println(input)
//	}
func (r *InputRepository) FindAllByServiceAndProvider(provider, service string) ([]*entity.Input, error) {
	query := bson.M{"metadata.provider": provider, "metadata.service": service}
	return r.find(query)
}

// FindAllBySourceAndProvider retrieves all Input documents that match the given provider and source.
//
// Parameters:
//   - provider: The provider name to match.
//   - source: The source name to match.
//
// Returns:
//   - A slice of Input entities.
//   - An error if the documents cannot be decoded.
//
// Example:
//
//	inputs, err := repository.FindAllBySourceAndProvider("myprovider", "mysource")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, input := range inputs {
//		fmt.Println(input)
//	}
func (r *InputRepository) FindAllBySourceAndProvider(provider, service string) ([]*entity.Input, error) {
	query := bson.M{"metadata.provider": provider, "metadata.source": service}
	return r.find(query)
}

// FindAllByServiceAndSourceAndProvider retrieves all Input documents that match the given provider, service and source.
//
// Parameters:
//   - provider: The provider name to match.
//   - service: The service name to match.
//   - source: The source name to match.
//
// Returns:
//   - A slice of Input entities.
//   - An error if the documents cannot be decoded.
//
// Example:
//
//	inputs, err := repository.FindAllByServiceAndSourceAndProvider("myprovider", "myservice", "mysource")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, input := range inputs {
//		fmt.Println(input)
//	}
func (r *InputRepository) FindAllByServiceAndSourceAndProvider(provider, service, source string) ([]*entity.Input, error) {
	query := bson.M{"metadata.provider": provider, "metadata.service": service, "metadata.source": source}
	return r.find(query)
}

// FindAllByStatusAndServiceAndProvider retrieves all Input documents that match the given provider, service and status.
//
// Parameters:
//   - provider: The provider name to match.
//   - service: The service name to match.
//   - status: The status code.
//
// Returns:
//   - A slice of Input entities.
//   - An error if the documents cannot be decoded.
//
// Example:
//
//	inputs, err := repository.FindAllByStatusAndServiceAndProvider("myprovider", "myservice", 0)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, input := range inputs {
//		fmt.Println(input)
//	}
func (r *InputRepository) FindAllByStatusAndServiceAndProvider(provider, service string, status int) ([]*entity.Input, error) {
	query := bson.M{"metadata.provider": provider, "metadata.service": service, "status.code": status}
	return r.find(query)
}

// FindAllByStatusAndSourceAndProvider retrieves all Input documents that match the given provider, source and status.
//
// Parameters:
//   - provider: The provider name to match.
//   - source: The source name to match.
//   - status: The status code.
//
// Returns:
//   - A slice of Input entities.
//   - An error if the documents cannot be decoded.
//
// Example:
//
//	inputs, err := repository.FindAllByStatusAndSourceAndProvider("myprovider", "mysource", 0)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, input := range inputs {
//		fmt.Println(input)
//	}
func (r *InputRepository) FindAllByStatusAndSourceAndProvider(provider, source string, status int) ([]*entity.Input, error) {
	query := bson.M{"metadata.provider": provider, "metadata.source": source, "status.code": status}
	return r.find(query)
}

// FindAllByStatusAndServiceAndSourceAndProvider retrieves all Input documents that match the given provider, service, source and status.
//
// Parameters:
//   - provider: The provider name to match.
//   - service: The service name to match.
//   - source: The source name to match.
//   - status: The status code.
//
// Returns:
//   - A slice of Input entities.
//   - An error if the documents cannot be decoded.
//
// Example:
//
//	inputs, err := repository.FindAllByStatusAndServiceAndSourceAndProvider("myprovider", "myservice", "mysource", 0)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, input := range inputs {
//		fmt.Println(input)
//	}
func (r *InputRepository) FindAllByStatusAndServiceAndSourceAndProvider(provider, service, source string, status int) ([]*entity.Input, error) {
	query := bson.M{"metadata.provider": provider, "metadata.service": service, "metadata.source": source, "status.code": status}
	return r.find(query)
}
