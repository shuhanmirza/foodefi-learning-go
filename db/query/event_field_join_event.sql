-- name: ListEvents :many
SELECT event_fields.id as event_field_id,
       event_id,
       "name",
       "type",
       "value",
       recorder,
       created_at,
       blockchain_id,
       block_number,
       event_name
FROM event_fields
         INNER JOIN events on events.id = event_fields.event_id
ORDER BY event_id LIMIT  $1
OFFSET $2;

-- name: ListEventsByBlockchainId :many
SELECT event_fields.id as event_field_id,
       event_id,
       "name",
       "type",
       "value",
       recorder,
       created_at,
       block_number,
       event_name
FROM event_fields
         INNER JOIN events on events.id = event_fields.event_id
WHERE blockchain_id = $3
ORDER BY event_id LIMIT  $1
OFFSET $2;

-- name: ListEventsByBlockchainIdBlockNumberRange :many
SELECT event_fields.id as event_field_id,
       event_id,
       "name",
       "type",
       "value",
       recorder,
       created_at,
       block_number,
       event_name
FROM event_fields
         INNER JOIN events on events.id = event_fields.event_id
WHERE (blockchain_id = $3)
  and (block_number between sqlc.arg(from_block) and sqlc.arg(to_block))
ORDER BY event_id LIMIT  $1
OFFSET $2;


-- name: ListEventsByBlockchainIdEventName :many
SELECT event_fields.id as event_field_id,
       event_id,
       "name",
       "type",
       "value",
       recorder,
       created_at,
       block_number
FROM event_fields
         INNER JOIN events on events.id = event_fields.event_id
WHERE blockchain_id = $3
  and events.event_name = $4
ORDER BY event_id LIMIT  $1
OFFSET $2;

-- name: ListEventsByBlockchainIdEventNameBlockNumberRange :many
SELECT event_fields.id as event_field_id,
       event_id,
       "name",
       "type",
       "value",
       recorder,
       created_at,
       block_number
FROM event_fields
         INNER JOIN events on events.id = event_fields.event_id
WHERE (blockchain_id = $3)
  and (block_number between sqlc.arg(from_block) and sqlc.arg(to_block))
  and (event_name = $4)
ORDER BY event_id LIMIT  $1
OFFSET $2;
