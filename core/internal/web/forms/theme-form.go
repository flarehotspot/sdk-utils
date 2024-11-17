package coreforms

import (
	"core/internal/config"
	"core/internal/plugins"
	sdkforms "sdk/api/forms"
	sdkplugin "sdk/api/plugin"
)

func GetThemeForm(g *plugins.CoreGlobals) (form sdkforms.Form, err error) {

	allPlugins := g.PluginMgr.All()
	adminThemes := []sdkplugin.PluginApi{}
	portalThemes := []sdkplugin.PluginApi{}

	cfg, err := config.ReadThemesConfig()
	if err != nil {
		return
	}

	for _, p := range allPlugins {
		features := p.Features()
		for _, f := range features {
			if f == "theme:admin" {
				adminThemes = append(adminThemes, p)
			}

			if f == "theme:portal" {
				portalThemes = append(portalThemes, p)
			}
		}
	}

	form = sdkforms.Form{
		Name:          "themes",
		CallbackRoute: "admin:themes:save",
		Sections: []sdkforms.FormSection{
			{
				Name: "themes",
				Fields: []sdkforms.FormField{
					sdkforms.ListField{
						Name:       "portal_theme",
						Label:      "Select Portal Theme",
						Type:       sdkforms.FormFieldTypeText,
						DefaultVal: cfg.PortalThemePkg,
						Options: func() []sdkforms.ListOption {
							opts := []sdkforms.ListOption{}
							for _, p := range portalThemes {
								opts = append(opts, sdkforms.ListOption{
									Label: p.Name(),
									Value: p.Pkg(),
								})
							}
							return opts
						},
					},
					sdkforms.ListField{
						Name:       "admin_theme",
						Label:      "Select Admin Theme",
						Type:       sdkforms.FormFieldTypeText,
						DefaultVal: cfg.AdminThemePkg,
						Options: func() []sdkforms.ListOption {
							opts := []sdkforms.ListOption{}
							for _, p := range adminThemes {
								opts = append(opts, sdkforms.ListOption{
									Label: p.Name(),
									Value: p.Pkg(),
								})
							}
							return opts
						},
					},
				},
			},
		},
	}

	return
}
