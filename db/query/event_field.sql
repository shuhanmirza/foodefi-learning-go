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

-- name: GetEventFieldByEventIdFieldName :one
SELECT *
from event_fields
WHERE event_id = $1
  and name = $2
LIMIT 1;

-- name: ListEventFields :many
SELECT *
FROM event_fields
ORDER BY event_id;

-- name: UpdateEventField :exec
UPDATE event_fields
SET value = $3
WHERE event_id = $1
  and name = $2
RETURNING *;

-- name: DeleteEventField :exec
DELETE
FROM event_fields
WHERE id = $1;

-- name: DeleteEventFieldByEventIdFieldName :exec
DELETE
FROM event_fields
WHERE event_id = $1
  and name = $2;

-- name: DeleteEventFieldByEventId :exec
DELETE
FROM event_fields
WHERE event_id = $1;