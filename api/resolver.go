package api

import (
	"github.com/robojones/graphql/api/mutation"
	"github.com/robojones/graphql/api/query"
	"github.com/robojones/graphql/api/super"
	"github.com/robojones/graphql/gqlgen"
	"github.com/robojones/graphql/prisma"
)

func New(client *prisma.Client) *Resolver {
	return &Resolver{
		Resolver: &super.Resolver{
			Prisma: client,
		},
	}
}

type Resolver struct {
	*super.Resolver
}

func (r *Resolver) Mutation() gqlgen.MutationResolver {
	return mutation.New(r.Resolver)
}

func (r *Resolver) Query() gqlgen.QueryResolver {
	return query.New(r.Resolver)
}
