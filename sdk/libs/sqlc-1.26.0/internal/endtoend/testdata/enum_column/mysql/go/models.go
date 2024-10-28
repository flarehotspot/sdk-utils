// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package querytest

import (
	"database/sql/driver"
	"fmt"
)

type AuthorsAddItem string

const (
	AuthorsAddItemOk    AuthorsAddItem = "ok"
	AuthorsAddItemAdded AuthorsAddItem = "added"
)

func (e *AuthorsAddItem) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AuthorsAddItem(s)
	case string:
		*e = AuthorsAddItem(s)
	default:
		return fmt.Errorf("unsupported scan type for AuthorsAddItem: %T", src)
	}
	return nil
}

type NullAuthorsAddItem struct {
	AuthorsAddItem AuthorsAddItem
	Valid          bool // Valid is true if AuthorsAddItem is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAuthorsAddItem) Scan(value interface{}) error {
	if value == nil {
		ns.AuthorsAddItem, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.AuthorsAddItem.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAuthorsAddItem) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.AuthorsAddItem), nil
}

type AuthorsAdded string

const (
	AuthorsAddedOk AuthorsAdded = "ok"
)

func (e *AuthorsAdded) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AuthorsAdded(s)
	case string:
		*e = AuthorsAdded(s)
	default:
		return fmt.Errorf("unsupported scan type for AuthorsAdded: %T", src)
	}
	return nil
}

type NullAuthorsAdded struct {
	AuthorsAdded AuthorsAdded
	Valid        bool // Valid is true if AuthorsAdded is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAuthorsAdded) Scan(value interface{}) error {
	if value == nil {
		ns.AuthorsAdded, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.AuthorsAdded.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAuthorsAdded) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.AuthorsAdded), nil
}

type AuthorsBar string

const (
	AuthorsBarOk AuthorsBar = "ok"
)

func (e *AuthorsBar) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AuthorsBar(s)
	case string:
		*e = AuthorsBar(s)
	default:
		return fmt.Errorf("unsupported scan type for AuthorsBar: %T", src)
	}
	return nil
}

type NullAuthorsBar struct {
	AuthorsBar AuthorsBar
	Valid      bool // Valid is true if AuthorsBar is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAuthorsBar) Scan(value interface{}) error {
	if value == nil {
		ns.AuthorsBar, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.AuthorsBar.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAuthorsBar) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.AuthorsBar), nil
}

type AuthorsFoo string

const (
	AuthorsFooOk AuthorsFoo = "ok"
)

func (e *AuthorsFoo) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AuthorsFoo(s)
	case string:
		*e = AuthorsFoo(s)
	default:
		return fmt.Errorf("unsupported scan type for AuthorsFoo: %T", src)
	}
	return nil
}

type NullAuthorsFoo struct {
	AuthorsFoo AuthorsFoo
	Valid      bool // Valid is true if AuthorsFoo is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAuthorsFoo) Scan(value interface{}) error {
	if value == nil {
		ns.AuthorsFoo, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.AuthorsFoo.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAuthorsFoo) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.AuthorsFoo), nil
}

type AuthorsRemoveItem string

const (
	AuthorsRemoveItemOk AuthorsRemoveItem = "ok"
)

func (e *AuthorsRemoveItem) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AuthorsRemoveItem(s)
	case string:
		*e = AuthorsRemoveItem(s)
	default:
		return fmt.Errorf("unsupported scan type for AuthorsRemoveItem: %T", src)
	}
	return nil
}

type NullAuthorsRemoveItem struct {
	AuthorsRemoveItem AuthorsRemoveItem
	Valid             bool // Valid is true if AuthorsRemoveItem is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAuthorsRemoveItem) Scan(value interface{}) error {
	if value == nil {
		ns.AuthorsRemoveItem, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.AuthorsRemoveItem.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAuthorsRemoveItem) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.AuthorsRemoveItem), nil
}

type BooksFoo string

const (
	BooksFooOk BooksFoo = "ok"
)

func (e *BooksFoo) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = BooksFoo(s)
	case string:
		*e = BooksFoo(s)
	default:
		return fmt.Errorf("unsupported scan type for BooksFoo: %T", src)
	}
	return nil
}

type NullBooksFoo struct {
	BooksFoo BooksFoo
	Valid    bool // Valid is true if BooksFoo is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullBooksFoo) Scan(value interface{}) error {
	if value == nil {
		ns.BooksFoo, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.BooksFoo.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullBooksFoo) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.BooksFoo), nil
}

type Author struct {
	ID         int64
	Foo        AuthorsFoo
	Bar        AuthorsBar
	Added      AuthorsAdded
	AddItem    AuthorsAddItem
	RemoveItem AuthorsRemoveItem
}

type Book struct {
	ID  int64
	Foo BooksFoo
}
