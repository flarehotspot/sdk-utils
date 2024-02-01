package sdkconnmgr

// SessionsMgrApi is used to manage client devices.
type SessionsMgrApi interface {

	// Connects a client device to the internet.
	Connect(clnt ClientDevice) error

	// Disconnects a client device from the internet.
	// If notify is not nil, then the client device will be notified of the disconnection.
	Disconnect(clnt ClientDevice, notify error) error

	// Checks if a client device is connected to the internet.
	IsConnected(clnt ClientDevice) (connected bool)

	// Get the current session of a client device.
	CurrSession(clnt ClientDevice) (cs ClientSession, ok bool)

	// Emits a socket event to a client device.
	// The event will be propagated to the client's browser via server-sent events.
	SocketEmit(clnt ClientDevice, t string, d map[string]any)
}
