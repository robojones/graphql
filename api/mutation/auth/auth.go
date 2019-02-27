package auth

import "github.com/robojones/graphql/api/super"

type Auth struct {
	*super.Resolver
}

func New(super *super.Resolver) *Auth {
	return &Auth{
		Resolver: super,
	}
}
