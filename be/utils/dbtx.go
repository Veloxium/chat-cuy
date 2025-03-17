package utils

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(c context.Context, query string, args ...any) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...any) *sql.Row
}
