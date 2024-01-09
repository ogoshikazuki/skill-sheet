package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/ogoshikazuki/skill-sheet/di"
	"github.com/ogoshikazuki/skill-sheet/infrastructure/server"
)

func main() {
	logger := log.Default()

	if err := godotenv.Load(); err != nil {
		logger.Print(err)
	}

	if err := di.Di(); err != nil {
		logger.Fatal(err)
	}

	server.NewServer(logger).Start()
}
