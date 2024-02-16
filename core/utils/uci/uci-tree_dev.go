//go:build dev

package uci

import (
	"path/filepath"

	"github.com/flarehotspot/sdk/libs/go-uci"
	paths "github.com/flarehotspot/sdk/utils/paths"
)

var treeRoot = filepath.Join(paths.AppDir, "openwrt-files/etc/config")
var UciTree = uci.NewTree(treeRoot)
