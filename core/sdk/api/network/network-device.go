package sdknet

type NetDevType string

const (
	NetDevBridge NetDevType = "bridge"
	NetDevEther  NetDevType = "ethernet"
	NetDevWLAN   NetDevType = "wlan"
	NetDevVLAN   NetDevType = "vlan"
)

// INetworkDevice represents a network device.
type INetworkDevice interface {

	// Returns the name of the network device.
	Name() string

	// Returns the type of the network device.
	Type() NetDevType

	// Returns the MAC address of the network device.
	MacAddr() string

	// Returns the status of the network device.
	Up() bool

	// Returns the link speed of the network device.
	Speed() string

	// Returns the names of bridge member ports.
	BridgeMembers() []string

	// Returns the current receive bytes of the network device.
	RxBytes() uint

	// Returns the current transmit bytes of the network device.
	TxBytes() uint
}
