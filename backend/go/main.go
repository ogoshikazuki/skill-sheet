package main

import (
	"log"

	"github.com/ogoshikazuki/skill-sheet/config"
	"github.com/ogoshikazuki/skill-sheet/di"
	"github.com/ogoshikazuki/skill-sheet/infrastructure/server"
)

func main() {
	cfg := config.NewConfig()

	logger := log.Default()

	if err := di.Di(cfg); err != nil {
		logger.Fatal(err)
	}

	server.NewServer(cfg, logger).Start()
}
