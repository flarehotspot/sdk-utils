package sdkconnmgr

import "context"

// IClientDevice represents a client device connected to the network.
type IClientDevice interface {
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
	ValidSession(ctx context.Context) (IClientSession, error)

	// Returns true if the client device has a valid session.
	HasSession(ctx context.Context) (ok bool)
}
