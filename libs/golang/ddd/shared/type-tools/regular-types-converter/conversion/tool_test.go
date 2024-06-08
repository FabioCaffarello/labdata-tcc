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

func (suite *RegularTypeToolSuite) TestConvertFromMapStringToEntity() {
	type TestEntity struct {
		Field1 string `bson:"field1"`
		Field2 int    `bson:"field2"`
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

func (suite *RegularTypeToolSuite) TestConvertFromMapStringToEntityWithNestedStructureTypes() {
	type TestEntity struct {
		Field1 string `bson:"field1"`
		Field2 int    `bson:"field2"`
	}

	type TestEntityNested struct {
		Field3 string     `bson:"field3"`
		Field4 TestEntity `bson:"field4"`
	}

	testData := map[string]interface{}{
		"field3": "value3",
		"field4": map[string]interface{}{
			"field1": "value1",
			"field2": 123,
		},
	}

	entity, err := ConvertFromMapStringToEntity(reflect.TypeOf(TestEntityNested{}), testData)

	assert.NoError(suite.T(), err)
	assert.IsType(suite.T(), TestEntityNested{}, entity)

	testEntityNested := entity.(TestEntityNested)
	assert.Equal(suite.T(), "value3", testEntityNested.Field3)
	assert.Equal(suite.T(), "value1", testEntityNested.Field4.Field1)
	assert.Equal(suite.T(), 123, testEntityNested.Field4.Field2)
}

func (suite *RegularTypeToolSuite) TestConvertFromArrayMapStringToEntities() {
	type TestEntity struct {
		Field1 string `bson:"field1"`
		Field2 int    `bson:"field2"`
	}

	testDataArray := []map[string]interface{}{
		{
			"field1": "value1",
			"field2": 123,
		},
		{
			"field1": "value2",
			"field2": 456,
		},
	}

	entities, err := ConvertFromArrayMapStringToEntities(reflect.TypeOf(TestEntity{}), testDataArray)

	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), entities, 2)

	testEntity1 := entities[0].(TestEntity)
	assert.Equal(suite.T(), "value1", testEntity1.Field1)
	assert.Equal(suite.T(), 123, testEntity1.Field2)

	testEntity2 := entities[1].(TestEntity)
	assert.Equal(suite.T(), "value2", testEntity2.Field1)
	assert.Equal(suite.T(), 456, testEntity2.Field2)
}

func (suite *RegularTypeToolSuite) TestConvertFromArrayMapStringToEntitiesWithNestedStructureTypes() {
	type TestEntity struct {
		Field1 string `bson:"field1"`
		Field2 int    `bson:"field2"`
	}

	type TestEntityNested struct {
		Field3 string     `bson:"field3"`
		Field4 TestEntity `bson:"field4"`
	}

	testDataArray := []map[string]interface{}{
		{
			"field3": "value3",
			"field4": map[string]interface{}{
				"field1": "value1",
				"field2": 123,
			},
		},
		{
			"field3": "value4",
			"field4": map[string]interface{}{
				"field1": "value2",
				"field2": 456,
			},
		},
	}

	entities, err := ConvertFromArrayMapStringToEntities(reflect.TypeOf(TestEntityNested{}), testDataArray)

	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), entities, 2)

	testEntityNested1 := entities[0].(TestEntityNested)
	assert.Equal(suite.T(), "value3", testEntityNested1.Field3)
	assert.Equal(suite.T(), "value1", testEntityNested1.Field4.Field1)
	assert.Equal(suite.T(), 123, testEntityNested1.Field4.Field2)

	testEntityNested2 := entities[1].(TestEntityNested)
	assert.Equal(suite.T(), "value4", testEntityNested2.Field3)
	assert.Equal(suite.T(), "value2", testEntityNested2.Field4.Field1)
	assert.Equal(suite.T(), 456, testEntityNested2.Field4.Field2)
}

func (suite *RegularTypeToolSuite) TestConvertFromEntityToMapString() {
	type TestEntity struct {
		Field1 string `bson:"field1"`
		Field2 int    `bson:"field2"`
	}

	testEntity := TestEntity{
		Field1: "value1",
		Field2: 123,
	}

	testData, err := ConvertFromEntityToMapString(testEntity)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "value1", testData["field1"])
	assert.Equal(suite.T(), 123, testData["field2"])
}

func (suite *RegularTypeToolSuite) TestConvertFromEntityToMapStringWithNestedStructureTypes() {
	type TestEntity struct {
		Field1 string `bson:"field1"`
		Field2 int    `bson:"field2"`
	}

	type TestEntityNested struct {
		Field3 string     `bson:"field3"`
		Field4 TestEntity `bson:"field4"`
	}

	testEntityNested := TestEntityNested{
		Field3: "value3",
		Field4: TestEntity{
			Field1: "value1",
			Field2: 123,
		},
	}

	testData, err := ConvertFromEntityToMapString(testEntityNested)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "value3", testData["field3"])
	assert.Equal(suite.T(), "value1", testData["field4"].(map[string]interface{})["field1"])
	assert.Equal(suite.T(), 123, testData["field4"].(map[string]interface{})["field2"])
}
