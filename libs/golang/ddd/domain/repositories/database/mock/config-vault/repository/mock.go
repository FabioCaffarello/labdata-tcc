package mockrepository

import (
	"libs/golang/ddd/domain/entities/config-vault/entity"

	"github.com/stretchr/testify/mock"
)

// ConfigRepositoryMock is a mock implementation of ConfigRepositoryInterface
type ConfigRepositoryMock struct {
	mock.Mock
}

// Create is a mock implementation of ConfigRepositoryInterface's Create method
func (m *ConfigRepositoryMock) Create(config *entity.Config) error {
	args := m.Called(config)
	return args.Error(0)
}

// FindByID is a mock implementation of ConfigRepositoryInterface's FindByID method
func (m *ConfigRepositoryMock) FindByID(id string) (*entity.Config, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Config), args.Error(1)
}

// FindAll is a mock implementation of ConfigRepositoryInterface's FindAll method
func (m *ConfigRepositoryMock) FindAll() ([]*entity.Config, error) {
	args := m.Called()
	return args.Get(0).([]*entity.Config), args.Error(1)
}

// Update is a mock implementation of ConfigRepositoryInterface's Update method
func (m *ConfigRepositoryMock) Update(config *entity.Config) error {
	args := m.Called(config)
	return args.Error(0)
}

// Delete is a mock implementation of ConfigRepositoryInterface's Delete method
func (m *ConfigRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// FindAllByService is a mock implementation of ConfigRepositoryInterface's FindAllByService method
func (m *ConfigRepositoryMock) FindAllByService(service string) ([]*entity.Config, error) {
	args := m.Called(service)
	return args.Get(0).([]*entity.Config), args.Error(1)
}

// FindAllBySource is a mock implementation of ConfigRepositoryInterface's FindAllBySource method
func (m *ConfigRepositoryMock) FindAllBySource(source string) ([]*entity.Config, error) {
	args := m.Called(source)
	return args.Get(0).([]*entity.Config), args.Error(1)
}

// FindAllByProvider is a mock implementation of ConfigRepositoryInterface's FindAllByProvider method
func (m *ConfigRepositoryMock) FindAllByServiceAndSource(service, source string) ([]*entity.Config, error) {
	args := m.Called(service, source)
	return args.Get(0).([]*entity.Config), args.Error(1)
}

// FindAllByServiceAndSourceAndProvider is a mock implementation of ConfigRepositoryInterface's FindAllByServiceAndSourceAndProvider method
func (m *ConfigRepositoryMock) FindAllByServiceAndSourceAndProvider(service, source, provider string) ([]*entity.Config, error) {
	args := m.Called(service, source, provider)
	return args.Get(0).([]*entity.Config), args.Error(1)
}

// FindAllByServiceAndProviderAndActive is a mock implementation of ConfigRepositoryInterface's FindAllByServiceAndProviderAndActive method
func (m *ConfigRepositoryMock) FindAllByServiceAndProviderAndActive(service, provider string, active bool) ([]*entity.Config, error) {
	args := m.Called(service, provider, active)
	return args.Get(0).([]*entity.Config), args.Error(1)
}

// FindAllByDependsOn is a mock implementation of ConfigRepositoryInterface's FindAllByDependsOn method
func (m *ConfigRepositoryMock) FindAllByDependsOn(dependsOn map[string]interface{}) ([]*entity.Config, error) {
	args := m.Called(dependsOn)
	return args.Get(0).([]*entity.Config), args.Error(1)
}
