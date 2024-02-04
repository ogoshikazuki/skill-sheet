package main

import (
	"log"

	"github.com/ogoshikazuki/skill-sheet/config"
	"github.com/ogoshikazuki/skill-sheet/infrastructure/postgres"
	"github.com/ogoshikazuki/skill-sheet/infrastructure/server"
)

func main() {
	cfg := config.NewConfig()

	logger := log.Default()

	sqlhandler, err := postgres.NewSqlHandler(
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDbname,
	)
	if err != nil {
		logger.Fatal(err)
	}

	server.NewServer(cfg, logger, sqlhandler).Start()
}
