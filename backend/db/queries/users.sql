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

-- name: GetUserByID :one
select
    "id",
    "role",
    "email",
    "username",
    "first_name",
    "last_name",
    "location"
from users
where "id" = $1;

-- name: GetUserByUsername :one
select
    "id",
    "role",
    "email",
    "username",
    "first_name",
    "last_name",
    "location"
from users
where "username" = $1;

-- name: GetUserByEmail :one
select
    "id",
    "role",
    "email",
    "username",
    "first_name",
    "last_name",
    "location"
from users
where "email" = $1;

-- name: UpdateUser :one
update users
set
    "role" = $1,
    "email" = $2,
    "first_name" = $3,
    "last_name" = $4,
    "location" = $5
where "id" = $6
returning *;

-- name: ChangePassword :exec
update users
set "password" = $1
where "id" = $2;

-- name: PromoteToAdmin :one
update users
set "role" = 'admin'
where "id" = $1
returning *;

-- name: DeleteUser :exec
delete from users
where "id" = $1;

-- name: Login :one
select
    "id",
    "role",
    "username",
    "first_name",
    "last_name",
    "email",
    "location",
    "password"
from users
where "username" = $1;
