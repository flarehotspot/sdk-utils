package sdkcfg

// Database is the database configuration.
type Database struct {
	// Host is the hostname or IP address of the database server.
	Host string

	// Username is the username to use when connecting to the database.
	Username string

	// Password is the password to use when connecting to the database.
	Password string

	// Database is the name of the database to use.
	Database string
}

// IDatabaseCfg is the interface for database configuration.
type IDatabaseCfg interface {
	// Read reads the database configuration.
	Read() (*Database, error)

	// Write writes the database configuration.
	Write(*Database) error
}
