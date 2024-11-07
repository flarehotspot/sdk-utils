package cfgfields

import (
	"core/internal/plugins"
	"errors"
	"path/filepath"
	sdkfields "sdk/api/config/fields"

	sdkfs "github.com/flarehotspot/go-utils/fs"
	sdkpaths "github.com/flarehotspot/go-utils/paths"
)

func NewPluginConfig(api *plugins.PluginApi, sec []sdkfields.Section) *PluginConfig {
	savePath := filepath.Join(sdkpaths.ConfigDir, "plugins", api.Pkg(), "data.json")
	return &PluginConfig{
		api:      api,
		sections: sec,
		datapath: savePath,
	}
}

type PluginConfig struct {
	api        *plugins.PluginApi
	datapath   string
	sections   []sdkfields.Section
	parsedData []SectionData
}

func (p *PluginConfig) GetConfig(v *[]sdkfields.Section) (err error) {
	if v == nil {
		return errors.New("v cannot be nil")
	}

	if !sdkfs.Exists(p.datapath) {
		return
	}

	var parsedData []SectionData
	if err = sdkfs.ReadJson(p.datapath, &parsedData); err != nil {
		return
	}

	vptr := *v

	for sidx, s := range p.sections {
		vsec := &vptr[sidx]
		for _, dataSec := range parsedData {
			if dataSec.Title == s.Title {
				for fidx, f := range s.Fields {
					for _, dataField := range dataSec.Fields {
						if dataField.Name == f.GetName() {
							switch f.GetType() {

							// Text field
							case sdkfields.FieldTypeText:
								textField := f.(sdkfields.TextField)
								v, ok := dataField.Value.(string)
								if ok {
									textField.Value = v
									vsec.Fields[fidx] = textField
								}

								// Number field
							case sdkfields.FieldTypeNumber:
								numField := f.(sdkfields.NumberField)
								v, ok := dataField.Value.(int)
								if ok {
									numField.Value = v
									vsec.Fields[fidx] = numField
								}
							}
						}
					}
				}
			}
		}
	}

	return
}
