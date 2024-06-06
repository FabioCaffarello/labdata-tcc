package rabbitmqwrapper

import (
	"errors"
	gorabbitmq "libs/golang/clients/resources/go-rabbitmq/client"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockClient is a mock implementation of the gorabbitmq.Client.
type MockClient struct {
	mock.Mock
}

func (m *MockClient) Close() error {
	args := m.Called()
	return args.Error(0)
}

// MockClientFactory is a mock implementation of the ClientFactory interface.
type MockClientFactory struct {
	mock.Mock
}

func (f *MockClientFactory) NewClient(config gorabbitmq.Config) (*gorabbitmq.Client, error) {
	args := f.Called(config)
	client, _ := args.Get(0).(*gorabbitmq.Client)
	return client, args.Error(1)
}

func TestRabbitMQWrapper_Init(t *testing.T) {
	// Set environment variables for testing
	os.Setenv("RABBITMQ_USER", "testuser")
	os.Setenv("RABBITMQ_PASSWORD", "testpassword")
	os.Setenv("RABBITMQ_HOST", "localhost")
	os.Setenv("RABBITMQ_PORT", "5672")
	os.Setenv("RABBITMQ_PROTOCOL", "amqp")
	os.Setenv("RABBITMQ_EXCHANGE_NAME", "testexchange")
	os.Setenv("RABBITMQ_EXCHANGE_TYPE", "direct")

	t.Run("Successful Init", func(t *testing.T) {
		mockFactory := new(MockClientFactory)
		mockFactory.On("NewClient", mock.Anything).Return(&gorabbitmq.Client{}, nil)

		wrapper := &RabbitMQWrapper{factory: mockFactory}
		err := wrapper.Init()
		assert.NoError(t, err)
		assert.NotNil(t, wrapper.client)

		mockFactory.AssertExpectations(t)
	})

	t.Run("Failed Init", func(t *testing.T) {
		mockFactory := new(MockClientFactory)
		mockFactory.On("NewClient", mock.Anything).Return((*gorabbitmq.Client)(nil), errors.New("connection error"))

		wrapper := &RabbitMQWrapper{factory: mockFactory}
		err := wrapper.Init()
		assert.Error(t, err)
		assert.Nil(t, wrapper.client)

		mockFactory.AssertExpectations(t)
	})
}

func TestRabbitMQWrapper_GetClient(t *testing.T) {
	wrapper := &RabbitMQWrapper{}
	mockClient := &gorabbitmq.Client{}
	wrapper.client = mockClient

	client := wrapper.GetClient()
	assert.Equal(t, mockClient, client)
}
