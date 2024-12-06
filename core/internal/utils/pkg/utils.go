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
	sdkpaths "github.com/flarehotspot/go-utils/paths"
	sdkpkg "github.com/flarehotspot/go-utils/pkg"
)

var (
	ErrNotInstalled = errors.New("Plugin is not installed")
)

func IsDefInList(defs []sdkpkg.PluginSrcDef, def sdkpkg.PluginSrcDef) bool {
	for _, i := range defs {
		if i.Equal(def) {
			return true
		}
	}
	return false
}

func AllPluginSrcDefs() []sdkpkg.PluginSrcDef {
	list := InstalledPluginsDef()
	localPlugins := LocalPluginSrcDefs()
	systemPlugins := SystemPluginSrcDefs()
	alldefs := append(systemPlugins, localPlugins...)

	for _, loc := range alldefs {
		if !IsDefInList(list, loc) {
			list = append(list, loc)
		}
	}

	return list
}

func LocalPluginSrcDefs() []sdkpkg.PluginSrcDef {
	list := []sdkpkg.PluginSrcDef{}
	paths := SearchPluginDirs(filepath.Join(sdkpaths.AppDir, "plugins/local"))
	for _, p := range paths {
		list = append(list, sdkpkg.PluginSrcDef{
			Src:       sdkpkg.PluginSrcLocal,
			LocalPath: p,
		})
	}
	log.Println("local plugins list: ", list)
	return list
}

func SystemPluginSrcDefs() []sdkpkg.PluginSrcDef {
	list := []sdkpkg.PluginSrcDef{}
	paths := SearchPluginDirs(filepath.Join(sdkpaths.AppDir, "plugins/system"))
	for _, pluginPath := range paths {
		list = append(list, sdkpkg.PluginSrcDef{
			Src:       sdkpkg.PluginSrcSystem,
			LocalPath: pluginPath,
		})
	}
	log.Println("system plugins list: ", list)
	return list
}

func InstalledPluginsDef() []sdkpkg.PluginSrcDef {
	list := []sdkpkg.PluginSrcDef{}
	paths := InstalledPluginDirs()
	for _, p := range paths {
		info, err := sdkpkg.GetInfoFromPath(p)
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

func SearchPluginDirs(searchPath string) (pluginDirs []string) {
	var list []string
	if err := sdkfs.LsDirs(searchPath, &list, false); err != nil {
		log.Println("Error listing directories in ", searchPath, ": ", err)
		return
	}
	for _, p := range list {
		if err := ValidateSrcPath(p); err == nil {
			pluginDirs = append(pluginDirs, p)
		} else {
			fmt.Println("Error validating source path: ", p, err)
		}
	}
	return
}

// InstalledPluginDirs returns the list of installed plugins in the plugins directory. The path of each plugin is an aboslute path.
func InstalledPluginDirs() (pluginDirs []string) {
	installedPluginsPath := filepath.Join(sdkpaths.PluginsDir, "installed")

	// check if plugins/installed directory exists before traversing
	if !(sdkfs.Exists(installedPluginsPath)) {
		return
	}

	// this lists all directories inside paths.PluginsDir/installed
	var list []string
	if err := sdkfs.LsDirs(installedPluginsPath, &list, false); err != nil {
		fmt.Printf("Error listing directories in %s: %v\n", installedPluginsPath, err)
		return
	}

	for _, p := range list {
		if err := ValidateInstallPath(p); err == nil {
			pluginDirs = append(pluginDirs, p)
		} else {
			fmt.Println("Error validating install path: ", p, err)
		}
	}

	return
}

func GetMetaDataPath(pkg string) string {
	return filepath.Join(sdkpaths.ConfigDir, "plugins", pkg, "metadata.json")
}

func WriteMetadata(def sdkpkg.PluginSrcDef, pkg string) error {
	cfg, err := config.ReadPluginsConfig()
	if err != nil {
		return err
	}

	meta := sdkpkg.PluginMetadata{
		Package: pkg,
		Def:     def,
	}

	for i, m := range cfg.Metadata {
		if m.Package == pkg {
			cfg.Metadata[i] = meta
			return config.WritePluginsConfig(cfg)
		}
	}

	cfg.Metadata = append(cfg.Metadata, meta)

	return config.WritePluginsConfig(cfg)
}

func ReadMetadata(pkg string) (metadata sdkpkg.PluginMetadata, err error) {
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

func IsSrcDefInstalled(def sdkpkg.PluginSrcDef) bool {
	installPath, ok := FindDefInstallPath(def)
	if !ok {
		return false
	}

	err := ValidateInstallPath(installPath)
	return err == nil
}

func InstalledPluginsList() (list []sdkpkg.PluginMetadata) {
	cfg, err := config.ReadPluginsConfig()
	if err != nil {
		return list
	}

	list = []sdkpkg.PluginMetadata{}
	for _, m := range cfg.Metadata {
		if IsSrcDefInstalled(m.Def) {
			list = append(list, m)
		}
	}

	return
}

func NeedsRecompile(def sdkpkg.PluginSrcDef) bool {
	if env.GO_ENV == env.ENV_DEV && (def.Src == sdkpkg.PluginSrcLocal || def.Src == sdkpkg.PluginSrcSystem) {
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
	requiredFiles := []string{"plugin.json", "go.mod", "main.go", "LICENSE.txt"}

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

func FindDefInstallPath(def sdkpkg.PluginSrcDef) (installPath string, ok bool) {
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

func GetAuthorNameFromGitUrl(def sdkpkg.PluginSrcDef) string {
	return strings.Split(strings.TrimPrefix(def.GitURL, "https://github.com/"), "/")[0]
}

func GetRepoFromGitUrl(def sdkpkg.PluginSrcDef) string {
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
			for _, pluginPath := range list {
				if err := ValidateSrcPath(pluginPath); err == nil {
					pluginDirs = append(pluginDirs, pluginPath)
				}
			}
		}
	}

	return pluginDirs
}
