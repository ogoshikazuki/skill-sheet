package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ogoshikazuki/skill-sheet/adapter/repository"
	"github.com/ogoshikazuki/skill-sheet/config"
	"github.com/ogoshikazuki/skill-sheet/entity"
	"github.com/ogoshikazuki/skill-sheet/graph"
	"github.com/ogoshikazuki/skill-sheet/infrastructure/server/middleware"
	"github.com/ogoshikazuki/skill-sheet/usecase"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type Server struct {
	cfg        config.Config
	logger     *log.Logger
	sqlHandler repository.SqlHandler
	handler    http.Handler
}

func NewServer(cfg config.Config, logger *log.Logger, sqlHandler repository.SqlHandler) Server {
	return Server{
		cfg:        cfg,
		logger:     logger,
		sqlHandler: sqlHandler,
		handler:    http.DefaultServeMux,
	}
}

func (s Server) Start() {
	s.handleHealth()
	s.handleGraphQL()

	s.applyMiddleware()

	s.listen()
}

func (s *Server) handleHealth() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}

func (s *Server) handleGraphQL() {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolverRoot(
		usecase.NewFindBasicInformationUsecase(repository.NewBasicInformationRepository(s.sqlHandler)),
		usecase.NewFindProjectUsecase(repository.NewProjectRepository(s.sqlHandler)),
		usecase.NewSearchProjectsUsecase(repository.NewProjectRepository(s.sqlHandler)),
	)}))
	srv.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
		if errors.Is(err, entity.ErrInternal) {
			return graphql.DefaultErrorPresenter(ctx, errors.New("internal server error"))
		}

		return graphql.DefaultErrorPresenter(ctx, err)
	})
	http.Handle("/query", srv)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
}

func (s *Server) applyMiddleware() {
	middlewares := [](func(http.Handler) http.Handler){
		middleware.CORS(middleware.WithCORSAllowedOrigins(s.cfg.CORSAllowedOrigins)),
		middleware.Dataloader(s.sqlHandler),
	}

	for _, middleware := range middlewares {
		s.handler = middleware(s.handler)
	}
}

func (s *Server) listen() {
	s.logger.Printf("connect to http://localhost:%s/ for GraphQL playground", s.cfg.Port)
	s.logger.Fatal(http.ListenAndServe(":"+s.cfg.Port, s.handler))
}
