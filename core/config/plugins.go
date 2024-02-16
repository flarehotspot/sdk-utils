package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	fs "github.com/flarehotspot/core/sdk/utils/fs"
	paths "github.com/flarehotspot/core/sdk/utils/paths"
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

// func PluginsDefaultList() (PluginList, error) {
// 	var pluginList []string
// 	if err := fs.LsDirs(paths.PluginsDir, &pluginList, false); err != nil {
// 		panic(err)
// 	}
// 	log.Println("Plugin List: ")
// 	for _, p := range pluginList {
// 		log.Println("\t" + p)
// 	}

// 	return pluginList, nil
// }

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

// PluginDirList returns the list of installed plugins in the plugins directory.
func PluginDirList() []string {
	var pluginList []string

	if err := fs.LsDirs(paths.SystemDir, &pluginList, false); err != nil {
		panic(err)
	}

	if err := fs.LsDirs(paths.PluginsDir, &pluginList, false); err != nil {
		panic(err)
	}

	log.Println("Plugin List: ")
	for _, p := range pluginList {
		log.Println("\t" + p)
	}

	return pluginList
}
