//go:build !dev

package encdisk

import (
	"core/internal/utils/cmd"
	"fmt"
	"os"
)

func (d *EncryptedDisk) Mount() error {
	if err := cmd.Exec(fmt.Sprintf(`echo -n "%s" | cryptsetup luksFormat %s -`, d.pass, d.file), nil); err != nil {
		return err
	}

	if err := cmd.Exec(fmt.Sprintf(`echo -n "%s" | cryptsetup luksOpen %s %s -`, d.pass, d.file, d.name), nil); err != nil {
		return err
	}

	if err := cmd.Exec("mkfs.ext4 /dev/mapper/"+d.name, nil); err != nil {
		return err
	}

	if err := os.MkdirAll(d.mountpath, 0755); err != nil {
		return err
	}

	if err := cmd.Exec(fmt.Sprintf("mount /dev/mapper/%s %s", d.name, d.mountpath), nil); err != nil {
		return err
	}

	return nil
}

func (d *EncryptedDisk) Unmount() error {
	if err := cmd.Exec(fmt.Sprintf("umount %s", d.mountpath), nil); err != nil {
		return err
	}
	if err := cmd.Exec(fmt.Sprintf("cryptsetup luksClose %s", d.name), nil); err != nil {
		return err
	}
	return nil
}
