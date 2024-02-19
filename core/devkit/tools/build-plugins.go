package tools

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	sdkfs "github.com/flarehotspot/core/sdk/utils/fs"
	sdkpaths "github.com/flarehotspot/core/sdk/utils/paths"
)

func BuildPlugin(dir string) error {
	if dir == "" {
		return errors.New("No plugin path provided")
	}

	mainFile := filepath.Join(dir, "main.go")
	pluginSo := filepath.Join(dir, "plugin.so")
	if !sdkfs.Exists(mainFile) && sdkfs.Exists(pluginSo) {
		fmt.Println("Plugin already built: " + sdkpaths.Strip(dir) + "/plugin.so")
		return nil
	}

	gofile := "main.go"
	outfile := "plugin.so"
	err := BuildGoModule(gofile, outfile, dir, "-buildmode=plugin")

	return err
}

func BuildGoModule(gofile string, outfile string, workDir string, extraArgs ...string) error {
	fmt.Println("Building go module: " + sdkpaths.Strip(filepath.Join(workDir, gofile)))

	goBin := GoBin()
	buildArgs := BuildArgs()
	buildArgs = append(buildArgs, extraArgs...)

	buildCmd := []string{"build"}
	buildCmd = append(buildCmd, buildArgs...)
	buildCmd = append(buildCmd, "-o", outfile, gofile)

	cmd := exec.Command(goBin, buildCmd...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = workDir

	fmt.Printf("Executing: %s %s\n", goBin, strings.Join(buildCmd, " "))
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error building go module " + sdkpaths.Strip(workDir) + ":" + err.Error())
		return err
	}

	fmt.Println("Module built successfully: " + sdkpaths.Strip(filepath.Join(workDir, outfile)))
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
