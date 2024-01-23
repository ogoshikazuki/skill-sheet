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
	s.handleGraphQL()
	s.handleGraphQLPlayground()

	s.listen()
}

func (s Server) handleGraphQL() {
	c := cors.New(cors.Options{
		AllowedOrigins:   s.cfg.CorsAllowdOrigins,
		AllowCredentials: true,
	})
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	srv.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
		s.logger.Printf("%+v", errors.Unwrap(err))
		return graphql.DefaultErrorPresenter(ctx, errors.New("internal server error"))
	})
	http.Handle("/query", c.Handler(srv))
}

func (s Server) handleGraphQLPlayground() {
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
}

func (s Server) listen() {
	s.logger.Printf("connect to http://localhost:%s/ for GraphQL playground", s.cfg.Port)
	s.logger.Fatal(http.ListenAndServe(":"+s.cfg.Port, nil))
}
