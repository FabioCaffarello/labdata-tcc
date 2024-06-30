package mockrepository

import (
	"libs/golang/ddd/domain/entities/input-broker/entity"

	"github.com/stretchr/testify/mock"
)

// InputRepositoryMock is a mock implementation of InputRepositoryInterface
type InputRepositoryMock struct {
	mock.Mock
}

// Create is a mock implementation of InputRepositoryInterface's Create method
func (m *InputRepositoryMock) Create(input *entity.Input) error {
	args := m.Called(input)
	return args.Error(0)
}

// FindByID is a mock implementation of InputRepositoryInterface's FindByID method
func (m *InputRepositoryMock) FindByID(id string) (*entity.Input, error) {
	args := m.Called(id)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*entity.Input), args.Error(1)
}

// FindAll is a mock implementation of InputRepositoryInterface's FindAll method
func (m *InputRepositoryMock) FindAll() ([]*entity.Input, error) {
	args := m.Called()
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*entity.Input), args.Error(1)
}

// Update is a mock implementation of InputRepositoryInterface's Update method
func (m *InputRepositoryMock) Update(input *entity.Input) error {
	args := m.Called(input)
	return args.Error(0)
}

// Delete is a mock implementation of InputRepositoryInterface's Delete method
func (m *InputRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// FindAllByServiceAndProvider is a mock implementation of InputRepositoryInterface's FindAllByServiceAndProvider method
func (m *InputRepositoryMock) FindAllByServiceAndProvider(provider, service string) ([]*entity.Input, error) {
	args := m.Called(provider, service)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*entity.Input), args.Error(1)
}

// FindAllBySourceAndProvider is a mock implementation of InputRepositoryInterface's FindAllBySourceAndProvider method
func (m *InputRepositoryMock) FindAllBySourceAndProvider(provider, source string) ([]*entity.Input, error) {
	args := m.Called(provider, source)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*entity.Input), args.Error(1)
}

// FindAllByServiceAndSourceAndProvider is a mock implementation of InputRepositoryInterface's FindAllByServiceAndSourceAndProvider method
func (m *InputRepositoryMock) FindAllByServiceAndSourceAndProvider(provider, service, source string) ([]*entity.Input, error) {
	args := m.Called(provider, service, source)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*entity.Input), args.Error(1)
}

// FindAllByStatusAndProvider is a mock implementation of InputRepositoryInterface's FindAllByStatusAndProvider method
func (m *InputRepositoryMock) FindAllByStatusAndProvider(provider string, status int) ([]*entity.Input, error) {
	args := m.Called(provider, status)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*entity.Input), args.Error(1)
}

// FindAllByStatusAndServiceAndProvider is a mock implementation of InputRepositoryInterface's FindAllByStatusAndServiceAndProvider method
func (m *InputRepositoryMock) FindAllByStatusAndServiceAndProvider(provider, service string, status int) ([]*entity.Input, error) {
	args := m.Called(provider, service, status)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*entity.Input), args.Error(1)
}

// FindAllByStatusAndSourceAndProvider is a mock implementation of InputRepositoryInterface's FindAllByStatusAndSourceAndProvider method
func (m *InputRepositoryMock) FindAllByStatusAndSourceAndProvider(provider, source string, status int) ([]*entity.Input, error) {
	args := m.Called(provider, source, status)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*entity.Input), args.Error(1)
}

// FindAllByStatusAndServiceAndSourceAndProvider is a mock implementation of InputRepositoryInterface's FindAllByStatusAndServiceAndSourceAndProvider method
func (m *InputRepositoryMock) FindAllByStatusAndServiceAndSourceAndProvider(provider, service, source string, status int) ([]*entity.Input, error) {
	args := m.Called(provider, service, source, status)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*entity.Input), args.Error(1)
}
