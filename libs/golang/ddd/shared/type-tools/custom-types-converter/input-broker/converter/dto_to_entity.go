package converter

import (
	"libs/golang/ddd/domain/entities/input-broker/entity"
	shareddto "libs/golang/ddd/dtos/input-broker/shared"
)

// ConvertMetadataDTOToEntity converts a MetadataDTO to a Metadata entity.
// This function maps the MetadataDTO fields to the corresponding Metadata entity fields.
//
// Parameters:
//
//	metadataDTO: A shareddto.MetadataDTO to be converted.
//
// Returns:
//
//	A entity.Metadata containing the converted data.
func ConvertMetadataDTOToEntity(metadataDTO shareddto.MetadataDTO) entity.Metadata {
	return entity.Metadata{
		Provider:            metadataDTO.Provider,
		Service:             metadataDTO.Service,
		Source:              metadataDTO.Source,
		ProcessingID:        metadataDTO.ProcessingID,
		ProcessingTimestamp: metadataDTO.ProcessingTimestamp,
	}
}

// ConvertStatusDTOToEntity converts a StatusDTO to a Status entity.
// This function maps the StatusDTO fields to the corresponding Status entity fields.
//
// Parameters:
//
//	statusDTO: A shareddto.StatusDTO to be converted.
//
// Returns:
//
//	A entity.Status containing the converted data.
func ConvertStatusDTOToEntity(statusDTO shareddto.StatusDTO) entity.Status {
	return entity.Status{
		Code:   statusDTO.Code,
		Detail: statusDTO.Detail,
	}
}
