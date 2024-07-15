package usecase

import (
	"libs/golang/ddd/domain/entities/input-broker/entity"
	outputdto "libs/golang/ddd/dtos/input-broker/output"
	shareddto "libs/golang/ddd/dtos/input-broker/shared"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/input-broker/converter"
	"time"
)

// UpdateStatusInputUseCase is the use case for updating status of an existing input.
type UpdateStatusInputUseCase struct {
	InputRepository entity.InputRepositoryInterface
}

// NewUpdateStatusInputUseCase initializes a new instance of UpdateStatusInputUseCase with the provided InputRepositoryInterface.
//
// Parameters:
//
//	inputRepository: The repository interface for managing Input entities.
//
// Returns:
//
//	A pointer to an instance of UpdateStatusInputUseCase.
func NewUpdateStatusInputUseCase(
	inputRepository entity.InputRepositoryInterface,
) *UpdateStatusInputUseCase {
	return &UpdateStatusInputUseCase{
		InputRepository: inputRepository,
	}
}

// Execute updates status of an existing input entity based on the provided input DTO and saves it using the repository.
// It then converts the updated entity to an output DTO and returns it.
//
// Parameters:
//
//	input: The input DTO containing the input data.
//
// Returns:
//
//	An output DTO containing the updated input data, and an error if any occurred during the process.
func (uc *UpdateStatusInputUseCase) Execute(id string, status shareddto.StatusDTO) (outputdto.InputDTO, error) {
	entityInput, err := uc.InputRepository.FindByID(id)
	if err != nil {
		return outputdto.InputDTO{}, err
	}

	entityInput.SetStatus(status.Code, status.Detail)
	entityInput.SetUpdatedAt(time.Now().Format(entity.DateLayout))
	entityInput.SetProcessingTimestamp(time.Now())

	err = uc.InputRepository.Update(entityInput)
	if err != nil {
		return outputdto.InputDTO{}, err
	}

	return outputdto.InputDTO{
		ID:        string(entityInput.ID),
		Data:      entityInput.Data,
		Metadata:  converter.ConvertMetadataEntityToDTO(entityInput.Metadata),
		Status:    converter.ConvertStatusEntityToDTO(entityInput.Status),
		CreatedAt: entityInput.CreatedAt,
		UpdatedAt: entityInput.UpdatedAt,
	}, nil
}
