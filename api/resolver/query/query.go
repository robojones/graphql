package query

import (
	"context"
	"github.com/robojones/graphql/lib/session_context"
	"github.com/robojones/graphql/prisma"
)

type Resolver struct {
	Prisma *prisma.Client
}

func New(client *prisma.Client) *Resolver {
	return &Resolver{
		Prisma: client,
	}
}

func (r *Resolver) User(ctx context.Context) (prisma.User, error) {
	user, err := session_context.User(ctx)

	if err != nil {
		return prisma.User{}, err
	}

	return *user, err
}
