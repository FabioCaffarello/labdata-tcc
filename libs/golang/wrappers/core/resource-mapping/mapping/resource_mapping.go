package resourcemapping

import (
	"fmt"
	resourceImpl "libs/golang/wrappers/core/resource-contract"
	"sync"
)

// Resources struct holds the initialized resources
type Resources struct {
	mu        sync.RWMutex
	resources map[string]resourceImpl.Resource
}

var (
	instance *Resources
	once     sync.Once
)

// NewResourceMapping creates a singleton instance of Resources
func NewResourceMapping() *Resources {
	once.Do(func() {
		instance = &Resources{
			resources: make(map[string]resourceImpl.Resource),
		}
	})

	return instance
}

// RegisterResource registers a resource with a given key
func (r *Resources) RegisterResource(key string, resource resourceImpl.Resource) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.resources[key] = resource
}

// GetResource retrieves a resource by key
func (r *Resources) GetResource(key string) (resourceImpl.Resource, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	resource, exists := r.resources[key]
	if !exists {
		return nil, fmt.Errorf("resource %s not found", key)
	}
	return resource, nil
}
