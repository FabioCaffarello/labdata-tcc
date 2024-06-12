package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"libs/golang/ddd/domain/entities/config-vault/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	configCollection = "configs"
	dateLayout       = "2006-01-02 15:04:05"
)

// ConfigRepository manages the operations on the Config collection in MongoDB.
type ConfigRepository struct {
	log        *log.Logger
	client     *mongo.Client
	database   string
	collection *mongo.Collection
}

// NewConfigRepository creates a new ConfigRepository instance.
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
func NewConfigRepository(client *mongo.Client, database string) *ConfigRepository {
	return &ConfigRepository{
		log:        log.New(os.Stdout, "[CONFIG-REPOSITORY] ", log.LstdFlags),
		client:     client,
		database:   database,
		collection: client.Database(database).Collection(configCollection),
	}
}

// getOneByID retrieves a single Config document by its ID.
//
// Parameters:
//   - id: The ID of the Config document.
//
// Returns:
//   - A pointer to the Config entity.
//   - An error if the document is not found or cannot be decoded.
//
// Example:
//
//	config, err := repository.getOneByID("60d5ec49e17e8e304c8f5310")
//	if err != nil {
//	    log.Fatal(err)
//	}
func (r *ConfigRepository) getOneByID(id string) (*entity.Config, error) {
	filter := bson.M{"_id": id}
	document := r.collection.FindOne(context.Background(), filter)
	if document.Err() != nil {
		return nil, document.Err()
	}

	var config entity.Config
	if err := document.Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

// Create inserts a new Config document into the collection.
//
// Parameters:
//   - config: The Config entity to be inserted.
//
// Returns:
//   - An error if the document already exists or cannot be inserted.
//
// Example:
//
//	err := repository.Create(newConfig)
//	if err != nil {
//	    log.Fatal(err)
//	}
func (r *ConfigRepository) Create(config *entity.Config) error {
	r.log.Printf("Saving config: %+v to collection: %s\n", config, configCollection)
	configMap, err := config.ToMap()
	if err != nil {
		return err
	}
	entityID := config.GetEntityID()
	_, err = r.getOneByID(entityID)
	if err == nil {
		r.log.Printf("Config with ID: %s already exists\n", entityID)
		return fmt.Errorf("config with ID: %s already exists", entityID)
	}

	doc, err := r.collection.InsertOne(context.Background(), configMap)
	if err != nil {
		return err
	}
	r.log.Printf("Config saved with ID: %s\n", doc.InsertedID)

	return nil
}

// FindByID retrieves a single Config document by its ID.
//
// Parameters:
//   - id: The ID of the Config document.
//
// Returns:
//   - A pointer to the Config entity.
//   - An error if the document is not found or cannot be decoded.
//
// Example:
//
//	config, err := repository.FindByID("60d5ec49e17e8e304c8f5310")
//	if err != nil {
//	    log.Fatal(err)
//	}
func (r *ConfigRepository) FindByID(id string) (*entity.Config, error) {
	return r.getOneByID(id)
}

// FindAll retrieves all Config documents in the collection.
//
// Returns:
//   - A slice of pointers to Config entities.
//   - An error if the query fails.
//
// Example:
//
//	configs, err := repository.FindAll()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, config := range configs {
//	    fmt.Printf("Config: %+v\n", config)
//	}
func (r *ConfigRepository) FindAll() ([]*entity.Config, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var configs []*entity.Config
	for cursor.Next(context.Background()) {
		var config entity.Config
		if err := cursor.Decode(&config); err != nil {
			return nil, err
		}
		configs = append(configs, &config)
	}

	return configs, nil
}

// Update modifies an existing Config document in the collection.
//
// Parameters:
//   - config: The Config entity with updated data.
//
// Returns:
//   - An error if the document is not found or cannot be updated.
//
// Example:
//
//	err := repository.Update(updatedConfig)
//	if err != nil {
//	    log.Fatal(err)
//	}
func (r *ConfigRepository) Update(config *entity.Config) error {
	r.log.Printf("Updating config: %+v\n", config)

	configID := config.GetEntityID()
	configStored, err := r.getOneByID(configID)
	if err != nil {
		r.log.Printf("Config with ID: %s not found\n", configID)
		return fmt.Errorf("config with ID: %s not found", configID)
	}

	config.SetCreatedAt(configStored.CreatedAt)
	config.SetUpdatedAt(time.Now().Format(dateLayout))

	configMap, err := config.ToMap()
	if err != nil {
		return err
	}

	filter := bson.M{"_id": configID}
	update := bson.M{"$set": configMap}
	_, err = r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	r.log.Printf("Config updated with ID: %s\n", configID)
	return nil
}

// Delete removes a Config document from the collection by its ID.
//
// Parameters:
//   - id: The ID of the Config document to be deleted.
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
func (r *ConfigRepository) Delete(id string) error {
	r.log.Printf("Deleting config with ID: %s\n", id)
	filter := bson.M{"_id": id}
	_, err := r.getOneByID(id)
	if err != nil {
		r.log.Printf("Config with ID: %s not found\n", id)
		return fmt.Errorf("config with ID: %s not found", id)
	}
	_, err = r.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	r.log.Printf("Config deleted with ID: %s\n", id)
	return nil
}

// find executes a query on the collection and returns the matching Config documents.
//
// Parameters:
//   - query: The BSON query to execute.
//
// Returns:
//   - A slice of pointers to Config entities.
//   - An error if the query fails or if the collection does not exist.
//
// Example:
//
//	query := bson.M{"service": "myservice"}
//	configs, err := repository.find(query)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, config := range configs {
//	    fmt.Printf("Config: %+v\n", config)
//	}
func (r *ConfigRepository) find(query bson.M) ([]*entity.Config, error) {
	cursor, err := r.collection.Find(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var configs []*entity.Config = []*entity.Config{}
	for cursor.Next(context.Background()) {
		var config entity.Config
		if err := cursor.Decode(&config); err != nil {
			return nil, err
		}
		configs = append(configs, &config)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return configs, nil
}

// FindAllByService retrieves all Config documents that match the given service.
//
// Parameters:
//   - service: The service name to match.
//
// Returns:
//   - A slice of pointers to Config entities.
//   - An error if the query fails.
//
// Example:
//
//	configs, err := repository.FindAllByService("myservice")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, config := range configs {
//	    fmt.Printf("Config: %+v\n", config)
//	}
func (r *ConfigRepository) FindAllByService(service string) ([]*entity.Config, error) {
	query := bson.M{"service": service}
	return r.find(query)
}

// FindAllBySource retrieves all Config documents that match the given source.
//
// Parameters:
//   - source: The source name to match.
//
// Returns:
//   - A slice of pointers to Config entities.
//   - An error if the query fails.
//
// Example:
//
//	configs, err := repository.FindAllBySource("mysource")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, config := range configs {
//	    fmt.Printf("Config: %+v\n", config)
//	}
func (r *ConfigRepository) FindAllBySource(source string) ([]*entity.Config, error) {
	query := bson.M{"source": source}
	return r.find(query)
}

// FindAllByServiceAndSource retrieves all Config documents that match the given service and source.
//
// Parameters:
//   - service: The service name to match.
//   - source: The source name to match.
//
// Returns:
//   - A slice of pointers to Config entities.
//   - An error if the query fails.
//
// Example:
//
//	configs, err := repository.FindAllByServiceAndSource("myservice", "mysource")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, config := range configs {
//	    fmt.Printf("Config: %+v\n", config)
//	}
func (r *ConfigRepository) FindAllByServiceAndSource(service, source string) ([]*entity.Config, error) {
	query := bson.M{"service": service, "source": source}
	return r.find(query)
}

// FindAllByServiceAndSourceAndProvider retrieves all Config documents that match the given service, source, and provider.
//
// Parameters:
//   - service: The service name to match.
//   - source: The source name to match.
//   - provider: The provider name to match.
//
// Returns:
//   - A slice of pointers to Config entities.
//   - An error if the query fails.
//
// Example:
//
//	configs, err := repository.FindAllByServiceAndSourceAndProvider("myservice", "mysource", "myprovider")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, config := range configs {
//	    fmt.Printf("Config: %+v\n", config)
//	}
func (r *ConfigRepository) FindAllByServiceAndSourceAndProvider(service, source, provider string) ([]*entity.Config, error) {
	query := bson.M{"service": service, "source": source, "provider": provider}
	return r.find(query)
}

// FindAllByServiceAndProviderAndActive retrieves all Config documents that match the given service, provider, and active status.
//
// Parameters:
//   - service: The service name to match.
//   - provider: The provider name to match.
//   - active: The active status to match.
//
// Returns:
//   - A slice of pointers to Config entities.
//   - An error if the query fails.
//
// Example:
//
//	configs, err := repository.FindAllByServiceAndProviderAndActive("myservice", "myprovider", true)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, config := range configs {
//	    fmt.Printf("Config: %+v\n", config)
//	}
func (r *ConfigRepository) FindAllByServiceAndProviderAndActive(service, provider string, active bool) ([]*entity.Config, error) {
	query := bson.M{"service": service, "provider": provider, "active": active}
	return r.find(query)
}

// FindAllByDependsOn retrieves all Config documents that have dependencies matching the given service and source.
//
// Parameters:
//   - service: The service name to match in dependencies.
//   - source: The source name to match in dependencies.
//
// Returns:
//   - A slice of pointers to Config entities.
//   - An error if the query fails.
//
// Example:
//
//	configs, err := repository.FindAllByDependsOn("dep_service", "dep_source")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, config := range configs {
//	    fmt.Printf("Config: %+v\n", config)
//	}
func (r *ConfigRepository) FindAllByDependsOn(service, source string) ([]*entity.Config, error) {
	query := bson.M{"depends_on": bson.M{"$elemMatch": bson.M{"service": service, "source": source}}}
	return r.find(query)
}
