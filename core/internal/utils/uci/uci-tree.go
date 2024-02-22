//go:build !dev
package uci

import "github.com/flarehotspot/sdk/libs/go-uci"

var UciTree = uci.NewTree("/etc/config")
