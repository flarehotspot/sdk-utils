package tools

import (
	"core/internal/utils/pkg"
	"fmt"
	"os"
	"path/filepath"

	sdkplugin "sdk/api/plugin"
	sdkfs "sdk/utils/fs"
	sdkruntime "sdk/utils/runtime"
	sdkstr "sdk/utils/strings"
)

func CreatePlugin(pack string, name string, desc string) {
	info := sdkplugin.PluginInfo{
		Name:        name,
		Package:     pack,
		Description: desc,
		Version:     "0.0.1",
	}

	goVersion := sdkruntime.GOVERSION
	pluginDir := filepath.Join("plugins/local", pack)
	if sdkfs.Exists(pluginDir) {
		fmt.Printf("Plugin already exists at %s\n", pluginDir)
		os.Exit(1)
	}

	sdkfs.EnsureDir(pluginDir)

	modPath := filepath.Join(pluginDir, "go.mod")
	modUri := fmt.Sprintf("com.mydomain.%s", sdkstr.Rand(8))
	goMod := fmt.Sprintf("module %s\n\ngo %s", modUri, goVersion)
	if err := os.WriteFile(modPath, []byte(goMod), 0644); err != nil {
		panic(err)
	}

	pluginJson := filepath.Join(pluginDir, "plugin.json")
	if err := sdkfs.WriteJson(pluginJson, &info); err != nil {
		panic(err)
	}

	mainPath := filepath.Join(pluginDir, "main.go")

	goMain := `
package main

import (
    sdkplugin "sdk/api/plugin"
)

func main() {}

func Init(api sdkplugin.PluginApi) {
    // Your plugin code here
}
    `

	if err := os.WriteFile(mainPath, []byte(goMain), 0644); err != nil {
		panic(err)
	}

	gitIgnorePath := filepath.Join(pluginDir, ".gitignore")
	gitIgnore := "# Ignore main_mono.go\nmain_mono.go\n"
	if err := os.WriteFile(gitIgnorePath, []byte(gitIgnore), 0644); err != nil {
		panic(err)
	}

	pkg.CreateGoWorkspace()

	fmt.Printf("\n\nPlugin created at %s\n", pluginDir)
}
