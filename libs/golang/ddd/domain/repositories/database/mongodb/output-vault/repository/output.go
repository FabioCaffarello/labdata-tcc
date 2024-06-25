package repository

import (
	"context"
	"fmt"
	"libs/golang/ddd/domain/entities/output-vault/entity"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	schemaCollection = "outputs"
)

// OutputRepository manages the operations on the outputs collection in MongoDB
type OutputRepository struct {
	log        *log.Logger
	client     *mongo.Client
	database   string
	collection *mongo.Collection
}

// NewOutputRepository creates a new OutputRepository instance.
// It initializes the collection for the specified database.
//
// Parameters:
//   - client: The MongoDB client.
//   - database: The name of the database.
//
// Returns:
//   - A pointer to a OutputRepository instance.
//
// Example:
//
//	client := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
//	repository := NewOutputRepository(client, "testdb")
func NewOutputRepository(client *mongo.Client, database string) *OutputRepository {
	return &OutputRepository{
		log:        log.New(log.Writer(), "[OUTPUT-REPOSITORY] ", log.LstdFlags),
		client:     client,
		database:   database,
		collection: client.Database(database).Collection(schemaCollection),
	}
}

// getOneByID retrieves a single Output document by its ID.
//
// Parameters:
//   - id: The ID of the Output document.
//
// Returns:
//   - A pointer to the Output entity.
//   - An error if the document is not found or cannot be decoded.
//
// Example:
//
//	output, err := repository.getOneByID("5f7b3b3b7b3b3b3b3b3b3b3b")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(output)
func (r *OutputRepository) getOneByID(id string) (*entity.Output, error) {
	filter := bson.M{"_id": id}
	document := r.collection.FindOne(context.Background(), filter)
	if document.Err() != nil {
		return nil, document.Err()
	}

	var output entity.Output
	if err := document.Decode(&output); err != nil {
		return nil, err
	}

	return &output, nil
}

// Create inserts a new Output document into the collection.
//
// Parameters:
//   - output: The Output entity to insert.
//
// Returns:
//   - An error if the document already exists or cannot be inserted.
//
// Example:
//
//	err := repository.Create(newOutput)
//	if err != nil {
//		log.Fatal(err)
//	}
func (r *OutputRepository) Create(output *entity.Output) error {
	r.log.Printf("Saving output: %+v to collection: %s\n", output, schemaCollection)
	outputMap, err := output.ToMap()
	if err != nil {
		return err
	}
	entityID := output.GetEntityID()
	_, err = r.getOneByID(entityID)
	if err == nil {
		r.log.Printf("Output with ID: %s already exists\n", entityID)
		return fmt.Errorf("output with ID: %s already exists", entityID)
	}

	doc, err := r.collection.InsertOne(context.Background(), outputMap)
	if err != nil {
		return err
	}
	r.log.Printf("Output saved with ID: %s\n", doc.InsertedID)

	return nil
}

// FindByID retrieves a single Output document by its ID.
//
// Parameters:
//   - id: The ID of the Output document.
//
// Returns:
//   - A pointer to the Output entity.
//   - An error if the document is not found or cannot be decoded.
//
// Example:
//
//	output, err := repository.FindByID("5f7b3b3b7b3b3b3b3b3b3b3b")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(output)
func (r *OutputRepository) FindByID(id string) (*entity.Output, error) {
	return r.getOneByID(id)
}

