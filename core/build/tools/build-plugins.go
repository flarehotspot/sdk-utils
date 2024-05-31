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
)

type GoBuildArgs struct {
	WorkDir   string
	Env       []string
	ExtraArgs []string
}

func BuildPlugin(dir string) error {
	if dir == "" {
		return errors.New("No plugin path provided")
	}

	var info sdkplugin.PluginInfo

	mainFile := filepath.Join(dir, "main.go")
	pluginSo := filepath.Join(dir, "plugin.so")
	err := sdkfs.ReadJson(filepath.Join(dir, "plugin.json"), &info)
	if err != nil {
		return err
	}

	if !sdkfs.Exists(mainFile) && sdkfs.Exists(pluginSo) {
		fmt.Printf("Plugin '%s' is already built. Skipping...\n", info.Package)
		return nil
	}

	gofile := "main.go"
	outfile := "plugin.so"
	args := &GoBuildArgs{WorkDir: dir, ExtraArgs: []string{"-buildmode=plugin"}}
	err = BuildGoModule(gofile, outfile, args)

	return err
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

func BuildCore() error {
	return BuildPlugin(sdkpaths.CoreDir)
}

func BuildAllPlugins() error {
	pluginPaths := PluginPathList()
	for _, pluginPath := range pluginPaths {
		if err := BuildPlugin(pluginPath); err != nil {
			return err
		}
	}
	return nil
}
