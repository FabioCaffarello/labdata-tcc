package servicediscovery_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"libs/golang/service-discovery/sd"
	resourceImpl "libs/golang/wrappers/core/resource-contract"
)

// MockResource is a mock implementation of the Resource interface
type MockResource struct {
	mock.Mock
}

func (m *MockResource) Init() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockResource) GetClient() interface{} {
	args := m.Called()
	return args.Get(0)
}

// ServiceDiscoverySuite defines the test suite for ServiceDiscovery
type ServiceDiscoverySuite struct {
	suite.Suite
	serviceDiscovery *servicediscovery.ServiceDiscovery
	mockMongo        *MockResource
	mockMinio        *MockResource
	mockRabbitMQ     *MockResource
}

func (suite *ServiceDiscoverySuite) SetupTest() {
	// Create mock resources
	suite.mockMongo = new(MockResource)
	suite.mockMinio = new(MockResource)
	suite.mockRabbitMQ = new(MockResource)

	// Set up the mocks to return no error on Init
	suite.mockMongo.On("Init").Return(nil)
	suite.mockMinio.On("Init").Return(nil)
	suite.mockRabbitMQ.On("Init").Return(nil)

	// Replace the actual resource initializer with the mock initializer
	servicediscovery.SetResourceInitializer(func(wrapper resourceImpl.Resource) {
		if mockResource, ok := wrapper.(*MockResource); ok {
			mockResource.Init()
		}
	})

	// Initialize the service discovery with the mock resources
	suite.serviceDiscovery = servicediscovery.NewServiceDiscovery()

	// Replace actual resources with mocks in the resource mapping
	suite.serviceDiscovery.RegisterResource("mongodb", suite.mockMongo)
	suite.serviceDiscovery.RegisterResource("minio", suite.mockMinio)
	suite.serviceDiscovery.RegisterResource("rabbitmq", suite.mockRabbitMQ)
}

func (suite *ServiceDiscoverySuite) TestSingletonInstance() {
	instance1 := servicediscovery.NewServiceDiscovery()
	instance2 := servicediscovery.NewServiceDiscovery()
	assert.Equal(suite.T(), instance1, instance2, "Expected singleton instances to be equal")
}

func (suite *ServiceDiscoverySuite) TestRegisterResources() {
	// Assert that resources are registered
	mongoClientWrapper, err := suite.serviceDiscovery.GetResource("mongodb")
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), mongoClientWrapper, "Expected MongoDB resource to be registered")

	minioClientWrapper, err := suite.serviceDiscovery.GetResource("minio")
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), minioClientWrapper, "Expected Minio resource to be registered")

	rabbitMQClientWrapper, err := suite.serviceDiscovery.GetResource("rabbitmq")
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), rabbitMQClientWrapper, "Expected RabbitMQ resource to be registered")
}

func (suite *ServiceDiscoverySuite) TestGetResourceNotFound() {
	_, err := suite.serviceDiscovery.GetResource("nonexistent")
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "resource nonexistent not found", err.Error())
}

func TestServiceDiscoverySuite(t *testing.T) {
	suite.Run(t, new(ServiceDiscoverySuite))
}
