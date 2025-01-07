package tools

import (
	"core/internal/utils/pkg"
	"fmt"
	"log"
	"os"
	"path/filepath"

	sdkfs "github.com/flarehotspot/go-utils/fs"
	sdkpkg "github.com/flarehotspot/go-utils/pkg"
	sdkruntime "github.com/flarehotspot/go-utils/runtime"
)

func CreateGoWorkspace() {
	goVersion := sdkruntime.GO_SHORT_VERSION
	goWork := fmt.Sprintf(`go %s

use (
    ./main
    ./core
    ./sdk/api
    ./sdk/utils`, goVersion)

	// insert libs paths
	libs := []string{}
	if err := sdkfs.LsDirs("sdk/libs", &libs, false); err != nil {
		log.Println(err)
	}

	for _, lib := range libs {
		goWork += "\n    ./" + lib
	}

	// insert plugin paths
	pluginSearchPaths := []string{"plugins/system", "plugins/local"}
	for _, searchPath := range pluginSearchPaths {
		if sdkfs.Exists(searchPath) {
			var entries []string
			if err := sdkfs.LsDirs(searchPath, &entries, false); err != nil {
				continue
			}

			for _, entry := range entries {
				pluginDir, err := sdkpkg.FindPluginSrc(entry)
				if err != nil {
					fmt.Printf("%s is not a valid plugin path, skipping...\n", pluginDir)
					continue
				}

				if pkg.ValidateSrcPath(pluginDir) == nil {
					goWork += "\n    ./" + pluginDir
				}
			}
		}
	}

	goWork += "\n)"

	if err := os.WriteFile(filepath.Join("go.work"), []byte(goWork), 0644); err != nil {
		panic(err)
	}

	// fmt.Printf("go.work file created: \n%s\n", goWork)
	fmt.Println("go.work file created.")
}
