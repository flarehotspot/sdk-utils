// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package datatype

import (
	"context"
)

const noop = `-- name: Noop :one
SELECT 1
`

func (q *Queries) Noop(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, noop)
	var column_1 int64
	err := row.Scan(&column_1)
	return column_1, err
}
