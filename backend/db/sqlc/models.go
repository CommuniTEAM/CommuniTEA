// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type LocationsCity struct {
	ID    pgtype.UUID
	Name  string
	State string
}

type LocationsState struct {
	Name         string
	Abbreviation string
}

type User struct {
	ID   int32
	Name string
}
