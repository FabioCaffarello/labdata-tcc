package usecase

import (
	"fmt"
	"libs/golang/ddd/domain/entities/events-router/entity"
	outputdto "libs/golang/ddd/dtos/events-router/output"
	inputdto "libs/golang/ddd/dtos/input-broker/output"

	"encoding/json"
	usecaseActions "libs/golang/ddd/usecases/events-router/usecase/actions"
	events "libs/golang/shared/go-events/amqp_events"
	"log"
)

var (
	errorQueue      = "error.created.pre-processing"
	baseRoutingKey  = "input.pre-processed"
	processStage    = "pre-processed"
	inputSchemaType = "input"

	invalidSchemaStatus = 401
	invalidSchemaDetail = "invalid schema"
)

// PreProcessingUseCase handles the pre-processing of input messages, including
// dispatching errors and processing orders.
type PreProcessingUseCase struct {
	EventOrderRepository entity.EventOrderRepositoryInterface
	ErrorCreated         events.EventInterface
	ProcessOrderCreated  events.EventInterface
	EventDispatcher      events.EventDispatcherInterface
}

// NewPreProcessingUseCase creates a new instance of PreProcessingUseCase.
//
// Parameters:
//   - eventOrderRepository: The repository interface for event orders.
//   - errorCreated: The event interface for error creation events.
//   - processOrderCreated: The event interface for process order creation events.
//   - eventDispatcher: The event dispatcher interface.
//
// Returns:
//   - A new instance of PreProcessingUseCase.
func NewPreProcessingUseCase(
	eventOrderRepository entity.EventOrderRepositoryInterface,
	errorCreated events.EventInterface,
	processOrderCreated events.EventInterface,
	eventDispatcher events.EventDispatcherInterface,
) *PreProcessingUseCase {
	return &PreProcessingUseCase{
		EventOrderRepository: eventOrderRepository,
		ErrorCreated:         errorCreated,
		ProcessOrderCreated:  processOrderCreated,
		EventDispatcher:      eventDispatcher,
	}
}

// dispatchError dispatches an error event with the provided error message, original message, and listener tag.
//
// Parameters:
//   - err: The error to be dispatched.
//   - msg: The original message that caused the error.
//   - listenerTag: The tag of the listener that processed the message.
func (uc *PreProcessingUseCase) dispatchError(err error, msg []byte, listenerTag string) {
	errMsg := outputdto.ErrMsgDTO{
		Err:         err,
		Msg:         msg,
		ListenerTag: listenerTag,
	}
	uc.ErrorCreated.SetPayload(errMsg)
	uc.EventDispatcher.Dispatch(uc.ErrorCreated, errorQueue)
}

// ProcessMessageChannel processes messages from the provided channel and dispatches them for further processing.
//
// Parameters:
//   - msgCh: The channel from which messages are received.
//   - listenerTag: The tag of the listener processing the messages.
func (uc *PreProcessingUseCase) ProcessMessageChannel(msgCh <-chan []byte, listenerTag string) {
	for msg := range msgCh {
		var msgDTO inputdto.InputDTO
		err := json.Unmarshal(msg, &msgDTO)
		if err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			uc.dispatchError(err, msg, listenerTag)
		} else {
			log.Printf("Message received: %v", msgDTO)
			err = uc.execute(msgDTO)

			if err != nil {
				log.Printf("Error processing message: %v", err)
				uc.dispatchError(err, msg, listenerTag)
			}
		}
	}
}

// execute processes the input message and dispatches the processed order.
//
// Parameters:
//   - msgDTO: The input message DTO to be processed.
//
// Returns:
//   - An error if the processing fails, otherwise nil.
func (uc *PreProcessingUseCase) execute(msgDTO inputdto.InputDTO) error {
	eventOrcerProps := entity.EventOrderProps{
		Service:      msgDTO.Metadata.Service,
		Source:       msgDTO.Metadata.Source,
		Provider:     msgDTO.Metadata.Provider,
		InputID:      msgDTO.ID,
		ProcessingID: msgDTO.Metadata.ProcessingID,
		Stage:        processStage,
		Data:         msgDTO.Data,
	}

	eventOrder, err := entity.NewEventOrder(eventOrcerProps)
	if err != nil {
		return err
	}

	err = uc.EventOrderRepository.Create(eventOrder)
	if err != nil {
		return err
	}

	dto := outputdto.ProcessOrderDTO{
		ID:           eventOrder.GetEntityID(),
		ProcessingID: eventOrder.ProcessingID,
		Service:      eventOrder.Service,
		Source:       eventOrder.Source,
		Provider:     eventOrder.Provider,
		Stage:        eventOrder.Stage,
		Data:         eventOrder.Data,
	}

	err = uc.prepareInputToProcess(dto)
	if err != nil {
		return err
	}

	uc.ProcessOrderCreated.SetPayload(dto)
	routingKey := fmt.Sprintf("%s.%s.%s.%s", baseRoutingKey, dto.Provider, dto.Service, dto.Source)
	uc.EventDispatcher.Dispatch(uc.ProcessOrderCreated, routingKey)
	uc.EventOrderRepository.Delete(eventOrder.GetEntityID())
	return nil
}

func (uc *PreProcessingUseCase) prepareInputToProcess(inputMsg outputdto.ProcessOrderDTO) error {
	log.Printf("Preparing input to process: %v", inputMsg)
	// TODO: create pre-processing methods
	// 1. Validate input
	validateSchemaAction := usecaseActions.NewValidateSchemaAction()
	err := validateSchemaAction.Execute(inputMsg, inputSchemaType)
	if err != nil {
		updateInputStatusAction := usecaseActions.NewUpdateInputStatusAction()
		err := updateInputStatusAction.Execute(inputMsg, invalidSchemaStatus, invalidSchemaDetail)
		return err
	}

	// 2. List Configs by dependencies
	dependenciesAction := usecaseActions.NewListAllByDependenciesAction()
	dependencie, err := dependenciesAction.Execute(inputMsg.Provider, inputMsg.Service, inputMsg.Source)
	if err != nil {
		return err
	}

	for _, dep := range dependencie {
		log.Printf("Dependency: %v", dep)
	}
	// 3. Create processing staging
	// 4. Create processing lineage ??
	return nil
}
