// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package querytest

import (
	"context"
)

const testRecursive = `-- name: TestRecursive :one
WITH t1 AS (
    select 1 as foo
)
SELECT foo FROM t1
`

func (q *Queries) TestRecursive(ctx context.Context) (int32, error) {
	row := q.db.QueryRow(ctx, testRecursive)
	var foo int32
	err := row.Scan(&foo)
	return foo, err
}
