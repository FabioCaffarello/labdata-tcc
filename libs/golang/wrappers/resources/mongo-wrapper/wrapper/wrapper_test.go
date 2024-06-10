package mongowrapper

import (
	"context"
	"errors"
	gomongodb "libs/golang/clients/resources/go-mongo/client"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockClient is a mock implementation of the gomongodb.Client.
type MockClient struct {
	mock.Mock
}

func (m *MockClient) Disconnect(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// MockClientFactory is a mock implementation of the ClientFactory interface.
type MockClientFactory struct {
	mock.Mock
}

func (f *MockClientFactory) NewClient(config gomongodb.Config) (*gomongodb.Client, error) {
	args := f.Called(config)
	client, _ := args.Get(0).(*gomongodb.Client)
	return client, args.Error(1)
}

func TestMongoDBWrapperInit(t *testing.T) {
	// Set environment variables for testing
	os.Setenv("MONGODB_USER", "testuser")
	os.Setenv("MONGODB_PASSWORD", "testpassword")
	os.Setenv("MONGODB_HOST", "localhost")
	os.Setenv("MONGODB_PORT", "27017")
	os.Setenv("MONGODB_DBNAME", "testdb")

	t.Run("Successful Init", func(t *testing.T) {
		mockFactory := new(MockClientFactory)
		mockFactory.On("NewClient", mock.Anything).Return(&gomongodb.Client{}, nil)

		wrapper := &MongoDBWrapper{factory: mockFactory}
		err := wrapper.Init()
		assert.NoError(t, err)
		assert.NotNil(t, wrapper.client)

		mockFactory.AssertExpectations(t)
	})

	t.Run("Failed Init", func(t *testing.T) {
		mockFactory := new(MockClientFactory)
		mockFactory.On("NewClient", mock.Anything).Return((*gomongodb.Client)(nil), errors.New("connection error"))

		wrapper := &MongoDBWrapper{factory: mockFactory}
		err := wrapper.Init()
		assert.Error(t, err)
		assert.Nil(t, wrapper.client)

		mockFactory.AssertExpectations(t)
	})
}

func TestMongoDBWrapperGetClient(t *testing.T) {
	wrapper := &MongoDBWrapper{}
	mockClient := &gomongodb.Client{}
	wrapper.client = mockClient

	client := wrapper.GetClient()
	assert.Equal(t, mockClient.Client, client)
}
