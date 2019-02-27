package query

import (
	"context"
	"github.com/robojones/graphql/api/super"
	"github.com/robojones/graphql/lib/session_context"
	"github.com/robojones/graphql/prisma"
)

type Resolver struct {
	*super.Resolver
}

func New(super *super.Resolver) *Resolver {
	return &Resolver{
		Resolver: super,
	}
}

func (r *Resolver) User(ctx context.Context) (prisma.User, error) {
	user, err := session_context.User(ctx)

	if err != nil {
		return prisma.User{}, err
	}

	return *user, err
}
