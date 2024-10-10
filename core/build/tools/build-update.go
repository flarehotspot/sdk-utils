package tools

import (
	"core/internal/utils/pkg"
	"fmt"
	"os"
	"runtime"
)

type FlareUpdateBuild struct {
	GOOS   string
	GOARCH string
	File   string
}

// Builds the flare system updater
func BuildSysUp() {
	builds := []FlareUpdateBuild{
		{
			GOOS:   "windows",
			GOARCH: runtime.GOARCH,
			File:   "./bin/update.exe",
		},
		{
			GOOS:   "linux",
			GOARCH: runtime.GOARCH,
			File:   "./bin/update",
		},
	}

	for _, b := range builds {
		fmt.Println("Building flare system updater...")
		sysupFile := "core/internal/sysup/main.go"
		sysupPath := b.File
		workDir, _ := os.Getwd()
		args := &pkg.GoBuildArgs{
			WorkDir: workDir,
			Env:     []string{"GOOS=" + b.GOOS, "GOARCH=" + b.GOARCH},
		}
		err := pkg.BuildGoModule(sysupFile, sysupPath, args)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Flare System Updater built at: %s\n", sysupPath)
	}
}
