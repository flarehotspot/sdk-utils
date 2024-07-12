package crypt

import "core/internal/utils/cmd"

type EncryptedDisk struct {
	file string
	name string
	pass string
}

func NewEncrypedDisk(file string, name string, pass string) *EncryptedDisk {
	return &EncryptedDisk{
		file: file,
		name: name,
		pass: pass,
	}
}

func (d *EncryptedDisk) Initialize() error {
    err := cmd.ExecShell("cryptsetup")
    return err
}
