package migrate

import "database/sql"


func Init(db *sql.DB) error {
	q := `
  CREATE TABLE IF NOT EXISTS migrations (
    id INT AUTO_INCREMENT PRIMARY KEY,
    file VARCHAR(255) NOT NULL,
    executed_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
  )
  `
	_, err := db.Exec(q)
	if err != nil {
		return err
	}
	return nil
}

