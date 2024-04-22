package tools

import (
	"fmt"
	"os"
)

func BuildArgs() []string {
	tags := os.Getenv("GO_TAGS")
	if tags == "" {
		tags = os.Getenv("GOTAGS")
	}

	args := []string{}
	args = append(args, "-tags", tags)
	args = append(args, "-ldflags", "-s -w", "-trimpath")

	fmt.Println("Build args: ", args)

	return args
}
