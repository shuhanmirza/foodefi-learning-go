// Code generated by sqlc. DO NOT EDIT.
// source: event.sql

package db

import (
	"context"
)

const createEvent = `-- name: CreateEvent :one
INSERT INTO events (blockchain_id,
                    block_number,
                    event_name)
VALUES ($1, $2, $3)
RETURNING id, blockchain_id, block_number, event_name
`

type CreateEventParams struct {
	BlockchainID int64  `json:"blockchain_id"`
	BlockNumber  int64  `json:"block_number"`
	EventName    string `json:"event_name"`
}

func (q *Queries) CreateEvent(ctx context.Context, arg CreateEventParams) (Events, error) {
	row := q.db.QueryRowContext(ctx, createEvent, arg.BlockchainID, arg.BlockNumber, arg.EventName)
	var i Events
	err := row.Scan(
		&i.ID,
		&i.BlockchainID,
		&i.BlockNumber,
		&i.EventName,
	)
	return i, err
}

const deleteEvent = `-- name: DeleteEvent :exec
DELETE
FROM events
WHERE id = $1
`

func (q *Queries) DeleteEvent(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteEvent, id)
	return err
}

const getEvent = `-- name: GetEvent :one
SELECT id, blockchain_id, block_number, event_name
from events
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetEvent(ctx context.Context, id int64) (Events, error) {
	row := q.db.QueryRowContext(ctx, getEvent, id)
	var i Events
	err := row.Scan(
		&i.ID,
		&i.BlockchainID,
		&i.BlockNumber,
		&i.EventName,
	)
	return i, err
}

const listEvents = `-- name: ListEvents :many
SELECT id, blockchain_id, block_number, event_name
FROM events
ORDER BY blockchain_id
`

func (q *Queries) ListEvents(ctx context.Context) ([]Events, error) {
	rows, err := q.db.QueryContext(ctx, listEvents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Events
	for rows.Next() {
		var i Events
		if err := rows.Scan(
			&i.ID,
			&i.BlockchainID,
			&i.BlockNumber,
			&i.EventName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
