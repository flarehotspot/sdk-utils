package cfgfields

import (
	"core/internal/plugins"
	"errors"
	"fmt"
	"path/filepath"
	sdkfields "sdk/api/config/fields"

	sdkfs "github.com/flarehotspot/go-utils/fs"
	sdkpaths "github.com/flarehotspot/go-utils/paths"
)

func NewPluginConfig(api *plugins.PluginApi, sec []sdkfields.Section) *PluginConfig {
	savePath := filepath.Join(sdkpaths.ConfigDir, "plugins", api.Pkg(), "data.json")
	return &PluginConfig{
		api:        api,
		sections:   sec,
		datapath:   savePath,
		parsedData: nil,
	}
}

type PluginConfig struct {
	api        *plugins.PluginApi
	datapath   string
	sections   []sdkfields.Section
	parsedData []SectionData
}

func (p *PluginConfig) LoadConfig() {
	if !sdkfs.Exists(p.datapath) {
		return
	}
	fmt.Println("Loading config from", p.datapath)
	if err := sdkfs.ReadJson(p.datapath, &p.parsedData); err != nil {
		p.parsedData = nil
	}
}

func (p *PluginConfig) GetSection(title string) (sec sdkfields.Section, ok bool) {
	for _, s := range p.sections {
		if s.Title == title {
			return s, true
		}
	}
	return
}

func (p *PluginConfig) GetField(title string, name string) (f sdkfields.ConfigField, ok bool) {
	for _, s := range p.sections {
		if s.Title == title {
			for _, fld := range s.Fields {
				if fld.GetName() == name {
					return fld, true
				}
			}
		}
	}
	return
}

func (p *PluginConfig) GetParsedSection(title string) (sec SectionData, ok bool) {
	if p.parsedData == nil {
		return
	}

	for _, s := range p.parsedData {
		if s.Title == title {
			return s, true
		}
	}
	return
}

func (p *PluginConfig) GetParsedField(title string, name string) (fld FieldData, ok bool) {
	if s, ok := p.GetParsedSection(title); ok {
		for _, f := range s.Fields {
			if f.Name == name {
				return f, true
			}
		}
	}
	return
}

func (p *PluginConfig) GetParsedFieldValue(title string, name string) (val interface{}, ok bool) {
	if f, ok := p.GetParsedField(title, name); ok {
		return f.Value, true
	}
	return
}

func (p *PluginConfig) GetDefaultValue(title string, name string) (val interface{}, err error) {
	if f, ok := p.GetField(title, name); ok {
		return f.GetDefaultValue(), nil
	}
	return nil, errors.New(fmt.Sprintf("section %s, field %s default value not found", title, name))
}

func (p *PluginConfig) GetFieldValue(title string, name string) (val interface{}, err error) {
	if v, ok := p.GetParsedFieldValue(title, name); ok {
		return v, nil
	}

	return p.GetDefaultValue(title, name)
}

func (p *PluginConfig) GetStringValue(title string, name string) (val string, err error) {
	v, err := p.GetFieldValue(title, name)
	if err != nil {
		return "", err
	}
	str, ok := v.(string)
	if !ok {
		return "", errors.New(fmt.Sprintf("section %s, field %s is not a string", title, name))
	}
	return str, nil
}

func (p *PluginConfig) GetIntValue(title string, name string) (val int, err error) {
	v, err := p.GetFieldValue(title, name)
	if err != nil {
		return 0, err
	}
	num, ok := v.(int)
	if !ok {
		return 0, errors.New(fmt.Sprintf("section %s, field %s is not an int", title, name))
	}
	return num, nil
}
