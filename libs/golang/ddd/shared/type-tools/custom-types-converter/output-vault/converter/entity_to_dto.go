package converter

import (
	"libs/golang/ddd/domain/entities/output-vault/entity"
	shareddto "libs/golang/ddd/dtos/output-vault/shared"
)

func ConvertMetadataEntityToDTO(metadata entity.Metadata) shareddto.MetadataDTO {
	return shareddto.MetadataDTO{
		InputID: metadata.InputID,
		Input: shareddto.InputDTO{
			Data:                metadata.Input.Data,
			ProcessingID:        metadata.Input.ProcessingID,
			ProcessingTimestamp: metadata.Input.ProcessingTimestamp,
		},
	}
}
