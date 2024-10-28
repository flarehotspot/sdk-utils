package main

import (
	"core/internal/utils/pkg"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	goBin := pkg.GoBin()
	buildArgs := pkg.BuildArgs()
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
