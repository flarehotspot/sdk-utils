package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"core/devkit"
	"core/devkit/tools"
	sdkfs "sdk/utils/fs"
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

	case "create-migration":
		name, err := tools.AskCmdInput("Enter migration name, e.g. create_users_table")
		if err != nil {
			panic(err)
		}
		tools.MigrationCreate("core", name)
		return

    case "sync-version":
        SyncVersion()
        return

	case "create-devkit":
		SyncVersion()
		devkit.CreateDevkit()
		return

	case "build-cli":
		tools.BuildFlareCLI()
		return

	case "server":
		SyncVersion() // sync core version to package.json
		Server()
		return
	}

	fmt.Println(Usage())
}

func Server() {
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

func SyncVersion() {
	version := tools.CoreInfo().Version
	packageJson := "package.json"
	var pkg map[string]interface{}
	err := sdkfs.ReadJson(packageJson, &pkg)
	if err != nil {
		panic(err)
	}
	pkg["version"] = version
	err = sdkfs.WriteJson(packageJson, pkg)
	if err != nil {
		panic(err)
	}
}

func Usage() string {
	return `
Available commands:
    server              Start the flare server

    make-mono           Create mono-repo files

    create-migration    Create new migration files

    sync-version        Sync core version to package.json version

    create-devkit       Generate devkit files

    build-cli           Build the flare executable CLI
`
}
