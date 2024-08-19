package pkg

import (
	"core/internal/config"
	jobque "core/internal/utils/job-que"
	"errors"
	"log"
	"os"
	"path/filepath"
	"sdk/libs/go-json"

	fs "sdk/utils/fs"
	sdkfs "sdk/utils/fs"
	paths "sdk/utils/paths"
	sdkpaths "sdk/utils/paths"
)

var (
	ErrNotInstalled = errors.New("Plugin is not installed")
)

const (
	PluginSrcGit          string = "git"
	PluginSrcStore        string = "store"
	PluginSrcSystem       string = "system"
	PluginSrcLocal        string = "local"
	pluginsConfigJsonFile string = "plugins.json"
)

var (
	markQue              = jobque.NewJobQue()
	installedPluginsJson = filepath.Join(sdkpaths.CacheDir, "installed_plugins.json")
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
	newList := []PluginInstalledMark{}
	plugins := InstalledPluginsList()
	found := false

	for _, p := range plugins {
		switch def.Src {
		case PluginSrcLocal, PluginSrcSystem:
			if def.Src == p.Def.Src && def.LocalPath == p.Def.LocalPath {
				found = true
				p.InstallPath = installPath
				p.Installed = true
			}
		}
		newList = append(newList, p)
	}

	if !found {
		newList = append(newList, PluginInstalledMark{Def: def, Installed: true, InstallPath: installPath})
	}

	_, err := markQue.Exec(func() (interface{}, error) {
		err := sdkfs.WriteJson(installedPluginsJson, newList)
		return nil, err
	})

	return err
}

func IsPluginInstalled(def PluginSrcDef) bool {
	_, ok := FindPluginInstallPath(def)
	return ok
}

func InstalledPluginsList() []PluginInstalledMark {
	result, err := markQue.Exec(func() (interface{}, error) {
		plugins := []PluginInstalledMark{}
		if err := sdkfs.ReadJson(installedPluginsJson, &plugins); err != nil {
			return nil, err
		}
		return plugins, nil
	})

	if err != nil {
		return []PluginInstalledMark{}
	}

	return result.([]PluginInstalledMark)
}

func NeedsRecompile(def PluginSrcDef) bool {
	cfg, err := config.ReadPluginsConfig()
	if err != nil {
		log.Println("Error reading plugins config: ", err)
		return true
	}

	path, ok := FindPluginInstallPath(def)
	if !ok {
		log.Println("Plugin is not installed: ", def.LocalPath)
		return true
	}

	info, err := GetSrcInfo(path)
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

func HasPendingUpdate(pkg string) bool {
	return sdkfs.Exists(filepath.Join(sdkpaths.PluginsDir, "update", pkg))
}

func MovePendingUpdate(pkg string) error {
	updatePath := GetPendingUpdatePath(pkg)
	if err := CreateBackup(pkg); err != nil {
		return err
	}
	if err := sdkfs.Copy(updatePath, GetInstallPath(pkg)); err != nil {
		return err
	}
	if err := os.RemoveAll(updatePath); err != nil {
		return err
	}
	return nil
}

func CreateBackup(pkg string) error {
	installPath := GetInstallPath(pkg)
	backupPath := GetBackupPath(pkg)
	return sdkfs.Copy(installPath, backupPath)
}

func HasBackup(pkg string) bool {
	return sdkfs.Exists(GetBackupPath(pkg))
}

func RestoreBackup(pkg string) error {
	backupPath := GetBackupPath(pkg)
	if err := sdkfs.Copy(backupPath, GetInstallPath(pkg)); err != nil {
		return err
	}
	if err := os.RemoveAll(backupPath); err != nil {
		return err
	}
	return nil
}

func RemoveBackup(pkg string) error {
	return os.RemoveAll(GetBackupPath(pkg))
}
