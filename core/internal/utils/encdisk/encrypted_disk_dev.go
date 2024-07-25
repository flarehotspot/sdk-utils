//go:build dev

package encdisk

import (
	"log"
	"os"
	sdkfs "sdk/utils/fs"
)

func (d *EncryptedDisk) Mount() error {
	log.Println("Initializing encrypted disk: ", d.mountpath)
	return sdkfs.EmptyDir(d.mountpath)
}

func (d *EncryptedDisk) Unmount() error {
	return os.RemoveAll(d.mountpath)
}
