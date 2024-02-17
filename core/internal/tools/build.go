package tools

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	sdkpaths "github.com/flarehotspot/sdk/utils/paths"
	sdktools "github.com/flarehotspot/sdk/utils/tools"
)

func BuildCore() error {
	return sdktools.BuildPlugin(sdkpaths.CoreDir)
}

func BuildMain() error {
	goBin := sdktools.GoBin()
	buildArgs := sdktools.BuildArgs()
	mainDir := filepath.Join(sdkpaths.AppDir, "main")

	buildCmd := []string{"build"}
	buildCmd = append(buildCmd, buildArgs...)
	buildCmd = append(buildCmd, "-o", sdktools.MainPath(), "main.go")

	cmd := exec.Command(goBin, buildCmd...)
	cmd.Dir = mainDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Building main: " + sdktools.MainPath())
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error building main: " + err.Error())
		return err
	}

	fmt.Println("Main built successfully: " + sdktools.MainPath())
	return nil
}

func BuildPlugins() error {
	return sdktools.BuildAllPlugins()
}
