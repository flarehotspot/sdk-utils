package devkit

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/flarehotspot/core/devkit/tools"
	"github.com/flarehotspot/core/env"
	sdkcfg "github.com/flarehotspot/sdk/api/config"
	sdkfs "github.com/flarehotspot/sdk/utils/fs"
	sdkpaths "github.com/flarehotspot/sdk/utils/paths"
	sdkstr "github.com/flarehotspot/sdk/utils/strings"
)

const (
	GOARCH = runtime.GOARCH
)

var (
	coreInfo     = tools.CoreInfo()
	RELEASE_DIR  = filepath.Join(sdkpaths.AppDir, "devkit-release", fmt.Sprintf("devkit-%s-%s", coreInfo.Version, GOARCH))
	DEVKIT_FILES = []string{
		"go",
		"bin",
		"sdk",
		"main/go.mod",
		"config/.defaults",
		"core/go.mod",
		"core/plugin.so",
		"core/plugin.json",
		"core/resources",
		"core/go-version",
	}
)

func CreateDevkit() {
	PrepareCleanup()
	tools.InstallGo("./go")
	tools.BuildFlareCLI()
	tools.BuildCore()
	tools.CloneDefaultPlugins(RELEASE_DIR)
	CopyDevkitFiles()
	CopyDevkitExtras()
	CopyDefaultWorksapce()
	CreateApplicationConfig()
	ZipDevkitRelease()
	CleanUpRelease()
}

func CreateApplicationConfig() {
	cfgPath := filepath.Join(RELEASE_DIR, "config/application.json")
	appcfg := sdkcfg.AppCfg{
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

	fmt.Println("Application config created: ", sdkpaths.StripRoot(cfgPath))
}

func CopyDevkitFiles() {
	for _, entry := range DEVKIT_FILES {
		srcPath := filepath.Join(sdkpaths.AppDir, entry)
		destPath := filepath.Join(RELEASE_DIR, entry)
		fmt.Println("Copying: ", sdkpaths.StripRoot(srcPath), " -> ", sdkpaths.StripRoot(destPath))

		if sdkfs.IsFile(srcPath) {
			err := sdkfs.CopyFile(srcPath, destPath)
			if err != nil {
				panic(err)
			}

		} else if sdkfs.IsDir(srcPath) {
			err := sdkfs.CopyDir(srcPath, destPath, nil)
			if err != nil {
				panic(err)
			}
		} else {
			fmt.Println("Warning: ", srcPath, " is not a file or directory. Skipping.")
		}
	}
}

func CopyDevkitExtras() {
	extrasPath := filepath.Join(sdkpaths.AppDir, "build/devkit/extras")
	fmt.Printf("Copying:  %s -> %s\n", sdkpaths.StripRoot(extrasPath), sdkpaths.StripRoot(RELEASE_DIR))
	err := sdkfs.CopyDir(extrasPath, RELEASE_DIR, nil)
	if err != nil {
		panic(err)
	}
}

func CopyDefaultWorksapce() {
	dst := filepath.Join(RELEASE_DIR, "go.work")
	def := "go.work.default"
	fmt.Println("Copying: ", sdkpaths.StripRoot(def), " -> ", sdkpaths.StripRoot(dst))
	sdkfs.CopyFile(def, dst)
}

func ZipDevkitRelease() {
	basename := filepath.Base(RELEASE_DIR) + "-" + sdkstr.Slugify(env.BuildTags, "-") + ".zip"
	dir := filepath.Dir(RELEASE_DIR)
	zipFile := filepath.Join(dir, basename)
	fmt.Printf("Zipping devkit release: %s...\n", zipFile)
	cmd := exec.Command("zip", "-r", zipFile, ".")
	cmd.Dir = RELEASE_DIR
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Devkit created: ", sdkpaths.StripRoot(zipFile))
}

func PrepareCleanup() {
	dirsToRemove := []string{"devkit-release"}
	for _, dir := range dirsToRemove {
		fmt.Println("Removing: ", filepath.Join(sdkpaths.AppDir, dir))
		err := os.RemoveAll(filepath.Join(sdkpaths.AppDir, dir))
		if err != nil {
			fmt.Println("Error removing: ", err)
			panic(err)
		}
	}
}

func CleanUpRelease() {
	fmt.Printf("Cleaning up release directory: %s\n", sdkpaths.StripRoot(RELEASE_DIR))
	os.RemoveAll(RELEASE_DIR)
}
