// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package datatype

import (
	"context"
)

const test = `-- name: Test :one
SELECT 1
`

func (q *Queries) Test(ctx context.Context) (int32, error) {
	row := q.db.QueryRow(ctx, test)
	var column_1 int32
	err := row.Scan(&column_1)
	return column_1, err
}
