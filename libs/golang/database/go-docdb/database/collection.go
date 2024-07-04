package database

import (
	"errors"
	"log"
	"sync"
)

// DocumentID represents the ID of a document in the collection.
type DocumentID string

// Document represents a document with key-value pairs.
type Document map[string]interface{}

// Collection represents a collection of documents with thread-safe operations.
type Collection struct {
	data map[string]Document
	mu   sync.RWMutex
}

// NewCollection creates a new collection and initializes its data map.
//
// Returns:
//   - A pointer to the newly created Collection instance.
func NewCollection() *Collection {
	return &Collection{
		data: make(map[string]Document),
	}
}

// InsertOne inserts a new document into the collection.
//
// Parameters:
//   - document: The document to insert, which must contain an "_id" field.
//
// Returns:
//   - An error if the document already exists or if the "_id" field is missing.
func (c *Collection) InsertOne(
	document Document,
) error {
	c.mu.Lock() // Lock for writing
	defer c.mu.Unlock()

	documentID, ok := document["_id"]
	log.Printf("type of documentID: %T", documentID)
	if !ok {
		return errors.New("_id field is required")
	}
	_, ok = c.data[documentID.(string)]
	if ok {
		return errors.New("document already exists")
	}
	c.data[documentID.(string)] = document
	return nil
}

// FindOne retrieves a document by its ID.
//
// Parameters:
//   - id: The ID of the document to retrieve.
//
// Returns:
//   - The document if found.
//   - An error if the document does not exist.
func (c *Collection) FindOne(
	id string,
) (Document, error) {
	c.mu.RLock() // Lock for reading
	defer c.mu.RUnlock()

	document, ok := c.data[id]
	if !ok {
		return nil, errors.New("document not found")
	}
	return document, nil
}

// FindAll retrieves all documents in the collection.
//
// Returns:
//   - A slice of all documents in the collection.
func (c *Collection) FindAll() []Document {
	c.mu.RLock() // Lock for reading
	defer c.mu.RUnlock()
	documents := make([]Document, 0, len(c.data))
	for _, document := range c.data {
		documents = append(documents, document)
	}
	return documents
}

// matchesQuery checks if a document matches the query criteria.
//
// Parameters:
//   - document: The document to check against the query.
//   - query: The query criteria to match against the document.
//
// Returns:
//   - A boolean indicating if the document matches the query.
func matchesQuery(document, query map[string]interface{}) bool {
	for key, value := range query {
		docValue, exists := document[key]
		if !exists {
			return false
		}

		// If the value is a map, recurse into it.
		if queryMap, ok := value.(map[string]interface{}); ok {
			docMap, ok := docValue.(map[string]interface{})
			if !ok || !matchesQuery(docMap, queryMap) {
				return false
			}
		} else {
			// Direct comparison for non-map values
			if docValue != value {
				return false
			}
		}
	}
	return true
}

// Find searches for documents matching a given query.
//
// Parameters:
//   - query: The query criteria to match documents against.
//
// Returns:
//   - A slice of documents that match the query.
func (c *Collection) Find(query map[string]interface{}) []Document {
	c.mu.RLock() // Lock for reading
	defer c.mu.RUnlock()
	documents := make([]Document, 0, len(c.data))
	for _, document := range c.data {
		if matchesQuery(document, query) {
			documents = append(documents, document)
		}
	}
	return documents
}

// DeleteOne deletes a document by its ID.
//
// Parameters:
//   - id: The ID of the document to delete.
//
// Returns:
//   - An error if the document does not exist.
func (c *Collection) DeleteOne(
	id string,
) error {
	c.mu.Lock() // Lock for writing
	defer c.mu.Unlock()
	_, ok := c.data[id]
	if !ok {
		return errors.New("document not found")
	}
	delete(c.data, id)
	return nil
}

// UpdateOne updates a document by its ID.
//
// Parameters:
//   - id: The ID of the document to update.
//   - update: The document fields to update.
//
// Returns:
//   - An error if the document does not exist.
func (c *Collection) UpdateOne(
	id string,
	update Document,
) error {
	c.mu.Lock() // Lock for writing
	defer c.mu.Unlock()
	_, ok := c.data[id]
	if !ok {
		return errors.New("document not found")
	}
	for key, value := range update {
		c.data[id][key] = value
	}
	return nil
}

// DeleteAll deletes all documents in the collection.
//
// Returns:
//   - An error if there is an issue during deletion.
func (c *Collection) DeleteAll() error {
	c.mu.Lock() // Lock for writing
	defer c.mu.Unlock()
	c.data = make(map[string]Document)
	return nil
}
