package migrate

import (
	"context"
	"database/sql"
)

func commitFile(path string, ctx context.Context, db *sql.DB) error {
	q := `INSERT INTO migrations (file) VALUES (?)`
	_, err := db.ExecContext(ctx, q, path)
	if err != nil {
		return err
	}
	return nil
}
