package query

import (
	"context"
	"github.com/robojones/graphql/lib/session_context"
	"github.com/robojones/graphql/prisma"
)

func (resolver *Query) Session(ctx context.Context) (*prisma.Session, error) {
	token, err := session_context.Token(ctx)

	if err != nil {
		return nil, err
	}

	return resolver.Prisma.Session(prisma.SessionWhereUniqueInput{
		Token: &token,
	}).Exec(ctx)
}
