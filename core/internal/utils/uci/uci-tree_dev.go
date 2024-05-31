//go:build dev

package uci

import (
	"path/filepath"

	"sdk/libs/go-uci"
	paths "sdk/utils/paths"
)

var treeRoot = filepath.Join(paths.AppDir, "openwrt-files/etc/config")
var UciTree = uci.NewTree(treeRoot)
