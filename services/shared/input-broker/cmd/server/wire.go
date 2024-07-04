// go:build wireinject
//go:build wireinject
// +build wireinject

package main

import (
	webHandler "libs/golang/ddd/adapters/http/handlers/input-broker/handlers"
	"libs/golang/ddd/domain/entities/input-broker/entity"
	"libs/golang/ddd/domain/repositories/database/mongodb/input-broker/repository"
	event "libs/golang/ddd/events/input-broker/event"
	events "libs/golang/shared/go-events/amqp_events"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

var setInputRepositoryDependency = wire.NewSet(
	repository.NewInputRepository,
	wire.Bind(
		new(entity.InputRepositoryInterface),
		new(*repository.InputRepository),
	),
)

var setInputCreatedEvent = wire.NewSet(
	event.NewInputCreated,
	wire.Bind(new(events.EventInterface), new(*event.InputCreated)),
)

func NewWebServiceInputHandler(client *mongo.Client, eventDispatcher events.EventDispatcherInterface, database string) *webHandler.WebInputHandler {
	wire.Build(
		setInputRepositoryDependency,
		setInputCreatedEvent,
		webHandler.NewWebInputHandler,
	)
	return &webHandler.WebInputHandler{}
}
