package tools

import (
	"fmt"
	"os"
	"runtime"
)

type FlareCliBuild struct {
	GOOS   string
	GOARCH string
	File   string
}

func BuildFlareCLI() {
	builds := []FlareCliBuild{
		{
			GOOS:   "windows",
			GOARCH: runtime.GOARCH,
			File:   "./bin/flare.exe",
		},
		{
			GOOS:   "linux",
			GOARCH: runtime.GOARCH,
			File:   "./bin/flare",
		},
	}

	for _, b := range builds {
		fmt.Println("Building flare CLI...")
		cliFile := "core/internal/cli/main.go"
		cliPath := b.File
		workDir, _ := os.Getwd()
		args := &GoBuildArgs{
			WorkDir: workDir,
			Env:     []string{"GOOS=" + b.GOOS, "GOARCH=" + b.GOARCH},
		}
		err := BuildGoModule(cliFile, cliPath, args)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Flare CLI built at: %s\n", cliPath)
	}
}
