//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/robojones/graphql/api"
	"github.com/robojones/graphql/prisma"
	"github.com/robojones/graphql/server"
)

func Initialize() (*server.Server, error) {
	wire.Build(
		prisma.Provider,
		api.Provider,
		server.Provider,
	)
	return &server.Server{}, nil
}
