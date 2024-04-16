package controllers

import (
	"testing"

	"github.com/flarehotspot/core/internal/plugins"
)

func BenchmarkGetLogs(b *testing.B) {
	g := plugins.NewGlobals()

	GetLogs(g)
}
