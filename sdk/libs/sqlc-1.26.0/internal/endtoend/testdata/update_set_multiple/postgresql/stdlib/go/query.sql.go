// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package querytest

import (
	"context"
)

const updateSetMultiple = `-- name: UpdateSetMultiple :exec
UPDATE foo SET (name, slug) = ($2, $1)
`

type UpdateSetMultipleParams struct {
	Slug string
	Name string
}

func (q *Queries) UpdateSetMultiple(ctx context.Context, arg UpdateSetMultipleParams) error {
	_, err := q.db.ExecContext(ctx, updateSetMultiple, arg.Slug, arg.Name)
	return err
}
