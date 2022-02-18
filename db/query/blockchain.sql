-- name: CreateBlockchain :one
INSERT INTO blockchains (name)
VALUES ($1)
RETURNING *;

-- name: GetBlockchain :one
SELECT *
from blockchains
WHERE id = $1
LIMIT 1;

-- name: ListBlockchains :many
SELECT *
FROM blockchains
ORDER BY id;

-- name: DeleteBlockchain :exec
DELETE
FROM blockchains
WHERE id = $1;

-- name: UpdateBlockchain :exec
UPDATE blockchains
SET name = $2
WHERE id = $1
RETURNING *;
