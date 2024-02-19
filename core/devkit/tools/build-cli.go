package tools

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	sdkfs "github.com/flarehotspot/sdk/utils/fs"
)

func BuildFlareCLI() {
	fmt.Println("Building flare CLI...")
	goBin := GoBin()
	sdkfs.EnsureDir("bin")

	cliPath := "flare"
    if runtime.GOOS == "windows" {
        cliPath += ".exe"
    }

	cmd := exec.Command(goBin, "build", "-o", cliPath, "sdk/cli/flare.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Flare CLI built at: %s\n", cliPath)
}
