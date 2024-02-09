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
select
    "id",
    "role",
    "username",
    "first_name",
    "last_name",
    "location"
from users
where "id" = $1;

-- name: PromoteToAdmin :one
update users
set "role" = 'admin'
where "id" = $1
returning *;

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
