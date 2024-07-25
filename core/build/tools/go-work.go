package tools

import (
	"fmt"
	"os"
	"path/filepath"

	sdkfs "sdk/utils/fs"
	sdkruntime "sdk/utils/runtime"
)

func CreateGoWorkspace() {
	goVersion := sdkruntime.GO_SHORT_VERSION
	goWork := fmt.Sprintf(`go %s

use (
    ./core
    ./sdk
    ./main`, goVersion)

	pluginSearchPaths := []string{"plugins/local"}

	for _, searchPath := range pluginSearchPaths {
		if sdkfs.Exists(searchPath) {
			entries, err := os.ReadDir(searchPath)
			if err != nil {
				continue
			}

			for _, entry := range entries {
                pluginDir := filepath.Join(searchPath, entry.Name())
				// pluginDir := searchPath + "/" + entry.Name()
				jsonFile := filepath.Join(pluginDir, "plugin.json")
				if entry.IsDir() && sdkfs.Exists(jsonFile) {
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
