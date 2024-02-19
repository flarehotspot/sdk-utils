package tools

import (
	"fmt"
	"os"
)

var (
	prod = os.Getenv("GO_ENV") != "development"
	tags = os.Getenv("GOTAGS")
)

func BuildArgs() []string {
	var args []string

	if prod {
		args = append(args, `-ldflags`, "-s -w", "-trimpath")
	} else {
		if tags == "" {
			tags = "dev"
		}
		args = append(args, "-tags", tags)
	}

	fmt.Println("Build args: ", args)

	return args
}
