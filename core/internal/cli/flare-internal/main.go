package main

import (
	"fmt"
	"os"

	"github.com/flarehotspot/flarehotspot/core/internal/devkit"
	"github.com/flarehotspot/flarehotspot/core/internal/tools"
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
