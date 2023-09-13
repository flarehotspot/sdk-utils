package migrate

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
)

func execFile(path string, ctx context.Context, db *sql.DB) error {
	b, err := ioutil.ReadFile(path)
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
