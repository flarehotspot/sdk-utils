package pkg

import (
	"fmt"

	"core/env"
)

func BuildArgs() []string {
	args := []string{}
	args = append(args, fmt.Sprintf(`-tags="%s"`, env.BuildTags))
	args = append(args, `-ldflags="-s -w"`, "-trimpath")

	fmt.Println("Build args: ", args)

	return args
}
