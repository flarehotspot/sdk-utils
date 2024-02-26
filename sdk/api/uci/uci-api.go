/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkuci

import "github.com/flarehotspot/sdk/libs/go-uci"

type UciApi interface {
	Uci() uci.Tree
	Network() NetworkApi
	Dhcp() DhcpApi
	Wireless() WirelessApi
}
