package models

import (
	"context"
	"database/sql"
	"time"
)

// ISession represents a record in sessions table.
type ISession interface {
	// Returns the session id.
	Id() int64

	// Returns teh device id it belongs to.
	DeviceId() int64

	// Returns the session type.
	SessionType() uint8

	// Returns the session time in seconds.
	TimeSecs() uint

	// Return the session data in megabytes.
	DataMbyte() float64

	// Returns the consumed time in seconds.
	TimeConsumed() uint

	// Returns the consumed data in megabytes.
	DataConsumed() float64

	// Returns the time when the session started.
	StartedAt() *time.Time

	// Returns the expiration time in days.
	ExpDays() *uint

	// Returns the expiration time.
	// This value is computed only after the session has started.
	ExpiresAt() *time.Time

	// Returns the downlink speed in Mbits.
	DownMbits() int

	// Returns the uplink speed in Mbits.
	UpMbits() int

	// Returns true if the session uses the global link speed limit.
	UseGlobal() bool

	// Updates the session with the given values using a database transaction.
	UpdateTx(tx *sql.Tx, ctx context.Context, devId int64, t uint8, timeSecs uint, dataMbytes float64, timeCons uint, dataCons float64, started *time.Time, exp *uint, downMbit int, upMbit int, g bool) error

	// Saves the session to the database using a database transaction.
	SaveTx(tx *sql.Tx, ctx context.Context) error

	// Updaets the session with the given values.
	Update(ctx context.Context, devId int64, t uint8, timeSecs uint, dataMbytes float64, timeCons uint, dataCons float64, started *time.Time, exp *uint, downMbit int, upMbit int, g bool) error

	// Saves the session to the database.
	Save(ctx context.Context) error
}
