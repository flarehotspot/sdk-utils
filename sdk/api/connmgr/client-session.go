/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkconnmgr

import (
	"context"
	"time"
)

// ClientSession represents a client's internet connection session.
type ClientSession interface {
	// Returns the session's id.
	Id() int64

	// Returns the session's device id.
	DeviceId() int64

	// Returns the session type.
	Type() uint8

	// Return the session's available time in seconds.
	TimeSecs() (sec uint)

	// Returns the session's available data in megabytes.
	DataMb() (mbytes float64)

	// Returns the session's time consumption in seconds.
	TimeConsumption() (sec uint)

	// Returns the session's data consumption in megabytes.
	DataConsumption() (mbytes float64)

	// Returns the time when session was started.
	StartedAt() *time.Time

	// Returns the session's expiration time in days.
	// If session has no expiration, it returns nil.
	ExpDays() *uint

	// Returns the time when session will expire.
	// If session has no expiration, it returns nil.
	// Expiration time is calculated from the time when session was started.
	ExpiresAt() *time.Time

	// Returns the session's download speed limit in megabits per second.
	DownMbits() int

	// Returns the session's upload speed limit in megabits per second.
	UpMbits() int

	// Returns whether session uses global speed limits.
	UseGlobal() bool

	// Sets the session's available time in seconds.
	// This value is not saved until Save() method is called.
	SetTimeSecs(sec uint)

	// Sets the session's available data in megabytes.
	// This value is not saved until Save() method is called.
	SetDataMb(mbytes float64)

	// Sets the session's time consumption in seconds.
	// This value is not saved until Save() method is called.
	SetTimeCons(sec uint)

	// Sets the session's data consumption in megabytes.
	// This value is not saved until Save() method is called.
	SetDataCons(mbytes float64)

	// Sets the time when session was started.
	// This value is not saved until Save() method is called.
	SetStartedAt(started *time.Time)

	// Sets the session's expiration time in days.
	// This value is not saved until Save() method is called.
	SetExpDays(exp *uint)

	// Sets the session's download speed limit in megabits per second.
	// This value is not saved until Save() method is called.
	SetDownMbits(mbits int)

	// Sets the session's upload speed limit in megabits per second.
	// This value is not saved until Save() method is called.
	SetUpMbits(mbits int)

	// Sets whether session uses global speed limits.
	// This value is not saved until Save() method is called.
	SetUseGlobals(bool)

	// Increases the session's time consumption in seconds.
	// This value is not saved until Save() method is called.
	IncTimeCons(sec uint)

	// Increases the session's data consumption in megabytes.
	// This value is not saved until Save() method is called.
	IncDataCons(mbytes float64)

	// Saves the session's changes.
	Save(ctx context.Context) error

	// Reloads the session's data from the database.
	Reload(ctx context.Context) error
}
