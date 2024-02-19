//go:build !mono

package main

import (
	"github.com/flarehotspot/core/internal/boot"
	"github.com/flarehotspot/core/internal/plugins"
)

func main() {}

func Init() {
	g := plugins.New()
	boot.Init(g)
}
