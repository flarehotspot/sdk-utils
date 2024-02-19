package tools

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	sdkpaths "github.com/flarehotspot/sdk/utils/paths"
)

func BuildCore() error {
	return BuildPlugin(sdkpaths.CoreDir)
}

func BuildMain() error {
	goBin := GoBin()
	buildArgs := BuildArgs()
	mainDir := filepath.Join(sdkpaths.AppDir, "main")

	buildCmd := []string{"build"}
	buildCmd = append(buildCmd, buildArgs...)
	buildCmd = append(buildCmd, "-o", MainPath(), "main.go")

	cmd := exec.Command(goBin, buildCmd...)
	cmd.Dir = mainDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Building main: " + MainPath())
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error building main: " + err.Error())
		return err
	}

	fmt.Println("Main built successfully: " + MainPath())
	return nil
}

func BuildPlugins() error {
	return BuildAllPlugins()
}
