package tools

import (
	"core/env"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	sdkfs "sdk/utils/fs"
	sdkpaths "sdk/utils/paths"
	sdkstr "sdk/utils/strings"
	sdkzip "sdk/utils/zip"
)

var (
	coreReleaseDir   = ""
	coreReleaseFiles = []string{
		"bin/flare",
		"config/.defaults",
		"core/go-version",
		"core/go.mod",
		"core/plugin.json",
		"core/plugin.so",
		"core/resources",
		"main/go.mod",
		"sdk",
	}
)

func init() {
	goversion, err := GoVersion()
	if err != nil {
		panic(err)
	}

	tags := sdkstr.Slugify(env.BuildTags, "-")
	coreReleaseDir = filepath.Join(sdkpaths.AppDir, "output/core", fmt.Sprintf("core-%s-%s-go%s-%s", CoreInfo().Version, runtime.GOARCH, goversion, tags))
}

func CreateRelease() {
	fmt.Println("Cleaning up", sdkpaths.StripRoot(coreReleaseDir), "...")
	sdkfs.RmEmpty(coreReleaseDir)

	BuildFlareCLI()
	BuildCore()
	CopyCoreReleaseFiles()
	ZipCoreRelease()
}

func CopyCoreReleaseFiles() {
	for _, entry := range coreReleaseFiles {
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
			panic("Unknown file type: " + srcPath)
		}
	}
}

func ZipCoreRelease() {
	basename := filepath.Base(coreReleaseDir) + ".zip"
	dir := filepath.Dir(coreReleaseDir)
	zipFile := filepath.Join(dir, basename)

	fmt.Printf("Zipping core release: %s...\n", sdkpaths.StripRoot(zipFile))
	if err := sdkzip.Zip(coreReleaseDir, zipFile); err != nil {
		panic(err)
	}

	outputTxtPath := filepath.Join(sdkpaths.AppDir, "output/core-release.txt")
	if err := os.WriteFile(outputTxtPath, []byte(zipFile), sdkfs.PermFile); err != nil {
		panic(err)
	}

	fmt.Println("Devkit created: ", sdkpaths.StripRoot(zipFile))
}
