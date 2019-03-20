package mutation

import (
	"github.com/robojones/graphql/api/resolver/mutation/auth"
	"github.com/robojones/graphql/api/resolver/mutation/project"
	"github.com/robojones/graphql/prisma"
)

type Mutation struct {
	Prisma *prisma.Client
	*auth.Auth
	*project.Project
}

func New(
	client *prisma.Client,
	auth *auth.Auth,
	project *project.Project,
) *Mutation {
	return &Mutation{
		Prisma:  client,
		Auth:    auth,
		Project: project,
	}
}
