-- name: CreateUser :one
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
returning *;

-- name: GetUser :one
select "id",
    "role",
    "username",
    "first_name",
    "last_name",
    "email",
    "location"
from users
where username = $1 limit 1;

-- name: Login :one
select
    "id",
    "role",
    "username",
    "first_name",
    "last_name",
    "location",
    "password"
from users
where "username" = $1;
