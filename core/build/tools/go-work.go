package tools

import (
	"core/internal/utils/pkg"
	"fmt"
	"os"
	"path/filepath"

	sdkfs "sdk/utils/fs"
	sdkruntime "sdk/utils/runtime"
)

func CreateGoWorkspace() {
	goVersion := sdkruntime.GO_VERSION
	goWork := fmt.Sprintf(`go %s

use (
    ./core
    ./sdk
    ./main`, goVersion)

	pluginSearchPaths := []string{"plugins/system", "plugins/local"}

	for _, searchPath := range pluginSearchPaths {
		if sdkfs.Exists(searchPath) {
			entries, err := os.ReadDir(searchPath)
			if err != nil {
				continue
			}

			for _, entry := range entries {
				pluginDir := filepath.Join(searchPath, entry.Name())
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
