//go:build !mono

package main

import (
	"github.com/flarehotspot/core/boot"
	"github.com/flarehotspot/core/plugins"
)

func main() {}

func Init() {
	g := plugins.New()
	boot.Init(g)
}
