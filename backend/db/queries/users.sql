-- name: CreateUser :one
insert into users
values ($1, $2, $3, $4, $5, $6, $7, $8)
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
