// go:build wireinject
//go:build wireinject
// +build wireinject

package main

import (
	webHandler "libs/golang/ddd/adapters/http/handlers/schema-vault/handlers"
	"libs/golang/ddd/domain/entities/schema-vault/entity"
	"libs/golang/ddd/domain/repositories/database/mongodb/schema-vault/repository"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

var setSchemaRepositoryDependency = wire.NewSet(
	repository.NewSchemaRepository,
	wire.Bind(
		new(entity.SchemaRepositoryInterface),
		new(*repository.SchemaRepository),
	),
)

func NewWebServiceSchemaHandler(client *mongo.Client, database string) *webHandler.WebSchemaHandler {
	wire.Build(
		setSchemaRepositoryDependency,
		webHandler.NewWebSchemaHandler,
	)
	return &webHandler.WebSchemaHandler{}
}
