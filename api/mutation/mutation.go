package mutation

import (
	"github.com/robojones/graphql/api/mutation/auth"
	"github.com/robojones/graphql/api/super"
)

type Mutation struct {
	*super.Resolver
	*auth.Auth
}

func New(super *super.Resolver) *Mutation {
	return &Mutation{
		Resolver: super,
		Auth:     auth.New(super),
	}
}
