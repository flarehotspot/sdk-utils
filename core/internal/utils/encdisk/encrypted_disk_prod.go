//go:build !dev

package encdisk

import (
	"core/internal/utils/cmd"
	"fmt"
	"log"
	"os"
)

func (d *EncryptedDisk) Mount() error {
	log.Printf("creating virtual disk at: %s", d.file)
	if err := cmd.ExecAsh("dd if=/dev/zero " + "of=" + d.file + " bs=1M count=50"); err != nil {
		return err
	}

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

func (d *EncryptedDisk) Unmount() error {
	if err := cmd.ExecAsh(fmt.Sprintf("umount %s", d.mountpath)); err != nil {
		return err
	}
	if err := cmd.ExecAsh(fmt.Sprintf("cryptsetup luksClose %s", d.name)); err != nil {
		return err
	}
	return nil
}
