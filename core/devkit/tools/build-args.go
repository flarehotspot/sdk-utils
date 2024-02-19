package tools

import (
	"os"
)

var (
	prod          = os.Getenv("GO_ENV") != "development"
	isDevkitBuild = os.Getenv("DEVKIT_BUILD") != ""
)

func BuildArgs() []string {
	var args []string

	if prod {
		args = append(args, `-ldflags`, "-s -w", "-trimpath")
		if isDevkitBuild {
			args = append(args, "-tags", "dev")
		}
	} else {
		args = append(args, "-tags", "mono dev")
	}

	return args
}
