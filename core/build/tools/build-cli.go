package tools

import (
	"core/env"
	"core/internal/utils/pkg"
	"fmt"
	"os"
	"runtime"

	sdkpkg "github.com/flarehotspot/go-utils/pkg"
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
		workdir, _ := os.Getwd()
		envs := []string{"GOOS=" + b.GOOS, "GOARCH=" + b.GOARCH}
		opts := sdkpkg.GoBuildOpts{
			GoBinPath: pkg.GoBin(),
			WorkDir:   workdir,
			Env:       envs,
			BuildTags: env.BuildTags,
		}

		if err := sdkpkg.BuildGoModule(cliFile, cliPath, opts); err != nil {
			panic(err)
		}

		fmt.Printf("Flare CLI built at: %s\n", cliPath)
	}
}