// FindAll retrieves all Output documents from the collection.
//
// Returns:
//   - A slice of Output entities.
//   - An error if the documents cannot be decoded.
//
// Example:
//
//	outputs, err := repository.FindAll()
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, output := range outputs {
//		fmt.Println(output)
//	}
func (r *OutputRepository) FindAll() ([]*entity.Output, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var outputs []*entity.Output
	for cursor.Next(context.Background()) {
		var output entity.Output
		if err := cursor.Decode(&output); err != nil {
			return nil, err
		}
		outputs = append(outputs, &output)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(outputs) == 0 {
		return []*entity.Output{}, nil
	}

	return outputs, nil
}

// Update modifies an existing Output document in the collection.
//
// Parameters:
//   - output: The Output entity to update.
//
// Returns:
//   - An error if the document does not exist or cannot be updated.
//
// Example:
//
//	err := repository.Update(updatedOutput)
//	if err != nil {
//		log.Fatal(err)
//	}
func (r *OutputRepository) Update(output *entity.Output) error {
	r.log.Printf("Updating output: %+v in collection: %s\n", output, schemaCollection)

	outputID := output.GetEntityID()
	outputStored, err := r.getOneByID(outputID)
	if err != nil {
		r.log.Printf("Output with ID: %s not found\n", outputID)
		return fmt.Errorf("output with ID: %s not found", outputID)
	}

	output.SetCreatedAt(outputStored.CreatedAt)

	outputMap, err := output.ToMap()
	if err != nil {
		return err
	}

	filter := bson.M{"_id": outputID}
	update := bson.M{"$set": outputMap}
	_, err = r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	r.log.Printf("Output updated with ID: %s\n", outputID)
	return nil
}

// Delete removes a Output document from the collection by its ID.
//
// Parameters:
//   - id: The ID of the Output document.
//
// Returns:
//   - An error if the document does not exist or cannot be removed.
//
// Example:
//
//	err := repository.Delete("5f7b3b3b7b3b3b3b3b3b3b3b")
//	if err != nil {
//		log.Fatal(err)
//	}
func (r *OutputRepository) Delete(id string) error {
	r.log.Printf("Deleting output with ID: %s from collection: %s\n", id, schemaCollection)
	filter := bson.M{"_id": id}
	_, err := r.getOneByID(id)
	if err != nil {
		r.log.Printf("Output with ID: %s not found\n", id)
		return fmt.Errorf("output with ID: %s not found", id)
	}
	_, err = r.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	r.log.Printf("Output deleted with ID: %s\n", id)
	return nil
}

// find executes a query on the collection and returns the matching Outputs documents.
//
// Parameters:
//   - query: The BSON query to execute.
//
// Returns:
//   - A slice of Output entities.
//   - An error if the documents cannot be decoded.
//
// Example:
//
//	outputs, err := repository.find(bson.M{"service": "test"})
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, output := range outputs {
//		fmt.Println(output)
//	}
func (r *OutputRepository) find(query bson.M) ([]*entity.Output, error) {
	cursor, err := r.collection.Find(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var outputs []*entity.Output
	for cursor.Next(context.Background()) {
		var output entity.Output
		if err := cursor.Decode(&output); err != nil {
			return nil, err
		}
		outputs = append(outputs, &output)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(outputs) == 0 {
		return []*entity.Output{}, nil
	}

	return outputs, nil
}

// FindAllByServiceAndProvider retrieves all Outputs documents that match the given provider and service.
//
// Parameters:
//   - service: The service name to match.
//   - provider: The provider name to match.
//
// Returns:
//   - A slice of pointers to Output entities.
//   - An error if the query fails.
//
// Example:
//
//	outputs, err := repository.FindAllByServiceAndProvider("myprovider", "myservice")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, output := range outputs {
//		fmt.Println(output)
//	}
func (r *OutputRepository) FindAllByServiceAndProvider(provider, service string) ([]*entity.Output, error) {
	return r.find(bson.M{"provider": provider, "service": service})
}

// FindAllBySourceAndProvider retrieves all Outputs documents that match the given provider and source.
//
// Parameters:
//   - source: The source name to match.
//   - provider: The provider name to match.
//
// Returns:
//   - A slice of pointers to Output entities.
//   - An error if the query fails.
//
// Example:
//
//	outputs, err := repository.FindAllBySourceAndProvider("myprovider", "mysource")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, output := range outputs {
//		fmt.Println(output)
//	}
func (r *OutputRepository) FindAllBySourceAndProvider(provider, source string) ([]*entity.Output, error) {
	return r.find(bson.M{"provider": provider, "source": source})
}

// FindAllByServiceAndSourceAndProvider retrieves all Outputs documents that match the given provider, service and source.
//
// Parameters:
//   - service: The service name to match.
//   - source: The source name to match.
//   - provider: The provider name to match.
//
// Returns:
//   - A slice of pointers to Output entities.
//   - An error if the query fails.
//
// Example:
//
//	 outputs, err := repository.FindAllByServiceAndSourceAndProvider("myprovider", "myservice", "mysource")
//		if err != nil {
//		    log.Fatal(err)
//		}
//		for _, output := range outputs {
//			fmt.Println(output)
//		}
func (r *OutputRepository) FindAllByServiceAndSourceAndProvider(provider, service, source string) ([]*entity.Output, error) {
	return r.find(bson.M{"provider": provider, "service": service, "source": source})
}
