package query

import (
	"context"
	"github.com/robojones/graphql/lib/session_context"
	"github.com/robojones/graphql/prisma"
)

func (r *Query) User(ctx context.Context) (*prisma.User, error) {
	return session_context.User(ctx)
}
