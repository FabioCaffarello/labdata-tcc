package regulartypetool

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RegularTypeToolSuite struct {
	suite.Suite
}

func TestRegularTypeToolSuite(t *testing.T) {
	suite.Run(t, new(RegularTypeToolSuite))
}

func (suite *RegularTypeToolSuite) TestConvertFromMapStringStringToEntity() {
	type TestEntity struct {
		Field1 string `json:"field1"`
		Field2 int    `json:"field2"`
	}

	testData := map[string]interface{}{
		"field1": "value1",
		"field2": 123,
	}

	entity, err := ConvertFromMapStringToEntity(reflect.TypeOf(TestEntity{}), testData)

	assert.NoError(suite.T(), err)
	assert.IsType(suite.T(), TestEntity{}, entity)

	testEntity := entity.(TestEntity)
	assert.Equal(suite.T(), "value1", testEntity.Field1)
	assert.Equal(suite.T(), 123, testEntity.Field2)
}
