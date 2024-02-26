/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkconnmgr

import "context"

// SessionsMgr is used to manage client devices.
type SessionsMgr interface {

	// Connects a client device to the internet.
	Connect(clnt ClientDevice) error

	// Disconnects a client device from the internet.
	// If notify is not nil, then the client device will be notified of the disconnection.
	Disconnect(clnt ClientDevice, notify error) error

	// Checks if a client device is connected to the internet.
	IsConnected(clnt ClientDevice) (connected bool)

	// Get the current session of a client device.
	CurrSession(clnt ClientDevice) (cs ClientSession, ok bool)

	CreateSession(ctx context.Context, devId int64, t uint8, timeSecs uint, dataMbytes float64, expDays *uint, downMbits int, upMbits int, useGlobal bool) error
}
