package sdknet

// Network is used to get network data from the system.
type Network interface {

	// Returns a list of all network devices.
	ListDevices() ([]NetworkDevice, error)

	// Returns a list of all network interfaces.
	ListInterfaces() ([]NetworkInterface, error)

	// Returns data of a single network device.
	GetDevice(name string) (NetworkDevice, error)

	// Returns data of a single network interface.
	GetInterface(name string) (NetworkInterface, error)

	// Returns data of a single network interface by its IP address.
	// The clientIp parameter is the IP address of the client that is connected to the system.
	FindByIp(clientIp string) (NetworkInterface, error)

	// Returns the network traffic API.
	Traffic() TrafficApi
}
