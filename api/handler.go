package api

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/robojones/graphql/api/resolver/root"
	"github.com/robojones/graphql/gqlgen"
	"github.com/robojones/graphql/lib/auth"
	"github.com/robojones/graphql/lib/handler_adapter"
	"github.com/robojones/graphql/prisma"
	"net/http"
)

type Handler struct {
	Next http.Handler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Next.ServeHTTP(w, r)
}

func New(client *prisma.Client, resolver *root.Resolver) *Handler {
	schema := gqlgen.NewExecutableSchema(gqlgen.Config{Resolvers: resolver})

	return &Handler{
		Next: &auth.Handler{
			Prisma: client,
			Next: &handler_adapter.HandlerFuncAdapter{
				NextFunc: handler.GraphQL(schema),
			},
		},
	}
}
