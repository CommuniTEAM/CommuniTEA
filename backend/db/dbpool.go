package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
)

// newDbPool establishes a concurrency-safe database
// connection pool and returns its reference
func NewDBPool(connString string) (*pgxpool.Pool, error) {
	pgxConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("could not parse db conn string: %w", err)
	}

	pgxConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxUUID.Register(conn.TypeMap())
		return nil
	}

	pgxDBPool, err := pgxpool.NewWithConfig(context.TODO(), pgxConfig)
	if err != nil {
		return nil, fmt.Errorf("could not create new db pool: %w", err)
	}

	return pgxDBPool, nil
}
