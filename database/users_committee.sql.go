// Code generated by sqlc. DO NOT EDIT.
// source: users_committee.sql

package database

import (
	"context"
)

const getPositionUsers = `-- name: GetPositionUsers :many
SELECT id, user_id, committee_id, from_date, till_date
FROM users_committee
WHERE committee_id = $1
`

func (q *Queries) GetPositionUsers(ctx context.Context, committeeID int32) ([]UsersCommittee, error) {
	rows, err := q.db.QueryContext(ctx, getPositionUsers, committeeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UsersCommittee
	for rows.Next() {
		var i UsersCommittee
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CommitteeID,
			&i.FromDate,
			&i.TillDate,
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

const getUserPositions = `-- name: GetUserPositions :many
SELECT id, user_id, committee_id, from_date, till_date
FROM users_committee
WHERE user_id = $1
`

func (q *Queries) GetUserPositions(ctx context.Context, userID int32) ([]UsersCommittee, error) {
	rows, err := q.db.QueryContext(ctx, getUserPositions, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UsersCommittee
	for rows.Next() {
		var i UsersCommittee
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CommitteeID,
			&i.FromDate,
			&i.TillDate,
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
