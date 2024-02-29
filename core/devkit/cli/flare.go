package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"plugin"
	"strconv"
	"strings"

	tools "github.com/flarehotspot/core/devkit/tools"
	"github.com/flarehotspot/core/env"
	sdkpaths "github.com/flarehotspot/sdk/utils/paths"
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

	case "fix-workspace":
		tools.CreateGoWorkspace()
		return

	case "install-go":
		InstallGo()
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
	pluginPaths := tools.PluginPathList()
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
		err = tools.BuildAllPlugins()
	} else {
		pluginPath := os.Args[2]
		err = tools.BuildPlugin(pluginPath)
	}
	if err != nil {
		fmt.Println("Error building plugin: " + err.Error())
		os.Exit(1)
	}
}

func InstallGo() {
	var installPath string
	if len(os.Args) > 2 {
		installPath = os.Args[2]
	}
	if installPath == "" {
		installPath = os.Getenv("GO_CUSTOM_PATH")
	}
	if installPath == "" {
		installPath = filepath.Join("go")
	}
	tools.InstallGo(installPath)
}

func Server() {
	if env.GoEnv == env.ENV_DEV {
		tools.CreateGoWorkspace()
		tools.BuildAllPlugins()
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

func Usage() string {
	return `
Usage: flare <command> [options]

list of commands:

    server                              Start the flare server

    create-plugin                       Create a new plugin

    create-migration                    Create a new migration

    build-plugin <plugin path>          Build plugin.so file. If no plugin path is provided, all plugins will be built.

    fix-workspace                       Re-generate the go.work file

    install-go  <install path>          Install Go to the given path. If install path argument is not defined, then it will install in
                                        the "$GO_CUSTOM_PATH" if defined, else it will install in "go" directory under the
                                        current working directory.
`
}
