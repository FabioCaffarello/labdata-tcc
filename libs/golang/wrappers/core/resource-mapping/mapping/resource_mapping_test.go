package resourcemapping

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ResourceMappingSuite struct {
	suite.Suite
}

func TestResourceMappingSuite(t *testing.T) {
	suite.Run(t, new(ResourceMappingSuite))
}

func (suite *ResourceMappingSuite) TestNewResourceMapping() {
	resources := NewResourceMapping()

	if resources == nil {
		assert.Fail(suite.T(), "expected non-nil Resources instance, got nil")
	}
}

func (suite *ResourceMappingSuite) TestRegisterResource() {
	resources := NewResourceMapping()
	mockResource := &MockResource{client: "mockClient"}

	resources.RegisterResource("mock", mockResource)

	retrievedResource, err := resources.GetResource("mock")
	if err != nil {
		assert.Fail(suite.T(), "expected no error, got %v", err)
	}

	if retrievedResource != mockResource {
		assert.Fail(suite.T(), "expected %v, got %v", mockResource, retrievedResource)
	}
}

func (suite *ResourceMappingSuite) TestGetResourceNotFound() {
	resources := NewResourceMapping()

	_, err := resources.GetResource("nonexistent")
	if err == nil {
		assert.Fail(suite.T(), "expected error, got nil")
	}

	expectedErrMsg := "resource nonexistent not found"
	if err.Error() != expectedErrMsg {
		assert.Fail(suite.T(), "expected %v, got %v", expectedErrMsg, err.Error())
	}
}

func (suite *ResourceMappingSuite) TestInitResource() {
	mockResource := &MockResource{client: "mockClient"}

	err := mockResource.Init()
	if err != nil {
		assert.Fail(suite.T(), "expected no error, got %v", err)
	}

	if !mockResource.initialized {
		assert.Fail(suite.T(), "expected resource to be initialized, got false")
	}
}

func (suite *ResourceMappingSuite) TestFailingInitResource() {
	failingMockResource := &FailingMockResource{}

	err := failingMockResource.Init()
	if err == nil {
		assert.Fail(suite.T(), "expected error, got nil")
	}

	expectedErrMsg := "initialization error"
	if err.Error() != expectedErrMsg {
		assert.Fail(suite.T(), "expected %v, got %v", expectedErrMsg, err.Error())
	}
}

func (suite *ResourceMappingSuite) TestGetClient() {
	mockResource := &MockResource{client: "mockClient"}

	client := mockResource.GetClient()
	if client != "mockClient" {
		assert.Fail(suite.T(), "expected %v, got %v", "mockClient", client)
	}
}
