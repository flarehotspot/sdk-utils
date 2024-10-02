// NOTE: Do not use "go mod tidy" to prevent coupling of dependencies.

module core

go 1.19

require (
	github.com/go-sql-driver/mysql v1.7.0
	github.com/gorilla/mux v1.8.0
	github.com/stretchr/testify v1.8.2
	github.com/twitchtv/twirp v8.1.3+incompatible
	google.golang.org/protobuf v1.34.2
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
