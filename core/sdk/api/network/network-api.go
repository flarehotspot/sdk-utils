package sdknet

// INetwork is used to get network data from the system.
type INetwork interface {

	// Returns a list of all network devices.
	ListDevices() ([]INetworkDevice, error)

	// Returns a list of all network interfaces.
	ListInterfaces() ([]INetworkInterface, error)

	// Returns data of a single network device.
	GetDevice(name string) (INetworkDevice, error)

	// Returns data of a single network interface.
	GetInterface(name string) (INetworkInterface, error)

	// Returns data of a single network interface by its IP address.
	// The clientIp parameter is the IP address of the client that is connected to the system.
	FindByIp(clientIp string) (INetworkInterface, error)

	// Returns the network traffic API.
	Traffic() ITrafficApi
}
