package main

import (
	"core/env"
	"core/internal/utils/pkg"
	"fmt"
	"os"
	"os/exec"

	sdkutils "github.com/flarehotspot/sdk-utils"
)

func main() {
	goBin := pkg.GoBin()
	buildArgs := sdkutils.DefaultGoBuildArgs(env.BuildTags)
	runCmd := []string{"run"}
	runCmd = append(runCmd, buildArgs...)
	runCmd = append(runCmd, "main/main.go")

	commandstr := goBin
	for _, arg := range runCmd {
		commandstr += " " + arg
	}

	fmt.Printf("Executing: %s\n", commandstr)

	cmd := exec.Command("sh", "-c", commandstr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
