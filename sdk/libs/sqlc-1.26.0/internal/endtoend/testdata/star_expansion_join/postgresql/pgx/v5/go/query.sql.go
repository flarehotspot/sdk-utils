// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const starExpansionJoin = `-- name: StarExpansionJoin :many
SELECT a, b, c, d FROM foo, bar
`

type StarExpansionJoinRow struct {
	A pgtype.Text
	B pgtype.Text
	C pgtype.Text
	D pgtype.Text
}

func (q *Queries) StarExpansionJoin(ctx context.Context) ([]StarExpansionJoinRow, error) {
	rows, err := q.db.Query(ctx, starExpansionJoin)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []StarExpansionJoinRow
	for rows.Next() {
		var i StarExpansionJoinRow
		if err := rows.Scan(
			&i.A,
			&i.B,
			&i.C,
			&i.D,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
