//go:build !mono

package main

import (
	paths "github.com/flarehotspot/core/sdk/utils/paths"
	"log"
	"path/filepath"
	"plugin"
)

func main() {
	log.Println("App dir: ", paths.AppDir)
	corePath := filepath.Join(paths.AppDir, "core/core.so")
	log.Println("Core path: ", corePath)
	p, err := plugin.Open(corePath)
	if err != nil {
		log.Println("Error loading core.so:", err)
		panic(err)
	}
	symInit, _ := p.Lookup("Init")
	initFn := symInit.(func())
	initFn()
}
