-- name: CreateUser :one
INSERT INTO users (username,
                   password,
                   role)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUser :one
SELECT *
from users
WHERE username = $1
LIMIT 1;

-- name: ListUsers :many
SELECT *
FROM users
ORDER BY username;

-- name: DeleteUser :exec
DELETE
FROM users
WHERE username = $1;

-- name: UpdateUserPassword :exec
UPDATE users
SET password = $2
WHERE username = $1
RETURNING *;

-- name: UpdateUserRole :one
UPDATE users
SET role = $2
WHERE username = $1
RETURNING *;