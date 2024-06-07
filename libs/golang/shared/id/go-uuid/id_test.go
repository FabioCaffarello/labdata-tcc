package gouuid

import (
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GoUUIDSuite struct {
	suite.Suite
}

func TestGoUUIDSuite(t *testing.T) {
	suite.Run(t, new(GoUUIDSuite))
}

func (suite *GoUUIDSuite) TestHashConfig() {
	data := []byte("test")
	expectedHash := sha256.Sum256(data)
	actualHash := hashData(data)
	assert.Equal(suite.T(), expectedHash[:], actualHash)
}

func (suite *GoUUIDSuite) TestGenerateUUIDFromHash() {
	hash := hashData([]byte("test"))
	expectedUUID := "70db741d-cea6-50d9-ae7a-de207c2758b5"
	actualUUID := generateUUIDFromHash(hash)
	assert.Equal(suite.T(), expectedUUID, actualUUID)
}

func (suite *GoUUIDSuite) TestGenerateUUIDFromMap() {
	testMap := map[string]interface{}{
		"key1": "value1",
		"key2": 123,
	}

	uuid, err := GenerateUUIDFromMap(testMap)

	assert.NoError(suite.T(), err)
	expectedUUID := "64ac103b-4dce-5f5b-ab40-57a292efe74e"
	assert.Equal(suite.T(), expectedUUID, uuid)
}
