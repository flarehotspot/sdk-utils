package tools

import (
	"core/env"
	"core/internal/utils/pkg"
	"fmt"
	"path/filepath"
	sdkpaths "sdk/utils/paths"
	sdkruntime "sdk/utils/runtime"
	sdkstr "sdk/utils/strings"
)

func BuildCoreBins() {
	BuildFlareCLI()
	BuildCore()

	goversion := sdkruntime.GO_VERSION
	tags := sdkstr.Slugify(env.BuildTags, "-")

	build := &BuildOutput{
		OutputDirName: filepath.Join("core-binaries", fmt.Sprintf("%s-%s-go%s-%s", pkg.CoreInfo().Version, sdkruntime.GOARCH, goversion, tags)),
		Files: []string{
			"bin/flare",
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
