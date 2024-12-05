package tools

import (
	"core/internal/utils/pkg"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	sdkfs "github.com/flarehotspot/go-utils/fs"
	sdkpaths "github.com/flarehotspot/go-utils/paths"
	sdkpkg "github.com/flarehotspot/go-utils/pkg"
	sdkstr "github.com/flarehotspot/go-utils/strings"
)

type PluginModule struct {
	PluginImportVar   string
	PluginModuleUri   string
	PluginPackageName string
}

func CreateMonoFiles() {
	CreateGoWorkspace()

	localDefs := pkg.LocalPluginDefs()
	systemDefs := pkg.SystemPluginsDefs()

	pluginDirs := []string{filepath.Join(sdkpaths.AppDir, "core")}
	for _, def := range append(systemDefs, localDefs...) {
		pluginDirs = append(pluginDirs, def.LocalPath)
	}

	for _, p := range pluginDirs {
		MakePluginMainMono(p)
	}

	MakePluginInitMono()
}

func MakePluginInitMono() {
	pluginPaths := []string{"core"}
	pluginDirs := []string{}

	localDefs := pkg.LocalPluginDefs()
	systemDefs := pkg.SystemPluginsDefs()
	for _, def := range append(systemDefs, localDefs...) {
		pluginDirs = append(pluginDirs, def.LocalPath)
	}

	pluginPaths = append(pluginPaths, pluginDirs...)
	coreInfo := pkg.GetCoreInfo()

	pluginMods := []PluginModule{}
	for _, dir := range pluginDirs {
		modVar := sdkstr.Slugify(filepath.Base(dir), "_")
		modPath := getGoModule(dir)
		pkgName := getPackage(dir)
		mod := PluginModule{modVar, modPath, pkgName}
		pluginMods = append(pluginMods, mod)
	}

	importModules := ""
	for _, mod := range pluginMods {
		importModules += fmt.Sprintf("\n\t"+`%s "%s"`, mod.PluginImportVar, mod.PluginModuleUri)
	}

	pluginSwitchCases := ""
	for _, mod := range pluginMods {
		pluginSwitchCases += fmt.Sprintf("\n\t\tcase \"%s\":\n\t\t\t%s.Init(p)", mod.PluginPackageName, mod.PluginImportVar)
	}

	pluginMonoInit := fmt.Sprintf(`//go:build mono

%s

package plugins
import (
    "log"
    %s
)

func (p *PluginApi) Init() error {
    switch p.Pkg() {
        case "%s":
            log.Println("core package, skipping plugin.Init()...")
%s
        default:
            log.Println("Unable to load plugin: " + p.dir)
    }
    return nil
}`, AUTO_GENERATED_HEADER, importModules, coreInfo.Package, pluginSwitchCases)

	pluginInitMonoPath := filepath.Join("core/internal/plugins/plugin-init_mono.go")
	err := os.WriteFile(pluginInitMonoPath, []byte(pluginMonoInit), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println(pluginInitMonoPath, "has been created.")
}

func getGoModule(pluginDir string) string {
	goModFile := filepath.Join(pluginDir, "go.mod")
	modContent, err := sdkfs.ReadFile(goModFile)
	if err != nil {
		panic(err)
	}

	regx := regexp.MustCompile(`module\s+([\w\/.-]+)`)
	matches := regx.FindStringSubmatch(string(modContent))
	if len(matches) > 0 && len(matches[0]) > 0 {
		return strings.Split(matches[0], " ")[1]
	}

	panic("Error: go.mod file does not contain module name")
}

func getPackage(pluginDir string) string {
	info, err := sdkpkg.GetInfoFromPath(pluginDir)
	if err != nil {
		panic(err)
	}
	return info.Package
}
