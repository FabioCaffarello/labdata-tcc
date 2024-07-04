package converter

import (
	"libs/golang/ddd/domain/entities/output-vault/entity"
	shareddto "libs/golang/ddd/dtos/output-vault/shared"
)

// ConvertOutputDTOToEntity converts an OutputDTO to an Output entity.
// This function maps the fields of the OutputDTO to the corresponding Output entity fields.
func ConvertMetadataDTOToEntity(metadataDTO shareddto.MetadataDTO) entity.Metadata {
	return entity.Metadata{
		InputID: metadataDTO.InputID,
		Input: entity.Input{
			Data:                metadataDTO.Input.Data,
			ProcessingID:        metadataDTO.Input.ProcessingID,
			ProcessingTimestamp: metadataDTO.Input.ProcessingTimestamp,
		},
	}
}

// ConvertOutputDTOToMap converts an OutputDTO to a map.
// This function maps the fields of the OutputDTO to the corresponding map fields.
func ConvertMetadataDTOToMap(metadataDTO shareddto.MetadataDTO) map[string]interface{} {
	return map[string]interface{}{
		"input_id": metadataDTO.InputID,
		"input": map[string]interface{}{
			"data":                 metadataDTO.Input.Data,
			"processing_id":        metadataDTO.Input.ProcessingID,
			"processing_timestamp": metadataDTO.Input.ProcessingTimestamp,
		},
	}
}
