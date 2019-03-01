package auth

import (
	"github.com/robojones/graphql/prisma"
)

type Auth struct {
	Prisma *prisma.Client
}

func New(client *prisma.Client) *Auth {
	return &Auth{
		Prisma: client,
	}
}
