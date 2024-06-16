package gominio

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// Client wraps the Minio client to provide additional functionality.
type Client struct {
	*minio.Client
}

// Config holds the configuration for connecting to a Minio instance.
type Config struct {
	Port      string // Minio server port
	Host      string // Minio server host
	AccessKey string // Access key for authentication
	SecretKey string // Secret key for authentication
	UseSSL    bool   // Use SSL connection
}

// NewClient creates a new Minio client with the given configuration.
// It returns the Client and an error if any occurred during connection.
//
// Example:
//
//		config := Config{
//	     Port:      "9000",
//	     Host:      "localhost",
//		    AccessKey: "minioaccesskey",
//		    SecretKey: "miniosecretkey",
//		    UseSSL:    false,
//		}
//		client, err := NewClient(config)
//		if err != nil {
//		    log.Fatal(err)
//		}
func NewClient(config Config) (*Client, error) {
	endpoint := fmt.Sprintf("%s:%s", config.Host, config.Port)
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.SecretKey, ""),
		Secure: config.UseSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create Minio client: %w", err)
	}
	return &Client{Client: client}, nil
}

// GetObject retrieves an object from the specified bucket and returns its content as a byte slice.
func (c *Client) GetObject(bucketName, fileName string) ([]byte, error) {
	object, err := c.Client.GetObject(context.Background(), bucketName, fileName, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get object: %w", err)
	}
	defer object.Close()

	var buf bytes.Buffer
	if _, err = io.Copy(&buf, object); err != nil {
		return nil, fmt.Errorf("failed to read object content: %w", err)
	}

	return buf.Bytes(), nil
}

// DownloadFile downloads a file from the specified URI and returns its content as a byte slice.
func (c *Client) DownloadFile(uri string) ([]byte, error) {
	parsedURI, err := url.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URI: %w", err)
	}

	pathParts := strings.Split(strings.TrimPrefix(parsedURI.Path, "/"), "/")
	if len(pathParts) < 2 {
		return nil, errors.New("invalid URI path")
	}

	bucketName := pathParts[0]
	objectKey := strings.Join(pathParts[1:], "/")

	objContent, err := c.GetObject(bucketName, objectKey)
	if err != nil {
		return nil, err
	}

	return objContent, nil
}

// UploadFile uploads a file to the specified bucket and returns the path to the uploaded file.
func (c *Client) UploadFile(bucketName, fileName string, fileContent []byte) (string, error) {
	reader := bytes.NewReader(fileContent)
	_, err := c.Client.PutObject(context.Background(), bucketName, fileName, reader, int64(len(fileContent)), minio.PutObjectOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}
	return fileName, nil
}

// UploadFileWithChunks uploads a file in chunks of a specified size to the specified bucket and returns the path to the uploaded parts.
func (c *Client) UploadFileWithChunks(bucketName, fileName string, fileContent []byte, partSize int64) (string, error) {
	fileSize := int64(len(fileContent))

	extension := strings.Split(fileName, ".")[1]
	fileNameClean := strings.Split(fileName, ".")[0]

	partPath := fileNameClean

	var offset int64
	var partNumber int

	for partNumber = 1; offset < fileSize; partNumber++ {
		if partSize > (fileSize - offset) {
			partSize = fileSize - offset
		}

		reader := bytes.NewReader(fileContent[offset : offset+partSize])
		partName := fmt.Sprintf("%s/part-%d.%s", partPath, partNumber, extension)
		_, err := c.Client.PutObject(context.Background(), bucketName, partName, reader, partSize, minio.PutObjectOptions{})
		if err != nil {
			return "", fmt.Errorf("failed to upload part: %w", err)
		}
		offset += partSize
	}

	if offset != fileSize {
		return "", errors.New("upload incomplete: offset does not match fileSize")
	}

	return partPath, nil
}

// RemoveObject removes an object from the specified bucket.
func (c *Client) RemoveObject(bucketName, objectName string) error {
	err := c.Client.RemoveObject(context.Background(), bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to remove object: %w", err)
	}
	return nil
}

// RemoveAllObjectsFromBucket removes all objects from the specified bucket.
func (c *Client) RemoveAllObjectsFromBucket(bucketName string) error {
	objectsCh := c.Client.ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{
		Recursive: true,
	})
	for object := range objectsCh {
		if object.Err != nil {
			return fmt.Errorf("failed to list objects: %w", object.Err)
		}
		if err := c.RemoveObject(bucketName, object.Key); err != nil {
			return err
		}
	}
	return nil
}
