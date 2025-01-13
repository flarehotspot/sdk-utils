package coreforms

import (
	"core/internal/plugins"
)

func RegisterForms(g *plugins.CoreGlobals) {
	err := RegisterThemesForm(g)
	if err != nil {
		panic(err)
	}
}
