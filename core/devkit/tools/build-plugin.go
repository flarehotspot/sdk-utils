package tools

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	sdkfs "github.com/flarehotspot/sdk/utils/fs"
	sdkpaths "github.com/flarehotspot/sdk/utils/paths"
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

	goBin := GoBin()
	buildArgs := BuildArgs()

	buildCmd := []string{"build", "-buildmode=plugin"}
	buildCmd = append(buildCmd, buildArgs...)
	buildCmd = append(buildCmd, "-o", "plugin.so", "main.go")

	cmd := exec.Command(goBin, buildCmd...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Building plugin: " + sdkpaths.Strip(dir))
	err := cmd.Run()
	if err != nil {
		return errors.New("Error building plugin " + sdkpaths.Strip(dir) + ":" + err.Error())
	}

	fmt.Println("Plugin built successfully: " + sdkpaths.Strip(dir) + "/plugin.so")
	return nil
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
