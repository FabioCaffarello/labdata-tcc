# go-docdb/client

`go-docdb` is a Go library that provides a client wrapper for interacting with an in-memory document-based database. This library facilitates creating and managing collections and documents within the database.

## Features

- Create and drop collections
- Insert, find, update, and delete documents
- List all collections
- Convert maps to documents with required fields

## Usage

### Creating a Client

```go
package main

import (
	"fmt"
	"libs/golang/database/go-docdb/database"
	"libs/golang/clients/resources/go-docdb/client"
)

func main() {
	// Create a new in-memory document database
	db := database.NewInMemoryDocBD("MyDatabase")

	// Create a new client
	c := client.NewClient(db)

	fmt.Println("Client created successfully")
}
```

### Creating a Collection

```go
package main

import (
	"fmt"
	"log"
	"libs/golang/database/go-docdb/database"
	"libs/golang/clients/resources/go-docdb/client"
)

func main() {
	db := database.NewInMemoryDocBD("MyDatabase")
	c := client.NewClient(db)

	err := c.CreateCollection("MyCollection")
	if err != nil {
		log.Fatalf("Failed to create collection: %v", err)
	}

	fmt.Println("Collection created successfully")
}
```

### Inserting a Document

```go
package main

import (
	"fmt"
	"log"
	"libs/golang/database/go-docdb/database"
	"libs/golang/clients/resources/go-docdb/client"
)

func main() {
	db := database.NewInMemoryDocBD("MyDatabase")
	c := client.NewClient(db)

	document := map[string]interface{}{
		"_id": "123",
		"name": "John Doe",
		"age": 30,
	}

	err := c.InsertOne("MyCollection", document)
	if err != nil {
		log.Fatalf("Failed to insert document: %v", err)
	}

	fmt.Println("Document inserted successfully")
}
```

### Finding a Document by ID

```go
package main

import (
	"fmt"
	"log"
	"libs/golang/database/go-docdb/database"
	"libs/golang/clients/resources/go-docdb/client"
)

func main() {
	db := database.NewInMemoryDocBD("MyDatabase")
	c := client.NewClient(db)

	doc, err := c.FindOne("MyCollection", "123")
	if err != nil {
		log.Fatalf("Failed to find document: %v", err)
	}

	fmt.Printf("Found document: %v\n", doc)
}
```

### Updating a Document

```go
package main

import (
	"fmt"
	"log"
	"libs/golang/database/go-docdb/database"
	"libs/golang/clients/resources/go-docdb/client"
)

func main() {
	db := database.NewInMemoryDocBD("MyDatabase")
	c := client.NewClient(db)

	update := map[string]interface{}{
		"age": 31,
	}

	err := c.UpdateOne("MyCollection", "123", update)
	if err != nil {
		log.Fatalf("Failed to update document: %v", err)
	}

	fmt.Println("Document updated successfully")
}
```

### Deleting a Document

```go
package main

import (
	"fmt"
	"log"
	"libs/golang/database/go-docdb/database"
	"libs/golang/clients/resources/go-docdb/client"
)

func main() {
	db := database.NewInMemoryDocBD("MyDatabase")
	c := client.NewClient(db)

	err := c.DeleteOne("MyCollection", "123")
	if err != nil {
		log.Fatalf("Failed to delete document: %v", err)
	}

	fmt.Println("Document deleted successfully")
}
```

### Listing All Collections

```go
package main

import (
	"fmt"
	"libs/golang/database/go-docdb/database"
	"libs/golang/clients/resources/go-docdb/client"
)

func main() {
	db := database.NewInMemoryDocBD("MyDatabase")
	c := client.NewClient(db)

	collections := c.ListCollections()
	for _, name := range collections {
		fmt.Println(name)
	}
}
```

## Testing

To run the tests for the `client` package, use the following command:

```sh
npx nx test libs-golang-clients-resources-go-docdb
```
