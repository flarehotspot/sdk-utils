package tools

import (
	"fmt"
	sdkfs "sdk/utils/fs"
)

func SyncVersion() {
	version := CoreInfo().Version
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
