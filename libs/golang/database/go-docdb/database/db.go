package database

import "errors"

// InMemoryDocBD represents an in-memory document-based database.
// It stores collections of documents in memory and provides methods
// to interact with these collections.
type InMemoryDocBD struct {
	Name        string
	Collections map[string]*Collection
}

// NewInMemoryDocBD creates a new in-memory document-based database
// with the specified name and initializes its collections map.
//
// Parameters:
//   - name: The name of the in-memory database.
//
// Returns:
//   - A pointer to the newly created InMemoryDocBD instance.
func NewInMemoryDocBD(name string) *InMemoryDocBD {
	return &InMemoryDocBD{
		Name:        name,
		Collections: make(map[string]*Collection),
	}
}

// GetCollection retrieves a collection by its name.
//
// Parameters:
//   - collectionName: The name of the collection to retrieve.
//
// Returns:
//   - A pointer to the Collection if found, otherwise nil.
//   - An error if the collection does not exist.
func (d *InMemoryDocBD) GetCollection(
	collectionName string,
) (*Collection, error) {
	collection, ok := d.Collections[collectionName]
	if !ok {
		return nil, errors.New("collection not found")
	}
	return collection, nil
}

// CreateCollection creates a new collection with the specified name.
//
// Parameters:
//   - collectionName: The name of the collection to create.
//
// Returns:
//   - An error if the collection already exists.
func (d *InMemoryDocBD) CreateCollection(
	collectionName string,
) error {
	if _, ok := d.Collections[collectionName]; ok {
		return errors.New("collection already exists")
	}
	d.Collections[collectionName] = NewCollection()
	return nil
}

// DropCollection removes a collection by its name.
//
// Parameters:
//   - collectionName: The name of the collection to remove.
//
// Returns:
//   - An error if the collection does not exist.
func (d *InMemoryDocBD) DropCollection(
	collectionName string,
) error {
	if _, ok := d.Collections[collectionName]; !ok {
		return errors.New("collection not found")
	}
	delete(d.Collections, collectionName)
	return nil
}

// ListCollections returns a list of all collection names.
//
// Returns:
//   - A slice of strings containing the names of all collections.
func (d *InMemoryDocBD) ListCollections() []string {
	collectionNames := make([]string, 0, len(d.Collections))
	for collectionName := range d.Collections {
		collectionNames = append(collectionNames, collectionName)
	}
	return collectionNames
}
