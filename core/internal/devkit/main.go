package devkit

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/flarehotspot/flarehotspot/core/config"
	"github.com/flarehotspot/flarehotspot/core/internal/tools"
	sdkfs "github.com/flarehotspot/sdk/utils/fs"
	sdkpaths "github.com/flarehotspot/sdk/utils/paths"
	sdkstr "github.com/flarehotspot/sdk/utils/strings"
	sdktools "github.com/flarehotspot/sdk/utils/tools"
)

const GOARCH = runtime.GOARCH

var (
	coreInfo     = sdktools.CoreInfo()
	RELEASE_DIR  = filepath.Join(sdkpaths.AppDir, "devkit-release", fmt.Sprintf("devkit-%s-%s", coreInfo.Version, GOARCH))
	DEVKIT_FILES = []string{
		"main/go.mod",
		"main/main.app",
		"config/.defaults",
		"core/plugin.so",
		"core/go.mod",
		"core/go.sum",
		"core/plugin.json",
		"core/resources",
		"core/go-version",
		"package.json",
		"package-lock.json",
		"sdk",
		"system",
	}
)

func CreateDevkit() {
	CleanUpDevkit()
	tools.BuildCore()
	tools.BuildMain()
	CopyDevkitFiles()
	CopyDevkitExtras()
	CopyDefaultWorksapce()
	CreateApplicationConfig()
}

func CreateApplicationConfig() {
	cfgPath := filepath.Join(RELEASE_DIR, "config/application.json")
	appcfg := config.AppConfig{
		Lang:     "en",
		Currency: "php",
		Secret:   sdkstr.Rand(16),
	}

	b, err := json.MarshalIndent(appcfg, "", "  ")
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(cfgPath, b, 0644); err != nil {
		panic(err)
	}

	fmt.Println("Application config created: ", sdkpaths.Strip(cfgPath))
}

func CopyDevkitFiles() {
	for _, entry := range DEVKIT_FILES {
		srcPath := filepath.Join(sdkpaths.AppDir, entry)
		destPath := filepath.Join(RELEASE_DIR, entry)
		fmt.Println("Copying: ", sdkpaths.Strip(srcPath), " -> ", sdkpaths.Strip(destPath))

		if sdkfs.IsFile(srcPath) {
			err := sdkfs.CopyFile(srcPath, destPath)
			if err != nil {
				panic(err)
			}

		} else if sdkfs.IsDir(srcPath) {
			copyOpts := sdkfs.CopyOpts{Recursive: true}
			err := sdkfs.CopyDir(srcPath, destPath, &copyOpts)
			if err != nil {
				panic(err)
			}
		} else {
			fmt.Println("Warning: ", srcPath, " is not a file or directory. Skipping.")
		}
	}
}

func CopyDevkitExtras() {
	extrasPath := filepath.Join(sdkpaths.AppDir, "build/devkit-extras")
	opts := sdkfs.CopyOpts{Recursive: true}
    fmt.Printf("Copying:  %s -> %s\n", sdkpaths.Strip(extrasPath), sdkpaths.Strip(RELEASE_DIR))
	err := sdkfs.CopyDir(extrasPath, RELEASE_DIR, &opts)
	if err != nil {
		panic(err)
	}
}

func CopyDefaultWorksapce() {
	dst := filepath.Join(RELEASE_DIR, "go.work")
	def := "go.work.default"
	fmt.Println("Copying: ", sdkpaths.Strip(def), " -> ", sdkpaths.Strip(dst))
	sdkfs.CopyFile(def, dst)
}

func CleanUpDevkit() {
	dirsToRemove := []string{".cache/assets", ".tmp", "public", "devkit-release"}
	for _, dir := range dirsToRemove {
		fmt.Println("Removing: ", filepath.Join(sdkpaths.AppDir, dir))
		os.RemoveAll(filepath.Join(sdkpaths.AppDir, dir))
	}
}
