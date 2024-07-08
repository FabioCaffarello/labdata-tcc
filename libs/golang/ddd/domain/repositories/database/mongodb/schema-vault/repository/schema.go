package repository

import (
	"context"
	"fmt"
	"log"
	"os"

	"libs/golang/ddd/domain/entities/schema-vault/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	schemaCollection = "schemas"
)

// SchemaRepository manages the operations on the schemas collection in MongoDB
type SchemaRepository struct {
	log        *log.Logger
	client     *mongo.Client
	database   string
	collection *mongo.Collection
}

// NewSchemaRepository creates a new SchemaRepository instance.
// It initializes the collection for the specified database.
//
// Parameters:
//   - client: The MongoDB client.
//   - database: The name of the database.
//
// Returns:
//   - A pointer to a ConfigRepository instance.
//
// Example:
//
//	client := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
//	repository := NewConfigRepository(client, "testdb")
func NewSchemaRepository(client *mongo.Client, database string) *SchemaRepository {
	return &SchemaRepository{
		log:        log.New(os.Stdout, "[SCHEMA-REPOSITORY] ", log.LstdFlags),
		client:     client,
		database:   database,
		collection: client.Database(database).Collection(schemaCollection),
	}
}

// getOneByID retrieves a single Schema document by its ID.
//
// Parameters:
//   - id: The ID of the Schema document.
//
// Returns:
//   - A pointer to the Schema entity.
//   - An error if the document is not found or cannot be decoded.
//
// Example:
//
//	schema, err := repository.getOneByID("60d5ec49e17e8e304c8f5310")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(schema)
func (r *SchemaRepository) getOneByID(id string) (*entity.Schema, error) {
	filter := bson.M{"_id": id}
	document := r.collection.FindOne(context.Background(), filter)
	if document.Err() != nil {
		return nil, document.Err()
	}

	var schema entity.Schema
	if err := document.Decode(&schema); err != nil {
		return nil, err
	}

	return &schema, nil
}

// Create inserts a new Schema document into the collection.
//
// Parameters:
//   - schema: The Schema entity to be inserted.
//
// Returns:
//   - An error if the document already exists or cannot be inserted.
//
// Example:
//
//	err := repository.Create(newSchema)
//	if err != nil {
//	    log.Fatal(err)
//	}
func (r *SchemaRepository) Create(schema *entity.Schema) error {
	r.log.Printf("Saving schema: %+v to collection: %s\n", schema, schemaCollection)
	schemaMap, err := schema.ToMap()
	if err != nil {
		return err
	}
	entityID := schema.GetEntityID()
	_, err = r.getOneByID(entityID)
	if err == nil {
		r.log.Printf("Schema with ID: %s already exists\n", entityID)
		return fmt.Errorf("schema with ID: %s already exists", entityID)
	}
	log.Printf("Schema map created: %+v\n", schemaMap)
	if jsonSchema, ok := schemaMap["json_schema"].(map[string]interface{}); ok {
		if required, ok := jsonSchema["required"].([]interface{}); ok {
			strRequired := make([]string, len(required))
			for i, v := range required {
				strRequired[i] = v.(string)
			}
			jsonSchema["required"] = strRequired
		}
		log.Printf("Required: %+v\n", jsonSchema["required"])
	}

	doc, err := r.collection.InsertOne(context.Background(), schemaMap)
	if err != nil {
		return err
	}
	r.log.Printf("Schema saved with ID: %s\n", doc.InsertedID)

	return nil
}

// FindByID retrieves a single Schema document by its ID.
//
// Parameters:
//   - id: The ID of the Schema document.
//
// Returns:
//   - A pointer to the Schema entity.
//   - An error if the document is not found or cannot be decoded.
//
// Example:
//
//	schema, err := repository.FindByID("60d5ec49e17e8e304c8f5310")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(schema)
func (r *SchemaRepository) FindByID(id string) (*entity.Schema, error) {
	return r.getOneByID(id)
}

