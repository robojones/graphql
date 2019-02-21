package query

import (
	"context"
	"github.com/robojones/graphql/api/super"
	"github.com/robojones/graphql/prisma"
)

type Resolver struct {
	*super.Resolver
}

func (r *Resolver) User(ctx context.Context) (prisma.User, error) {
	email := "alice@prisma.io"

	user, err := r.Prisma.User(prisma.UserWhereUniqueInput{
		Email: &email,
	}).Exec(ctx)

	return *user, err
}
