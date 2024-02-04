package repository_test

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/ogoshikazuki/skill-sheet/infrastructure/postgres"
	"github.com/ogoshikazuki/skill-sheet/repository"
)

var postgresHost, postgresPort, postgresUser, postgresPassword, postgresDbname string

func init() {
	postgresHost = "localhost"
	postgresPort = "5432"
	postgresUser = "postgres"
	postgresPassword = "postgres"
	postgresDbname = "postgres"
}

var testDbname string = "test"

var (
	sqlHandler    repository.SqlHandler
	errSqlHandler repository.SqlHandler
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	sqlhandlerToManageTestDb := newSqlhandlerToManageTestDb()
	createTestDbAndSetSqlHandler(ctx, sqlhandlerToManageTestDb)
	defer dropTestDb(ctx, sqlhandlerToManageTestDb)
	initTestDb(ctx)
	setErrSqlHandler()

	m.Run()
}

func newSqlhandlerToManageTestDb() repository.SqlHandler {
	sqlhandlerToManageTestDb, err := postgres.NewSqlHandler(
		postgresHost,
		postgresPort,
		postgresUser,
		postgresPassword,
		postgresDbname,
	)
	if err != nil {
		panic(err)
	}

	return sqlhandlerToManageTestDb
}

func createTestDbAndSetSqlHandler(ctx context.Context, sqlhandlerToManageTestDb repository.SqlHandler) {
	if _, err := sqlhandlerToManageTestDb.ExecContext(ctx, fmt.Sprintf("DROP DATABASE IF EXISTS %s", testDbname)); err != nil {
		panic(err)
	}
	if _, err := sqlhandlerToManageTestDb.ExecContext(ctx, fmt.Sprintf("CREATE DATABASE %s", testDbname)); err != nil {
		panic(err)
	}

	var err error
	sqlHandler, err = postgres.NewSqlHandler(
		postgresHost,
		postgresPort,
		postgresUser,
		postgresPassword,
		testDbname,
	)
	if err != nil {
		panic(err)
	}
}

func initTestDb(ctx context.Context) {
	if _, err := sqlHandler.ExecContextFromFile(ctx, filepath.Join("..", "..", "..", "postgres", "init")); err != nil {
		panic(err)
	}
}

func setErrSqlHandler() {
	var err error
	errSqlHandler, err = postgres.NewSqlHandler(
		postgresHost,
		postgresPort,
		"",
		"",
		"",
	)
	if err != nil {
		panic(err)
	}
}

func dropTestDb(ctx context.Context, sqlHandlerToManageTestDb repository.SqlHandler) {
	sqlHandler.Close()

	if _, err := sqlHandlerToManageTestDb.ExecContext(ctx, fmt.Sprintf("DROP DATABASE IF EXISTS %s", testDbname)); err != nil {
		panic(err)
	}
}
