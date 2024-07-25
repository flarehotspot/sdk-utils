package config

import (
	"core/internal/config/plugincfg"
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	fs "sdk/utils/fs"
	sdkfs "sdk/utils/fs"
	paths "sdk/utils/paths"
	sdkpaths "sdk/utils/paths"
)

const (
	PluginSrcGit          PluginSrc = "git"
	PluginSrcStore        PluginSrc = "store"
	pluginsConfigJsonFile string    = "plugins.json"
)

type PluginSrc string

// A plugin can be from store or from a git repo.
type PluginSrcDef struct {
	Src          PluginSrc `json:"src"`           // git | strore
	StorePackage string    `json:"store_pacakge"` // if src is "store"
	StoreVersion string    `json:"store_version"` // if src is "store"
	GitURL       string    `json:"git_url"`       // if src is "git"
	GitRef       string    `json:"git_ref"`       // can be a branch, tag or commit hash
}

type PluginList []*PluginSrcDef

func PluginsUserList() PluginList {
	configFile := filepath.Join(paths.ConfigDir, pluginsConfigJsonFile)
	bytes, err := os.ReadFile(configFile)
	if err != nil {
		return PluginList{}
	}

	var userJson PluginList

	err = json.Unmarshal(bytes, &userJson)
	if err != nil {
		return PluginList{}
	}

	return userJson
}

func AllPluginSrc() PluginList {
	// defaultPlugins, err := PluginsDefaultList()
	// if err != nil {
	// 	log.Println("Failed to load default plugins:", err)
	// }

	userPlugins := PluginsUserList()
	return userPlugins
	// return append(defaultPlugins, userPlugins...)
}

// InstalledDirList returns the list of installed plugins in the plugins directory.
func InstalledDirList() []string {
	var pluginList []string

	if err := fs.LsDirs(filepath.Join(paths.PluginsDir, "installed"), &pluginList, false); err != nil {
		panic(err)
	}

	log.Println("Plugin List: ")
	for _, p := range pluginList {
		log.Println("\t" + p)
	}

	return pluginList
}
