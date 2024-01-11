package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ogoshikazuki/skill-sheet/config"
	"github.com/ogoshikazuki/skill-sheet/graph"
	"github.com/rs/cors"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type Server struct {
	cfg    config.Config
	logger *log.Logger
}

func NewServer(cfg config.Config, logger *log.Logger) Server {
	return Server{
		cfg:    cfg,
		logger: logger,
	}
}

func (s Server) Start() {
	c := cors.New(cors.Options{
		AllowedOrigins:   s.cfg.CorsAllowdOrigins,
		AllowCredentials: true,
	})
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	srv.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
		s.logger.Printf("%+v", errors.Unwrap(err))
		return graphql.DefaultErrorPresenter(ctx, err)
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", c.Handler(srv))

	s.logger.Printf("connect to http://localhost:%s/ for GraphQL playground", s.cfg.Port)
	s.logger.Fatal(http.ListenAndServe(":"+s.cfg.Port, nil))
}
