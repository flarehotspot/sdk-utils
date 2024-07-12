package cmd

import (
	"core/env"
	"fmt"
	"os"
	"path/filepath"
	sdkfs "sdk/utils/fs"
	sdkpaths "sdk/utils/paths"
	sdkstr "sdk/utils/strings"
)

func ExecShell(command string) error {
	file := filepath.Join(sdkpaths.TmpDir, "shell-"+sdkstr.Rand(16))
	var shell_bin string
	if env.GoEnv == env.ENV_DEV {
		shell_bin = "/bin/bash"
	} else {
		shell_bin = "/bin/ash"
	}

	content := fmt.Sprintf("#!%s\n%s\n", shell_bin, command)
	if err := os.WriteFile(file, []byte(content), sdkfs.PermFile); err != nil {
		return err
	}

	err := Exec(shell_bin + " " + file)

	// clean up file
	os.RemoveAll(file)

	return err
}
