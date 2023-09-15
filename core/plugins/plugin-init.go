//go:build !mono

package plugins

import (
	"path/filepath"
	"plugin"

	sdk "github.com/flarehotspot/core/sdk/api/plugin"
)

func (api *PluginApi) Init() error {
	pluginLib := filepath.Join(api.dir, "plugin.so")
	p, err := plugin.Open(pluginLib)
	if err != nil {
		return err
	}

	initSym, err := p.Lookup("Init")
	if err != nil {
		return err
	}

	initFn := initSym.(func(sdk.IPluginApi))
	initFn(api)

	return nil
}
