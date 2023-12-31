package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/cockroachdb/errors"
	"github.com/ogoshikazuki/skill-sheet/adapter/repository"

	_ "github.com/lib/pq"
)

type sqlHandler struct {
	db *sql.DB
}

func (handler *sqlHandler) ExecContext(ctx context.Context, query string, args ...any) (repository.Result, error) {
	r, err := handler.db.ExecContext(ctx, query, args)
	if err != nil {
		return nil, err
	}

	return &result{r}, nil
}

func (handler *sqlHandler) QueryContext(ctx context.Context, query string, args ...any) (repository.Rows, error) {
	r, err := handler.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &rows{rows: r}, nil
}

func (handler *sqlHandler) Close() error {
	return handler.db.Close()
}

func NewSqlHandler() (repository.Sqlhandler, error) {
	db, err := sql.Open("postgres", getConnectionString())
	if err != nil {
		return nil, err
	}

	return &sqlHandler{db: db}, nil
}

type result struct {
	result sql.Result
}

func (r *result) LastInsertId() (int64, error) {
	return r.LastInsertId()
}

func (r *result) RowsAffected() (int64, error) {
	return r.RowsAffected()
}

type rows struct {
	rows *sql.Rows
}

func (r *rows) Scan(dest any) error {
	err := r.rows.Scan(dest)

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *rows) Next() bool {
	return r.rows.Next()
}

func (r *rows) Close() error {
	return r.rows.Close()
}

func getConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_DBNAME"),
	)
}
