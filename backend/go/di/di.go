package di

import (
	"github.com/ogoshikazuki/skill-sheet/adapter/repository"
	"github.com/ogoshikazuki/skill-sheet/infrastructure/postgres"
	"github.com/ogoshikazuki/skill-sheet/usecase"
)

type usecases struct {
	FindBasicInformationUsecase usecase.FindBasicInformationUsecase
}

var Usecases usecases

func Di() error {
	sqlhandler, err := postgres.NewSqlHandler()
	if err != nil {
		return err
	}

	basicInformationRepository := repository.NewBasicInformationRepository(sqlhandler)

	Usecases.FindBasicInformationUsecase = usecase.NewFindBasicInformationUsecase(basicInformationRepository)

	return nil
}
