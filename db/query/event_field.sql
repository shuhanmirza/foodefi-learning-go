-- name: CreateEventField :one
INSERT INTO event_fields (event_id,
                          name,
                          type,
                          value,
                          recorder)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetEventFieldById :one
SELECT *
from event_fields
WHERE id = $1
LIMIT 1;

-- name: ListEventFields :many
SELECT *
FROM event_fields
ORDER BY event_id;

-- name: DeleteEventField :exec
DELETE
FROM event_fields
WHERE id = $1;