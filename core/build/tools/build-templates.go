package tools

import (
	"core/env"
	"core/internal/utils/pkg"
)

func BuildTemplates() {
	includeCore := false
	if env.GO_ENV == env.ENV_DEV {
		includeCore = true
	}

	pluginDirs := pkg.ListPluginDirs(includeCore)
	for _, p := range pluginDirs {
		pkg.BuildTemplates(p)
	}
}
