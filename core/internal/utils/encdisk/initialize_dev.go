//go:build dev

package encdisk

import (
	"log"
	sdkfs "sdk/utils/fs"
)

func (d *EncryptedDisk) Initialize() error {
	log.Println("Initializing encrypted disk: ", d.mountpath)
	return sdkfs.EmptyDir(d.mountpath)
}
