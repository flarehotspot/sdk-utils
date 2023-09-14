//go:build !debug && mono

package core

import (
	"github.com/flarehotspot/core/boot"
	"github.com/flarehotspot/core/globals"
)

func Init() {
	g := globals.New()
	boot.Init(g)
}
