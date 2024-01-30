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
	"github.com/ogoshikazuki/skill-sheet/infrastructure/server/middleware"
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
	s.handleHealth()
	s.handleGraphQL()

	s.listen()
}

func (s Server) handleHealth() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}

func (s Server) handleGraphQL() {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	srv.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
		s.logger.Printf("%+v", errors.Unwrap(err))
		return graphql.DefaultErrorPresenter(ctx, errors.New("internal server error"))
	})
	http.Handle("/query", srv)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
}

func (s Server) listen() {
	handler := middleware.CORS(middleware.WithCORSAllowedOrigins(s.cfg.CORSAllowedOrigins))(http.DefaultServeMux)

	s.logger.Printf("connect to http://localhost:%s/ for GraphQL playground", s.cfg.Port)
	s.logger.Fatal(http.ListenAndServe(":"+s.cfg.Port, handler))
}
