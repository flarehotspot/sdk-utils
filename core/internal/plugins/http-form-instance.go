package plugins

import (
	formsutl "core/internal/utils/forms"
	formsview "core/resources/views/forms"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"reflect"
	sdkhttp "sdk/api/http"
	"strconv"

	"github.com/a-h/templ"
	sdkfs "github.com/flarehotspot/go-utils/fs"
	sdkpaths "github.com/flarehotspot/go-utils/paths"
)

var (
	ErrFieldMulti = errors.New("field type is multifield")
)

func NewHttpForm(api *PluginApi, key string, sections []sdkhttp.Section) (*HttpFormInstance, error) {
	configDir := filepath.Join(sdkpaths.ConfigDir, "plugins", api.Pkg(), key)

	if err := sdkfs.EnsureDir(configDir); err != nil {
		return nil, err
	}

	form := &HttpFormInstance{
		api:        api,
		configdir:  configDir,
		sections:   sections,
		parsedData: nil,
	}

	if err := form.writeConfig(); err != nil {
		return nil, err
	}

	form.LoadData()

	return form, nil
}

func LoadHttpForm(api *PluginApi, configDir string) (*HttpFormInstance, error) {
	form := &HttpFormInstance{
		api:        api,
		configdir:  configDir,
		sections:   nil,
		parsedData: nil,
	}

	if err := form.LoadConfig(); err != nil {
		return nil, err
	}

	return form, nil
}

type HttpFormInstance struct {
	api        *PluginApi
	configdir  string
	sections   []sdkhttp.Section
	parsedData formsutl.ConfigData
}

func (p *HttpFormInstance) Template(r *http.Request) templ.Component {
	csrfTag := p.api.HttpAPI.Helpers().CsrfHtmlTag(r)
	return formsview.HtmlForm(csrfTag, p.sections, p.parsedData)
}

func (p *HttpFormInstance) LoadConfig() (err error) {
	p.sections = nil
	if err := sdkfs.ReadJson(p.configPath(), &p.sections); err != nil {
		return err
	}

	fmt.Printf("config: %+v\n", p.sections)
	return
}

func (p *HttpFormInstance) LoadData() {
	if !sdkfs.Exists(p.dataPath()) {
		return
	}
	fmt.Println("Loading values from", p.dataPath())
	if err := sdkfs.ReadJson(p.dataPath(), &p.parsedData); err != nil {
		p.parsedData = nil
	}
}

func (p *HttpFormInstance) GetFormData() formsutl.ConfigData {
	return p.parsedData
}

func (p *HttpFormInstance) SaveForm(r *http.Request) (err error) {
	parsedData := make([]formsutl.SectionData, len(p.sections))

	for sidx, sec := range p.sections {
		sectionData := formsutl.SectionData{
			Name:   sec.Name,
			Fields: make([]formsutl.FieldData, len(sec.Fields)),
		}

		for fidx, fld := range sec.Fields {
			field := formsutl.FieldData{Name: fld.Name}

			if _, ok := fld.DefaultVal.(string); ok {
				val := r.FormValue(sec.Name + "::" + fld.Name)
				field.Value = val
			}

			if _, ok := fld.DefaultVal.(int); ok {
				val := r.FormValue(sec.Name + "::" + fld.Name)
				v, err := strconv.Atoi(val)
				if err != nil {
					return err
				}
				field.Value = v
			}

			if _, ok := fld.DefaultVal.([][]interface{}); ok {
				multifld, err := p.parseMultiField(r, sec, fld)
				if err != nil {
					return err
				}
				field.Value = multifld
			}

			sectionData.Fields[fidx] = field
		}

		parsedData[sidx] = sectionData
	}

	if err = sdkfs.WriteJson(p.dataPath(), parsedData); err != nil {
		return
	}

	fmt.Printf("parsedData: %+v\n", parsedData)

	p.parsedData = parsedData

	return
}

