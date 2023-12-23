-- name: CreateUser :one
insert into users
values ($1, $2, $3, $4, $5, $6, $7, $8)
returning *;
