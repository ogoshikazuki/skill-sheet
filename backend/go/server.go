package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/cockroachdb/errors"
	"github.com/joho/godotenv"
	"github.com/ogoshikazuki/skill-sheet/di"
	"github.com/ogoshikazuki/skill-sheet/graph"
	"github.com/rs/cors"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const defaultPort = "8080"

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if err := di.Di(); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ","),
		AllowCredentials: true,
	})
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	srv.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
		log.Printf("%+v", errors.Unwrap(err))
		return graphql.DefaultErrorPresenter(ctx, err)
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", c.Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
