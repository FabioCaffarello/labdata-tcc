// go:build wireinject
//go:build wireinject
// +build wireinject

package main

import (
	webHandler "libs/golang/ddd/adapters/http/handlers/output-vault/handlers"
	"libs/golang/ddd/domain/entities/output-vault/entity"
	"libs/golang/ddd/domain/repositories/database/mongodb/output-vault/repository"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

var setOutputRepositoryDependency = wire.NewSet(
	repository.NewOutputRepository,
	wire.Bind(
		new(entity.OutputRepositoryInterface),
		new(*repository.OutputRepository),
	),
)

func NewWebServiceOutputHandler(client *mongo.Client, database string) *webHandler.WebOutputHandler {
	wire.Build(
		setOutputRepositoryDependency,
		webHandler.NewWebOutputHandler,
	)
	return &webHandler.WebOutputHandler{}
}
