package can

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/robojones/graphql/api/directive/root"
	"github.com/robojones/graphql/gqlgen"
	"github.com/robojones/graphql/prisma"
)



func New(prisma *prisma.Client) root.Can {
	return func (ctx context.Context, obj interface{}, next graphql.Resolver, operation gqlgen.Operation, scope *gqlgen.Scope) (res interface{}, err error) {
		// TODO: write stuff
		return next(ctx)
	}
}
