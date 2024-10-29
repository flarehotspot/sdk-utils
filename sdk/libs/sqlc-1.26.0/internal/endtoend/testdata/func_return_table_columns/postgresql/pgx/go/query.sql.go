// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const testFuncSelectBlog = `-- name: TestFuncSelectBlog :many
select id, name from test_select_blog($1)
`

type TestFuncSelectBlogRow struct {
	ID   pgtype.Int4
	Name pgtype.Text
}

func (q *Queries) TestFuncSelectBlog(ctx context.Context, pID int32) ([]TestFuncSelectBlogRow, error) {
	rows, err := q.db.Query(ctx, testFuncSelectBlog, pID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TestFuncSelectBlogRow
	for rows.Next() {
		var i TestFuncSelectBlogRow
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
