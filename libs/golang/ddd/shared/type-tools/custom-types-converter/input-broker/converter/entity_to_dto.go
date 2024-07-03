package converter

import (
	"libs/golang/ddd/domain/entities/input-broker/entity"
	shareddto "libs/golang/ddd/dtos/input-broker/shared"
)

// ConvertMetadataEntityToDTO converts a Metadata entity to a MetadataDTO.
// This function maps the Metadata entity fields to the corresponding MetadataDTO fields.
//
// Parameters:
//
//	metadata: A entity.Metadata to be converted.
//
// Returns:
//
//	A shareddto.MetadataDTO containing the converted data.
func ConvertMetadataEntityToDTO(metadata entity.Metadata) shareddto.MetadataDTO {
	return shareddto.MetadataDTO{
		Provider:            metadata.Provider,
		Service:             metadata.Service,
		Source:              metadata.Source,
		ProcessingID:        metadata.ProcessingID,
		ProcessingTimestamp: metadata.ProcessingTimestamp,
	}
}

// ConvertStatusEntityToDTO converts a Status entity to a StatusDTO.
// This function maps the Status entity fields to the corresponding StatusDTO fields.
//
// Parameters:
//
//	status: A entity.Status to be converted.
//
// Returns:
//
//	A shareddto.StatusDTO containing the converted data.
func ConvertStatusEntityToDTO(status entity.Status) shareddto.StatusDTO {
	return shareddto.StatusDTO{
		Code:   status.Code,
		Detail: status.Detail,
	}
}
