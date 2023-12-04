/* These are example queries and do not reflect
the columns in the production users table */

-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1 LIMIT 1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (name)
VALUES ($1)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
SET name = $2
WHERE id = $1
RETURNING *;
