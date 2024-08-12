package pkg

import (
	"core/internal/config"
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	sdkplugin "sdk/api/plugin"
	fs "sdk/utils/fs"
	sdkfs "sdk/utils/fs"
	paths "sdk/utils/paths"
	sdkpaths "sdk/utils/paths"
)

const (
	PluginSrcGit          string = "git"
	PluginSrcStore        string = "store"
	PluginSrcSystem       string = "system"
	PluginSrcLocal        string = "local"
	pluginsConfigJsonFile string = "plugins.json"
)

var (
	installedPluginsJson = filepath.Join(sdkpaths.CacheDir, "installed_plugins.json")

	PLUGIN_FILES = []string{
		"plugin.json",
		"plugin.so",
		"resources",
		"go.mod",
		"LICENSE.txt",
	}
)

type PluginSrc string

type PluginInstalledMark struct {
	Def         PluginSrcDef
	Installed   bool
	InstallPath string
}

type PluginDefList []PluginSrcDef

func PluginsUserList() PluginDefList {
	configFile := filepath.Join(paths.ConfigDir, pluginsConfigJsonFile)
	bytes, err := os.ReadFile(configFile)
	if err != nil {
		return PluginDefList{}
	}

	var userJson PluginDefList

	err = json.Unmarshal(bytes, &userJson)
	if err != nil {
		return PluginDefList{}
	}

	return userJson
}

func AllPluginDef() PluginDefList {
	// defaultPlugins, err := PluginsDefaultList()
	// if err != nil {
	// 	log.Println("Failed to load default plugins:", err)
	// }

	localPlugins := LocalPlugins()
	return localPlugins
	// return append(defaultPlugins, userPlugins...)
}

func LocalPlugins() PluginDefList {
	var list PluginDefList
	paths := LocalPluginPaths()
	for _, p := range paths {
		list = append(list, PluginSrcDef{Src: PluginSrcLocal, LocalPath: p})
	}
	log.Println("local plugins list: ", list)
	return list
}

// LocalPluginPaths returns a list of plugin absolute source paths
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

// InstalledDirList returns the list of installed plugins in the plugins directory.
func InstalledDirList() []string {
	var pluginList []string

	installedPluginsPath := filepath.Join(paths.PluginsDir, "installed")

	// check if plugins/installed directory exists before traversing
	if !(fs.Exists(installedPluginsPath)) {
		return pluginList
	}

	// this lists all directories inside paths.PluginsDir/installed
	if err := fs.LsDirs(installedPluginsPath, &pluginList, false); err != nil {
		panic(err)
	}

	log.Println("Plugin List: ")
	for _, p := range pluginList {
		log.Println("\t" + p)
	}

	return pluginList
}

func MarkPluginAsInstalled(def PluginSrcDef, installPath string) error {
	installedPlugins := InstalledPluginsList()
	for i, p := range installedPlugins {
		switch def.Src {
		case PluginSrcLocal, PluginSrcSystem:
			if p.Def.LocalPath == def.LocalPath {
				installedPlugins[i].InstallPath = installPath
			}
		default:
			log.Println("MarkPluginAsInstalled: invalid source type")
			return nil
		}
	}

	installedPlugins = append(installedPlugins, PluginInstalledMark{
		Def:         def,
		Installed:   true,
		InstallPath: installPath,
	})

	return sdkfs.WriteJson(installedPluginsJson, installedPlugins)
}

func IsPluginInstalled(def PluginSrcDef) (ok bool, path string) {
	installedPlugins := InstalledPluginsList()
	for _, p := range installedPlugins {
		if (def.Src == PluginSrcLocal || def.Src == PluginSrcSystem) && p.Def.LocalPath == def.LocalPath {
			return sdkfs.Exists(p.InstallPath), p.InstallPath
		}
	}
	return false, ""
}

func InstalledPluginsList() []PluginInstalledMark {
	installedPlugins := []PluginInstalledMark{}
	if err := sdkfs.ReadJson(installedPluginsJson, &installedPlugins); err != nil {
		return installedPlugins
	}
	return installedPlugins
}

func PluginInstallPath(info sdkplugin.PluginInfo) string {
	return filepath.Join(sdkpaths.PluginsDir, "installed", info.Package)
}

func NeedsRecompile(def PluginSrcDef) bool {
	cfg, err := config.ReadPluginsConfig()
	if err != nil {
		log.Println("Error reading plugins config: ", err)
		return true
	}

	ok, path := IsPluginInstalled(def)
	if !ok {
		log.Println("Plugin is not installed: ", def.LocalPath)
		return true
	}

	info, err := PluginInfo(path)
	if err != nil {
		return true
	}

	for _, pkg := range cfg.Recompile {
		if info.Package == pkg {
			return true
		}
	}

	return false
}

func PluginInfo(path string) (sdkplugin.PluginInfo, error) {
	pluginInfo := sdkplugin.PluginInfo{}
	if err := sdkfs.ReadJson(filepath.Join(path, "plugin.json"), &pluginInfo); err != nil {
		return sdkplugin.PluginInfo{}, err
	}
	return pluginInfo, nil
}

func CoreInfo() sdkplugin.PluginInfo {
	pluginJsonPath := filepath.Join(sdkpaths.CoreDir, "plugin.json")
	var pluginDef sdkplugin.PluginInfo
	if err := sdkfs.ReadJson(pluginJsonPath, &pluginDef); err != nil {
		panic(err)
	}
	return pluginDef
}
