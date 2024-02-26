/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkconnmgr

import "context"

// ClientDevice represents a client device connected to the network.
type ClientDevice interface {
	// Returns the database id of the device.
	Id() int64

	// Returns the hostname of the device.
	Hostname() string

	// Returns the IP address of the device.
	IpAddr() string

	// Returns the MAC address of the device.
	MacAddr() string

	// Updates the client device.
	Update(ctx context.Context, mac string, ip string, hostname string) error

	// Returns valid session for the client device.
	ValidSession(ctx context.Context) (ClientSession, error)

	// Returns true if the client device has a valid session.
	HasSession(ctx context.Context) (ok bool)

	// Emits a socket event to a client device.
	// The event will be propagated to the client's browser via server-sent events.
	// SocketEmit(clnt ClientDevice, t string, d map[string]any)
}
