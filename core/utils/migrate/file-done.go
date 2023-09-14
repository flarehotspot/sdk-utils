package migrate

import (
	"context"
	"database/sql"
)

func fileDone(f string, ctx context.Context, db *sql.DB) (exists bool, err error) {
	var id int
	q := `SELECT id FROM migrations WHERE file = ? LIMIT 1`
	row := db.QueryRowContext(ctx, q, f)
	err = row.Scan(&id)

	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	if err != nil && err == sql.ErrNoRows {
		return false, nil
	}

	return true, nil
}
