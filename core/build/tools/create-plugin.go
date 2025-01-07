package tools

import (
	"core/internal/utils/pkg"
	"fmt"
	"os"
	"path/filepath"

	sdkfs "github.com/flarehotspot/go-utils/fs"
	sdkpkg "github.com/flarehotspot/go-utils/pkg"
	sdkruntime "github.com/flarehotspot/go-utils/runtime"
	sdkstr "github.com/flarehotspot/go-utils/strings"
)

func CreatePlugin(pack string, name string, desc string) {
	info := sdkpkg.PluginInfo{
		Name:        name,
		Package:     pack,
		Description: desc,
		Version:     "0.0.1",
	}

	goVersion := sdkruntime.GO_SHORT_VERSION
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
	gitIgnore := `
.DS_Store
/node_modules
/resources/assets/dist
*.so
main_mono.go
*_templ.go
`
	if err := os.WriteFile(gitIgnorePath, []byte(gitIgnore), 0644); err != nil {
		panic(err)
	}

	licenseFile := filepath.Join(pluginDir, "LICENSE.txt")
	licenseTxt := `# No License Chosen

This software does not currently have a license.

By default, all rights are reserved. This means:
- You may view the code.
- You may not use, modify, or distribute this software for any purpose without explicit written permission from the copyright holder.

The license for this software is still under consideration and will be added in the future. Until then, please contact [YOUR CONTACT INFORMATION] for any inquiries about usage or licensing.
`
	if err := os.WriteFile(licenseFile, []byte(licenseTxt), sdkfs.PermFile); err != nil {
		panic(err)
	}

	CreateGoWorkspace()

	if err := pkg.ValidateSrcPath(pluginDir); err != nil {
		panic("Error validating newly created plugin: " + err.Error())
	}

	fmt.Printf("\n\nPlugin created at %s\n", pluginDir)
}
