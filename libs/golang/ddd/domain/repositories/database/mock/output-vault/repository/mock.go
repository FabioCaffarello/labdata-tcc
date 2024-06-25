package mockrepository

import (
	"libs/golang/ddd/domain/entities/output-vault/entity"

	"github.com/stretchr/testify/mock"
)

// OutputRepositoryMock is a mock implementation of OutputRepositoryInterface
type OutputRepositoryMock struct {
	mock.Mock
}

// Create is a mock implementation of OutputRepositoryInterface's Create method
func (m *OutputRepositoryMock) Create(output *entity.Output) error {
	args := m.Called(output)
	return args.Error(0)
}

// FindByID is a mock implementation of OutputRepositoryInterface's FindByID method
func (m *OutputRepositoryMock) FindByID(id string) (*entity.Output, error) {
	args := m.Called(id)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*entity.Output), args.Error(1)
}

// FindAll is a mock implementation of OutputRepositoryInterface's FindAll method
func (m *OutputRepositoryMock) FindAll() ([]*entity.Output, error) {
	args := m.Called()
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*entity.Output), args.Error(1)
}

// Update is a mock implementation of OutputRepositoryInterface's Update method
func (m *OutputRepositoryMock) Update(output *entity.Output) error {
	args := m.Called(output)
	return args.Error(0)
}

// Delete is a mock implementation of OutputRepositoryInterface's Delete method
func (m *OutputRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// FindAllByServiceAndProvider is a mock implementation of OutputRepositoryInterface's FindAllByServiceAndProvider method
func (m *OutputRepositoryMock) FindAllByServiceAndProvider(provider, service string) ([]*entity.Output, error) {
	args := m.Called(provider, service)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*entity.Output), args.Error(1)
}

// FindAllBySourceAndProvider is a mock implementation of OutputRepositoryInterface's FindAllBySourceAndProvider method
func (m *OutputRepositoryMock) FindAllBySourceAndProvider(provider, source string) ([]*entity.Output, error) {
	args := m.Called(provider, source)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*entity.Output), args.Error(1)
}

// FindAllByServiceAndSourceAndProvider is a mock implementation of OutputRepositoryInterface's FindAllByServiceAndSourceAndProvider method
func (m *OutputRepositoryMock) FindAllByServiceAndSourceAndProvider(service, source, provider string) ([]*entity.Output, error) {
	args := m.Called(service, source, provider)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*entity.Output), args.Error(1)
}
