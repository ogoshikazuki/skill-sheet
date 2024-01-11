package di

import (
	"github.com/ogoshikazuki/skill-sheet/adapter/repository"
	"github.com/ogoshikazuki/skill-sheet/config"
	"github.com/ogoshikazuki/skill-sheet/infrastructure/postgres"
	"github.com/ogoshikazuki/skill-sheet/usecase"
)

type usecases struct {
	FindBasicInformationUsecase usecase.FindBasicInformationUsecase
}

var Usecases usecases

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

	Usecases.FindBasicInformationUsecase = usecase.NewFindBasicInformationUsecase(basicInformationRepository)

	return nil
}
