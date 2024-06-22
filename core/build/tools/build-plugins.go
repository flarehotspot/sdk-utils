package tools

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	sdkplugin "sdk/api/plugin"
	sdkfs "sdk/utils/fs"
	sdkpaths "sdk/utils/paths"
	sdkruntime "sdk/utils/runtime"
)

type GoBuildArgs struct {
	WorkDir   string
	Env       []string
	ExtraArgs []string
}

func BuildPlugin(pluginSrcDir string, workdir string) error {
	if pluginSrcDir == "" {
		return errors.New("Build plugin error: no plugin source path")
	}

	if workdir == "" {
		return errors.New("Build plugin error: no build path")
	}

	var info sdkplugin.PluginInfo

	pluginSoPath := filepath.Join(pluginSrcDir, "plugin.so")
	if err := sdkfs.ReadJson(filepath.Join(pluginSrcDir, "plugin.json"), &info); err != nil {
		return err
	}

	buildpath := filepath.Join(workdir, "plugins", info.Package)

	if sdkfs.Exists(pluginSoPath) {
		if err := os.Remove(pluginSoPath); err != nil {
			return err
		}
	}

	if err := sdkfs.EmptyDir(workdir); err != nil {
		return err
	}
	defer os.RemoveAll(workdir)

	if err := sdkfs.EnsureDir(filepath.Join(workdir, "plugins")); err != nil {
		return err
	}

	if err := os.Symlink(pluginSrcDir, buildpath); err != nil {
		return err
	}

	if err := os.Symlink(filepath.Join(sdkpaths.AppDir, "sdk"), filepath.Join(workdir, "sdk")); err != nil {
		return err
	}

	goWork := fmt.Sprintf(`
go %s

use (
    ./sdk
    ./plugins/%s
)
    `, sdkruntime.GO_SHORT_VERSION, info.Package)

	if err := os.WriteFile(filepath.Join(workdir, "go.work"), []byte(goWork), sdkfs.PermFile); err != nil {
		return err
	}

	gofile := "main.go"
	outfile := "plugin.so"
	args := &GoBuildArgs{WorkDir: buildpath, ExtraArgs: []string{"-buildmode=plugin"}}
	if err := BuildGoModule(gofile, outfile, args); err != nil {
		return err
	}

	pluginSoOut := filepath.Join(buildpath, "plugin.so")
	fmt.Printf("Copying '%s' to '%s'\n", sdkpaths.StripRoot(pluginSoOut), sdkpaths.StripRoot(pluginSoPath))
	if err := sdkfs.CopyFile(pluginSoOut, pluginSoPath); err != nil {
		return err
	}

	return nil
}

func BuildGoModule(gofile string, outfile string, params *GoBuildArgs) error {
	if params == nil {
		params = &GoBuildArgs{}
	}

	fmt.Println("Building go module: " + sdkpaths.StripRoot(filepath.Join(params.WorkDir, gofile)))

	goBin := GoBin()
	buildArgs := BuildArgs()
	buildArgs = append(buildArgs, params.ExtraArgs...)

	buildCmd := []string{"build"}
	buildCmd = append(buildCmd, buildArgs...)
	buildCmd = append(buildCmd, "-o", outfile, gofile)

	fmt.Printf(`Building go module in path "%s"...`+"\n", sdkpaths.StripRoot(params.WorkDir))
	cmd := exec.Command(goBin, buildCmd...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), params.Env...)
	cmd.Dir = params.WorkDir

	fmt.Printf("Executing: %s %s\n", goBin, strings.Join(buildCmd, " "))
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error building go module " + sdkpaths.StripRoot(params.WorkDir) + ":" + err.Error())
		return err
	}

	fmt.Println("Module built successfully: " + sdkpaths.StripRoot(filepath.Join(params.WorkDir, outfile)))
	return nil
}

func BuildCore() {
	workdir := filepath.Join(sdkpaths.TmpDir, "builds/core")
	if err := BuildPlugin(sdkpaths.CoreDir, workdir); err != nil {
		panic(err)
	}
}

func BuildAllPlugins() error {
	pluginPaths := PluginPathList()
	for _, pluginPath := range pluginPaths {
		workdir := filepath.Join(sdkpaths.TmpDir, "builds", filepath.Base(pluginPath))
		if err := BuildPlugin(pluginPath, workdir); err != nil {
			return err
		}
	}
	return nil
}
