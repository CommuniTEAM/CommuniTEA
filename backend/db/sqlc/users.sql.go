// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: users.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
insert into users (
    "id",
    "role",
    "username",
    "first_name",
    "last_name",
    "email",
    "password",
    "location"
)
values ($1, $2, $3, $4, $5, $6, $7, $8)
returning id, role, username, first_name, last_name, email, password, location
`

type CreateUserParams struct {
	ID        uuid.UUID   `json:"id"`
	Role      string      `json:"role"`
	Username  string      `json:"username"`
	FirstName pgtype.Text `json:"first_name"`
	LastName  pgtype.Text `json:"last_name"`
	Email     pgtype.Text `json:"email"`
	Password  []byte      `json:"password"`
	Location  uuid.UUID   `json:"location"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ID,
		arg.Role,
		arg.Username,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Password,
		arg.Location,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Role,
		&i.Username,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.Location,
	)
	return i, err
}

const login = `-- name: Login :one
select
    "id",
    "role",
    "username",
    "first_name",
    "last_name",
    "location",
    "password"
from users
where "username" = $1
`

type LoginRow struct {
	ID        uuid.UUID   `json:"id"`
	Role      string      `json:"role"`
	Username  string      `json:"username"`
	FirstName pgtype.Text `json:"first_name"`
	LastName  pgtype.Text `json:"last_name"`
	Location  uuid.UUID   `json:"location"`
	Password  []byte      `json:"password"`
}

func (q *Queries) Login(ctx context.Context, username string) (LoginRow, error) {
	row := q.db.QueryRow(ctx, login, username)
	var i LoginRow
	err := row.Scan(
		&i.ID,
		&i.Role,
		&i.Username,
		&i.FirstName,
		&i.LastName,
		&i.Location,
		&i.Password,
	)
	return i, err
}
