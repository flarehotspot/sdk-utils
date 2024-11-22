// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package querytest

import (
	"database/sql/driver"
	"fmt"
)

type Level string

const (
	LevelDEBUG Level = "DEBUG"
	LevelINFO  Level = "INFO"
	LevelWARN  Level = "WARN"
	LevelERROR Level = "ERROR"
	LevelFATAL Level = "FATAL"
)

func (e *Level) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Level(s)
	case string:
		*e = Level(s)
	default:
		return fmt.Errorf("unsupported scan type for Level: %T", src)
	}
	return nil
}

type NullLevel struct {
	Level Level
	Valid bool // Valid is true if Level is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullLevel) Scan(value interface{}) error {
	if value == nil {
		ns.Level, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Level.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullLevel) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Level), nil
}

type NewEvent string

const (
	NewEventSTART NewEvent = "START"
	NewEventSTOP  NewEvent = "STOP"
)

func (e *NewEvent) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = NewEvent(s)
	case string:
		*e = NewEvent(s)
	default:
		return fmt.Errorf("unsupported scan type for NewEvent: %T", src)
	}
	return nil
}

type NullNewEvent struct {
	NewEvent NewEvent
	Valid    bool // Valid is true if NewEvent is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullNewEvent) Scan(value interface{}) error {
	if value == nil {
		ns.NewEvent, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.NewEvent.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullNewEvent) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.NewEvent), nil
}

type LogLine struct {
	ID     int64
	Status NewEvent
	Level  Level
}