func (p *HttpFormInstance) GetStringValue(secname string, name string) (val string, err error) {
	v, err := p.getFieldValue(secname, name)
	if err != nil {
		return val, err
	}
	str, ok := v.(string)
	if !ok {
		return val, errors.New(fmt.Sprintf("section %s, field %s is not a string slice", secname, name))
	}
	return str, nil
}

func (p *HttpFormInstance) GetStringValues(secname string, name string) (val []string, err error) {
	v, err := p.getFieldValue(secname, name)
	if err != nil {
		return nil, err
	}
	str, ok := v.([]string)
	if !ok {
		return nil, errors.New(fmt.Sprintf("section %s, field %s is not a string slice", secname, name))
	}
	return str, nil
}

func (p *HttpFormInstance) GetIntValue(secname string, name string) (val int, err error) {
	v, err := p.getFieldValue(secname, name)
	if err != nil {
		return 0, err
	}
	num, ok := v.(int)
	if !ok {
		return 0, errors.New(fmt.Sprintf("section %s, field %s is not an int", secname, name))
	}
	return num, nil
}

func (p *HttpFormInstance) GetIntValues(secname string, name string) (val []int, err error) {
	v, err := p.getFieldValue(secname, name)
	if err != nil {
		return val, err
	}
	num, ok := v.([]int)
	if !ok {
		return val, errors.New(fmt.Sprintf("section %s, field %s is not an int", secname, name))
	}
	return num, nil
}

func (p *HttpFormInstance) GetBoolValue(secname string, name string) (val bool, err error) {
	v, err := p.getFieldValue(secname, name)
	if err != nil {
		return
	}
	if val, ok := v.(bool); ok {
		return val, nil
	}
	return false, errors.New(fmt.Sprintf("section %s, field %s is not a boolean", secname, name))
}

func (p *HttpFormInstance) GetBoolValues(secname string, name string) (val []bool, err error) {
	v, err := p.getFieldValue(secname, name)
	if err != nil {
		return
	}
	if val, ok := v.([]bool); ok {
		return val, nil
	}
	return val, errors.New(fmt.Sprintf("section %s, field %s is not a boolean", secname, name))
}

func (p *HttpFormInstance) GetFloatValue(secname string, name string) (val float64, err error) {
	v, err := p.getFieldValue(secname, name)
	if err != nil {
		return
	}
	if val, ok := v.(float64); ok {
		return val, nil
	}
	return val, errors.New(fmt.Sprintf("section %s, field %s is not a boolean", secname, name))
}

func (p *HttpFormInstance) GetFloatValues(secname string, name string) (val []float64, err error) {
	v, err := p.getFieldValue(secname, name)
	if err != nil {
		return
	}
	if val, ok := v.([]float64); ok {
		return val, nil
	}
	return val, errors.New(fmt.Sprintf("section %s, field %s is not a boolean", secname, name))
}

func (p *HttpFormInstance) GetMultiValue(secname string, name string) (val MultiFieldData, err error) {
	v, err := p.getFieldValue(secname, name)
	if err != nil {
		return
	}

	fields, ok := v.(MultiFieldData)
	if !ok {
		return val, errors.New(fmt.Sprintf("section %s, field %s is not a multi-field", secname, name))
	}

	return fields, nil
}

// start private funcs---------------------

func (p *HttpFormInstance) configPath() string {
	return filepath.Join(p.configdir, "config.json")
}

func (p *HttpFormInstance) dataPath() string {
	return filepath.Join(p.configdir, "data.json")
}

func (p *HttpFormInstance) writeConfig() error {
	return sdkfs.WriteJson(p.configPath(), p.sections)
}

func (p *HttpFormInstance) writeData() error {
	return sdkfs.WriteJson(p.dataPath(), p.parsedData)
}

func (p *HttpFormInstance) parseFieldValue(fld sdkhttp.Field, valstr string) (val interface{}, err error) {
	if _, ok := fld.DefaultVal.(string); ok {
		return valstr, nil
	}

	if _, ok := fld.DefaultVal.(int); ok {
		return strconv.Atoi(valstr)
	}

	if _, ok := fld.DefaultVal.([][]interface{}); ok {
		return nil, ErrFieldMulti
	}

	t := reflect.TypeOf(fld.DefaultVal)

	return nil, errors.New(fmt.Sprintf("field type %s not supported", t.Kind().String()))
}

