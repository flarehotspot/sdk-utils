package sdknet

import "net"

// NetworkInterface represents a network interface in the system.
type NetworkInterface interface {

	// Returns the name of the interface.
	Ifname() string

	// Returns the device used for this interface.
	Device() (NetworkDevice, error)

	// Returns the status of the network interface.
	Up() bool

	// Returns the IPv4 address of the network interface.
	IpV4Addr() (*NetworkIpv4, error)

	// Returns the ip net value of the network interface.
	IPNet() (*net.IPNet, error)
}
