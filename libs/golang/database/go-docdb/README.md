# go-docdb

`go-docdb` library provides an in-memory document-based database implementation in Go. It allows you to create, read, update, and delete collections and documents with thread-safe operations.

## Features

- Create and manage collections of documents.
- Insert, find, update, and delete documents.
- Supports querying documents based on specified criteria.
- Thread-safe operations using `sync.RWMutex`.

## Usage

### Create a new in-memory document database

```go
import (
    "libs/golang/database/go-docdb/database"
)

db := database.NewInMemoryDocBD("MyDatabase")
```

### Create a new collection

```go
err := db.CreateCollection("MyCollection")
if err != nil {
    log.Fatal(err)
}
```

### Insert a document into a collection

```go
collection, err := db.GetCollection("MyCollection")
if err != nil {
    log.Fatal(err)
}

document := database.Document{
    "_id": "123",
    "name": "John Doe",
    "age": 30,
}
err = collection.InsertOne(document)
if err != nil {
    log.Fatal(err)
}
```

### Find a document by ID

```go
doc, err := collection.FindOne("123")
if err != nil {
    log.Fatal(err)
}
fmt.Println(doc)
```

### Find all documents in a collection

```go
docs := collection.FindAll()
for _, doc := range docs {
    fmt.Println(doc)
}
```

### Update a document

```go
update := database.Document{
    "age": 31,
}
err = collection.UpdateOne("123", update)
if err != nil {
    log.Fatal(err)
}
```

### Delete a document

```go
err = collection.DeleteOne("123")
if err != nil {
    log.Fatal(err)
}
```

### List all collections

```go
collections := db.ListCollections()
for _, name := range collections {
    fmt.Println(name)
}
```

### Delete all documents in a collection

```go
err = collection.DeleteAll()
if err != nil {
    log.Fatal(err)
}
```

## Documentation

### Package `database`

#### Types

- `InMemoryDocBD`: Represents an in-memory document-based database.
- `DocumentID`: Represents the ID of a document.
- `Document`: Represents a document with key-value pairs.
- `Collection`: Represents a collection of documents.

#### Functions and Methods

- `NewInMemoryDocBD(name string) *InMemoryDocBD`: Creates a new in-memory document-based database.
- `(*InMemoryDocBD) GetCollection(collectionName string) (*Collection, error)`: Retrieves a collection by its name.
- `(*InMemoryDocBD) CreateCollection(collectionName string) error`: Creates a new collection.
- `(*InMemoryDocBD) DropCollection(collectionName string) error`: Removes a collection by its name.
- `(*InMemoryDocBD) ListCollections() []string`: Returns a list of all collection names.
- `NewCollection() *Collection`: Creates a new collection.
- `(*Collection) InsertOne(document Document) error`: Inserts a new document into the collection.
- `(*Collection) FindOne(id string) (Document, error)`: Retrieves a document by its ID.
- `(*Collection) FindAll() []Document`: Retrieves all documents in the collection.
- `(*Collection) Find(query map[string]interface{}) []Document`: Searches for documents matching a given query.
- `(*Collection) DeleteOne(id string) error`: Deletes a document by its ID.
- `(*Collection) UpdateOne(id string, update Document) error`: Updates a document by its ID.
- `(*Collection) DeleteAll() error`: Deletes all documents in the collection.
