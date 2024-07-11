package converter

import (
	"libs/golang/ddd/domain/entities/config-vault/entity"
	shareddto "libs/golang/ddd/dtos/config-vault/shared"
	"log"
)

// ConvertJobDependenciesEntityToDTO converts a slice of JobDependencies entities to a slice of JobDependenciesDTO.
// This function iterates over each JobDependencies entity and maps its fields to the corresponding JobDependenciesDTO fields.
//
// Parameters:
//
//	dependsOn: A slice of entity.JobDependencies to be converted.
//
// Returns:
//
//	A slice of shareddto.JobDependenciesDTO containing the converted data.
func ConvertJobDependenciesEntityToDTO(dependsOn []entity.JobDependencies) []shareddto.JobDependenciesDTO {
	dtoDeps := make([]shareddto.JobDependenciesDTO, len(dependsOn))
	for i, dep := range dependsOn {
		dtoDeps[i] = shareddto.JobDependenciesDTO{
			Service: dep.Service,
			Source:  dep.Source,
		}
	}
	return dtoDeps
}

// ConvertJobParametersEntityToDTO converts a JobParameters entity to a JobParametersDTO.
// This function maps the fields of the JobParameters entity to the corresponding fields of the JobParametersDTO.
//
// Parameters:
//
// jobParams: An entity.JobParameters to be converted.
//
// Returns:
//
//	A shareddto.JobParametersDTO containing the converted data.
func ConvertJobParametersEntityToDTO(jobParams entity.JobParameters) shareddto.JobParametersDTO {
	log.Printf("Converting JobParameters entity to DTO: %+v", jobParams)
	return shareddto.JobParametersDTO{
		ParserModule: jobParams.ParserModule,
	}
}
