package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Port              string
	CorsAllowdOrigins []string
	PostgresHost      string
	PostgresPort      string
	PostgresUser      string
	PostgresPassword  string
	PostgresDbname    string
}

func NewConfig() Config {
	godotenv.Load()

	return Config{
		Port:              os.Getenv("PORT"),
		CorsAllowdOrigins: strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ","),
		PostgresHost:      os.Getenv("POSTGRES_HOST"),
		PostgresPort:      os.Getenv("POSTGRES_PORT"),
		PostgresUser:      os.Getenv("POSTGRES_USER"),
		PostgresPassword:  os.Getenv("POSTGRES_PASSWORD"),
		PostgresDbname:    os.Getenv("POSTGRES_DBNAME"),
	}
}
