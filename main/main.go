//go:build !dev

package main

import (
	"github.com/flarehotspot/interface/fs/paths"
	"path/filepath"
	"plugin"
)

func main() {
	p, _ := plugin.Open(filepath.Join(paths.AppDir, "core/core.so"))
	symInit, _ := p.Lookup("Init")
	initFn := symInit.(func())
	initFn()
}
