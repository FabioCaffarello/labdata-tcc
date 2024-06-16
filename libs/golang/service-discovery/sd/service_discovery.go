package servicediscovery

import (
	"sync"

	resourceImpl "libs/golang/wrappers/core/resource-contract"
	resourcemapping "libs/golang/wrappers/core/resource-mapping/mapping"
	miniowrapper "libs/golang/wrappers/resources/minio-wrapper/wrapper"
	mongowrapper "libs/golang/wrappers/resources/mongo-wrapper/wrapper"
	rabbitmqWrapper "libs/golang/wrappers/resources/rabbitmq-wrapper/wrapper"
)

// ServiceDiscovery holds the service discovery logic, managing resource mappings and services.
type ServiceDiscovery struct {
	resourceMapping *resourcemapping.Resources
	services        map[string]interface{}
}

var (
	instance            *ServiceDiscovery
	once                sync.Once
	resourceInitializer = defaultResourceInitializer
)

// defaultResourceInitializer initializes the given resource and panics if an error occurs.
func defaultResourceInitializer(wrapper resourceImpl.Resource) {
	if err := wrapper.Init(); err != nil {
		panic(err)
	}
}

// SetResourceInitializer sets the resource initializer function, typically used for testing.
func SetResourceInitializer(initializer func(resourceImpl.Resource)) {
	resourceInitializer = initializer
}

// NewServiceDiscovery creates and returns a singleton instance of ServiceDiscovery.
func NewServiceDiscovery() *ServiceDiscovery {
	once.Do(func() {
		resourceMapping := resourcemapping.NewResourceMapping()
		instance = &ServiceDiscovery{
			resourceMapping: resourceMapping,
			services:        make(map[string]interface{}),
		}
		instance.registerResources()
	})
	return instance
}

// registerResources registers all necessary resources with their respective keys.
func (s *ServiceDiscovery) registerResources() {
	// Register MongoDB resource
	mongoWrapper := mongowrapper.NewMongoDBWrapper()
	s.InitResourceWrapper(mongoWrapper)
	s.resourceMapping.RegisterResource("mongodb", mongoWrapper)

	// Register Minio resource
	minioWrapper := miniowrapper.NewMinioWrapper()
	s.InitResourceWrapper(minioWrapper)
	s.resourceMapping.RegisterResource("minio", minioWrapper)

	// Register RabbitMQ resource
	rabbitMQWrapper := rabbitmqWrapper.NewRabbitMQWrapper()
	s.InitResourceWrapper(rabbitMQWrapper)
	s.resourceMapping.RegisterResource("rabbitmq", rabbitMQWrapper)
}

// InitResourceWrapper initializes the given resource wrapper using the configured initializer.
func (s *ServiceDiscovery) InitResourceWrapper(wrapper resourceImpl.Resource) {
	resourceInitializer(wrapper)
}

// RegisterResource registers a resource with the given key.
func (s *ServiceDiscovery) RegisterResource(key string, resource resourceImpl.Resource) {
	s.resourceMapping.RegisterResource(key, resource)
}

// GetResource retrieves a resource by key from the resource mapping.
func (s *ServiceDiscovery) GetResource(key string) (resourceImpl.Resource, error) {
	return s.resourceMapping.GetResource(key)
}
