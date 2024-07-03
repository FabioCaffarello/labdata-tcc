package handler

import "github.com/stretchr/testify/mock"

// MockRabbitMQNotifier is a mock implementation of RabbitMQNotifier for testing purposes.
type MockRabbitMQNotifier struct {
	mock.Mock
}

// Notify is the mock implementation of the Notify method.
func (m *MockRabbitMQNotifier) Notify(message []byte, routingKey string) error {
	args := m.Called(message, routingKey)
	return args.Error(0)
}