// FindAll retrieves all Schema documents in the collection.
//
// Returns:
//   - A slice of pointers to Schema entities.
//   - An error if the query fails.
//
// Example:
//
//	schemas, err := repository.FindAll()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, schema := range schemas {
//	    fmt.Printf("Schema: %+v\n", schema)
//	}
func (r *SchemaRepository) FindAll() ([]*entity.Schema, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var schemas []*entity.Schema
	for cursor.Next(context.Background()) {
		var schema entity.Schema
		if err := cursor.Decode(&schema); err != nil {
			return nil, err
		}
		schemas = append(schemas, &schema)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(schemas) == 0 {
		return []*entity.Schema{}, nil
	}

	return schemas, nil
}

// Update modifies an existing Schema document in the collection.
//
// Parameters:
//   - schema: The Schema entity with updated data.
//
// Returns:
//   - An error if the document is not found or cannot be updated.
//
// Example:
//
//	err := repository.Update(updatedSchema)
//	if err != nil {
//	    log.Fatal(err)
//	}
func (r *SchemaRepository) Update(schema *entity.Schema) error {
	r.log.Printf("Updating schema: %+v\n in collection: %s\n", schema, schemaCollection)

	schemaID := schema.GetEntityID()
	schemaStored, err := r.getOneByID(schemaID)
	if err != nil {
		r.log.Printf("Schema with ID: %s not found\n", schemaID)
		return fmt.Errorf("schema with ID: %s not found", schemaID)
	}

	schema.SetCreatedAt(schemaStored.CreatedAt)

	schemaMap, err := schema.ToMap()
	if err != nil {
		return err
	}

	if jsonSchema, ok := schemaMap["json_schema"].(map[string]interface{}); ok {
		if required, ok := jsonSchema["required"].([]interface{}); ok {
			strRequired := make([]string, len(required))
			for i, v := range required {
				strRequired[i] = v.(string)
			}
			jsonSchema["required"] = strRequired
		}
	}

	filter := bson.M{"_id": schemaID}
	update := bson.M{"$set": schemaMap}
	_, err = r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	r.log.Printf("Schema updated with ID: %s\n", schemaID)
	return nil
}

// Delete removes a Schema document from the collection by its ID.
//
// Parameters:
//   - id: The ID of the Schema document to be deleted.
//
// Returns:
//   - An error if the document is not found or cannot be deleted.
//
// Example:
//
//	err := repository.Delete("60d5ec49e17e8e304c8f5310")
//	if err != nil {
//	    log.Fatal(err)
//	}
func (r *SchemaRepository) Delete(id string) error {
	r.log.Printf("Deleting schema with ID: %s from collection: %s\n", id, schemaCollection)
	filter := bson.M{"_id": id}
	_, err := r.getOneByID(id)
	if err != nil {
		r.log.Printf("Schema with ID: %s not found\n", id)
		return fmt.Errorf("schema with ID: %s not found", id)
	}
	_, err = r.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	r.log.Printf("Schema deleted with ID: %s\n", id)
	return nil
}

// find executes a query on the collection and returns the matching Schema documents.
//
// Parameters:
//   - query: The BSON query to execute.
//
// Returns:
//   - A slice of pointers to Schema entities.
//   - An error if the query fails or if the collection does not exist.
//
// Example:
//
//	query := bson.M{"service": "myservice"}
//	schemas, err := repository.find(query)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, schema := range schemas {
//	    fmt.Printf("Schema: %+v\n", schema)
//	}
func (r *SchemaRepository) find(query bson.M) ([]*entity.Schema, error) {
	log.Printf("Query: %+v\n", query)
	cursor, err := r.collection.Find(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var schemas []*entity.Schema
	for cursor.Next(context.Background()) {
		var schema entity.Schema
		if err := cursor.Decode(&schema); err != nil {
			return nil, err
		}
		schemas = append(schemas, &schema)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(schemas) == 0 {
		return []*entity.Schema{}, nil
	}

	return schemas, nil
}

// FindAllByServiceAndProvider retrieves all Schema documents that match the given provider and service.
//
// Parameters:
//   - service: The service name to match.
//   - provider: The provider name to match.
//
// Returns:
//   - A slice of pointers to Schema entities.
//   - An error if the query fails.
//
// Example:
//
//	schemas, err := repository.FindAllByServiceAndProvider("myprovider", "myservice")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, schema := range schemas {
//	    fmt.Printf("Schema: %+v\n", schema)
//	}
func (r *SchemaRepository) FindAllByServiceAndProvider(provider, service string) ([]*entity.Schema, error) {
	return r.find(bson.M{"provider": provider, "service": service})
}

// FindAllBySourceAndProvider retrieves all Schema documents that match the given provider and source.
//
// Parameters:
//   - source: The source name to match.
//
// Returns:
//   - A slice of pointers to Schema entities.
//   - An error if the query fails.
//
// Example:
//
//	schemas, err := repository.FindAllBySourceAndProvider("myprovider", "mysource")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, schema := range schemas {
//	    fmt.Printf("Schema: %+v\n", schema)
//	}
func (r *SchemaRepository) FindAllBySourceAndProvider(provider, source string) ([]*entity.Schema, error) {
	return r.find(bson.M{"provider": provider, "source": source})
}

// FindAllByServiceAndSourceAndProvider retrieves all Schema documents that match the given service, source, and provider.
//
// Parameters:
//   - service: The service name to match.
//   - source: The source name to match.
//   - provider: The provider name to match.
//
// Returns:
//   - A slice of pointers to Schema entities.
//   - An error if the query fails.
//
// Example:
//
//	schemas, err := repository.FindAllByServiceAndSourceAndProvider("myservice", "mysource", "myprovider")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, schema := range schemas {
//	    fmt.Printf("Schema: %+v\n", schema)
//	}
func (r *SchemaRepository) FindAllByServiceAndSourceAndProvider(service, source, provider string) ([]*entity.Schema, error) {
	return r.find(bson.M{"provider": provider, "service": service, "source": source})
}

// FindAllByServiceAndSourceAndProviderAndSchemaType retrieves one Schema document that matches the given service, source, provider, and schema type.
//
// Parameters:
//   - service: The service name to match.
//   - source: The source name to match.
//   - provider: The provider name to match.
//   - schemaType: The schema type to match.
//
// Returns:
//   - A pointer to the Schema entity.
//   - An error if the query fails.
//
// Example:
//
//	schema, err := repository.FindOneByServiceAndSourceAndProviderAndSchemaType("myservice", "mysource", "myprovider", "myschematype")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(schema)
func (r *SchemaRepository) FindOneByServiceAndSourceAndProviderAndSchemaType(service, source, provider, schemaType string) (*entity.Schema, error) {
	schemas, err := r.find(bson.M{"provider": provider, "service": service, "source": source, "schema_type": schemaType})
	if err != nil {
		return nil, err
	}

	if len(schemas) == 0 {
		return nil, fmt.Errorf("schema not found")
	}

	return schemas[0], nil
}
