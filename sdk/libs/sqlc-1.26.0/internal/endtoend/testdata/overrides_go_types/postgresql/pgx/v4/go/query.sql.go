// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package override

import (
	"context"

	orm "database/sql"
	fuid "github.com/gofrs/uuid"
	uuid "github.com/gofrs/uuid"
	null "github.com/volatiletech/null/v8"
	null_v4 "gopkg.in/guregu/null.v4"
)

const loadFoo = `-- name: LoadFoo :many
SELECT id, other_id, age, balance, bio, about FROM foo WHERE id = $1
`

func (q *Queries) LoadFoo(ctx context.Context, id uuid.UUID) ([]Foo, error) {
	rows, err := q.db.Query(ctx, loadFoo, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Foo
	for rows.Next() {
		var i Foo
		if err := rows.Scan(
			&i.ID,
			&i.OtherID,
			&i.Age,
			&i.Balance,
			&i.Bio,
			&i.About,
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

const loadFooWithAliases = `-- name: LoadFooWithAliases :many
SELECT
  id AS aliased_id,
  other_id AS aliased_other_id,
  age AS aliased_age,
  balance AS aliased_balance,
  bio AS aliased_bio,
  about AS aliased_about
FROM foo
WHERE id = $1
`

type LoadFooWithAliasesRow struct {
	AliasedID      uuid.UUID
	AliasedOtherID fuid.UUID
	AliasedAge     orm.NullInt32
	AliasedBalance null.Float32
	AliasedBio     null_v4.String
	AliasedAbout   *string
}

func (q *Queries) LoadFooWithAliases(ctx context.Context, namedParameterID uuid.UUID) ([]LoadFooWithAliasesRow, error) {
	rows, err := q.db.Query(ctx, loadFooWithAliases, namedParameterID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []LoadFooWithAliasesRow
	for rows.Next() {
		var i LoadFooWithAliasesRow
		if err := rows.Scan(
			&i.AliasedID,
			&i.AliasedOtherID,
			&i.AliasedAge,
			&i.AliasedBalance,
			&i.AliasedBio,
			&i.AliasedAbout,
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
