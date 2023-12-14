package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// newDbPool establishes a concurrency-safe database

// connection pool and returns its reference

func NewDBPool(connString string) (*pgxpool.Pool, error) {

	pgxConfig, err := pgxpool.ParseConfig(connString)

	if err != nil {

		return nil, fmt.Errorf("could not parse db conn string: %w", err)

	}

	pgxDBPool, err := pgxpool.NewWithConfig(context.TODO(), pgxConfig)

	if err != nil {

		return nil, fmt.Errorf("could not create new db pool: %w", err)

	}

	return pgxDBPool, nil

}
