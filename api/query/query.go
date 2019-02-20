package query

import (
	"context"
	"github.com/robojones/graphql/prisma"
)

type QueryResolver struct {

}

func (*QueryResolver) User(ctx context.Context) (prisma.User, error) {
	return prisma.User{
		Email: "hi",
	}, nil
}

