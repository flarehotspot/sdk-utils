/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkconnmgr

import (
	"context"

	"github.com/google/uuid"
)

// SessionsMgrApi is used to manage client devices.
type SessionsMgrApi interface {

	// Connects a client device to the internet.
	Connect(ctx context.Context, clnt ClientDevice, notify string) error

	// Disconnects a client device from the internet.
	// If notify is not nil, then the client device will be notified of the disconnection.
	Disconnect(ctx context.Context, clnt ClientDevice, notify string) error

	// Checks if a client device is connected to the internet.
	IsConnected(clnt ClientDevice) (connected bool)

	// Create a session for the client device
	CreateSession(
		ctx context.Context,
		devId uuid.UUID,
		t uint8,
		timeSecs uint,
		dataMbytes float64,
		expDays *uint,
		downMbits int,
		upMbits int,
		useGlobal bool,
	) error

	// Get the current running session of a client device.
	CurrSession(clnt ClientDevice) (cs ClientSession, ok bool)

	// Returns unconsumed session (if any) for the client device.
	GetSession(ctx context.Context, clnt ClientDevice) (ClientSession, error)

	// Register a hook to find a session for a client device.
	RegisterSessionProvider(SessionProvider)
}
