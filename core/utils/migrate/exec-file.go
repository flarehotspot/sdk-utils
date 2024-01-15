package migrate

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
)

func execFile(path string, ctx context.Context, db *sql.DB) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	_, err = db.ExecContext(ctx, string(b))
	if err != nil {
		log.Println(fmt.Sprintf("Error migrating file: %s \n%+v", path, err))
		return err
	}

	return nil
}
