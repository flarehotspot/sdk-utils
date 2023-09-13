package plugins

import (
	cnet "github.com/flarehotspot/core/network"
	"github.com/flarehotspot/core/utils/ubus"
	sdknet "github.com/flarehotspot/core/sdk/api/network"
)

type NetworkApi struct {
	trfk *cnet.TrafficMgr
}

func NewNetworkApi(trfk *cnet.TrafficMgr) *NetworkApi {
	return &NetworkApi{trfk}
}

func (self *NetworkApi) ListDevices() (netdevs []sdknet.INetworkDevice, err error) {
	devices, err := ubus.GetNetworkDevices()
	if err != nil {
		return nil, err
	}

	netdevs = []sdknet.INetworkDevice{}
	for _, v := range devices {
		dev := cnet.NewNetworkDevice(v)
		netdevs = append(netdevs, dev)
	}

	return netdevs, nil
}

func (self *NetworkApi) ListInterfaces() (interfaces []sdknet.INetworkInterface, err error) {
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

func (self *NetworkApi) GetDevice(name string) (sdknet.INetworkDevice, error) {
	dev, err := ubus.GetNetworkDevice(name)
	if err != nil {
		return nil, err
	}
	return cnet.NewNetworkDevice(dev), nil
}

func (self *NetworkApi) GetInterface(name string) (sdknet.INetworkInterface, error) {
	_, err := ubus.GetNetworkInterface(name)
	if err != nil {
		return nil, err
	}
	return cnet.NewNetworkInterface(name), nil
}

func (self *NetworkApi) FindByIp(clientIp string) (sdknet.INetworkInterface, error) {
	iface, err := cnet.FindByIp(clientIp)
	if err != nil {
		return nil, err
	}

	return cnet.NewNetworkInterface(iface.Name()), nil
}

func (self *NetworkApi) Traffic() sdknet.ITrafficApi {
	return self.trfk
}
