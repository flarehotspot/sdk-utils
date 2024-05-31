//go:build !mono

package main

import (
	"core/internal/boot"
	"core/internal/plugins"
)

func main() {}

func Init() {
	g := plugins.NewGlobals()
	boot.Init(g)
	defer g.Db.SqlDB().Close()
}
