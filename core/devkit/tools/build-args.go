package tools

import (
	"fmt"
	"os"
)

var (
	prod = os.Getenv("GO_ENV") != "development"
	tags = os.Getenv("GO_TAGS")
)

func BuildArgs() []string {
	args := []string{"-ldflags", "-s -w", "-trimpath"}

	if !prod {
		if tags == "" {
			tags = "dev"
		}
		args = append(args, "-tags", tags)
	}

	fmt.Println("Build args: ", args)

	return args
}
