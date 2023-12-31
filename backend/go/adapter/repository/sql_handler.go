package repository

import "context"

type (
	Sqlhandler interface {
		ExecContext(ctx context.Context, query string, args ...any) (Result, error)
		QueryContext(ctx context.Context, query string, args ...any) (Rows, error)
		Close() error
	}

	Result interface {
		LastInsertId() (int64, error)
		RowsAffected() (int64, error)
	}

	Rows interface {
		Scan(dest any) error
		Next() bool
		Close() error
	}
)
