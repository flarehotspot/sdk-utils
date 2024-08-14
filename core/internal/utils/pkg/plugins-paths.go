package pkg

import (
	"math/rand"
	"path/filepath"

	sdkfs "sdk/utils/fs"
	sdkstr "sdk/utils/strings"
)

// LocalPluginPaths returns a list of plugin (absolute) paths to be compiled and installed
func LocalPluginPaths() []string {
	searchPaths := []string{"plugins/system", "plugins/local"}
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

func RandomPluginPath() string {
	paths := []string{"/etc", "/usr", "/var"}
	randname := sdkstr.Rand(6)
	randpath := paths[rand.Intn(len(paths))]

	return filepath.Join(randpath, randname)
}
