package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ogoshikazuki/skill-sheet/entity"
	"github.com/ogoshikazuki/skill-sheet/repository"
	"github.com/tanimutomo/sqlfile"

	_ "github.com/lib/pq"
)

type sqlHandler struct {
	db *sql.DB
}

func (handler *sqlHandler) ExecContext(ctx context.Context, query string, args ...any) (repository.Result, error) {
	r, err := handler.db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return &result{r}, nil
}

func (handler *sqlHandler) QueryContext(ctx context.Context, query string, args ...any) (repository.Rows, error) {
	r, err := handler.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, entity.NewInternalServerError(err)
	}

	return &rows{rows: r}, nil
}

func (handler *sqlHandler) Close() error {
	return handler.db.Close()
}

func (s *sqlHandler) ExecFromFile(path string) ([]repository.Result, error) {
	sf := sqlfile.New()

	if err := sf.Directory(path); err != nil {
		return nil, entity.NewInternalServerError(err)
	}

	res, err := sf.Exec(s.db)
	if err != nil {
		return nil, entity.NewInternalServerError(err)
	}

	var results []repository.Result
	for _, r := range res {
		results = append(results, &result{r})
	}

	return results, nil
}

func NewSqlHandler(host, port, user, password, dbname string) (repository.SqlHandler, error) {
	db, err := sql.Open("postgres", getConnectionString(host, port, user, password, dbname))
	if err != nil {
		return nil, err
	}

	return &sqlHandler{db: db}, nil
}

type result struct {
	result sql.Result
}

func (r *result) LastInsertId() (int64, error) {
	return r.result.LastInsertId()
}

func (r *result) RowsAffected() (int64, error) {
	return r.result.RowsAffected()
}

type rows struct {
	rows *sql.Rows
}

func (r *rows) Scan(dest ...any) error {
	err := r.rows.Scan(dest...)

	if err != nil {
		return entity.NewInternalServerError(err)
	}

	return nil
}

func (r *rows) Next() bool {
	return r.rows.Next()
}

func (r *rows) Close() error {
	return r.rows.Close()
}

func getConnectionString(host, port, user, password, dbname string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		password,
		host,
		port,
		dbname,
	)
}
