package migrate

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
)

func execFile(path string, ctx context.Context, db *sql.DB) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	content := string(b)
	queries := strings.Split(content, ";")

	for _, q := range queries {
		if strings.TrimSpace(q) != "" {
			_, err = db.ExecContext(ctx, q + ";")
			if err != nil {
                log.Println(fmt.Sprintf("Error migrating\nfile: %s \n%+v\nquery: %s", path, err, q))
				return err
			}
		}
	}

	return nil
}
