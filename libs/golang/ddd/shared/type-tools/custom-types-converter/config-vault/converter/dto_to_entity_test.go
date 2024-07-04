package converter

import (
	"libs/golang/ddd/domain/entities/config-vault/entity"
	shareddto "libs/golang/ddd/dtos/config-vault/shared"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ConfigConverterDTOToEntitySuite struct {
	suite.Suite
}

func TestConfigConverterDTOToEntitySuite(t *testing.T) {
	suite.Run(t, new(ConfigConverterDTOToEntitySuite))
}

func (suite *ConfigConverterDTOToEntitySuite) TestConvertJobDependenciesDTOToEntity_Empty() {
	dto := []shareddto.JobDependenciesDTO{}
	entityDeps := ConvertJobDependenciesDTOToEntity(dto)

	suite.Equal(0, len(entityDeps))
}

func (suite *ConfigConverterDTOToEntitySuite) TestConvertJobDependenciesDTOToEntity_Single() {
	dto := []shareddto.JobDependenciesDTO{
		{
			Service: "service1",
			Source:  "source1",
		},
	}

	expected := []entity.JobDependencies{
		{
			Service: "service1",
			Source:  "source1",
		},
	}

	entityDeps := ConvertJobDependenciesDTOToEntity(dto)

	assert.Equal(suite.T(), len(expected), len(entityDeps))
	assert.Equal(suite.T(), expected, entityDeps)
}

func (suite *ConfigConverterDTOToEntitySuite) TestConvertJobDependenciesDTOToEntity_Multiple() {
	dto := []shareddto.JobDependenciesDTO{
		{
			Service: "service1",
			Source:  "source1",
		},
		{
			Service: "service2",
			Source:  "source2",
		},
	}

	expected := []entity.JobDependencies{
		{
			Service: "service1",
			Source:  "source1",
		},
		{
			Service: "service2",
			Source:  "source2",
		},
	}

	entityDeps := ConvertJobDependenciesDTOToEntity(dto)

	assert.Equal(suite.T(), len(expected), len(entityDeps))
	assert.Equal(suite.T(), expected, entityDeps)
}

func (suite *ConfigConverterDTOToEntitySuite) TestConvertJobDependenciesDTOToMapWhenEmpty() {
	dto := []shareddto.JobDependenciesDTO{}
	entityMap := ConvertJobDependenciesDTOToMap(dto)

	assert.Equal(suite.T(), 0, len(entityMap))
}

func (suite *ConfigConverterDTOToEntitySuite) TestConvertJobDependenciesDTOToMapWhenSingle() {
	dto := []shareddto.JobDependenciesDTO{
		{
			Service: "service1",
			Source:  "source1",
		},
	}

	expected := []map[string]interface{}{
		{
			"service": "service1",
			"source":  "source1",
		},
	}

	entityMap := ConvertJobDependenciesDTOToMap(dto)

	assert.Equal(suite.T(), len(expected), len(entityMap))
	assert.Equal(suite.T(), expected, entityMap)
}

func (suite *ConfigConverterDTOToEntitySuite) TestConvertJobDependenciesDTOToMapWhenMultiple() {
	dto := []shareddto.JobDependenciesDTO{
		{
			Service: "service1",
			Source:  "source1",
		},
		{
			Service: "service2",
			Source:  "source2",
		},
	}

	expected := []map[string]interface{}{
		{
			"service": "service1",
			"source":  "source1",
		},
		{
			"service": "service2",
			"source":  "source2",
		},
	}

	entityMap := ConvertJobDependenciesDTOToMap(dto)

	assert.Equal(suite.T(), len(expected), len(entityMap))
	assert.Equal(suite.T(), expected, entityMap)
}
