package typetools

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TypeToolsSuite struct {
	suite.Suite
}

func TestTypeToolsSuite(t *testing.T) {
	suite.Run(t, new(TypeToolsSuite))
}

func (suite *TypeToolsSuite) TestMapInterfaceToString() {
	data := map[string]interface{}{
		"foo": "bar",
		"baz": 123,
	}
	expected := "baz:123;foo:bar;"
	assert.Equal(suite.T(), expected, MapInterfaceToString(data))
}

func (suite *TypeToolsSuite) TestMapStringToString() {
	data := map[string]string{
		"foo": "bar",
		"baz": "qux",
	}
	expected := "baz:qux;foo:bar;"
	assert.Equal(suite.T(), expected, MapStringToString(data))
}

func (suite *TypeToolsSuite) TestToString() {
	str, err := ToString("test")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "test", str)

	str, err = ToString(123.456)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "123.456000", str)

	str, err = ToString(map[string]interface{}{"foo": "bar", "baz": 123})
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "baz:123;foo:bar;", str)

	str, err = ToString(map[string]string{"foo": "bar", "baz": "qux"})
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "baz:qux;foo:bar;", str)
}

func (suite *TypeToolsSuite) TestToStringWithInt() {
	str, err := ToString(123)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "123", str)
}

func (suite *TypeToolsSuite) TestToStringWithBool() {
	str, err := ToString(true)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "true", str)

	str, err = ToString(false)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "false", str)
}

func (suite *TypeToolsSuite) TestToFloat64() {
	num, err := ToFloat64(123.456)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 123.456, num)

	num, err = ToFloat64("123.456")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 123.456, num)

	num, err = ToFloat64(123)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 123.0, num)

	num, err = ToFloat64(uint(123))
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 123.0, num)
}

func (suite *TypeToolsSuite) TestToFloat64WithInvalidInput() {
	_, err := ToFloat64("invalid")
	assert.Error(suite.T(), err)

	_, err = ToFloat64(true)
	assert.Error(suite.T(), err)
}

func (suite *TypeToolsSuite) TestToFloat64WithInt() {
	num, err := ToFloat64(123)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 123.0, num)
}

func (suite *TypeToolsSuite) TestToFloat64WithUint() {
	num, err := ToFloat64(uint(123))
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 123.0, num)
}

func (suite *TypeToolsSuite) TestToFloat64WithFloat() {
	num, err := ToFloat64(123.456)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 123.456, num)
}

func (suite *TypeToolsSuite) TestToFloat64WithFloatString() {
	num, err := ToFloat64("123.456")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 123.456, num)
}

func (suite *TypeToolsSuite) TestToFloat64WithIntString() {
	num, err := ToFloat64("123")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 123.0, num)
}

func (suite *TypeToolsSuite) TestToFloat64WithUintString() {
	num, err := ToFloat64("123")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 123.0, num)
}

func (suite *TypeToolsSuite) TestToFloat64WithInvalidString() {
	_, err := ToFloat64("invalid")
	assert.Error(suite.T(), err)
}

func (suite *TypeToolsSuite) TestParseDateWithFormat() {
	date, err := ParseDateWithFormat("2021-01-01", "2006-01-02")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "2021-01-01 00:00:00 +0000 UTC", date.String())
}

func (suite *TypeToolsSuite) TestParseDateWithFormatWithInvalidDate() {
	_, err := ParseDateWithFormat("invalid", "2006-01-02")
	assert.Error(suite.T(), err)
}

func (suite *TypeToolsSuite) TestParseDateWithFormatWithInvalidFormat() {
	_, err := ParseDateWithFormat("2021-01-01", "invalid")
	assert.Error(suite.T(), err)
}
