package main

import (
	"core/internal/utils/pkg"
	"fmt"
	"os"
	"os/exec"
	"strings"

	sdkpkg "github.com/flarehotspot/go-utils/pkg"
)

func main() {
	goBin := pkg.GoBin()
	buildArgs := sdkpkg.DefaultBuildArgs()
	runCmd := []string{"run"}
	runCmd = append(runCmd, buildArgs...)
	runCmd = append(runCmd, "main/main.go")
	fmt.Printf("Executing: %s %s\n", goBin, strings.Join(runCmd, " "))

	cmd := exec.Command(goBin, runCmd...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
