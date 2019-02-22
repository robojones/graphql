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

func (r *Resolver) User(ctx context.Context) (prisma.User, error) {
	user, err := session_context.User(ctx)

	// needed because the return type cannot be nil
	if err != nil {
		return prisma.User{}, err
	}

	return *user, err
}
