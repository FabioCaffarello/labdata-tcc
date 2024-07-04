package usecase

import (
	"libs/golang/ddd/domain/entities/input-broker/entity"
	outputdto "libs/golang/ddd/dtos/input-broker/output"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/input-broker/converter"
)

// ListAllByStatusAndServiceAndProviderInputUseCase is the use case for listing all inputs by service and provider.
type ListAllByStatusAndServiceAndProviderInputUseCase struct {
	InputRepository entity.InputRepositoryInterface
}

// NewListAllByStatusAndServiceAndProviderInputUseCase initializes a new instance of ListAllByStatusAndServiceAndProviderInputUseCase with the provided InputRepositoryInterface.
//
// Parameters:
//
//	inputRepository: The repository interface for managing Input entities.
//
// Returns:
//
//	A pointer to an instance of ListAllByStatusAndServiceAndProviderInputUseCase.
func NewListAllByStatusAndServiceAndProviderInputUseCase(
	inputRepository entity.InputRepositoryInterface,
) *ListAllByStatusAndServiceAndProviderInputUseCase {
	return &ListAllByStatusAndServiceAndProviderInputUseCase{
		InputRepository: inputRepository,
	}
}

// Execute retrieves all inputs by service and provider from the repository.
//
// Parameters:
//
//	provider: The provider name to filter inputs by.
//	service: The service name to filter inputs by.
//	status: The status to filter inputs by.
//
// Returns:
//
//	A slice of input DTOs containing the input data, and an error if any occurred during the process.
func (uc *ListAllByStatusAndServiceAndProviderInputUseCase) Execute(provider, service string, status int) ([]outputdto.InputDTO, error) {
	inputs, err := uc.InputRepository.FindAllByStatusAndServiceAndProvider(provider, service, status)
	if err != nil {
		return []outputdto.InputDTO{}, err
	}

	inputDTOs := make([]outputdto.InputDTO, 0, len(inputs))
	for _, input := range inputs {
		inputDTOs = append(inputDTOs, outputdto.InputDTO{
			ID:        string(input.ID),
			Metadata:  converter.ConvertMetadataEntityToDTO(input.Metadata),
			Status:    converter.ConvertStatusEntityToDTO(input.Status),
			Data:      input.Data,
			CreatedAt: input.CreatedAt,
			UpdatedAt: input.UpdatedAt,
		})
	}

	return inputDTOs, nil
}
