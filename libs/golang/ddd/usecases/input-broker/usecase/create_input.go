package usecase

import (
	"fmt"
	"libs/golang/ddd/domain/entities/input-broker/entity"
	inputdto "libs/golang/ddd/dtos/input-broker/input"
	outputdto "libs/golang/ddd/dtos/input-broker/output"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/input-broker/converter"
	events "libs/golang/shared/go-events/amqp_events"
)

var (
	exchangeName = "services"
	routingKey   = "input.created"
)

// CreateInputUseCase represents the use case for creating an input.
type CreateInputUseCase struct {
	InputRepository entity.InputRepositoryInterface
	InputCreated    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

// NewCreateInputUseCase creates a new CreateInputUseCase.
//
// Parameters:
//
//	inputRepository: The repository interface for managing Input entities.
//	inputCreated: The event to be dispatched when an input is created.
//	eventDispatcher: The event dispatcher to dispatch the input created event.
//
// Returns:
//
//	A pointer to an instance of CreateInputUseCase.
func NewCreateInputUseCase(
	inputRepository entity.InputRepositoryInterface,
	inputCreated events.EventInterface,
	eventDispatcher events.EventDispatcherInterface,
) *CreateInputUseCase {
	return &CreateInputUseCase{
		InputRepository: inputRepository,
		InputCreated:    inputCreated,
		EventDispatcher: eventDispatcher,
	}
}

// Execute creates a new input entity based on the provided input DTO and saves it using the repository.
// It then returns the created entity and an error if any occurred during the process.
//
// Parameters:
//
//	input: The input DTO containing the input data.
//
// Returns:
//
//	An input DTO containing the created input data, and an error if any occurred during the process.
func (uc *CreateInputUseCase) Execute(input inputdto.InputDTO) (outputdto.InputDTO, error) {
	inputProps := entity.InputProps{
		Provider: input.Provider,
		Service:  input.Service,
		Source:   input.Source,
		Data:     input.Data,
	}

	entityInput, err := entity.NewInput(inputProps)
	if err != nil {
		return outputdto.InputDTO{}, err
	}

	err = uc.InputRepository.Create(entityInput)
	if err != nil {
		return outputdto.InputDTO{}, err
	}

	dto := outputdto.InputDTO{
		ID:        string(entityInput.ID),
		Data:      input.Data,
		Metadata:  converter.ConvertMetadataEntityToDTO(entityInput.Metadata),
		Status:    converter.ConvertStatusEntityToDTO(entityInput.Status),
		CreatedAt: entityInput.CreatedAt,
		UpdatedAt: entityInput.UpdatedAt,
	}

	uc.InputCreated.SetPayload(dto)
	uc.EventDispatcher.Dispatch(uc.InputCreated, fmt.Sprintf("%s.%s.%s.%s", routingKey, input.Provider, input.Service, input.Source))

	return dto, nil
}
