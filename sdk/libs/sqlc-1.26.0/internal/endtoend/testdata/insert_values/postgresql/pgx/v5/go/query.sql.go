// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const insertMultipleValues = `-- name: InsertMultipleValues :exec
INSERT INTO foo (a, b) VALUES ($1, $2), ($3, $4)
`

type InsertMultipleValuesParams struct {
	A   pgtype.Text
	B   pgtype.Int4
	A_2 pgtype.Text
	B_2 pgtype.Int4
}

func (q *Queries) InsertMultipleValues(ctx context.Context, arg InsertMultipleValuesParams) error {
	_, err := q.db.Exec(ctx, insertMultipleValues,
		arg.A,
		arg.B,
		arg.A_2,
		arg.B_2,
	)
	return err
}

const insertValues = `-- name: InsertValues :exec
INSERT INTO foo (a, b) VALUES ($1, $2)
`

type InsertValuesParams struct {
	A pgtype.Text
	B pgtype.Int4
}

func (q *Queries) InsertValues(ctx context.Context, arg InsertValuesParams) error {
	_, err := q.db.Exec(ctx, insertValues, arg.A, arg.B)
	return err
}
