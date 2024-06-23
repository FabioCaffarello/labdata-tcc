package mockrepository

import (
	"libs/golang/ddd/domain/entities/schema-vault/entity"

	"github.com/stretchr/testify/mock"
)

// SchemaRepositoryMock is a mock implementation of SchemaRepositoryInterface
type SchemaRepositoryMock struct {
	mock.Mock
}

// Create is a mock implementation of SchemaRepositoryInterface's Create method
func (m *SchemaRepositoryMock) Create(config *entity.Schema) error {
	args := m.Called(config)
	return args.Error(0)
}

// FindByID is a mock implementation of SchemaRepositoryInterface's FindByID method
func (m *SchemaRepositoryMock) FindByID(id string) (*entity.Schema, error) {
	args := m.Called(id)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*entity.Schema), args.Error(1)
}

// FindAll is a mock implementation of SchemaRepositoryInterface's FindAll method
func (m *SchemaRepositoryMock) FindAll() ([]*entity.Schema, error) {
	args := m.Called()
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*entity.Schema), args.Error(1)
}

// Update is a mock implementation of SchemaRepositoryInterface's Update method
func (m *SchemaRepositoryMock) Update(config *entity.Schema) error {
	args := m.Called(config)
	return args.Error(0)
}

// Delete is a mock implementation of SchemaRepositoryInterface's Delete method
func (m *SchemaRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// FindAllByServiceAndProvider is a mock implementation of SchemaRepositoryInterface's FindAllByServiceAndProvider method
func (m *SchemaRepositoryMock) FindAllByServiceAndProvider(provider, service string) ([]*entity.Schema, error) {
	args := m.Called(provider, service)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*entity.Schema), args.Error(1)
}

// FindAllBySourceAndProvider is a mock implementation of SchemaRepositoryInterface's FindAllBySourceAndProvider method
func (m *SchemaRepositoryMock) FindAllBySourceAndProvider(provider, source string) ([]*entity.Schema, error) {
	args := m.Called(provider, source)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*entity.Schema), args.Error(1)
}

// FindAllByServiceAndSourceAndProvider is a mock implementation of SchemaRepositoryInterface's FindAllByServiceAndSourceAndProvider method
func (m *SchemaRepositoryMock) FindAllByServiceAndSourceAndProvider(service, source, provider string) ([]*entity.Schema, error) {
	args := m.Called(service, source, provider)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*entity.Schema), args.Error(1)
}
