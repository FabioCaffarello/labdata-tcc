package main

import (
	inMemoryDBClient "libs/golang/clients/resources/go-docdb/client"
	gorabbitmq "libs/golang/clients/resources/go-rabbitmq/client"
	inMemoryDB "libs/golang/database/go-docdb/database"
	inMemoryDBRepository "libs/golang/ddd/domain/repositories/database/in-memory/go-docdb/events-router/repository"
	event "libs/golang/ddd/events/events-router/event"
	eventHandlers "libs/golang/ddd/events/events-router/handlers"
	"libs/golang/ddd/usecases/events-router/usecase"
	amqpConsumer "libs/golang/server/events/amqp-consumer/consumer"
	eventServer "libs/golang/server/events/event-server/server"
	eventListener "libs/golang/server/events/listener/listener"
	servicediscovery "libs/golang/service-discovery/sd"
	events "libs/golang/shared/go-events/amqp_events"
	"log"
	"os"
)

var (
	dbName                  = os.Getenv("DOCDB_DBNAME")  // "events-order"
	consumerName            = os.Getenv("CONSUMER_NAME") // "events-router"
	preProcessingQueueName  = "pre-processing"
	preProcessingRoutingKey = "input.created.*"
)

func getRabbitMQResource(sd *servicediscovery.ServiceDiscovery) *gorabbitmq.Client {
	rabbitmq, err := sd.GetResource("rabbitmq")
	if err != nil {
		panic(err)
	}
	client, ok := rabbitmq.GetClient().(*gorabbitmq.Client)
	if !ok {
		panic("invalid RabbitMQ client type")
	}
	return client
}

func getRabbitMQNotifier(rmqClient *gorabbitmq.Client) *gorabbitmq.RabbitMQNotifier {
	return gorabbitmq.NewRabbitMQNotifier(rmqClient)
}

func main() {
	log.New(os.Stdout, "[EVENT-ROUTER] - ", log.LstdFlags)
	sd := servicediscovery.NewServiceDiscovery()
	db := inMemoryDB.NewInMemoryDocBD(dbName)
	dbClient := inMemoryDBClient.NewClient(db)
	eventOrderRepository := inMemoryDBRepository.NewEventOrderRepository(dbClient, dbName)

	rmq := getRabbitMQResource(sd)
	notifier := getRabbitMQNotifier(rmq)

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("ErrorCreated", &eventHandlers.ErrorCreatedHandler{
		Notifier: notifier,
	})
	eventDispatcher.Register("OrderedProcess", &eventHandlers.OrderedProcessHandler{
		Notifier: notifier,
	})

	errorEventHandler := event.NewErrorCreated()
	eventOrderEventHandler := event.NewOrderedProcess()

	eventOrderUsecase := usecase.NewPreProcessingUseCase(
		eventOrderRepository,
		errorEventHandler,
		eventOrderEventHandler,
		eventDispatcher,
	)

	listener := eventListener.NewEventListener()
	preProcessingConsumer := amqpConsumer.NewAmqpConsumer(rmq, preProcessingQueueName, consumerName, preProcessingRoutingKey)

	listener.AddListener(preProcessingConsumer, eventOrderUsecase)

	listenerServer := eventServer.NewListenerServer(listener)
	listenerServer.Start()
}
