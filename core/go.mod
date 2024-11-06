// NOTE: Do not use "go mod tidy" to prevent coupling of dependencies.

module core

go 1.21

toolchain go1.21.13

require (
	github.com/go-sql-driver/mysql v1.8.1
	github.com/gorilla/mux v1.8.0
	github.com/stretchr/testify v1.8.2
	github.com/twitchtv/twirp v8.1.3+incompatible
	google.golang.org/protobuf v1.34.2
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/evanw/esbuild v0.24.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.5.5 // indirect
	github.com/gorilla/csrf v1.7.2 // indirect
	github.com/gorilla/securecookie v1.1.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/sqlc-dev/sqlc v1.26.0 // indirect
	golang.org/x/crypto v0.20.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
