package usecase

import (
	"libs/golang/ddd/domain/entities/output-vault/entity"
	outputdto "libs/golang/ddd/dtos/output-vault/output"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/output-vault/converter"
)

// ListOneByIDOutputUseCase is the use case for listing a output by ID.
type ListOneByIDOutputUseCase struct {
	OutputRepository entity.OutputRepositoryInterface
}

// NewListOneByIDOutputUseCase initializes a new instance of ListOneByIDOutputUseCase with the provided OutputRepositoryInterface.
//
// Parameters:
//
//	outputRepository: The repository interface for managing Output entities.
//
// Returns:
//
//	A pointer to an instance of ListOneByIDOutputUseCase.
func NewListOneByIDOutputUseCase(
	outputRepository entity.OutputRepositoryInterface,
) *ListOneByIDOutputUseCase {
	return &ListOneByIDOutputUseCase{
		OutputRepository: outputRepository,
	}
}

// Execute retrieves a output by its ID from the repository and converts it to an output DTO.
//
// Parameters:
//
//	id: The ID of the output to retrieve.
//
// Returns:
//
//	An output DTO containing the output data, and an error if any occurred during the process.
func (uc *ListOneByIDOutputUseCase) Execute(id string) (outputdto.OutputDTO, error) {
	output, err := uc.OutputRepository.FindByID(id)
	if err != nil {
		return outputdto.OutputDTO{}, err
	}

	dto := outputdto.OutputDTO{
		ID:        string(output.ID),
		Service:   output.Service,
		Source:    output.Source,
		Provider:  output.Provider,
		Data:      output.Data,
		Metadata:  converter.ConvertMetadataEntityToDTO(output.Metadata),
		CreatedAt: output.CreatedAt,
		UpdatedAt: output.UpdatedAt,
	}

	return dto, nil
}
