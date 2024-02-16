package ifbutil

import (
	"sync/atomic"

	"github.com/flarehotspot/flarehotspot/core/utils/cmd"
)

var (
	supported atomic.Bool
)

func init() {
	err := cmd.Exec("modprobe ifb")
	supported.Store(err == nil)
}

// check if ifb interface is supported
func Supported() bool {
	return supported.Load()
}
