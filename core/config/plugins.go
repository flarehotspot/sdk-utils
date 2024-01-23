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

func PluginsDefaultList() (PluginList, error) {
	defaultsFile := filepath.Join(paths.DefaultsDir, pluginsConfigJsonFile)
	bytes, err := os.ReadFile(defaultsFile)
	if err != nil {
		return PluginList{}, err
	}

	var defaultsJson PluginList
	err = json.Unmarshal(bytes, &defaultsJson)

	return defaultsJson, err
}

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
	defaultPlugins, err := PluginsDefaultList()
	if err != nil {
		log.Println("Failed to load default plugins:", err)
	}

	userPlugins := PluginsUserList()
	return append(defaultPlugins, userPlugins...)
}

// ListDirs returns a list of plugins (aboslute path to plugin directory) from "vendor" directory.
// If same directory name exists in "plugins" directory, the absolute path from "plugins" directory is returned instead.
func PluginDirList() []string {
	vendorDirs := []string{}
	if err := fs.LsDirs(paths.VendorDir, &vendorDirs, false); err != nil {
		panic("Unable to list plugin directories.")
	}

	pluginDirs := []string{}
	if err := fs.LsDirs(paths.PluginsDir, &pluginDirs, false); err != nil {
		return vendorDirs
	}

	list := []string{}

	for _, vendorDir := range vendorDirs {
		var pluginDir *string

		for _, pdir := range pluginDirs {
			vname := filepath.Base(vendorDir)
			pname := filepath.Base(pdir)

			if pname == vname {
				pluginDir = &pdir
				break
			}
		}

		if pluginDir != nil {
			list = append(list, *pluginDir)
		} else {
			list = append(list, vendorDir)
		}
	}

	log.Println("Plugin List: ")
	for _, p := range list {
		log.Println("\t" + p)
	}

	return list
}
