package query

import (
	"context"
	"github.com/robojones/graphql/api/super"
	"github.com/robojones/graphql/prisma"
)

type Resolver struct {
	*super.Resolver
}

func (*Resolver) User(ctx context.Context) (prisma.User, error) {
	return prisma.User{
		Email: "hi",
	}, nil
}

