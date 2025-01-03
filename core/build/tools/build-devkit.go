package tools

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"core/env"
	"core/internal/utils/pkg"
	sdkcfg "sdk/api/config"

	sdkfs "github.com/flarehotspot/go-utils/fs"
	sdkpaths "github.com/flarehotspot/go-utils/paths"
	sdkruntime "github.com/flarehotspot/go-utils/runtime"
	sdkstr "github.com/flarehotspot/go-utils/strings"
	sdktargz "github.com/flarehotspot/go-utils/targz"

	"github.com/goccy/go-json"
)

var (
	devkitReleaseDir string
	devkitFiles      = []string{
		".go-version",
		"bin/flare",
		"config/.defaults",
		"core/go.mod",
		"core/plugin.so",
		"core/plugin.json",
		"core/resources",
		"go.work.default",
		"main/go.mod",
		"plugins/system",
		"scripts/install-tools.sh",
		"sdk",
	}
)

func init() {
	goversion := sdkruntime.GO_VERSION
	tags := sdkstr.Slugify(env.BuildTags, "-")
	devkitReleaseDir = filepath.Join(sdkpaths.AppDir, "output/devkit", fmt.Sprintf("devkit-%s-%s-go%s-%s", pkg.GetCoreInfo().Version, runtime.GOARCH, goversion, tags))
}

func CreateDevkit() {
	// Clean up output path
	if err := sdkfs.EmptyDir(filepath.Dir(devkitReleaseDir)); err != nil {
		panic(err)
	}

	// Build the bin/flare cli
	BuildFlareCLI()

	// Build core/plugin.so
	BuildCore()

	// Copy devkit files
	for _, entry := range devkitFiles {
		srcPath := filepath.Join(sdkpaths.AppDir, entry)
		destPath := filepath.Join(devkitReleaseDir, entry)
		fmt.Println("Copying: ", sdkpaths.StripRoot(srcPath), " -> ", sdkpaths.StripRoot(destPath))

		if err := sdkfs.Copy(srcPath, destPath); err != nil {
			panic(err)
		}
	}

	// Copy extra devkit files to the release directory
	extrasPath := filepath.Join(sdkpaths.AppDir, "core/build/devkit/extras")
	fmt.Printf("Copying:  %s -> %s\n", sdkpaths.StripRoot(extrasPath), sdkpaths.StripRoot(devkitReleaseDir))
	err := sdkfs.CopyDir(extrasPath, devkitReleaseDir, nil)
	if err != nil {
		panic(err)
	}

	// Generate default application config
	appConfigFile := filepath.Join(devkitReleaseDir, "config/application.json")
	appConfig := sdkcfg.AppCfg{
		Lang:     "en",
		Currency: "php",
		Secret:   sdkstr.Rand(16),
	}

	b, err := json.MarshalIndent(appConfig, "", "  ")
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(appConfigFile, b, 0644); err != nil {
		panic(err)
	}

	fmt.Println("Application config created: ", sdkpaths.StripRoot(appConfigFile))

	// Compress devkit release files
	tarname := filepath.Base(devkitReleaseDir) + ".tar.gz"
	dir := filepath.Dir(devkitReleaseDir)
	tarfile := filepath.Join(dir, tarname)
	err = sdktargz.TarGz(devkitReleaseDir, tarfile)
	if err != nil {
		panic(err)
	}

	fmt.Println("Devkit created: ", sdkpaths.StripRoot(tarfile))
}
