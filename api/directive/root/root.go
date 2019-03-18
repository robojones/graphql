package root

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/robojones/graphql/gqlgen"
)

type Can = func(ctx context.Context, obj interface{}, next graphql.Resolver, operation gqlgen.Operation, scope *gqlgen.Scope) (res interface{}, err error)

func New(canDirective Can) *gqlgen.DirectiveRoot {
	return &gqlgen.DirectiveRoot{
		Can: canDirective,
	}
}
