package query

import (
	"context"
	"github.com/robojones/graphql/lib/session_context"
	"github.com/robojones/graphql/prisma"
)

func (resolver *Query) Projects(ctx context.Context) ([]prisma.Project, error) {
	user, err := session_context.User(ctx)

	if err != nil {
		return nil, err
	}

	return resolver.Prisma.User(prisma.UserWhereUniqueInput{
		ID: &user.ID,
	}).Projects(nil).Exec(ctx)
}
