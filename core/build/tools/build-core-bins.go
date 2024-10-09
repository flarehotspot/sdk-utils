package tools

import (
	"core/env"
	"core/internal/utils/pkg"
	"fmt"
	"path/filepath"

	sdkpaths "github.com/flarehotspot/go-utils/paths"
	sdkruntime "github.com/flarehotspot/go-utils/runtime"
	sdkstr "github.com/flarehotspot/go-utils/strings"
)

func BuildCoreBins() {
	BuildFlareCLI()
	BuildCore()
	BuildSysUp()

	goversion := sdkruntime.GO_VERSION
	tags := sdkstr.Slugify(env.BuildTags, "-")

	build := &BuildOutput{
		OutputDirName: filepath.Join("core-binaries", fmt.Sprintf("%s-%s-go%s-%s", pkg.CoreInfo().Version, sdkruntime.GOARCH, goversion, tags)),
		Files: []string{
			"bin/flare",
			"bin/update",
			"core/plugin.so",
		},
	}

	if err := build.Run(); err != nil {
		panic(err)
	}
}

func BuildCore() {
	workdir := filepath.Join(sdkpaths.TmpDir, "builds/core")
	if err := pkg.BuildPlugin(sdkpaths.CoreDir, workdir); err != nil {
		panic(err)
	}
}
