// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"libs/golang/ddd/adapters/http/handlers/output-vault/handlers"
	"libs/golang/ddd/domain/entities/output-vault/entity"
	"libs/golang/ddd/domain/repositories/database/mongodb/output-vault/repository"
)

// Injectors from wire.go:

func NewWebServiceOutputHandler(client *mongo.Client, database string) *handlers.WebOutputHandler {
	outputRepository := repository.NewOutputRepository(client, database)
	webOutputHandler := handlers.NewWebOutputHandler(outputRepository)
	return webOutputHandler
}

// wire.go:

var setOutputRepositoryDependency = wire.NewSet(repository.NewOutputRepository, wire.Bind(
	new(entity.OutputRepositoryInterface),
	new(*repository.OutputRepository),
),
)
