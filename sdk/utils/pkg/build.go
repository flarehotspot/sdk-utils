package sdkpkg

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	sdkpaths "github.com/flarehotspot/go-utils/paths"
)

type GoBuildOpts struct {
	GoBinPath string
	WorkDir   string
	Env       []string
	ExtraArgs []string
}

func BuildGoModule(gofile string, outfile string, opts GoBuildOpts) error {
	if opts.GoBinPath == "" {
		opts.GoBinPath = "go"
	}

	fmt.Println("Building go module: " + sdkpaths.StripRoot(filepath.Join(opts.WorkDir, gofile)))

	goBin := opts.GoBinPath
	buildArgs := DefaultBuildArgs()
	buildArgs = append(buildArgs, opts.ExtraArgs...)

	buildCmd := []string{"build"}
	buildCmd = append(buildCmd, buildArgs...)
	buildCmd = append(buildCmd, "-o", outfile, gofile)

	cmdstr := goBin
	for _, arg := range buildCmd {
		cmdstr += " " + arg
	}

	fmt.Printf(`Build working directory: %s`+"\n", sdkpaths.StripRoot(opts.WorkDir))

	var stderr strings.Builder

	cmd := exec.Command("sh", "-c", cmdstr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = &stderr
	cmd.Env = append(os.Environ(), opts.Env...)
	cmd.Dir = opts.WorkDir
	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("Failed to build go module: %s\n%s", err, stderr.String())
	}

	fmt.Println("Module built successfully: " + sdkpaths.StripRoot(filepath.Join(opts.WorkDir, outfile)))
	return nil
}

func DefaultBuildArgs() []string {
	args := []string{}
	args = append(args, `-ldflags="-s -w"`, "-trimpath", "-buildvcs=false")

	fmt.Println("Build args: ", args)

	return args
}
