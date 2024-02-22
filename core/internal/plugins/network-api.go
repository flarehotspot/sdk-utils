package plugins

import (
	cnet "github.com/flarehotspot/core/internal/network"
	"github.com/flarehotspot/core/internal/utils/ubus"
	sdknet "github.com/flarehotspot/sdk/api/network"
)

func NewNetworkApi(trfk *cnet.TrafficMgr) *NetworkApi {
	return &NetworkApi{trfk}
}

type NetworkApi struct {
	trfk *cnet.TrafficMgr
}

func (self *NetworkApi) ListDevices() (netdevs []sdknet.NetworkDevice, err error) {
	devices, err := ubus.GetNetworkDevices()
	if err != nil {
		return nil, err
	}

	netdevs = []sdknet.NetworkDevice{}
	for _, v := range devices {
		dev := cnet.NewNetworkDevice(v)
		netdevs = append(netdevs, dev)
	}

	return netdevs, nil
}

func (self *NetworkApi) ListInterfaces() (interfaces []sdknet.NetworkInterface, err error) {
	ifaces, err := ubus.GetNetworkInterfaces()
	if err != nil {
		return nil, err
	}

	for ifname := range ifaces {
		iface := cnet.NewNetworkInterface(ifname)
		interfaces = append(interfaces, iface)
	}

	return interfaces, nil
}

func (self *NetworkApi) GetDevice(name string) (sdknet.NetworkDevice, error) {
	dev, err := ubus.GetNetworkDevice(name)
	if err != nil {
		return nil, err
	}
	return cnet.NewNetworkDevice(dev), nil
}

func (self *NetworkApi) GetInterface(name string) (sdknet.NetworkInterface, error) {
	_, err := ubus.GetNetworkInterface(name)
	if err != nil {
		return nil, err
	}
	return cnet.NewNetworkInterface(name), nil
}

func (self *NetworkApi) FindByIp(clientIp string) (sdknet.NetworkInterface, error) {
	iface, err := cnet.FindByIp(clientIp)
	if err != nil {
		return nil, err
	}

	return cnet.NewNetworkInterface(iface.Name()), nil
}

func (self *NetworkApi) Traffic() sdknet.TrafficApi {
	return self.trfk
}
