//go:generate wire

package root

import (
	"github.com/robojones/graphql/api/resolver/mutation"
	"github.com/robojones/graphql/api/resolver/query"
	"github.com/robojones/graphql/gqlgen"
	"github.com/robojones/graphql/prisma"
)

type Resolver struct {
	Prisma           *prisma.Client
	MutationResolver *mutation.Resolver
	QueryResolver    *query.Resolver
}

func New(client *prisma.Client, mutation *mutation.Resolver, query *query.Resolver) *Resolver {
	return &Resolver{
		Prisma:           client,
		MutationResolver: mutation,
		QueryResolver:    query,
	}
}

func (r *Resolver) Mutation() gqlgen.MutationResolver {
	return r.MutationResolver
}

func (r *Resolver) Query() gqlgen.QueryResolver {
	return r.QueryResolver
}
