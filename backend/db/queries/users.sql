/* These are example queries and do not reflect
the columns in the production users table */

-- name: GetUsers :many
select * from users;

-- name: GetUser :one
select *
from users
where id = $1 limit 1;

-- name: DeleteUser :exec
delete from users
where id = $1;

-- name: CreateUser :one
insert into users (name)
values ($1)
returning *;

-- name: UpdateUser :exec
update users
set name = $2
where id = $1
returning *;
