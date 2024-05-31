package tools

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"core/env"
	sdkcfg "sdk/api/config"
	sdkfs "sdk/utils/fs"
	sdkpaths "sdk/utils/paths"
	sdkstr "sdk/utils/strings"
)

var (
	coreReleaseDir = filepath.Join(sdkpaths.AppDir, "output/devkit", fmt.Sprintf("devkit-%s-%s", CoreInfo().Version, runtime.GOARCH))
	devkitFiles    = []string{
		"go",
		"bin",
		"config/.defaults",
		"core/go.mod",
		"core/plugin.so",
		"core/plugin.json",
		"core/resources",
		"core/go-version",
		"main/go.mod",
		"sdk",
	}
)

func CreateDevkit() {
	PrepareCleanup()
	InstallGo("./go")
	BuildFlareCLI()
	BuildCore()
	CloneDefaultPlugins(coreReleaseDir)
	CopyDevkitFiles()
	CopyDevkitExtras()
	CopyDefaultWorksapce()
	CreateApplicationConfig()
	ZipDevkitRelease()
	CleanUpRelease()
}

func CreateApplicationConfig() {
	cfgPath := filepath.Join(coreReleaseDir, "config/application.json")
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
	for _, entry := range devkitFiles {
		srcPath := filepath.Join(sdkpaths.AppDir, entry)
		destPath := filepath.Join(coreReleaseDir, entry)
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
	extrasPath := filepath.Join(sdkpaths.AppDir, "core/build/devkit/extras")
	fmt.Printf("Copying:  %s -> %s\n", sdkpaths.StripRoot(extrasPath), sdkpaths.StripRoot(coreReleaseDir))
	err := sdkfs.CopyDir(extrasPath, coreReleaseDir, nil)
	if err != nil {
		panic(err)
	}
}

func CopyDefaultWorksapce() {
	dst := filepath.Join(coreReleaseDir, "go.work")
	def := "go.work.default"
	fmt.Println("Copying: ", sdkpaths.StripRoot(def), " -> ", sdkpaths.StripRoot(dst))
	sdkfs.CopyFile(def, dst)
}

func ZipDevkitRelease() {
	basename := filepath.Base(coreReleaseDir) + "-" + sdkstr.Slugify(env.BuildTags, "-") + ".zip"
	dir := filepath.Dir(coreReleaseDir)
	zipFile := filepath.Join(dir, basename)
	fmt.Printf("Zipping devkit release: %s...\n", zipFile)
	cmd := exec.Command("zip", "-r", zipFile, ".")
	cmd.Dir = coreReleaseDir
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
	fmt.Printf("Cleaning up release directory: %s\n", sdkpaths.StripRoot(coreReleaseDir))
	os.RemoveAll(coreReleaseDir)
}
