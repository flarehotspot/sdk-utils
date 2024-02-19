package migrate

import (
	"context"
	"database/sql"
	"strings"
)

func MigrateUp(dir string, db *sql.DB) error {
	files, err := listFiles(dir, migration_Up)
	if err != nil {
		return err
	}

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	for _, f := range files {
		done, err := fileDone(f, ctx, db)
		if err != nil {
			return err
		}

		if !done {
			err = execFile(f, ctx, db)
			if err != nil {
				return err
			}
			err := commitFile(f, ctx, db)
			if err != nil {
				return err
			}
		}
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func MigrateDown(dir string, db *sql.DB) error {
	files, err := listFiles(dir, migration_Down)
	if err != nil {
		return err
	}

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	for _, downfile := range files {
		upfile := strings.ReplaceAll(downfile, ".down.sql", ".up.sql")
		done, err := fileDone(upfile, ctx, db)
		if err != nil {
			return err
		}

		if done {
			err = execFile(downfile, ctx, db)
			if err != nil {
				return err
			}
			err := uncommitFile(upfile, ctx, db)
			if err != nil {
				return err
			}
		}
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
