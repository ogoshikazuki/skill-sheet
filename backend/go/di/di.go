package di

import (
	"github.com/ogoshikazuki/skill-sheet/adapter/repository"
	"github.com/ogoshikazuki/skill-sheet/config"
	"github.com/ogoshikazuki/skill-sheet/entity"
	"github.com/ogoshikazuki/skill-sheet/infrastructure/postgres"
	"github.com/ogoshikazuki/skill-sheet/usecase"
)

type usecases struct {
	FindBasicInformationUsecase usecase.FindBasicInformationUsecase
	FindProjectUsecase          usecase.FindProjectUsecase
	SearchProjectsUsecase       usecase.SearchProjectsUsecase
}

var Usecases usecases

type repositories struct {
	TechnologyRepository entity.TechnologyRepository
}

var Repositories repositories

func Di(cfg config.Config) error {
	sqlhandler, err := postgres.NewSqlHandler(
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDbname,
	)
	if err != nil {
		return err
	}

	basicInformationRepository := repository.NewBasicInformationRepository(sqlhandler)
	projectRepository := repository.NewProjectRepository(sqlhandler)
	Repositories.TechnologyRepository = repository.NewTechnologyRepository(sqlhandler)

	Usecases.FindBasicInformationUsecase = usecase.NewFindBasicInformationUsecase(basicInformationRepository)
	Usecases.FindProjectUsecase = usecase.NewFindProjectUsecase(projectRepository)
	Usecases.SearchProjectsUsecase = usecase.NewSearchProjectsUsecase(projectRepository)

	return nil
}
