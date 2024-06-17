package converter

import (
	"libs/golang/ddd/domain/entities/config-vault/entity"
	shareddto "libs/golang/ddd/dtos/config-vault/shared"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ConfigConverterEntityToDTOSuite struct {
	suite.Suite
}

func TestConfigConverterEntityToDTOSuite(t *testing.T) {
	suite.Run(t, new(ConfigConverterEntityToDTOSuite))
}

func (suite *ConfigConverterEntityToDTOSuite) TestConvertJobDependenciesEntityToDTO_Empty() {
	entities := []entity.JobDependencies{}
	dtoDeps := ConvertJobDependenciesEntityToDTO(entities)

	suite.Equal(0, len(dtoDeps))
}

func (suite *ConfigConverterEntityToDTOSuite) TestConvertJobDependenciesEntityToDTO_Single() {
	entities := []entity.JobDependencies{
		{
			Service: "service1",
			Source:  "source1",
		},
	}

	expected := []shareddto.JobDependenciesDTO{
		{
			Service: "service1",
			Source:  "source1",
		},
	}

	dtoDeps := ConvertJobDependenciesEntityToDTO(entities)

	suite.Equal(len(expected), len(dtoDeps))
	suite.Equal(expected, dtoDeps)
}

func (suite *ConfigConverterEntityToDTOSuite) TestConvertJobDependenciesEntityToDTO_Multiple() {
	entities := []entity.JobDependencies{
		{
			Service: "service1",
			Source:  "source1",
		},
		{
			Service: "service2",
			Source:  "source2",
		},
	}

	expected := []shareddto.JobDependenciesDTO{
		{
			Service: "service1",
			Source:  "source1",
		},
		{
			Service: "service2",
			Source:  "source2",
		},
	}

	dtoDeps := ConvertJobDependenciesEntityToDTO(entities)

	suite.Equal(len(expected), len(dtoDeps))
	suite.Equal(expected, dtoDeps)
}
