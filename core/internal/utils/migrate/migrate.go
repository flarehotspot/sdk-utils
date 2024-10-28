package migrate

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"strings"
)

func MigrateUp(db *sql.DB, dir string) error {
	files, err := listFiles(dir, migration_Up)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		return nil
	}

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
			if err := execFile(f, ctx, db); err != nil {
				return err
			}

			if err := commitFile(f, ctx, db); err != nil {
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
