//go:build !dev

package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"
    sdkstr "sdk/utils/strings"
)

func ExecAsh(command string) error {
	f := sdkstr.Rand(16) + ".sh"
	script := filepath.Join(os.TempDir(), f)

	err := ioutil.WriteFile(script, []byte(command), 0644)
	if err != nil {
		return err
	}

	err = Exec("/bin/ash " + script)
	if err != nil {
		return err
	}

	os.Remove(script)

	return nil
}
