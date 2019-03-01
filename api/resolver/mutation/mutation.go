package mutation

import (
	"github.com/robojones/graphql/api/resolver/mutation/auth"
	"github.com/robojones/graphql/prisma"
)

type Resolver struct {
	Prisma *prisma.Client
	*auth.Auth
}

func New(client *prisma.Client, auth *auth.Auth) *Resolver {
	return &Resolver{
		Prisma: client,
		Auth:   auth,
	}
}
