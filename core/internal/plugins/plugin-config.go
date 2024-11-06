package plugins

import sdkfields "sdk/api/config/fields"

func NewPluginConfig(api *PluginApi, sec []sdkfields.Section) *PluginConfig {
	return &PluginConfig{
		api:      api,
		Sections: sec,
	}
}

type PluginConfig struct {
	api      *PluginApi
	Sections []sdkfields.Section
}
