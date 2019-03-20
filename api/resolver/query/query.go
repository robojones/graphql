package query

import (
	"github.com/robojones/graphql/prisma"
)

type Query struct {
	Prisma *prisma.Client
}

func New(client *prisma.Client) *Query {
	return &Query{
		Prisma: client,
	}
}
