// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const getUserByID = `-- name: GetUserByID :one
SELECT first_name, id, last_name FROM users WHERE id = ?
`

type GetUserByIDRow struct {
	FirstName string
	ID        int32
	LastName  sql.NullString
}

func (q *Queries) GetUserByID(ctx context.Context, targetID int32) (GetUserByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, targetID)
	var i GetUserByIDRow
	err := row.Scan(&i.FirstName, &i.ID, &i.LastName)
	return i, err
}

const insertNewUser = `-- name: InsertNewUser :exec
INSERT INTO users (first_name, last_name) VALUES (?, ?)
`

type InsertNewUserParams struct {
	FirstName string
	LastName  sql.NullString
}

func (q *Queries) InsertNewUser(ctx context.Context, arg InsertNewUserParams) error {
	_, err := q.db.ExecContext(ctx, insertNewUser, arg.FirstName, arg.LastName)
	return err
}

const limitSQLCArg = `-- name: LimitSQLCArg :many
select first_name, id FROM users LIMIT ?
`

type LimitSQLCArgRow struct {
	FirstName string
	ID        int32
}

func (q *Queries) LimitSQLCArg(ctx context.Context, limit int32) ([]LimitSQLCArgRow, error) {
	rows, err := q.db.QueryContext(ctx, limitSQLCArg, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []LimitSQLCArgRow
	for rows.Next() {
		var i LimitSQLCArgRow
		if err := rows.Scan(&i.FirstName, &i.ID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUserOrders = `-- name: ListUserOrders :many
SELECT
	users.id,
	users.first_name,
	orders.price
FROM
	orders
LEFT JOIN users ON orders.user_id = users.id
WHERE orders.price > ?
`

type ListUserOrdersRow struct {
	ID        sql.NullInt32
	FirstName sql.NullString
	Price     string
}

func (q *Queries) ListUserOrders(ctx context.Context, minPrice string) ([]ListUserOrdersRow, error) {
	rows, err := q.db.QueryContext(ctx, listUserOrders, minPrice)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListUserOrdersRow
	for rows.Next() {
		var i ListUserOrdersRow
		if err := rows.Scan(&i.ID, &i.FirstName, &i.Price); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUserParenExpr = `-- name: ListUserParenExpr :many
SELECT id, first_name, last_name, age, job_status FROM users WHERE (job_status = 'APPLIED' OR job_status = 'PENDING')
AND id > ?
ORDER BY id
LIMIT ?
`

type ListUserParenExprParams struct {
	ID    int32
	Limit int32
}

func (q *Queries) ListUserParenExpr(ctx context.Context, arg ListUserParenExprParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUserParenExpr, arg.ID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Age,
			&i.JobStatus,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsersByFamily = `-- name: ListUsersByFamily :many
SELECT first_name, last_name FROM users WHERE age < ? AND last_name = ?
`

type ListUsersByFamilyParams struct {
	MaxAge   int32
	InFamily sql.NullString
}

type ListUsersByFamilyRow struct {
	FirstName string
	LastName  sql.NullString
}

func (q *Queries) ListUsersByFamily(ctx context.Context, arg ListUsersByFamilyParams) ([]ListUsersByFamilyRow, error) {
	rows, err := q.db.QueryContext(ctx, listUsersByFamily, arg.MaxAge, arg.InFamily)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListUsersByFamilyRow
	for rows.Next() {
		var i ListUsersByFamilyRow
		if err := rows.Scan(&i.FirstName, &i.LastName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsersByID = `-- name: ListUsersByID :many
SELECT first_name, id, last_name FROM users WHERE id < ?
`

type ListUsersByIDRow struct {
	FirstName string
	ID        int32
	LastName  sql.NullString
}

func (q *Queries) ListUsersByID(ctx context.Context, id int32) ([]ListUsersByIDRow, error) {
	rows, err := q.db.QueryContext(ctx, listUsersByID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListUsersByIDRow
	for rows.Next() {
		var i ListUsersByIDRow
		if err := rows.Scan(&i.FirstName, &i.ID, &i.LastName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsersWithLimit = `-- name: ListUsersWithLimit :many
SELECT first_name, last_name FROM users LIMIT ?
`

type ListUsersWithLimitRow struct {
	FirstName string
	LastName  sql.NullString
}

func (q *Queries) ListUsersWithLimit(ctx context.Context, limit int32) ([]ListUsersWithLimitRow, error) {
	rows, err := q.db.QueryContext(ctx, listUsersWithLimit, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListUsersWithLimitRow
	for rows.Next() {
		var i ListUsersWithLimitRow
		if err := rows.Scan(&i.FirstName, &i.LastName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
