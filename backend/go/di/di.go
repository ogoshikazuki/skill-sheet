package di

import (
	"os"

	"github.com/ogoshikazuki/skill-sheet/adapter/repository"
	"github.com/ogoshikazuki/skill-sheet/infrastructure/postgres"
	"github.com/ogoshikazuki/skill-sheet/usecase"
)

type usecases struct {
	FindBasicInformationUsecase usecase.FindBasicInformationUsecase
}

var Usecases usecases

func Di() error {
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDbname := os.Getenv("POSTGRES_DBNAME")

	sqlhandler, err := postgres.NewSqlHandler(postgresHost, postgresPort, postgresUser, postgresPassword, postgresDbname)
	if err != nil {
		return err
	}

	basicInformationRepository := repository.NewBasicInformationRepository(sqlhandler)

	Usecases.FindBasicInformationUsecase = usecase.NewFindBasicInformationUsecase(basicInformationRepository)

	return nil
}
