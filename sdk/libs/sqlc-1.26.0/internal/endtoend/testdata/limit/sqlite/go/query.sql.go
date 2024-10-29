// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package querytest

import (
	"context"
)

const deleteLimit = `-- name: DeleteLimit :exec
DELETE FROM foo LIMIT ?
`

func (q *Queries) DeleteLimit(ctx context.Context, limit int64) error {
	_, err := q.db.ExecContext(ctx, deleteLimit, limit)
	return err
}

const limitMe = `-- name: LimitMe :many
SELECT bar FROM foo LIMIT ?
`

func (q *Queries) LimitMe(ctx context.Context, limit int64) ([]bool, error) {
	rows, err := q.db.QueryContext(ctx, limitMe, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []bool
	for rows.Next() {
		var bar bool
		if err := rows.Scan(&bar); err != nil {
			return nil, err
		}
		items = append(items, bar)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateLimit = `-- name: UpdateLimit :exec
UPDATE foo SET bar='baz' LIMIT ?
`

func (q *Queries) UpdateLimit(ctx context.Context, limit int64) error {
	_, err := q.db.ExecContext(ctx, updateLimit, limit)
	return err
}
