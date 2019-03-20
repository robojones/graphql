package project

import (
	"context"
	"github.com/robojones/graphql/prisma"
)

type Project struct {
	Prisma *prisma.Client

}

func (resolver *Project) Members(ctx context.Context, project *prisma.Project) ([]prisma.User, error) {
	return resolver.Prisma.Project(prisma.ProjectWhereUniqueInput{
		ID: &project.ID,
	}).Members(nil).Exec(ctx)
}

func New(prisma *prisma.Client) *Project {
	return &Project{
		Prisma: prisma,
	}
}
