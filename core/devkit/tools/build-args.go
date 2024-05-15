package tools

import (
	"fmt"

	"github.com/flarehotspot/core/env"
)

func BuildArgs() []string {
	args := []string{}
	args = append(args, "-tags", env.BuildTags)
	args = append(args, "-ldflags", "-s -w", "-trimpath")

	fmt.Println("Build args: ", args)

	return args
}
