# go-minio

`go-minio` is a Go library that provides a convenient interface for interacting with Minio, an object storage server compatible with Amazon S3 cloud storage service. This library wraps the Minio client to offer additional functionality and simplifies operations such as uploading, downloading, and managing objects in Minio.

## Features

- Create a new Minio client with configuration options.
- Upload files to a Minio bucket.
- Upload files in chunks to handle large files.
- Download files from a Minio bucket.
- Remove objects from a Minio bucket.
- Remove all objects from a Minio bucket.

## Usage

### Creating a Client

```go
package main

import (
    "log"
    "github.com/yourusername/gominio"
)

func main() {
    config := gominio.Config{
        Endpoint:  "localhost:9000",
        AccessKey: "minioaccesskey",
        SecretKey: "miniosecretkey",
        UseSSL:    false,
    }
    client, err := gominio.NewClient(config)
    if err != nil {
        log.Fatal(err)
    }
    // Use the client...
}
```

### Uploading a File

```go
package main

import (
    "log"
    "github.com/yourusername/gominio"
)

func main() {
    config := gominio.Config{
        Endpoint:  "localhost:9000",
        AccessKey: "minioaccesskey",
        SecretKey: "miniosecretkey",
        UseSSL:    false,
    }
    client, err := gominio.NewClient(config)
    if err != nil {
        log.Fatal(err)
    }

    fileContent := []byte("Hello, Minio!")
    bucketName := "example-bucket"
    fileName := "hello.txt"
    uploadedPath, err := client.UploadFile(bucketName, fileName, fileContent)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("File uploaded to:", uploadedPath)
}
```

### Uploading a File with Chunks

```go
package main

import (
    "log"
    "github.com/yourusername/gominio"
)

func main() {
    config := gominio.Config{
        Endpoint:  "localhost:9000",
        AccessKey: "minioaccesskey",
        SecretKey: "miniosecretkey",
        UseSSL:    false,
    }
    client, err := gominio.NewClient(config)
    if err != nil {
        log.Fatal(err)
    }

    fileContent := make([]byte, 150*1024*1024) // 150MB
    bucketName := "example-bucket"
    fileName := "largefile.txt"
    partSize := int64(50 * 1024 * 1024) // 50MB
    uploadedPath, err := client.UploadFileWithChunks(bucketName, fileName, fileContent, partSize)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("File uploaded in chunks to:", uploadedPath)
}
```

### Downloading a File

```go
package main

import (
    "log"
    "github.com/yourusername/gominio"
)

func main() {
    config := gominio.Config{
        Endpoint:  "localhost:9000",
        AccessKey: "minioaccesskey",
        SecretKey: "miniosecretkey",
        UseSSL:    false,
    }
    client, err := gominio.NewClient(config)
    if err != nil {
        log.Fatal(err)
    }

    uri := "http://localhost:9000/example-bucket/hello.txt"
    content, err := client.DownloadFile(uri)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("File content:", string(content))
}
```

### Removing an Object

```go
package main

import (
    "log"
    "github.com/yourusername/gominio"
)

func main() {
    config := gominio.Config{
        Endpoint:  "localhost:9000",
        AccessKey: "minioaccesskey",
        SecretKey: "miniosecretkey",
        UseSSL:    false,
    }
    client, err := gominio.NewClient(config)
    if err != nil {
        log.Fatal(err)
    }

    bucketName := "example-bucket"
    objectName := "hello.txt"
    err = client.RemoveObject(bucketName, objectName)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Object removed:", objectName)
}
```

### Removing All Objects from a Bucket

```go
package main

import (
    "log"
    "github.com/yourusername/gominio"
)

func main() {
    config := gominio.Config{
        Endpoint:  "localhost:9000",
        AccessKey: "minioaccesskey",
        SecretKey: "miniosecretkey",
        UseSSL:    false,
    }
    client, err := gominio.NewClient(config)
    if err != nil {
        log.Fatal(err)
    }

    bucketName := "example-bucket"
    err = client.RemoveAllObjectsFromBucket(bucketName)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("All objects removed from bucket:", bucketName)
}
```

## Testing

To run the tests for the `gominio` package, use the following command:

```sh
npx nx test libs-golang-clients-resources-go-minio
```
