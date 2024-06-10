package regulartypetool

import (
	"reflect"
	"testing"
	"time"

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
	assert.IsType(suite.T(), &TestEntity{}, entity)

	testEntity := entity.(*TestEntity)
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
	assert.IsType(suite.T(), &TestEntityNested{}, entity)

	testEntityNested := entity.(*TestEntityNested)
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

	testEntity1 := entities[0].(*TestEntity)
	assert.Equal(suite.T(), "value1", testEntity1.Field1)
	assert.Equal(suite.T(), 123, testEntity1.Field2)

	testEntity2 := entities[1].(*TestEntity)
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

	testEntityNested1 := entities[0].(*TestEntityNested)
	assert.Equal(suite.T(), "value3", testEntityNested1.Field3)
	assert.Equal(suite.T(), "value1", testEntityNested1.Field4.Field1)
	assert.Equal(suite.T(), 123, testEntityNested1.Field4.Field2)

	testEntityNested2 := entities[1].(*TestEntityNested)
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

func (suite *RegularTypeToolSuite) TestConvertFromMapStringToEntityWithPointerToStruct() {
	type TestEntity struct {
		Field1 string `bson:"field1"`
		Field2 int    `bson:"field2"`
	}

	testData := map[string]interface{}{
		"field1": "value1",
		"field2": 123,
	}

	entity, err := ConvertFromMapStringToEntity(reflect.TypeOf(&TestEntity{}).Elem(), testData)

	assert.NoError(suite.T(), err)
	assert.IsType(suite.T(), &TestEntity{}, entity)

	testEntity := entity.(*TestEntity)
	assert.Equal(suite.T(), "value1", testEntity.Field1)
	assert.Equal(suite.T(), 123, testEntity.Field2)
}

func (suite *RegularTypeToolSuite) TestConvertFromMapStringToEntityWithMissingFields() {
	type TestEntity struct {
		Field1 string `bson:"field1"`
		Field2 int    `bson:"field2"`
	}

	testData := map[string]interface{}{
		"field1": "value1",
	}

	_, err := ConvertFromMapStringToEntity(reflect.TypeOf(TestEntity{}), testData)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "field field2: missing value", err.Error())
}

func (suite *RegularTypeToolSuite) TestConvertFromMapStringToEntityWithInvalidFieldTypes() {
	type TestEntity struct {
		Field1 string `bson:"field1"`
		Field2 int    `bson:"field2"`
	}

	testData := map[string]interface{}{
		"field1": "value1",
		"field2": "notAnInt",
	}

	_, err := ConvertFromMapStringToEntity(reflect.TypeOf(TestEntity{}), testData)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "field field2: cannot convert string to int", err.Error())
}

func (suite *RegularTypeToolSuite) TestConvertFromArrayMapStringToEntitiesWithInvalidFieldTypes() {
	type TestEntity struct {
		Field1 string `bson:"field1"`
		Field2 int    `bson:"field2"`
	}

	testDataArray := []map[string]interface{}{
		{
			"field1": "value1",
			"field2": "notAnInt",
		},
		{
			"field1": "value2",
			"field2": 456,
		},
	}

	_, err := ConvertFromArrayMapStringToEntities(reflect.TypeOf(TestEntity{}), testDataArray)

	assert.Error(suite.T(), err)
}

func (suite *RegularTypeToolSuite) TestConvertFromEntityToMapStringWithUnexportedFields() {
	type TestEntity struct {
		Field1 string `bson:"field1"`
		field2 int    // unexported field
	}

	testEntity := TestEntity{
		Field1: "value1",
	}

	testData, err := ConvertFromEntityToMapString(testEntity)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "value1", testData["field1"])
	_, exists := testData["field2"]
	assert.False(suite.T(), exists)
}

func (suite *RegularTypeToolSuite) TestSetFieldValue() {
	type TestEntity struct {
		Field1 string    `bson:"field1"`
		Field2 int       `bson:"field2"`
		Field3 time.Time `bson:"field3"`
	}

	entity := reflect.New(reflect.TypeOf(TestEntity{})).Elem()
	err := setFieldValue(entity, reflect.TypeOf(TestEntity{}).Field(0), "value1")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "value1", entity.Field(0).Interface())

	err = setFieldValue(entity, reflect.TypeOf(TestEntity{}).Field(1), 123)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 123, entity.Field(1).Interface())

	timeStr := "2023-01-01 12:00:00"
	err = setFieldValue(entity, reflect.TypeOf(TestEntity{}).Field(2), timeStr)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC), entity.Field(2).Interface())
}

func (suite *RegularTypeToolSuite) TestSetStructField() {
	type TestEntityNested struct {
		Field1 string `bson:"field1"`
		Field2 int    `bson:"field2"`
	}
	type TestEntity struct {
		Field3 string           `bson:"field3"`
		Field4 TestEntityNested `bson:"field4"`
	}

	entity := reflect.New(reflect.TypeOf(TestEntity{})).Elem()

	nestedData := map[string]interface{}{
		"field1": "value1",
		"field2": 123,
	}
	err := setStructField(entity.FieldByName("Field4"), reflect.TypeOf(TestEntity{}).Field(1), nestedData)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "value1", entity.FieldByName("Field4").FieldByName("Field1").Interface())
	assert.Equal(suite.T(), 123, entity.FieldByName("Field4").FieldByName("Field2").Interface())
}

func (suite *RegularTypeToolSuite) TestSetSliceField() {
	type TestEntity struct {
		Field1 []int    `bson:"field1"`
		Field2 []string `bson:"field2"`
	}

	entity := reflect.New(reflect.TypeOf(TestEntity{})).Elem()

	sliceData := []interface{}{1, 2, 3}
	err := setSliceField(entity.FieldByName("Field1"), reflect.TypeOf(TestEntity{}).Field(0), sliceData)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), []int{1, 2, 3}, entity.FieldByName("Field1").Interface())

	stringSliceData := []interface{}{"a", "b", "c"}
	err = setSliceField(entity.FieldByName("Field2"), reflect.TypeOf(TestEntity{}).Field(1), stringSliceData)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), []string{"a", "b", "c"}, entity.FieldByName("Field2").Interface())
}

func (suite *RegularTypeToolSuite) TestSetBasicField() {
	type TestEntity struct {
		Field1 string `bson:"field1"`
		Field2 int    `bson:"field2"`
	}

	entity := reflect.New(reflect.TypeOf(TestEntity{})).Elem()
	err := setBasicField(entity.FieldByName("Field1"), reflect.TypeOf(TestEntity{}).Field(0), "value1")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "value1", entity.FieldByName("Field1").Interface())

	err = setBasicField(entity.FieldByName("Field2"), reflect.TypeOf(TestEntity{}).Field(1), 123)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 123, entity.FieldByName("Field2").Interface())
}

func (suite *RegularTypeToolSuite) TestSetTimeField() {
	type TestEntity struct {
		Field1 time.Time `bson:"field1"`
	}

	entity := reflect.New(reflect.TypeOf(TestEntity{})).Elem()

	timeStr := "2023-01-01 12:00:00"
	err := setTimeField(entity.FieldByName("Field1"), timeStr)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC), entity.FieldByName("Field1").Interface())

	timeValue := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
	err = setTimeField(entity.FieldByName("Field1"), timeValue)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), timeValue, entity.FieldByName("Field1").Interface())
}
