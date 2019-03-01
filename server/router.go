package server

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/robojones/graphql/api"
	"net/http"
)

func NewServeMux(api *api.Handler) *http.ServeMux {
	mux := &http.ServeMux{}

	mux.Handle("/", handler.Playground("GraphQL Playground", "/query"))
	mux.Handle("/query", api)

	return mux
}
