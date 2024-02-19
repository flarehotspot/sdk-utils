package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/flarehotspot/core/devkit"
	"github.com/flarehotspot/core/devkit/tools"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(Usage())
		return
	}

	command := os.Args[1]

	switch command {
	case "make-mono":
		tools.CreateMonoFiles()
		return

	case "create-devkit":
		devkit.CreateDevkit()
		return

	case "build-cli":
		tools.BuildFlareCLI()
		return

	case "server":
		Server()
		return
	}
	fmt.Println(Usage())
}

func Server() {
	buildArgs := tools.BuildArgs()
	runCmd := []string{"run"}
	runCmd = append(runCmd, buildArgs...)
	runCmd = append(runCmd, "core/internal/main/main.go")
	fmt.Println("Executing: go " + strings.Join(runCmd, " "))

	cmd := exec.Command("go", runCmd...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func Usage() string {
	return `
Available commands:
    server          Start the flare server

    make-mono       Create mono-repo files

    create-devkit   Generate devkit files

    build-cli       Build the flare executable CLI
`
}
