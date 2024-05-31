package main

import (
	"core/build/tools"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	goBin := tools.GoBin()
	serverBin := "./bin/debug-server"
	buildArgs := tools.BuildArgs()
	runCmd := []string{"build"}
	runCmd = append(runCmd, buildArgs...)
	runCmd = append(runCmd, "-o", serverBin, "main/main.go")
	fmt.Printf("Executing: %s %s\n", goBin, strings.Join(runCmd, " "))

	cmd := exec.Command(goBin, runCmd...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	fileInfo, err := os.Stat(serverBin)
	if err != nil {
		panic(err)
	}

	executableMode := fileInfo.Mode() | 0111
	err = os.Chmod(serverBin, executableMode)
	if err != nil {
		panic(err)
	}

	cmd = exec.Command(serverBin)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

}
