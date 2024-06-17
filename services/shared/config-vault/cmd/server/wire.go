// go:build wireinject
//go:build wireinject
// +build wireinject

package main

import (
	webHandler "libs/golang/ddd/adapters/http/handlers/config-vault/handlers"
	"libs/golang/ddd/domain/entities/config-vault/entity"
	"libs/golang/ddd/domain/repositories/database/mongodb/config-vault/repository"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

var setConfigRepositoryDependency = wire.NewSet(
	repository.NewConfigRepository,
	wire.Bind(
		new(entity.ConfigRepositoryInterface),
		new(*repository.ConfigRepository),
	),
)

func NewWebServiceConfigHandler(client *mongo.Client, database string) *webHandler.WebConfigHandler {
	wire.Build(
		setConfigRepositoryDependency,
		webHandler.NewWebConfigHandler,
	)
	return &webHandler.WebConfigHandler{}
}
