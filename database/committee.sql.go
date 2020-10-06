// Code generated by sqlc. DO NOT EDIT.
// source: committee.sql

package database

import (
	"context"
)

const getPosition = `-- name: GetPosition :one
SELECT id, name, alias, ordering, description, executive
FROM committee
WHERE id = $1
`

func (q *Queries) GetPosition(ctx context.Context, id int32) (Position, error) {
	row := q.db.QueryRowContext(ctx, getPosition, id)
	var i Position
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Alias,
		&i.Ordering,
		&i.Description,
		&i.Executive,
	)
	return i, err
}

const listPositions = `-- name: ListPositions :many
SELECT id, name, alias, ordering, description, executive
FROM committee
`

func (q *Queries) ListPositions(ctx context.Context) ([]Position, error) {
	rows, err := q.db.QueryContext(ctx, listPositions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Position
	for rows.Next() {
		var i Position
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Alias,
			&i.Ordering,
			&i.Description,
			&i.Executive,
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