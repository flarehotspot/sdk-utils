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

func NewPluginConfig(api *plugins.PluginApi) *PluginConfig {
	valuesPath := filepath.Join(sdkpaths.ConfigDir, "plugins", api.Pkg(), "values.json")
	configPath := filepath.Join(sdkpaths.ConfigDir, "plugins", api.Pkg(), "config.json")
	return &PluginConfig{
		api:        api,
		datapath:   valuesPath,
		configpath: configPath,
		parsedData: nil,
	}
}

type PluginConfig struct {
	api        *plugins.PluginApi
	datapath   string
	configpath string
	config     sdkfields.Config
	parsedData ConfigData
}

func (p *PluginConfig) LoadConfig() (err error) {
	p.config = nil

	if err := sdkfs.ReadJson(p.configpath, &p.config); err != nil {
		return err
	}

	fmt.Printf("config: %+v\n", p.config)
	return
}

func (p *PluginConfig) LoadValues() {
	if !sdkfs.Exists(p.datapath) {
		return
	}
	fmt.Println("Loading values from", p.datapath)
	if err := sdkfs.ReadJson(p.datapath, &p.parsedData); err != nil {
		p.parsedData = nil
	}
}

func (p *PluginConfig) SaveForm(r *http.Request) (err error) {
	parsedData := make([]SectionData, len(p.config))

	for sidx, sec := range p.config {
		sectionData := SectionData{
			Name:   sec.Name,
			Fields: make([]FieldData, len(sec.Fields)),
		}

		for fidx, fld := range sec.Fields {
			field := FieldData{Name: fld.Name}

			switch fld.InputType {
			case sdkfields.FieldTypeText:
				val := r.FormValue(sec.Name + "::" + fld.Name)
				field.Value = val
			case sdkfields.FieldTypeNumber:
				val := r.FormValue(sec.Name + "::" + fld.Name)
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

	if err = sdkfs.WriteJson(p.datapath, parsedData); err != nil {
		return
	}

	fmt.Printf("parsedData: %+v\n", parsedData)

	p.parsedData = parsedData

	return
}

func (p *PluginConfig) ParseField(r *http.Request, sec sdkfields.Section, fld sdkfields.Field, valstr string) (field FieldData, err error) {
	if fld.InputType == "" {
		return field, errors.New(fmt.Sprintf("field %s with value %s has no type", fld.Name, valstr))
	}

	field = FieldData{
		Name: fld.Name,
	}

	switch fld.InputType {
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
	default:
		return field, errors.New(fmt.Sprintf("field type %s not supported", fld.InputType))
	}

	fmt.Printf("parsed field: %+v\n", field)

	return field, nil
}

func (p *PluginConfig) ParseMultiField(r *http.Request, sec sdkfields.Section, fld sdkfields.Field) (field MultiFieldData, err error) {
	if len(fld.Columns) < 1 {
		return field, errors.New(fmt.Sprintf("multi-field %s has no columns", fld.Name))
	}

	fmt.Printf("multi field form: %+v\n", r.Form)

	col1 := sec.Name + "::" + fld.Name + "::" + fld.Columns[0].Name + "[]"
	numRows := len(r.Form[col1])

	fmt.Printf("numRows: %d\n", numRows)

	field = MultiFieldData{
		Name:   fld.Name,
		Fields: make([][]FieldData, numRows),
	}

	for ridx := 0; ridx < numRows; ridx++ {
		row := make([]FieldData, len(fld.Columns))
		for cidx, col := range fld.Columns {
			colfld := sdkfields.Field{
				Name:      col.Name,
				InputType: col.Type,
				Default:   col.Default,
			}
			inputName := sec.Name + "::" + fld.Name + "::" + colfld.Name + "[]"
			colarr := r.Form[inputName]
			fmt.Printf("colarr: %+v\n", colarr)
			valstr := colarr[ridx]
			fmt.Printf("valstr: %s\n", valstr)

			row[cidx], err = p.ParseField(r, sec, colfld, valstr)
			if err != nil {
				return field, err
			}
		}
		fmt.Printf("row: %+v\n", row)
		field.Fields[ridx] = row
	}

	fmt.Printf("parsed multi-field: %+v\n", field)

	return field, nil

}

func (p *PluginConfig) GetSection(secname string) (sec sdkfields.Section, ok bool) {
	for _, s := range p.config {
		if s.Name == secname {
			return s, true
		}
	}
	return
}

func (p *PluginConfig) GetField(secname string, name string) (f sdkfields.Field, ok bool) {
	for _, s := range p.config {
		if s.Name == secname {
			for _, fld := range s.Fields {
				if fld.Name == name {
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
		return f.Default, nil
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
