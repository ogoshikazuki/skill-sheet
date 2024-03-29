package server

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ogoshikazuki/skill-sheet/config"
	"github.com/ogoshikazuki/skill-sheet/entity"
	"github.com/ogoshikazuki/skill-sheet/graph"
	"github.com/ogoshikazuki/skill-sheet/infrastructure/server/middleware"
	"github.com/ogoshikazuki/skill-sheet/repository"
	"github.com/ogoshikazuki/skill-sheet/usecase"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type Server struct {
	cfg        config.Config
	logger     entity.Logger
	sqlHandler repository.SqlHandler
	handler    http.Handler
}

func NewServer(cfg config.Config, logger entity.Logger, sqlHandler repository.SqlHandler) Server {
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
	c := graph.Config{Resolvers: graph.NewResolverRoot(
		usecase.NewFindBasicInformationUsecase(repository.NewBasicInformationRepository(s.sqlHandler)),
		usecase.NewFindProjectUsecase(repository.NewProjectRepository(s.sqlHandler)),
		usecase.NewSearchProjectsUsecase(repository.NewProjectRepository(s.sqlHandler)),
		usecase.NewUpdateBasicInformationUsecase(repository.NewBasicInformationRepository(s.sqlHandler), repository.NewTransactionController(s.sqlHandler)),
	)}
	c.Directives.Admin = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		if !middleware.IsAuthenticated(ctx) {
			return nil, entity.ErrUnauthenticated
		}
		if !middleware.HasScope(ctx, "admin") {
			return nil, entity.ErrUnauthorized
		}

		return next(ctx)
	}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))
	srv.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
		var internalServerError *entity.InternalServerError
		if errors.As(err, &internalServerError) {
			s.logger.Errorf("%+v", internalServerError.ErrWithStack())
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
		middleware.Auth0(s.cfg.Auth0Domain, s.cfg.Auth0Audience),
	}

	for _, middleware := range middlewares {
		s.handler = middleware(s.handler)
	}
}

func (s *Server) listen() {
	s.logger.Errorf("connect to http://localhost:%s/ for GraphQL playground", s.cfg.Port)
	s.logger.Errorf("%+v", http.ListenAndServe(":"+s.cfg.Port, s.handler))
	os.Exit(1)
}
