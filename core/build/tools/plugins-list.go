package tools

import (
	"path/filepath"

	sdkfs "sdk/utils/fs"
)

// PluginPathList returns a list of plugin (absolute) paths
func PluginPathList() []string {
	searchPaths := []string{"plugins"}
	pluginPaths := []string{}

	for _, sp := range searchPaths {
		if sdkfs.Exists(sp) {
			var dirs []string
			if err := sdkfs.LsDirs(sp, &dirs, false); err != nil {
				continue
			}

			for _, dir := range dirs {
				pluginJson := filepath.Join(dir, "plugin.json")
				modFile := filepath.Join(dir, "go.mod")

				if sdkfs.Exists(pluginJson) && sdkfs.Exists(modFile) {
					pluginPaths = append(pluginPaths, dir)
				}
			}
		}
	}

	return pluginPaths
}
