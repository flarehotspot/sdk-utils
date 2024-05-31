package plugins

import (
	"core/internal/utils/uci"
	"sdk/api/uci"
	gouci "sdk/libs/go-uci"
)

type UciApi struct {
	networkApi  *uci.UciNetworkApi
	dhcpApi     *uci.UciDhcpApi
	wirelessApi *uci.UciWirelessApi
}

func NewUciApi() *UciApi {
	return &UciApi{}
}

func (self *UciApi) Network() sdkuci.NetworkApi {
	return self.networkApi
}

func (self *UciApi) Dhcp() sdkuci.DhcpApi {
	return self.dhcpApi
}

func (self *UciApi) Wireless() sdkuci.WirelessApi {
	return self.wirelessApi
}

func (self *UciApi) Uci() gouci.Tree {
	return uci.UciTree
}
