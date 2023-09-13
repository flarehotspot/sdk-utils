package uci

// INetworkApi is used to get/set network configuration
type INetworkApi interface {
	// device
	GetDevice(section string) (dev string, err error)
	GetDeviceSec(name string) (section string, err error)
	GetDeviceType(section string) (t string, err error)

	// bridge
	GetBridgeSecs() (sections []string)
	GetBridgeVlanFilter(section string) (enabled bool)
	GetBridgePorts(section string) (ports []string, err error)
	SetBridgePorts(section string, ports []string) error
	SetBridgeVlanFilter(section string, enabled bool) error

	// bridge-vlan
	CreateBrVlan(brvlan *BrVlan, ports []*BrVlanPort) error
	GetBrVlanSecs() (sections []string)
	GetBrVlanSec(brvlan *BrVlan) (section string, ok bool)
  GetBrVlanID(brvlan *BrVlan) (vlanid int, ok bool)
	GetBrVlanPorts(brvlan *BrVlan) (ports []*BrVlanPort, err error)
	SetBrVlanPorts(brvlan *BrVlan, ports []*BrVlanPort) error
	DeleteBrVlan(brvlan *BrVlan)

	// interface
	GetInterface(section string) (iface *NetIface, err error)
	GetInterfaceSecs() (sections []string)
  GetInterfaces() (ifaces []*NetIface, err error)
	SetInterface(section string, cfg *NetIface) error
}
