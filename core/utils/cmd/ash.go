//go:build !dev

package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/flarehotspot/core/sdk/libs/slug"
)

func ExecAsh(command string) error {
	f := slug.Make(command) + ".sh"
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
