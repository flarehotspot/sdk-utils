//go:build !mono

package main

import (
	"log"
	"path/filepath"
	"plugin"

	"github.com/flarehotspot/core/devkit/tools"
	"github.com/flarehotspot/core/devkit/env"
	paths "github.com/flarehotspot/core/sdk/utils/paths"
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
