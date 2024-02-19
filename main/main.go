//go:build !mono

package main

import (
	"log"
	"path/filepath"
	"plugin"

	"github.com/flarehotspot/flarehotspot/core/env"
	"github.com/flarehotspot/flarehotspot/core/utils/tools"
	paths "github.com/flarehotspot/sdk/utils/paths"
)

func main() {
	if env.GoEnv == env.ENV_DEV {
		tools.CreateGoWorkspace()
		tools.BuildAllPlugins()
	}

	corePath := filepath.Join(paths.AppDir, "core/plugin.so")
	p, err := plugin.Open(corePath)
	if err != nil {
		log.Println("Error loading core plugin:", err)
		panic(err)
	}
	symInit, _ := p.Lookup("Init")
	initFn := symInit.(func())
	initFn()
}
