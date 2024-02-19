package tools

import (
	"fmt"
	"os"
	"path/filepath"

	sdkfs "github.com/flarehotspot/core/sdk/utils/fs"
)

func CreateGoWorkspace() {
	goVersion, err := GoShortVersion()
	if err != nil {
		panic(err)
	}

	goWork := fmt.Sprintf(`go %s

use (
    ./core`, goVersion)

	pluginSearchPaths := []string{"plugins"}

	for _, searchPath := range pluginSearchPaths {
		if sdkfs.Exists(searchPath) {
			entries, err := os.ReadDir(searchPath)
			if err != nil {
				continue
			}

			for _, entry := range entries {
				pluginDir := filepath.Join(searchPath, entry.Name())
				jsonFile := filepath.Join(pluginDir, "plugin.json")
				if entry.IsDir() && sdkfs.Exists(jsonFile) {
					goWork += "\n    ./" + pluginDir
				}
			}
		}
	}

	goWork += "\n)"

	if err = os.WriteFile(filepath.Join("go.work"), []byte(goWork), 0644); err != nil {
		panic(err)
	}

    fmt.Printf("go.work file created: \n%s\n", goWork)
}
