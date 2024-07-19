//go:build !dev

package encdisk

import (
	"core/internal/utils/cmd"
	"fmt"
	"os"
)

func (d *EncryptedDisk) Initialize() error {
	if err := cmd.ExecAsh(fmt.Sprintf(`echo -n "%s" | cryptsetup luksFormat %s -`, d.pass, d.file)); err != nil {
		return err
	}

	if err := cmd.ExecAsh(fmt.Sprintf(`echo -n "%s" | cryptsetup luksOpen %s %s -`, d.pass, d.file, d.name)); err != nil {
		return err
	}

	if err := cmd.ExecAsh("mkfs.ext4 /dev/mapper/" + d.name); err != nil {
		return err
	}

	if err := os.MkdirAll(d.mountpath, 0755); err != nil {
		return err
	}

	if err := cmd.ExecAsh(fmt.Sprintf("mount /dev/mapper/%s %s", d.name, d.mountpath)); err != nil {
		return err
	}

	return nil
}
