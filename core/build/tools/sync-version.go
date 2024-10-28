package tools

import (
	"core/internal/utils/pkg"
	"fmt"

	sdkfs "github.com/flarehotspot/go-utils/fs"
)

func SyncVersion() {
	version := pkg.CoreInfo().Version
	packageJson := "package.json"
	var pkg map[string]interface{}
	err := sdkfs.ReadJson(packageJson, &pkg)
	if err != nil {
		panic(err)
	}
	pkg["version"] = version
	err = sdkfs.WriteJson(packageJson, pkg)
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated package.json version to", version)
}
