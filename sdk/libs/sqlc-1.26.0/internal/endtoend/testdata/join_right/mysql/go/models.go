// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package querytest

import (
	"database/sql"
)

type Bar struct {
	ID uint64
}

type Foo struct {
	ID    uint64
	BarID sql.NullInt32
}
