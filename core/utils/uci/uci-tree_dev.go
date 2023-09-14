//go:build dev

package uci

import (
	"path/filepath"

	"github.com/flarehotspot/core/sdk/libs/go-uci"
	"github.com/flarehotspot/core/sdk/utils/paths"
)

var treeRoot = filepath.Join(paths.AppDir, "mock-files/etc/config")
var UciTree = uci.NewTree(treeRoot)
