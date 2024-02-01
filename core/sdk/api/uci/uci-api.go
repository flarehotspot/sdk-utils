package sdkuci

import "github.com/flarehotspot/core/sdk/libs/go-uci"

type UciApi interface {
	Uci() uci.Tree
	Network() NetworkApi
	Dhcp() DhcpApi
	Wireless() WirelessApi
}
