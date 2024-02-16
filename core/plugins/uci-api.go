package plugins

import (
	"github.com/flarehotspot/flarehotspot/core/utils/uci"
	ucisdk "github.com/flarehotspot/flarehotspot/core/sdk/api/uci"
	gouci "github.com/flarehotspot/flarehotspot/core/sdk/libs/go-uci"
)

type UciApi struct {
	networkApi  *uci.UciNetworkApi
	dhcpApi     *uci.UciDhcpApi
	wirelessApi *uci.UciWirelessApi
}

func NewUciApi() *UciApi {
	return &UciApi{}
}

func (self *UciApi) Network() ucisdk.NetworkApi {
	return self.networkApi
}

func (self *UciApi) Dhcp() ucisdk.DhcpApi {
	return self.dhcpApi
}

func (self *UciApi) Wireless() ucisdk.WirelessApi {
	return self.wirelessApi
}

func (self *UciApi) Uci() gouci.Tree {
	return uci.UciTree
}
