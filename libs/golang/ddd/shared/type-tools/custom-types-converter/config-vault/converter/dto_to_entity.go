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
