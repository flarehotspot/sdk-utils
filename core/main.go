//go:build !mono

package main

import (
	"github.com/flarehotspot/core/boot"
	"github.com/flarehotspot/core/globals"
)

func main() {}

func Init() {
	g := globals.New()
	boot.Init(g)
}
