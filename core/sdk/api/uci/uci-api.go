package uci

import "github.com/flarehotspot/core/sdk/libs/go-uci"

type IUciApi interface {
	Uci() uci.Tree
	Network() INetworkApi
	Dhcp() IDhcpApi
	Wireless() IWirelessApi
}
