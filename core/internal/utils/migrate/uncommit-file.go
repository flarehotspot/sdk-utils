package migrate

import (
	"context"
	"database/sql"
)

func uncommitFile(path string, ctx context.Context, db *sql.DB) error {
	q := `DELTE FROM migrations WHERE file = "?" LIMIT 1`
	_, err := db.ExecContext(ctx, q, path)
	if err != nil {
		return err
	}
	return nil
}
