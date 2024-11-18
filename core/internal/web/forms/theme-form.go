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

	portalThemesField := sdkforms.ListField{
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
	}

	adminThemesField := sdkforms.ListField{
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
	}

	multiField := sdkforms.MultiField{
		Name:  "multi_field",
		Label: "Multi Field",
		Columns: func() []sdkforms.MultiFieldCol {
			cols := []sdkforms.MultiFieldCol{
				{
					Name:       "col1",
					Label:      "Column 1 (text)",
					Type:       sdkforms.FormFieldTypeText,
					DefaultVal: "default val 1",
				},
				{
					Name:       "col2",
					Label:      "Column 2 (number)",
					Type:       sdkforms.FormFieldTypeNumber,
					DefaultVal: 1,
				},
				{
					Name:       "col3",
					Label:      "Column 3 (boolean)",
					Type:       sdkforms.FormFieldTypeBoolean,
					DefaultVal: true,
				},
			}
			return cols
		},
		DefaultVal: sdkforms.MultiFieldData{},
	}

	listFieldTxt := sdkforms.ListField{
		Name:       "list_field_txt",
		Label:      "List Field (text)",
		Multiple:   true,
		Type:       sdkforms.FormFieldTypeText,
		DefaultVal: []string{"val1"},
		Options: func() []sdkforms.ListOption {
			return []sdkforms.ListOption{
				{
					Label: "Value 1",
					Value: "val1",
				},
				{
					Label: "Value 2",
					Value: "val2",
				},
			}
		},
	}

	listFieldNum := sdkforms.ListField{
		Name:       "list_field_num",
		Label:      "List Field (number)",
		Type:       sdkforms.FormFieldTypeNumber,
		DefaultVal: 100,
		Options: func() []sdkforms.ListOption {
			return []sdkforms.ListOption{
				{
					Label: "100",
					Value: 100,
				},
				{
					Label: "200",
					Value: 200,
				},
			}
		},
	}

	form = sdkforms.Form{
		Name:          "themes",
		CallbackRoute: "admin:themes:save",
		Sections: []sdkforms.FormSection{
			{
				Name: "themes",
				Fields: []sdkforms.FormField{
					portalThemesField,
					adminThemesField,
					multiField,
					listFieldTxt,
					listFieldNum,
				},
			},
		},
	}

	return
}
