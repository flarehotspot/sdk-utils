package tools

import (
	"core/env"
	"core/internal/utils/pkg"
	"fmt"
	"os"

	sdkpkg "github.com/flarehotspot/go-utils/pkg"
)

type FlareCliBuild struct {
	GOOS   string
	GOARCH string
	File   string
}

func BuildFlareCLI() {
	fmt.Println("Building flare CLI...")

	cliFile := "core/internal/cli/main.go"
	cliPath := "bin/flare"
	workdir, _ := os.Getwd()
	opts := sdkpkg.GoBuildOpts{
		GoBinPath: pkg.GoBin(),
		WorkDir:   workdir,
		Env:       os.Environ(),
		BuildTags: env.BuildTags,
	}

	if err := sdkpkg.BuildGoModule(cliFile, cliPath, opts); err != nil {
		panic(err)
	}

	fmt.Printf("Flare CLI built at: %s\n", cliPath)
}
