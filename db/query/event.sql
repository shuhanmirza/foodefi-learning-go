-- name: CreateEvent :one
INSERT INTO events (blockchain_id,
                    block_number,
                    event_name)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetEvent :one
SELECT *
from events
WHERE id = $1
LIMIT 1;

-- name: ListEvents :many
SELECT *
FROM events
ORDER BY blockchain_id;

-- name: DeleteEvent :exec
DELETE
FROM events
WHERE id = $1;