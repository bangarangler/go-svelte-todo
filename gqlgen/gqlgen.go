package gqlgen

import (
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/bangarangler/go-svelte-todo/pg"
)

func NewHandler(repo pg.Repository) http.Handler {
	return handler.GraphQL(NewExecutableSchema(Config{
		Resolvers: &Resolver{
			Repository: repo,
		},
	}))
}

func NewPlaygroundHandler(endpoint string) http.Handler {
	return handler.Playground("GraphQL Playground", endpoint)
}
