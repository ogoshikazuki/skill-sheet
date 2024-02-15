package repository

import "context"

type (
	SqlHandler interface {
		ExecContext(ctx context.Context, query string, args ...any) (Result, error)
		QueryContext(ctx context.Context, query string, args ...any) (Rows, error)
		Close() error
		BeginTx(ctx context.Context) (Tx, error)
		ExecFromFile(path string) ([]Result, error)
	}

	Result interface {
		LastInsertId() (int64, error)
		RowsAffected() (int64, error)
	}

	Rows interface {
		Scan(...any) error
		Next() bool
		Close() error
	}

	Tx interface {
		ExecContext(ctx context.Context, query string, args ...any) (Result, error)
		QueryContext(ctx context.Context, query string, args ...any) (Rows, error)
		Commit() error
		Rollback() error
	}
)
