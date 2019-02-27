package mutation

import (
	"github.com/robojones/graphql/api/mutation/auth"
	"github.com/robojones/graphql/api/super"
	"github.com/robojones/graphql/gqlgen"
)

type Mutation struct {
	*super.Resolver
	*auth.Auth
}

func New(super *super.Resolver) gqlgen.MutationResolver {
	return &Mutation{
		Resolver: super,
		Auth:     auth.New(super),
	}
}
