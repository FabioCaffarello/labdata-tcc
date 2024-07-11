package converter

import (
	"libs/golang/ddd/domain/entities/config-vault/entity"
	shareddto "libs/golang/ddd/dtos/config-vault/shared"
)

// ConvertJobDependenciesDTOToEntity converts a slice of JobDependenciesDTO to a slice of JobDependencies entities.
// This function iterates over each JobDependenciesDTO and maps its fields to the corresponding JobDependencies entity fields.
//
// Parameters:
//
//	dependsOnDTO: A slice of shareddto.JobDependenciesDTO to be converted.
//
// Returns:
//
//	A slice of entity.JobDependencies containing the converted data.
func ConvertJobDependenciesDTOToEntity(dependsOnDTO []shareddto.JobDependenciesDTO) []entity.JobDependencies {
	entityDeps := make([]entity.JobDependencies, len(dependsOnDTO))
	for i, dep := range dependsOnDTO {
		entityDeps[i] = entity.JobDependencies{
			Service: dep.Service,
			Source:  dep.Source,
		}
	}
	return entityDeps
}

// ConvertJobDependenciesDTOToMap converts a slice of JobDependenciesDTO to a map of JobDependencies entities.
// This function iterates over each JobDependenciesDTO and maps its fields to the corresponding JobDependencies entity fields.
//
// Parameters:
//
//	dependsOnDTO: A slice of shareddto.JobDependenciesDTO to be converted.
//
// Returns:
//
//	A map of entity.JobDependencies containing the converted data.
func ConvertJobDependenciesDTOToMap(dependsOnDTO []shareddto.JobDependenciesDTO) []map[string]interface{} {
	entityDeps := make([]map[string]interface{}, len(dependsOnDTO))
	for i, dep := range dependsOnDTO {
		entityDeps[i] = map[string]interface{}{
			"service": dep.Service,
			"source":  dep.Source,
		}
	}
	return entityDeps
}

// ConvertJobParametersDTOToEntity converts a JobParametersDTO to a JobParameters entity.
// This function maps the fields of the JobParametersDTO to the corresponding fields of the JobParameters entity.
//
// Parameters:
//
//	parmas: A shareddto.JobParametersDTO to be converted.
//
// Returns:
//
//	An entity.JobParameters containing the converted data.
func ConvertJobParametersDTOToEntity(params shareddto.JobParametersDTO) entity.JobParameters {
	return entity.JobParameters{
		ParserModule: params.ParserModule,
	}
}

// ConvertJobParametersDTOToMap converts a JobParametersDTO to a map of JobParameters entities.
// This function maps the fields of the JobParametersDTO to the corresponding fields of the JobParameters entity.
//
// Parameters:
//
//	parmas: A shareddto.JobParametersDTO to be converted.
//
// Returns:
//
//	A map of entity.JobParameters containing the converted data.
func ConvertJobParametersDTOToMap(params shareddto.JobParametersDTO) map[string]interface{} {
	entityParams := make(map[string]interface{})
	entityParams["parser_module"] = params.ParserModule
	return entityParams
}
