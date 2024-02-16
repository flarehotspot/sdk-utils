//go:build dev

package uci

import (
	"path/filepath"

	"github.com/flarehotspot/flarehotspot/core/sdk/libs/go-uci"
	paths "github.com/flarehotspot/flarehotspot/core/sdk/utils/paths"
)

var treeRoot = filepath.Join(paths.AppDir, "openwrt-files/etc/config")
var UciTree = uci.NewTree(treeRoot)
