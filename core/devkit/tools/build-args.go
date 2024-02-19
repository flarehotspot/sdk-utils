package tools

import (
	"fmt"
	"os"
)

func BuildArgs() []string {
	prod := os.Getenv("GO_ENV") != "development"
	tags := os.Getenv("GO_TAGS")
	if tags == "" {
		tags = os.Getenv("GOTAGS")
	}

	args := []string{}

	if !prod {
		if tags == "" {
			tags = "dev"
		}
		args = append(args, "-tags", tags)
	}

	args = append(args, "-ldflags", "-s -w", "-trimpath")

	fmt.Println("Build args: ", args)

	return args
}
