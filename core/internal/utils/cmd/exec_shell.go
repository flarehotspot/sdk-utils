package cmd

import (
	"errors"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	sdkfs "sdk/utils/fs"
	sdkpaths "sdk/utils/paths"
	sdkstr "sdk/utils/strings"
	"strings"
)

type ExecOpts struct {
	Stdout io.Writer
	Stderr io.Writer
	Dir    string
	Env    []string
}

func execShell(command string, opts *ExecOpts) (err error) {

	f := filepath.Join(sdkpaths.TmpDir, sdkstr.Rand(16)+".sh")
	if err = os.WriteFile(f, []byte(command), sdkfs.PermFile); err != nil {
		return err
	}

	defer os.Remove(f)

	var (
		shells = []string{"/bin/ash", "/bin/bash", "/bin/zsh"}
	)

	var shell string
	for _, s := range shells {
		if sdkfs.Exists(s) {
			shell = s
			break
		}
	}

	hasStderr := false
	cmd := exec.Command(shell, f)

	if opts != nil {
		if opts.Stdout != nil {
			cmd.Stdout = opts.Stdout
		}
		if opts.Stderr != nil {
			hasStderr = true
			cmd.Stderr = opts.Stderr
		}
		if opts.Dir != "" {
			cmd.Dir = opts.Dir
		}
		if len(opts.Env) > 0 {
			cmd.Env = opts.Env
		}
	}

	var stderr strings.Builder
	if !hasStderr {
		cmd.Stderr = &stderr
	}

	log.Printf("Executing '%s': %s\n", shell, command)

	if err = cmd.Run(); err != nil {
		if !hasStderr && stderr.String() != "" {
			err = errors.New(stderr.String())
		}
	}

	return err
}
