package api

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/swaggest/usecase/status"
)

// API is the foundation for the api endpoint functions that

// supplies access to the database.

type API struct {
	DBPool PgxPoolIface
}

// defaultInput defines input schema for endpoints that do not require

// a json body or query parameters. Any input schema that requires

// authentication should embed this struct.

type defaultInput struct {
	AccessToken string `cookie:"bearer-token" json:"-"`
}

// uuidInput defines input schema for endpoints that take a uuid query

// parameter. Also supports protected endpoints.

type uuidInput struct {
	defaultInput

	ID uuid.UUID `nullable:"false" path:"id"`
}

// genericOutput defines output schema for endpoints that do not return

// any data in a json body, such as DELETE methods.

type genericOutput struct {
	Message string `json:"message"`
}

type PgxPoolIface interface {
	Acquire(ctx context.Context) (*pgxpool.Conn, error)

	Close()

	Config() *pgxpool.Config
}

const (

	// internalErrMsg defines the response message for a 500 http response

	// code. Used by endpoints in status.Wrap()

	internalErrMsg string = "could not process request, please try again"

	adminRole string = "admin"
)

// dbConn is a helper function that establishes a database connection from

// the API's database pool or, if that fails, returns a pre-formatted error

// with a 500 http response code.

func (a *API) dbConn(ctx context.Context) (*pgxpool.Conn, error) {
	conn, err := a.DBPool.Acquire(ctx)

	if err != nil {
		log.Println(fmt.Errorf("could not acquire db connection: %w", err))

		return nil, status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
	}

	return conn, nil
}