func (p *HttpFormInstance) parseField(r *http.Request, sec sdkhttp.Section, fld sdkhttp.Field, valstr string) (field formsutl.FieldData, err error) {

	field = formsutl.FieldData{
		Name: fld.Name,
	}

	val, err := p.parseFieldValue(fld, valstr)
	if err != nil && errors.Is(err, ErrFieldMulti) {
		multifld, err := p.parseMultiField(r, sec, fld)
		if err != nil {
			return field, err
		}

		field.Value = multifld
		return field, nil
	}

	if err != nil {
		return field, err
	}

	field.Value = val

	fmt.Printf("parsed field: %+v\n", field)

	return field, nil
}

func (p *HttpFormInstance) parseMultiField(r *http.Request, sec sdkhttp.Section, fld sdkhttp.Field) (multifld MultiFieldData, err error) {
	if len(fld.Columns) < 1 {
		return multifld, errors.New(fmt.Sprintf("multi-field %s has no columns", fld.Name))
	}

	fmt.Printf("multi field form: %+v\n", r.Form)

	col1 := sec.Name + "::" + fld.Name + "::" + fld.Columns[0].Name + "[]"
	numRows := len(r.Form[col1])

	fmt.Printf("numRows: %d\n", numRows)

	multifld = MultiFieldData{
		Name:   fld.Name,
		Fields: make([][]formsutl.FieldData, numRows),
	}

	for ridx := 0; ridx < numRows; ridx++ {
		row := make([]formsutl.FieldData, len(fld.Columns))
		for cidx, col := range fld.Columns {
			colfld := sdkhttp.Field{
				Name:       col.Name,
				DefaultVal: col.DefaultVal,
			}
			inputName := sec.Name + "::" + fld.Name + "::" + colfld.Name + "[]"
			colarr := r.Form[inputName]
			fmt.Printf("colarr: %+v\n", colarr)
			valstr := colarr[ridx]
			fmt.Printf("valstr: %s\n", valstr)

			row[cidx], err = p.parseField(r, sec, colfld, valstr)
			if err != nil {
				return multifld, err
			}
		}
		fmt.Printf("row: %+v\n", row)
		multifld.Fields[ridx] = row
	}

	fmt.Printf("parsed multi-field: %+v\n", multifld)

	return multifld, nil

}

func (p *HttpFormInstance) getSection(secname string) (sec sdkhttp.Section, ok bool) {
	for _, s := range p.sections {
		if s.Name == secname {
			return s, true
		}
	}
	return
}

func (p *HttpFormInstance) getField(secname string, name string) (f sdkhttp.Field, ok bool) {
	for _, s := range p.sections {
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

func (p *HttpFormInstance) getParsedSection(secname string) (sec formsutl.SectionData, ok bool) {
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

func (p *HttpFormInstance) getParsedField(secname string, name string) (fld formsutl.FieldData, ok bool) {
	if s, ok := p.getParsedSection(secname); ok {
		for _, f := range s.Fields {
			if f.Name == name {
				return f, true
			}
		}

	}

	return
}

func (p *HttpFormInstance) getParsedFieldValue(secname string, name string) (val interface{}, ok bool) {
	if f, ok := p.getParsedField(secname, name); ok {
		return f.Value, true
	}
	return
}

func (p *HttpFormInstance) getDefaultValue(secname string, name string) (val interface{}, err error) {
	if f, ok := p.getField(secname, name); ok {
		return f.DefaultVal, nil
	}
	return nil, errors.New(fmt.Sprintf("section %s, field %s default value not found", secname, name))
}

func (p *HttpFormInstance) getFieldValue(secname string, name string) (val interface{}, err error) {
	if v, ok := p.getParsedFieldValue(secname, name); ok {
		return v, nil
	}

	return p.getDefaultValue(secname, name)
}
