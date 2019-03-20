package project

import (
	"context"
	"github.com/robojones/graphql/gqlgen"
	"github.com/robojones/graphql/prisma"
)

func (resolver *Project) UpdateProject(
	ctx context.Context,
	input gqlgen.UpdateProjectInput,
) (*gqlgen.UpdateProjectPayload, error) {
	project, err := resolver.Prisma.UpdateProject(prisma.ProjectUpdateParams{
		Where: prisma.ProjectWhereUniqueInput{
			ID: &input.ID,
		},
		Data: prisma.ProjectUpdateInput{
			Title: input.Title,
		},
	}).Exec(ctx)

	if err != nil {
		panic(err)
	}

	return &gqlgen.UpdateProjectPayload{
		Project: *project,
	}, nil
}
