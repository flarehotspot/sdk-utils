// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: copyfrom.go

package querytest

import (
	"context"
)

// iteratorForStageUserData implements pgx.CopyFromSource.
type iteratorForStageUserData struct {
	rows                 []StageUserDataParams
	skippedFirstNextCall bool
}

func (r *iteratorForStageUserData) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForStageUserData) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].IDParam,
		r.rows[0].UserParam,
	}, nil
}

func (r iteratorForStageUserData) Err() error {
	return nil
}

func (q *Queries) StageUserData(ctx context.Context, arg []StageUserDataParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"user_data"}, []string{"id", "user"}, &iteratorForStageUserData{rows: arg})
}
