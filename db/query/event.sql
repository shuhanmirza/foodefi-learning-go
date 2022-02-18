-- name: CreateEvent :one
INSERT INTO events (blockchain_id,
                    block_number,
                    event_name)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetEventById :one
SELECT *
from events
WHERE id = $1
LIMIT 1;

-- name: GetEventByAllData :one
SELECT *
from events
WHERE blockchain_id = $1
  and block_number = $2
  and event_name = $3
LIMIT 1;

-- name: ListEvents :many
SELECT *
FROM events
ORDER BY blockchain_id;

-- name: DeleteEvent :exec
DELETE
FROM events
WHERE id = $1;