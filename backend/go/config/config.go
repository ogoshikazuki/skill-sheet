package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

type Config struct {
	Port               string
	CORSAllowedOrigins []string
	PostgresHost       string
	PostgresPort       string
	PostgresUser       string
	PostgresPassword   string
	PostgresDbname     string
}

func NewConfig() Config {
	return Config{
		Port:               os.Getenv("PORT"),
		CORSAllowedOrigins: strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ","),
		PostgresHost:       os.Getenv("POSTGRES_HOST"),
		PostgresPort:       os.Getenv("POSTGRES_PORT"),
		PostgresUser:       os.Getenv("POSTGRES_USER"),
		PostgresPassword:   os.Getenv("POSTGRES_PASSWORD"),
		PostgresDbname:     os.Getenv("POSTGRES_DBNAME"),
	}
}
