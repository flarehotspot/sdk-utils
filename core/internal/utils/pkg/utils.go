package pkg

import (
	"core/env"
	"core/internal/config"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	sdkfs "github.com/flarehotspot/go-utils/fs"
	paths "github.com/flarehotspot/go-utils/paths"
	sdkpaths "github.com/flarehotspot/go-utils/paths"
)

var (
	ErrNotInstalled = errors.New("Plugin is not installed")
)

func IsDefInList(defs []config.PluginSrcDef, def config.PluginSrcDef) bool {
	for _, i := range defs {
		if i.Equal(def) {
			return true
		}
	}
	return false
}

func AllPluginDef() []config.PluginSrcDef {
	list := InsalledPluginsDef()
	localPlugins := LocalPlugins()
	for _, loc := range localPlugins {
		if !IsDefInList(list, loc) {
			list = append(list, loc)
		}
	}
	return list
}

func LocalPlugins() []config.PluginSrcDef {
	list := []config.PluginSrcDef{}
	paths := LocalPluginPaths()
	for _, p := range paths {
		list = append(list, config.PluginSrcDef{Src: config.PluginSrcLocal, LocalPath: p})
	}
	log.Println("local plugins list: ", list)
	return list
}

func InsalledPluginsDef() []config.PluginSrcDef {
	list := []config.PluginSrcDef{}
	paths := InstalledDirList()
	for _, p := range paths {
		info, err := GetInfoFromPath(p)
		if err != nil {
			log.Println("Error reading plugin info: ", err)
			continue
		}
		metadata, err := ReadMetadata(info.Package)
		if err != nil {
			log.Println("Error reading plugin metadata: ", err)
			continue
		}

		if info.Package == metadata.Package {
			list = append(list, metadata.Def)
		}
	}
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

// InstalledDirList returns the list of installed plugins in the plugins directory. The path of each plugin is an aboslute path.
func InstalledDirList() []string {
	var pluginList []string

	installedPluginsPath := filepath.Join(paths.PluginsDir, "installed")

	// check if plugins/installed directory exists before traversing
	if !(sdkfs.Exists(installedPluginsPath)) {
		return pluginList
	}

	// this lists all directories inside paths.PluginsDir/installed
	if err := sdkfs.LsDirs(installedPluginsPath, &pluginList, false); err != nil {
		panic(err)
	}

	return pluginList
}

func GetMetaDataPath(pkg string) string {
	return filepath.Join(sdkpaths.ConfigDir, "plugins", pkg, "metadata.json")
}

func WriteMetadata(def config.PluginSrcDef, pkg string, installPath string) error {
	cfg, err := config.ReadPluginsConfig()
	if err != nil {
		return err
	}

	meta := config.PluginMetadata{
		Def:         def,
		InstallPath: installPath,
	}

	cfg.Metadata = append(cfg.Metadata, meta)

	return config.WritePluginsConfig(cfg)
}

func ReadMetadata(pkg string) (metadata config.PluginMetadata, err error) {
	cfg, err := config.ReadPluginsConfig()
	if err != nil {
		return
	}

	for _, m := range cfg.Metadata {
		if m.Package == pkg {
			return m, nil
		}
	}

	return metadata, ErrNotInstalled
}

func IsPackageInstalled(pkg string) bool {
	installPath := GetInstallPath(pkg)
	err := ValidateInstallPath(installPath)
	return err == nil
}

func IsSrcDefInstalled(def config.PluginSrcDef) bool {
	installPath, ok := FindDefInstallPath(def)
	if !ok {
		return false
	}

	err := ValidateInstallPath(installPath)
	return err == nil
}

func InstalledPluginsList() (list []config.PluginMetadata) {
	cfg, err := config.ReadPluginsConfig()
	if err != nil {
		return list
	}

	list = []config.PluginMetadata{}
	for _, m := range cfg.Metadata {
		if IsSrcDefInstalled(m.Def) {
			list = append(list, m)
		}
	}

	return
}

func NeedsRecompile(def config.PluginSrcDef) bool {
	if env.GO_ENV == env.ENV_DEV && (def.Src == config.PluginSrcLocal || def.Src == config.PluginSrcSystem) {
		return true
	}

	info, err := GetInfoFromDef(def)
	if err != nil {
		return true
	}

	cfg, err := config.ReadPluginsConfig()
	if err != nil {
		log.Println("Error reading plugins config: ", err)
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
	updatepath := GetPendingUpdatePath(pkg)
	return ValidateInstallPath(updatepath) == nil
}

func MovePendingUpdate(pkg string) error {
	updatePath := GetPendingUpdatePath(pkg)
	if err := CreateBackup(pkg); err != nil {
		return err
	}
	if err := sdkfs.Copy(updatePath, GetInstallPath(pkg)); err != nil {
		if err := RestoreBackup(pkg); err != nil {
			return err
		}
		return err
	}
	if err := os.RemoveAll(updatePath); err != nil {
		return err
	}
	if HasBackup(pkg) {
		if err := RemoveBackup(pkg); err != nil {
			return err
		}
	}
	return nil
}

func CreateBackup(pkg string) error {
	installPath := GetInstallPath(pkg)
	backupPath := GetBackupPath(pkg)
	if sdkfs.Exists(backupPath) {
		if err := os.RemoveAll(backupPath); err != nil {
			return err
		}
	}
	return sdkfs.Copy(installPath, backupPath)
}

func HasBackup(pkg string) bool {
	backup := GetBackupPath(pkg)
	err := ValidateInstallPath(backup)
	return err == nil
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

func RemovePendingUpdate(pkg string) error {
	return os.RemoveAll(GetPendingUpdatePath(pkg))
}

func ValidateSrcPath(src string) error {
	requiredFiles := []string{"plugin.json", "go.mod", "main.go"}

	for _, f := range requiredFiles {
		if !sdkfs.Exists(filepath.Join(src, f)) {
			return errors.New(f + " not found in source path")
		}
	}
	return nil
}

func ValidateInstallPath(src string) error {
	requiredFiles := []string{"plugin.json", "go.mod", "plugin.so"}

	for _, f := range requiredFiles {
		if !sdkfs.Exists(filepath.Join(src, f)) {
			return errors.New(f + " not found in source path")
		}
	}
	return nil
}

func FindPluginSrc(dir string) (string, error) {
	files := []string{}
	err := sdkfs.LsFiles(dir, &files, true)
	if err != nil {
		return dir, err
	}

	for _, f := range files {
		if filepath.Base(f) == "plugin.json" {
			return filepath.Dir(f), nil
		}
	}

	return "", errors.New("Can't find plugin.json in " + paths.StripRoot(dir))
}

func FindDefInstallPath(def config.PluginSrcDef) (installPath string, ok bool) {
	cfg, err := config.ReadPluginsConfig()
	if err != nil {
		return
	}

	for _, meta := range cfg.Metadata {
		if def.Equal(meta.Def) {
			return GetInstallPath(meta.Package), true
		}
	}

	return "", false
}

func GetAuthorNameFromGitUrl(def config.PluginSrcDef) string {
	return strings.Split(strings.TrimPrefix(def.GitURL, "https://github.com/"), "/")[0]
}

func GetRepoFromGitUrl(def config.PluginSrcDef) string {
	return strings.Split(strings.TrimPrefix(def.GitURL, fmt.Sprintf("https://github.com/%s/", GetAuthorNameFromGitUrl(def))), "/")[0]
}

func GetInstallPath(pkg string) string {
	return filepath.Join(sdkpaths.PluginsDir, "installed", pkg)
}

func GetPendingUpdatePath(pkg string) string {
	return filepath.Join(sdkpaths.PluginsDir, "update", pkg)
}

func GetBackupPath(pkg string) string {
	return filepath.Join(sdkpaths.PluginsDir, "backup", pkg)
}

func ListPluginDirs(includeCore bool) []string {
	searchPaths := []string{"plugins/system", "plugins/local"}
	pluginDirs := []string{}

	if includeCore {
		pluginDirs = append(pluginDirs, "core")
	}

	for _, s := range searchPaths {
		var list []string
		if err := sdkfs.LsDirs(s, &list, false); err == nil {
			pluginDirs = append(pluginDirs, list...)
		}
	}

	return pluginDirs
}
