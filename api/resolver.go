package api

import (
	"github.com/robojones/graphql/api/query"
	"github.com/robojones/graphql/gqlgen"
	"github.com/robojones/graphql/prisma"
)

type Resolver struct {
	Prisma *prisma.Client
}

func (*Resolver) Mutation() gqlgen.MutationResolver {
	panic("implement me")
}

func (*Resolver) Query() gqlgen.QueryResolver {
	return &query.QueryResolver{}
}

