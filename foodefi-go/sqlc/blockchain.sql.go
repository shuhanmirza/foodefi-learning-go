// Code generated by sqlc. DO NOT EDIT.
// source: blockchain.sql

package db

import (
	"context"
)

const createBlockchain = `-- name: CreateBlockchain :one
INSERT INTO blockchains (name)
VALUES ($1)
RETURNING id, name, created_at
`

func (q *Queries) CreateBlockchain(ctx context.Context, name string) (Blockchains, error) {
	row := q.db.QueryRowContext(ctx, createBlockchain, name)
	var i Blockchains
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const deleteBlockchain = `-- name: DeleteBlockchain :exec
DELETE
FROM blockchains
WHERE id = $1
`

func (q *Queries) DeleteBlockchain(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteBlockchain, id)
	return err
}

const getBlockchain = `-- name: GetBlockchain :one
SELECT id, name, created_at
from blockchains
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetBlockchain(ctx context.Context, id int64) (Blockchains, error) {
	row := q.db.QueryRowContext(ctx, getBlockchain, id)
	var i Blockchains
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const listBlockchains = `-- name: ListBlockchains :many
SELECT id, name, created_at
FROM blockchains
ORDER BY id
`

func (q *Queries) ListBlockchains(ctx context.Context) ([]Blockchains, error) {
	rows, err := q.db.QueryContext(ctx, listBlockchains)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Blockchains{}
	for rows.Next() {
		var i Blockchains
		if err := rows.Scan(&i.ID, &i.Name, &i.CreatedAt); err != nil {
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

const updateBlockchain = `-- name: UpdateBlockchain :exec
UPDATE blockchains
SET name = $2
WHERE id = $1
RETURNING id, name, created_at
`

type UpdateBlockchainParams struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateBlockchain(ctx context.Context, arg UpdateBlockchainParams) error {
	_, err := q.db.ExecContext(ctx, updateBlockchain, arg.ID, arg.Name)
	return err
}
