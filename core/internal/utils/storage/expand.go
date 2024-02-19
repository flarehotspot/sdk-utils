package storage

import (
	"fmt"

	"github.com/flarehotspot/core/internal/utils/cmd"
)

func Expand(dev string, part string, partnum int) error {
	err := cmd.ExecAsh(fmt.Sprintf("parted /dev/%s resizepart %d 100%%", dev, partnum))
	if err != nil {
		return err
	}

	return cmd.ExecAsh(fmt.Sprintf("resize2fs %s", part))
}
