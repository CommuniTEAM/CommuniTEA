package api

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

const internalErrMsg = "could not process request, please try again"

type PgxPoolIface interface {
	Acquire(ctx context.Context) (*pgxpool.Conn, error)
	Close()
	Config() *pgxpool.Config
}
