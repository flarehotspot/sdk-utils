package storage

import (
	"fmt"

	"core/internal/utils/cmd"
)

func Expand(dev string, part string, partnum int) error {
	err := cmd.Exec(fmt.Sprintf("parted /dev/%s resizepart %d 100%%", dev, partnum), nil)
	if err != nil {
		return err
	}

	return cmd.Exec(fmt.Sprintf("resize2fs %s", part), nil)
}
