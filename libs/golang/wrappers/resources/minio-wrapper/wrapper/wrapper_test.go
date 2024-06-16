package miniowrapper

import (
	"context"
	"errors"
	gominio "libs/golang/clients/resources/go-minio/client"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockClient is a mock implementation of the gominio.Client.
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

func (f *MockClientFactory) NewClient(config gominio.Config) (*gominio.Client, error) {
	args := f.Called(config)
	client, _ := args.Get(0).(*gominio.Client)
	return client, args.Error(1)
}

func TestMinioWrapper_Init(t *testing.T) {
	// Set environment variables for testing
	os.Setenv("MINIO_PORT", "9000")
	os.Setenv("MINIO_HOST", "localhost")
	os.Setenv("MINIO_ACCESS_KEY", "minioaccesskey")
	os.Setenv("MINIO_SECRET_KEY", "miniosecretkey")
	os.Setenv("MINIO_USE_SSL", "false")

	t.Run("Successful Init", func(t *testing.T) {
		mockFactory := new(MockClientFactory)
		mockFactory.On("NewClient", mock.Anything).Return(&gominio.Client{}, nil)

		wrapper := &MinioWrapper{factory: mockFactory}
		err := wrapper.Init()
		assert.NoError(t, err)
		assert.NotNil(t, wrapper.client)

		mockFactory.AssertExpectations(t)
	})

	t.Run("Failed Init", func(t *testing.T) {
		mockFactory := new(MockClientFactory)
		mockFactory.On("NewClient", mock.Anything).Return((*gominio.Client)(nil), errors.New("connection error"))

		wrapper := &MinioWrapper{factory: mockFactory}
		err := wrapper.Init()
		assert.Error(t, err)
		assert.Nil(t, wrapper.client)

		mockFactory.AssertExpectations(t)
	})
}

func TestMinioWrapper_GetClient(t *testing.T) {
	wrapper := &MinioWrapper{}
	mockClient := &gominio.Client{}
	wrapper.client = mockClient

	client := wrapper.GetClient()
	assert.Equal(t, mockClient, client)
}
