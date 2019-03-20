package project

import (
	"context"
	"github.com/robojones/graphql/gqlgen"
	"github.com/robojones/graphql/lib/session_context"
	"github.com/robojones/graphql/prisma"
)

func (resolver *Project) CreateProject(
	ctx context.Context,
	input gqlgen.CreateProjectInput,
) (*gqlgen.CreateProjectPayload, error) {
	user, err := session_context.User(ctx)

	if err != nil {
		return nil, err
	}

	project, err := resolver.Prisma.CreateProject(prisma.ProjectCreateInput{
		Title: input.Title,
		Members: &prisma.UserCreateManyWithoutProjectsInput{
			Connect: []prisma.UserWhereUniqueInput{{
				ID: &user.ID,
			}},
		},
	}).Exec(ctx)

	if err != nil {
		panic(err)
	}

	return &gqlgen.CreateProjectPayload{
		Project: *project,
	}, nil
}
