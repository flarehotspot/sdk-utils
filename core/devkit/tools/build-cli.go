package tools

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	sdkfs "github.com/flarehotspot/core/sdk/utils/fs"
)

func BuildFlareCLI() {
	fmt.Println("Building flare CLI...")
	goBin := GoBin()
	sdkfs.EnsureDir("bin")
	buildArgs := BuildArgs()

	os.RemoveAll("bin")

	cliPath := "bin/flare"
	if runtime.GOOS == "windows" {
		cliPath += ".exe"
	}

	buildCmd := []string{"build"}
	buildCmd = append(buildCmd, buildArgs...)
	buildCmd = append(buildCmd, "-o", cliPath, "core/devkit/cli/flare.go")

	cmd := exec.Command(goBin, buildCmd...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Flare CLI built at: %s\n", cliPath)
}
