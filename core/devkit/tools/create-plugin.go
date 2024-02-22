package tools

import (
	"fmt"
	"os"
	"path/filepath"

	sdkplugin "github.com/flarehotspot/sdk/api/plugin"
	sdkfs "github.com/flarehotspot/sdk/utils/fs"
	sdkstr "github.com/flarehotspot/sdk/utils/strings"
)

func CreatePlugin(pkg string, name string, desc string) {
	info := sdkplugin.PluginInfo{
		Name:        name,
		Package:     pkg,
		Description: desc,
		Version:     "0.0.1",
	}

	goVersion, err := GoShortVersion()
	if err != nil {
		panic(err)
	}

	pluginDir := filepath.Join("plugins", pkg)
	if sdkfs.Exists(pluginDir) {
		fmt.Printf("Plugin already exists at %s\n", pluginDir)
		os.Exit(1)
	}

	sdkfs.EnsureDir(pluginDir)

	modPath := filepath.Join(pluginDir, "go.mod")
	modUri := fmt.Sprintf("github.com/your-account/my-plugin-%s", sdkstr.Rand(8))
	goMod := fmt.Sprintf("module %s\n\ngo %s", modUri, goVersion)
	err = os.WriteFile(modPath, []byte(goMod), 0644)
	if err != nil {
		panic(err)
	}

	pluginJson := filepath.Join(pluginDir, "plugin.json")
	err = sdkfs.WriteJson(pluginJson, &info)
	if err != nil {
		panic(err)
	}

	mainPath := filepath.Join(pluginDir, "main.go")

	goMain := `
package main

import (
    sdkplugin "github.com/flarehotspot/sdk/api/plugin"
)

func main() {}

func Init(api sdkplugin.PluginApi) {
    // Your plugin code here
}
    `

	err = os.WriteFile(mainPath, []byte(goMain), 0644)
	if err != nil {
		panic(err)
	}

	gitIgnorePath := filepath.Join(pluginDir, ".gitignore")
	gitIgnore := "# Ignore main_mono.go\nmain_mono.go\n"
	err = os.WriteFile(gitIgnorePath, []byte(gitIgnore), 0644)
	if err != nil {
		panic(err)
	}

	CreateGoWorkspace()

	fmt.Printf("Plugin created at %s\n", pluginDir)
}
