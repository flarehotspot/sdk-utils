package coreforms

import (
	"core/internal/config"
	"core/internal/plugins"
	sdkforms "sdk/api/forms"
	sdkplugin "sdk/api/plugin"
)

const (
	ThemesFormName = "themes"
)

func RegisterThemesForm(g *plugins.CoreGlobals) (err error) {
	allPlugins := g.PluginMgr.All()
	adminThemes := []sdkplugin.IPluginApi{}
	portalThemes := []sdkplugin.IPluginApi{}

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
		Name:  "portal_theme",
		Label: "Select Portal Theme",
		Type:  sdkforms.FormFieldTypeText,
		ValueFn: func() interface{} {
			cfg, err := config.ReadThemesConfig()
			if err != nil {
				return ""
			}
			return cfg.PortalThemePkg
		},
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
		Name:  "admin_theme",
		Label: "Select Admin Theme",
		Type:  sdkforms.FormFieldTypeText,
		ValueFn: func() interface{} {
			cfg, err := config.ReadThemesConfig()
			if err != nil {
				return ""
			}
			return cfg.AdminThemePkg
		},
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
					Name:  "col1",
					Label: "Column 1 (text)",
					Type:  sdkforms.FormFieldTypeText,
					ValueFn: func() interface{} {
						return "text value"
					},
				},
				{
					Name:  "col2",
					Label: "Column 2 (decimal)",
					Type:  sdkforms.FormFieldTypeDecimal,
					ValueFn: func() interface{} {
						return 100.1
					},
				},
				{
					Name:  "col3",
					Label: "Column 3 (integer)",
					Type:  sdkforms.FormFieldTypeInteger,
					ValueFn: func() interface{} {
						return 1
					},
				},
				{
					Name:  "col4",
					Label: "Column 4 (boolean)",
					Type:  sdkforms.FormFieldTypeBoolean,
					ValueFn: func() interface{} {
						return true
					},
				},
			}
			return cols
		},
	}

	listFieldTxt := sdkforms.ListField{
		Name:     "list_field_txt",
		Label:    "List Field (text)",
		Multiple: true,
		Type:     sdkforms.FormFieldTypeText,
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
		ValueFn: func() interface{} {
			return []string{"val1", "val2"}
		},
	}

	listFieldNum := sdkforms.ListField{
		Name:     "list_field_num",
		Label:    "List Field (number)",
		Type:     sdkforms.FormFieldTypeDecimal,
		Multiple: true,
		Options: func() []sdkforms.ListOption {
			return []sdkforms.ListOption{
				{
					Label: "100",
					Value: 100.0,
				},
				{
					Label: "200",
					Value: 200.0,
				},
			}
		},
		ValueFn: func() interface{} {
			return []float64{100.0, 200.0}
		},
	}

	textField := sdkforms.TextField{
		Name:  "text_field",
		Label: "Text Field",
		ValueFn: func() string {
			return "text value"
		},
	}

	intField := sdkforms.IntegerField{
		Name:  "int_field",
		Label: "Int Field",
		ValueFn: func() int64 {
			return 124
		},
	}

	decimalField := sdkforms.DecimalField{
		Name:      "decimal_field",
		Label:     "Decimal Field",
		Step:      0.1,
		Precision: 2,
		ValueFn: func() float64 {
			return 201.50
		},
	}

	boolField := sdkforms.BooleanField{
		Name:  "boolean_field",
		Label: "Boolean Field",
		ValueFn: func() bool {
			return true
		},
	}

	themesForm := sdkforms.Form{
		Name:          ThemesFormName,
		CallbackRoute: "admin:themes:save",
		SubmitLabel:   "Save",
		Sections: []sdkforms.FormSection{
			{
				Name: "themes",
				Fields: []sdkforms.IFormField{
					textField,
					intField,
					decimalField,
					boolField,
					portalThemesField,
					adminThemesField,
					multiField,
					listFieldTxt,
					listFieldNum,
				},
			},
		},
	}

	err = g.CoreAPI.HttpAPI.Forms().RegisterForms(themesForm)
	if err != nil {
		return err
	}

	return nil
}
