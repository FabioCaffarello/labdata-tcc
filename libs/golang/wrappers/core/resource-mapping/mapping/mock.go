package resourcemapping

import (
	"errors"
)

type MockResource struct {
	initialized bool
	client      interface{}
}

func (m *MockResource) Init() error {
	m.initialized = true
	return nil
}

func (m *MockResource) GetClient() interface{} {
	return m.client
}

type FailingMockResource struct{}

func (f *FailingMockResource) Init() error {
	return errors.New("initialization error")
}

func (f *FailingMockResource) GetClient() interface{} {
	return nil
}
