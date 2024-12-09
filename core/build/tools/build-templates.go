package tools

import (
	"core/internal/utils/pkg"
	"path/filepath"

	sdkpaths "github.com/flarehotspot/go-utils/paths"
	sdkpkg "github.com/flarehotspot/go-utils/pkg"
)

func BuildTemplates() {
	pluginDirs := []string{}

	defs := pkg.AllPluginSrcDefs()
	for _, def := range defs {
		if def.Src == sdkpkg.PluginSrcLocal || def.Src == sdkpkg.PluginSrcSystem {
			pluginDirs = append(pluginDirs, def.LocalPath)
		}
	}

	corePath := filepath.Join(sdkpaths.AppDir, "core")
	pluginDirs = append(pluginDirs, corePath)

	for _, p := range pluginDirs {
		pkg.BuildTemplates(p)
	}
}
