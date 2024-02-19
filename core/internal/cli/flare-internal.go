package main

import (
	"fmt"
	"os"

	"github.com/flarehotspot/core/devkit"
	"github.com/flarehotspot/core/devkit/tools"
)

func main() {
	command := os.Args[1]

	switch command {
	case "make-mono":
		tools.CreateMonoFiles()
		return
	case "create-devkit":
		devkit.CreateDevkit()
        return
	}

	fmt.Println(Usage())
}

func Usage() string {
	return `
Available commands:
    make-mono       Create mono-repo files
    create-devkit   Generate devkit files
`
}
