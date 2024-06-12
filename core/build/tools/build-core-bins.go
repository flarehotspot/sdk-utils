package tools

import (
	"core/env"
	"fmt"
	"path/filepath"
	sdkruntime "sdk/utils/runtime"
	sdkstr "sdk/utils/strings"
)

func BuildCoreBins() {
	BuildFlareCLI()
	BuildCore()

	goversion := sdkruntime.GOVERSION
	tags := sdkstr.Slugify(env.BuildTags, "-")

	build := &BuildOutput{
		OutputDirName: filepath.Join("core-binaries", fmt.Sprintf("%s-%s-go%s-%s", CoreInfo().Version, sdkruntime.GOARCH, goversion, tags)),
		Files: []string{
			"bin/flare",
			"core/plugin.so",
		},
	}

	if err := build.Run(); err != nil {
		panic(err)
	}
}
