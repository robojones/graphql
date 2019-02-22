package mutation

import (
	"context"
	"github.com/robojones/graphql/api/super"
	"github.com/robojones/graphql/prisma"
)

type Mutation struct {
	*super.Resolver
}

func (*Mutation) Signup(ctx context.Context, email string, name string) (prisma.User, error) {
	panic("implement me")
}
