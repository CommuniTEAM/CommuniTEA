package api

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// internalErrMsg defines the response message for a 500 http response
// code. Used by endpoints in status.Wrap()
const internalErrMsg = "could not process request, please try again"

// genericInput defines input schema for endpoints that do not require
// a json body or query parameters. It supports protected endpoints.
type genericInput struct {
	Cookie string `cookie:"bearer-token,httponly,secure,samesite=strict,path=/,max-age:3600" json:"-"`
}

type PgxPoolIface interface {
	Acquire(ctx context.Context) (*pgxpool.Conn, error)
	Close()
	Config() *pgxpool.Config
}
