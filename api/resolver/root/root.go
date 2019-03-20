//go:generate wire

package root

import (
	"github.com/robojones/graphql/api/resolver/mutation"
	"github.com/robojones/graphql/api/resolver/project"
	"github.com/robojones/graphql/api/resolver/query"
	"github.com/robojones/graphql/gqlgen"
)

type Root struct {
	MutationResolver *mutation.Mutation
	ProjectResolver  *project.Project
	QueryResolver    *query.Query
}

func New(mutation *mutation.Mutation, query *query.Query, project *project.Project) *Root {
	return &Root{
		ProjectResolver:  project,
		MutationResolver: mutation,
		QueryResolver:    query,
	}
}

func (r *Root) Mutation() gqlgen.MutationResolver {
	return r.MutationResolver
}

func (r *Root) Project() gqlgen.ProjectResolver {
	return r.ProjectResolver
}

func (r *Root) Query() gqlgen.QueryResolver {
	return r.QueryResolver
}
