//go:build !mono

package main

import (
	"log"
	"path/filepath"
	"plugin"

	"github.com/flarehotspot/flarehotspot/core/env"
	paths "github.com/flarehotspot/sdk/utils/paths"
	sdktools "github.com/flarehotspot/sdk/utils/tools"
)

func main() {
	if env.GoEnv == env.ENV_DEV {
		sdktools.CreateGoWorkspace()
		sdktools.BuildAllPlugins()
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
