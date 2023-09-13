//go:build !dev
package uci

import "github.com/flarehotspot/core/sdk/libs/go-uci"

var UciTree = uci.NewTree("/etc/config")
