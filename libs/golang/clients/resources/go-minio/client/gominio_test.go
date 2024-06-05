package gominio

import (
	"bytes"
	"context"
	"log"
	"testing"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MinioSuite struct {
	suite.Suite
	client *Client
	ctx    context.Context
	bucket string
}

func TestMinioSuite(t *testing.T) {
	suite.Run(t, new(MinioSuite))
}

func (suite *MinioSuite) SetupSuite() {
	suite.ctx = context.Background()
	config := Config{
		Endpoint:  "localhost:9000",
		AccessKey: "test-root-user",
		SecretKey: "test-root-password",
		UseSSL:    false,
	}

	var err error
	for i := 0; i < 10; i++ {
		suite.client, err = NewClient(config)
		if err == nil {
			break
		}
		log.Printf("Retrying Minio connection, attempt %d", i+1)
		time.Sleep(5 * time.Second)
	}
	assert.NoError(suite.T(), err, "Failed to connect to Minio")

	suite.bucket = "test-bucket"
	err = suite.client.MakeBucket(suite.ctx, suite.bucket, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := suite.client.BucketExists(suite.ctx, suite.bucket)
		assert.NoError(suite.T(), errBucketExists)
		assert.True(suite.T(), exists)
	}
}

func (suite *MinioSuite) TearDownSuite() {
	err := suite.client.RemoveAllObjectsFromBucket(suite.bucket)
	assert.NoError(suite.T(), err)

	err = suite.client.RemoveBucket(suite.ctx, suite.bucket)
	assert.NoError(suite.T(), err)
}

func (suite *MinioSuite) TestUploadAndDownloadFile() {
	content := []byte("test content")
	fileName := "testfile.txt"

	// Upload file
	uploadedPath, err := suite.client.UploadFile(suite.bucket, fileName, content)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), uploadedPath)

	// Download file
	uri := "http://" + suite.client.Client.EndpointURL().Host + "/" + suite.bucket + "/" + uploadedPath
	downloadedContent, err := suite.client.DownloadFile(uri)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), content, downloadedContent)
}

func (suite *MinioSuite) TestUploadFileWithChunks() {
	content := bytes.Repeat([]byte("a"), 150*1024*1024) // 150MB content
	fileName := "largefile.txt"
	partSize := int64(50 * 1024 * 1024) // 50MB

	// Upload file in chunks
	uploadedPath, err := suite.client.UploadFileWithChunks(suite.bucket, fileName, content, partSize)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), uploadedPath)
}

func (suite *MinioSuite) TestRemoveObject() {
	content := []byte("test content")
	fileName := "testfile.txt"

	// Upload file
	uploadedPath, err := suite.client.UploadFile(suite.bucket, fileName, content)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), uploadedPath)

	// Remove object
	err = suite.client.RemoveObject(suite.bucket, uploadedPath)
	assert.NoError(suite.T(), err)

	// Check if object was removed
	_, err = suite.client.GetObject(suite.bucket, uploadedPath)
	assert.Error(suite.T(), err)
}

func (suite *MinioSuite) TestRemoveAllObjectsFromBucket() {
	content := []byte("test content")
	fileName := "testfile.txt"

	// Upload files
	for i := 0; i < 5; i++ {
		uploadedPath, err := suite.client.UploadFile(suite.bucket, fileName, content)
		assert.NoError(suite.T(), err)
		assert.NotEmpty(suite.T(), uploadedPath)
	}

	// Upload a directory with a file (if required)
	largeFileContent := bytes.Repeat([]byte("a"), 150*1024*1024) // 150MB content
	_, err := suite.client.UploadFile(suite.bucket, "largefile/part-1.txt", largeFileContent)
	assert.NoError(suite.T(), err)

	// Remove all objects from bucket
	err = suite.client.RemoveAllObjectsFromBucket(suite.bucket)
	assert.NoError(suite.T(), err)

	// Check if bucket is empty
	objectsCh := suite.client.ListObjects(context.Background(), suite.bucket, minio.ListObjectsOptions{
		Recursive: true,
	})
	for object := range objectsCh {
		assert.Fail(suite.T(), "Bucket is not empty", object.Key)
	}
}
