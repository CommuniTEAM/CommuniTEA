package api

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// internalErrMsg defines the response message for a 500 http response
// code. Used by endpoints in status.Wrap()
const internalErrMsg = "could not process request, please try again"

// genericInput defines input schema for endpoints that do not require
// a json body or query parameters. Any input schema that requires
// authentication should embed this struct.
type genericInput struct {
	AccessToken string `cookie:"bearer-token" json:"-"`
}

// uuidInput defines input schema for endpoints that take a uuid query
// parameter. Also supports protected endpoints.
type uuidInput struct {
	genericInput
	ID uuid.UUID `nullable:"false" path:"id"`
}

type PgxPoolIface interface {
	Acquire(ctx context.Context) (*pgxpool.Conn, error)
	Close()
	Config() *pgxpool.Config
}
