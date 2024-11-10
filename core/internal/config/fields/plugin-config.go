package cfgfields

import (
	"core/internal/plugins"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	sdkfields "sdk/api/config/fields"
	"strconv"

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

func (p *PluginConfig) SaveForm(r *http.Request) (err error) {
	parsedData := make([]SectionData, len(p.sections))

	for sidx, sec := range p.sections {
		sectionData := SectionData{
			Name:   sec.Name,
			Fields: make([]FieldData, len(sec.Fields)),
		}

		for fidx, fld := range sec.Fields {
			field := FieldData{Name: fld.GetName()}

			switch fld.GetType() {
			case sdkfields.FieldTypeText:
				val := r.FormValue(sec.Name + "::" + fld.GetName())
				field.Value = val
			case sdkfields.FieldTypeNumber:
				val := r.FormValue(sec.Name + "::" + fld.GetName())
				v, err := strconv.Atoi(val)
				if err != nil {
					return err
				}
				field.Value = v
			case sdkfields.FieldTypeMulti:
				multifld, err := p.ParseMultiField(r, sec, fld)
				if err != nil {
					return err
				}
				field.Value = multifld
			}
			sectionData.Fields[fidx] = field
		}

		parsedData[sidx] = sectionData
	}

	p.parsedData = parsedData

	return
}

func (p *PluginConfig) ParseField(r *http.Request, sec sdkfields.Section, fld sdkfields.ConfigField, valstr string) (field FieldData, err error) {
	field = FieldData{
		Name: fld.GetName(),
	}

	switch fld.GetType() {
	case sdkfields.FieldTypeText:
		field.Value = valstr

	case sdkfields.FieldTypeNumber:
		val, err := strconv.Atoi(valstr)
		if err != nil {
			return field, err
		}
		field.Value = val

	case sdkfields.FieldTypeMulti:
		multifld, err := p.ParseMultiField(r, sec, fld)
		if err != nil {
			return field, err
		}
		field.Value = multifld.Fields
	}

	return field, nil
}

func (p *PluginConfig) ParseMultiField(r *http.Request, sec sdkfields.Section, fld sdkfields.ConfigField) (field MultiFieldData, err error) {
	multifld, ok := fld.(sdkfields.MultiField)
	if !ok {
		fmt.Printf("fld: %+v\n", fld)
		return field, errors.New("field is not a multi-field")
	}

	if len(multifld.Columns) < 1 {
		return field, errors.New(fmt.Sprintf("multi-field %s has no columns", fld.GetName()))
	}

	col1 := sec.Name + "::" + fld.GetName() + "::" + multifld.Columns[0].GetName() + "[]"
	numRows := len(r.Form[col1])

	field = MultiFieldData{
		Name:   fld.GetName(),
		Fields: make([][]FieldData, numRows),
	}

	for ridx := 0; ridx < numRows; ridx++ {
		row := make([]FieldData, len(multifld.Columns))
		for cidx, colfld := range multifld.Columns {
			inputName := sec.Name + "::" + fld.GetName() + "::" + colfld.GetName() + "[]"
			colarr := r.Form[inputName]
			valstr := colarr[ridx]
			row[cidx], err = p.ParseField(r, sec, colfld, valstr)
			if err != nil {
				return field, err
			}
		}
		field.Fields[ridx] = row
	}

	return field, nil

}

func (p *PluginConfig) GetSection(secname string) (sec sdkfields.Section, ok bool) {
	for _, s := range p.sections {
		if s.Name == secname {
			return s, true
		}
	}
	return
}

func (p *PluginConfig) GetField(secname string, name string) (f sdkfields.ConfigField, ok bool) {
	for _, s := range p.sections {
		if s.Name == secname {
			for _, fld := range s.Fields {
				if fld.GetName() == name {
					return fld, true
				}
			}
		}
	}
	return
}

func (p *PluginConfig) GetParsedSection(secname string) (sec SectionData, ok bool) {
	if p.parsedData == nil {
		return
	}

	for _, s := range p.parsedData {
		if s.Name == secname {
			return s, true
		}
	}

	return
}

func (p *PluginConfig) GetParsedField(secname string, name string) (fld FieldData, ok bool) {
	if s, ok := p.GetParsedSection(secname); ok {
		for _, f := range s.Fields {
			if f.Name == name {
				return f, true
			}
		}

	}

	return
}

func (p *PluginConfig) GetParsedFieldValue(secname string, name string) (val interface{}, ok bool) {
	if f, ok := p.GetParsedField(secname, name); ok {
		return f.Value, true
	}
	return
}

func (p *PluginConfig) GetDefaultValue(secname string, name string) (val interface{}, err error) {
	if f, ok := p.GetField(secname, name); ok {
		return f.GetDefaultValue(), nil
	}
	return nil, errors.New(fmt.Sprintf("section %s, field %s default value not found", secname, name))
}

func (p *PluginConfig) GetFieldValue(secname string, name string) (val interface{}, err error) {
	if v, ok := p.GetParsedFieldValue(secname, name); ok {
		return v, nil
	}

	return p.GetDefaultValue(secname, name)
}

func (p *PluginConfig) GetStringValue(secname string, name string) (val string, err error) {
	v, err := p.GetFieldValue(secname, name)
	if err != nil {
		return "", err
	}
	str, ok := v.(string)
	if !ok {
		return "", errors.New(fmt.Sprintf("section %s, field %s is not a string", secname, name))
	}
	return str, nil
}

func (p *PluginConfig) GetIntValue(secname string, name string) (val int, err error) {
	v, err := p.GetFieldValue(secname, name)
	if err != nil {
		return 0, err
	}
	num, ok := v.(int)
	if !ok {
		return 0, errors.New(fmt.Sprintf("section %s, field %s is not an int", secname, name))
	}
	return num, nil
}

func (p *PluginConfig) GetMultiValue(secname string, name string) (val MultiFieldData, err error) {
	v, err := p.GetFieldValue(secname, name)
	if err != nil {
		return
	}

	fields, ok := v.(MultiFieldData)
	if !ok {
		return val, errors.New(fmt.Sprintf("section %s, field %s is not a multi-field", secname, name))
	}

	return fields, nil
}
