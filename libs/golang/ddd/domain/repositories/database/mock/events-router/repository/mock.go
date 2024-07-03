package mockrepository

import (
	"libs/golang/ddd/domain/entities/events-router/entity"

	"github.com/stretchr/testify/mock"
)

type EventOrderRepositoryMock struct {
	mock.Mock
}

// Create is a mock implementation of EventOrderRepositoryInterface's Create method
func (m *EventOrderRepositoryMock) Create(event *entity.EventOrder) error {
	args := m.Called(event)
	return args.Error(0)
}

// FindByID is a mock implementation of EventOrderRepositoryInterface's FindByID method
func (m *EventOrderRepositoryMock) FindByID(id string) (*entity.EventOrder, error) {
	args := m.Called(id)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*entity.EventOrder), args.Error(1)
}

// FindAll is a mock implementation of EventOrderRepositoryInterface's FindAll method
func (m *EventOrderRepositoryMock) FindAll() ([]*entity.EventOrder, error) {
	args := m.Called()
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*entity.EventOrder), args.Error(1)
}

// Update is a mock implementation of EventOrderRepositoryInterface's Update method
func (m *EventOrderRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
