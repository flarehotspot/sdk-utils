package plugins

import (
	"github.com/flarehotspot/core/internal/utils/uci"
	"github.com/flarehotspot/sdk/api/uci"
	gouci "github.com/flarehotspot/sdk/libs/go-uci"
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
