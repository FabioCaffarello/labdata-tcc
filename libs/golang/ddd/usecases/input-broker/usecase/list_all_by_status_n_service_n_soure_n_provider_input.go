package usecase

import (
	"libs/golang/ddd/domain/entities/input-broker/entity"
	outputdto "libs/golang/ddd/dtos/input-broker/output"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/input-broker/converter"
)

// ListAllByStatusAndServiceAndSourceAndProviderInputUseCaseSuite is the use case for listing all inputs by service and provider.
type ListAllByStatusAndServiceAndSourceAndProviderInputUseCaseSuite struct {
	InputRepository entity.InputRepositoryInterface
}

// NewListAllByStatusAndServiceAndSourceAndProviderInputUseCaseSuite initializes a new instance of ListAllByStatusAndServiceAndSourceAndProviderInputUseCaseSuite with the provided InputRepositoryInterface.
//
// Parameters:
//
//	inputRepository: The repository interface for managing Input entities.
//
// Returns:
//
//	A pointer to an instance of ListAllByStatusAndServiceAndSourceAndProviderInputUseCaseSuite.
func NewListAllByStatusAndServiceAndSourceAndProviderInputUseCaseSuite(
	inputRepository entity.InputRepositoryInterface,
) *ListAllByStatusAndServiceAndSourceAndProviderInputUseCaseSuite {
	return &ListAllByStatusAndServiceAndSourceAndProviderInputUseCaseSuite{
		InputRepository: inputRepository,
	}
}

// Execute retrieves all inputs by service and provider from the repository.
//
// Parameters:
//
//		provider: The provider name to filter inputs by.
//		service: The service name to filter inputs by.
//	 source: The source name to filter inputs by.
//		status: The status to filter inputs by.
//
// Returns:
//
//	A slice of input DTOs containing the input data, and an error if any occurred during the process.
func (uc *ListAllByStatusAndServiceAndSourceAndProviderInputUseCaseSuite) Execute(provider, service, source string, status int) ([]outputdto.InputDTO, error) {
	inputs, err := uc.InputRepository.FindAllByStatusAndServiceAndSourceAndProvider(service, source, provider, status)
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
