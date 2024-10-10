package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"plugin"
	"strconv"
	"strings"
	"syscall"

	tools "core/build/tools"
	"core/env"
	"core/internal/utils/pkg"

	sdkpaths "github.com/flarehotspot/go-utils/paths"
)

var (
	gowork bool
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(Usage())
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "env":
		fmt.Println(GoEnvToString(env.GO_ENV))
		return

	case "server":
		Server()
		return

	case "create-plugin":
		CreatePlugin()
		return

	case "create-migration":
		CreateMigration()
		return

	case "build-plugin":
		BuildPlugin()
		return

	case "build-plugins":
		BuildPlugin()
		return

	case "fix-workspace":
		tools.CreateGoWorkspace()
		return

	case "install-go":
		var installPath string
		if len(os.Args) > 2 {
			installPath = os.Args[2]
		}
		tools.InstallGo(installPath)
		return

	case "help":
		fmt.Println(Usage())
		return

	case "-h":
		fmt.Println(Usage())
		return

	default:
		fmt.Println("Unrecognized command: " + command)
	}

	fmt.Println(Usage())
	os.Exit(1)
}

func CreatePlugin() {
	var (
		err        error
		pluginPkg  string
		pluginName string
		pluginDesc string
	)

	for len(strings.Split(pluginPkg, ".")) < 3 {
		pluginPkg, err = tools.AskCmdInput("Enter the plugin package name, e.g. com.mydomain.plugin")
		if err != nil {
			panic(err)
		}
		if len(strings.Split(pluginPkg, ".")) < 3 {
			fmt.Println("Error: Package name must be at least 3 segments. For example: com.my-domain.my-plugin")
		}
	}

	pluginName, err = tools.AskCmdInput("Enter the plugin name, e.g. MyPlugin")
	if err != nil {
		panic(err)
	}

	pluginDesc, err = tools.AskCmdInput("Enter the plugin description")
	if err != nil {
		panic(err)
	}

	tools.CreatePlugin(pluginPkg, pluginName, pluginDesc)
}

func CreateMigration() {
	pluginPaths := pkg.LocalPluginPaths()
	pluginPkgs := make([]string, len(pluginPaths))
	for i, pluginPath := range pluginPaths {
		pluginPkgs[i] = filepath.Base(pluginPath)
	}

	pluginNums := make([]string, len(pluginPkgs))
	for i, pluginPkg := range pluginPkgs {
		pluginNums[i] = fmt.Sprintf("%d. %s", i+1, pluginPkg)
	}

	selectPkgAsk := fmt.Sprintf("\nSelect the plugin to create the migration for:\n%s\n\nEnter the number of the corresponding plugin", strings.Join(pluginNums, "\n"))

	selectPkg, err := tools.AskCmdInput(selectPkgAsk)
	if err != nil {
		panic(err)
	}

	pluginIdx, err := strconv.Atoi(selectPkg)
	if err != nil {
		panic(err)
	}

	if pluginIdx < 1 || pluginIdx > len(pluginPkgs) {
		panic(fmt.Errorf("Invalid plugin number: %d", pluginIdx))
	}

	pluginPkg := pluginPkgs[pluginIdx-1]

	name, err := tools.AskCmdInput("Enter the migration name, e.g. create_users_table")
	if err != nil {
		panic(err)
	}

	pluginDir := filepath.Join("plugins", pluginPkg)
	tools.MigrationCreate(pluginDir, name)
}

func BuildPlugin() {
	var err error
	if len(os.Args) < 3 {
		err = pkg.BuildLocalPlugins()
	} else {
		pluginPath := os.Args[2]
		workdir := filepath.Join(sdkpaths.TmpDir, "builds", filepath.Base(pluginPath))
		err = pkg.BuildPlugin(pluginPath, workdir)
	}
	if err != nil {
		fmt.Println("Error building plugin: " + err.Error())
		os.Exit(1)
	}
}

func Server() {
	// TODO: kill sysup if running
	if err := killSysUp(); err != nil {
		log.Println("Error killing sysup:", err)
		return
	}

	corePath := filepath.Join(sdkpaths.AppDir, "core/plugin.so")
	p, err := plugin.Open(corePath)
	if err != nil {
		log.Println("Error loading core plugin:", err)
		panic(err)
	}
	symInit, _ := p.Lookup("Init")
	initFn := symInit.(func())
	initFn()
}

func killSysUp() error {
	// read env
	fromSysUp := os.Getenv("FROM_SYSUP")
	if strings.ToLower(fromSysUp) == "true" {
		fmt.Println("flare cli spawned from updater")

		// check if sysup is still running
		ppid := os.Getppid()
		pproc, err := os.FindProcess(ppid)
		if err != nil {
			log.Println("Error finding sysup process: ", err)
			return err
		}

		if isProcRunning(pproc) {
			// kill sysup
			fmt.Println("Killing sysup..")
			if err = pproc.Kill(); err != nil {
				log.Println("Error killing sysup proc:", err)
				return err
			}

		} else {
			fmt.Println("Updater is not running")
		}
	}

	return nil
}

// checks if the proc is running
func isProcRunning(proc *os.Process) bool {
	if err := proc.Signal(syscall.Signal(0)); err != nil {
		log.Println("Error:", err)
		return false
	}

	return true
}

func GoEnvToString(e int8) string {
	switch e {
	case env.ENV_DEV:
		return "development"
	case env.ENV_PRODUCTION:
		return "production"
	case env.ENV_SANDBOX:
		return "sandbox"
	}
	return "unknown"
}

func Usage() string {
	return `
Usage: flare <command> [options]

list of commands:
    env                                 Print the build environment

    server                              Start the flare server

    create-plugin                       Create a new plugin

    create-migration                    Create a new migration

    build-plugin <plugin path>          Build plugin.so file. If no plugin path is provided, all plugins will be built.

    build-plugins                       Build plugin.so of all the local and system plugins. Similar to build-plugin command without arguments.

    fix-workspace                       Re-generate the go.work file

    install-go  <install path>          Install Go to the given path. If install path argument is not defined, then it will install in
                                        the "$GO_CUSTOM_PATH" if defined, else it will install in "go" directory under the
                                        current working directory.
`
}
