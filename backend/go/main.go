package main

import (
	"os"

	"github.com/ogoshikazuki/skill-sheet/config"
	"github.com/ogoshikazuki/skill-sheet/infrastructure/logger"
	"github.com/ogoshikazuki/skill-sheet/infrastructure/postgres"
	"github.com/ogoshikazuki/skill-sheet/infrastructure/server"
)

func main() {
	cfg := config.NewConfig()

	logger := logger.NewStandardLogger()

	sqlhandler, err := postgres.NewSqlHandler(
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDbname,
	)
	if err != nil {
		logger.Errorf("%+v", err)
		os.Exit(1)
	}

	server.NewServer(cfg, logger, sqlhandler).Start()
}
