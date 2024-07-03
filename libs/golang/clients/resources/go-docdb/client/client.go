package client

import (
	"errors"
	"libs/golang/database/go-docdb/database"
)

// Client provides a higher-level interface for interacting with the
// in-memory document-based database. It wraps around the InMemoryDocBD
// instance to perform various database operations.
type Client struct {
	db *database.InMemoryDocBD
}

// NewClient creates a new Client instance.
//
// Parameters:
//  - db: A pointer to an InMemoryDocBD instance.
//
// Returns:
//  - A pointer to the newly created Client instance.
func NewClient(
	db *database.InMemoryDocBD,
) *Client {
	return &Client{
		db: db,
	}
}

// getCollection retrieves a collection by its name.
//
// Parameters:
//  - collectionName: The name of the collection to retrieve.
//
// Returns:
//  - A pointer to the Collection if found, otherwise nil.
//  - An error if the collection does not exist.
func (c *Client) getCollection(
	collectionName string,
) (*database.Collection, error) {
	collection, err := c.db.GetCollection(collectionName)
	if err != nil {
		return nil, err
	}
	return collection, nil
}

// CreateCollection creates a new collection with the specified name.
//
// Parameters:
//  - collectionName: The name of the collection to create.
//
// Returns:
//  - An error if the collection already exists.
func (c *Client) CreateCollection(
	collectionName string,
) error {
	err := c.db.CreateCollection(collectionName)
	if err != nil {
		return err
	}
	return nil
}

// DropCollection removes a collection by its name.
//
// Parameters:
//  - collectionName: The name of the collection to remove.
//
// Returns:
//  - An error if the collection does not exist.
func (c *Client) DropCollection(
	collectionName string,
) error {
	err := c.db.DropCollection(collectionName)
	if err != nil {
		return err
	}
	return nil
}

// ListCollections returns a list of all collection names.
//
// Returns:
//  - A slice of strings containing the names of all collections.
func (c *Client) ListCollections() []string {
	return c.db.ListCollections()
}

// ConvertToDocument converts a map into a Document type.
//
// Parameters:
//  - document: The map to convert.
//
// Returns:
//  - The converted Document.
//  - An error if the document is nil or does not contain an "_id" field.
func (c *Client) ConvertToDocument(
	document map[string]interface{},
) (database.Document, error) {
	if document == nil {
		return nil, errors.New("document is nil")
	}
	if _, ok := document["_id"]; !ok {
		return nil, errors.New("_id field is required")
	}
	return database.Document(document), nil
}

// InsertOne inserts a new document into the specified collection.
//
// Parameters:
//  - collectionName: The name of the collection to insert the document into.
//  - document: The document to insert.
//
// Returns:
//  - An error if the collection does not exist or if the document is invalid.
func (c *Client) InsertOne(
	collectionName string,
	document map[string]interface{},
) error {
	collection, err := c.getCollection(collectionName)
	if err != nil {
		return err
	}
	doc, err := c.ConvertToDocument(document)
	if err != nil {
		return err
	}
	return collection.InsertOne(doc)
}

// FindOne retrieves a document by its ID from the specified collection.
//
// Parameters:
//  - collectionName: The name of the collection to search.
//  - id: The ID of the document to retrieve.
//
// Returns:
//  - The document if found.
//  - An error if the collection or document does not exist.
func (c *Client) FindOne(
	collectionName string,
	id string,
) (map[string]interface{}, error) {
	collection, err := c.getCollection(collectionName)
	if err != nil {
		return nil, err
	}
	doc, err := collection.FindOne(id)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}(doc), nil
}

// FindAll retrieves all documents from the specified collection.
//
// Parameters:
//  - collectionName: The name of the collection to search.
//
// Returns:
//  - A slice of documents if found.
//  - An error if the collection does not exist.
func (c *Client) FindAll(
	collectionName string,
) ([]map[string]interface{}, error) {
	collection, err := c.getCollection(collectionName)
	if err != nil {
		return nil, err
	}
	docs := collection.FindAll()
	documents := make([]map[string]interface{}, 0, len(docs))
	for _, doc := range docs {
		documents = append(documents, map[string]interface{}(doc))
	}
	return documents, nil
}

// Find searches for documents matching a given query in the specified collection.
//
// Parameters:
//  - collectionName: The name of the collection to search.
//  - filter: The query criteria to match documents against.
//
// Returns:
//  - A slice of documents that match the query.
//  - An error if the collection does not exist.
func (c *Client) Find(
	collectionName string,
	filter map[string]interface{},
) ([]map[string]interface{}, error) {
	collection, err := c.getCollection(collectionName)
	if err != nil {
		return nil, err
	}
	docs := collection.Find(filter)
	documents := make([]map[string]interface{}, 0, len(docs))
	for _, doc := range docs {
		documents = append(documents, map[string]interface{}(doc))
	}
	return documents, nil
}

// UpdateOne updates a document by its ID in the specified collection.
//
// Parameters:
//  - collectionName: The name of the collection to update the document in.
//  - id: The ID of the document to update.
//  - update: The document fields to update.
//
// Returns:
//  - An error if the collection or document does not exist, or if the update is empty.
func (c *Client) UpdateOne(
	collectionName string,
	id string,
	update map[string]interface{},
) error {
	collection, err := c.getCollection(collectionName)
	if err != nil {
		return err
	}
	if update == nil || len(update) == 0 {
		return errors.New("update is empty")
	}
	return collection.UpdateOne(id, update)
}

// DeleteOne deletes a document by its ID from the specified collection.
//
// Parameters:
//  - collectionName: The name of the collection to delete the document from.
//  - id: The ID of the document to delete.
//
// Returns:
//  - An error if the collection or document does not exist.
func (c *Client) DeleteOne(
	collectionName string,
	id string,
) error {
	collection, err := c.getCollection(collectionName)
	if err != nil {
		return err
	}
	return collection.DeleteOne(id)
}

// DeleteAll deletes all documents from the specified collection.
//
// Parameters:
//  - collectionName: The name of the collection to delete all documents from.
//
// Returns:
//  - An error if the collection does not exist.
func (c *Client) DeleteAll(
	collectionName string,
) error {
	collection, err := c.getCollection(collectionName)
	if err != nil {
		return err
	}
	return collection.DeleteAll()
}
