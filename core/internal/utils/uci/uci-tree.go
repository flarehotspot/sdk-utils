//go:build !dev
package uci

import "sdk/libs/go-uci"

var UciTree = uci.NewTree("/etc/config")
