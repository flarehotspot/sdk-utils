package plugins

import (
	formsutl "core/internal/utils/forms"
	formsview "core/resources/views/forms/bootstrap5"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	sdkforms "sdk/api/forms"
	sdkhttp "sdk/api/http"
	"sync"

	"github.com/a-h/templ"
	sdkfs "github.com/flarehotspot/go-utils/fs"
	sdkpaths "github.com/flarehotspot/go-utils/paths"
)

var (
	ErrFieldMulti = errors.New("field type is multifield")
)

func NewHttpForm(api *PluginApi, form sdkforms.Form) *HttpFormInstance {
	httpForm := &HttpFormInstance{
		api:  api,
		form: form,
		data: nil,
	}
	httpForm.LoadFormData()
	return httpForm
}

type HttpFormInstance struct {
	mu   sync.RWMutex
	api  *PluginApi
	form sdkforms.Form
	data []formsutl.SectionData
}

func (self *HttpFormInstance) Template(r *http.Request) templ.Component {
	csrfTag := self.api.HttpAPI.Helpers().CsrfHtmlTag(r)
	return formsview.HtmlForm(self, csrfTag, self.getSubmitUrl())
}

func (self *HttpFormInstance) LoadFormData() {
	if !sdkfs.Exists(self.dataPath()) {
		return
	}

	self.mu.Lock()
	defer self.mu.Unlock()

	if err := sdkfs.ReadJson(self.dataPath(), &self.data); err != nil {
		self.data = nil
	}
}

func (self *HttpFormInstance) GetSections() []sdkforms.FormSection {
	return self.form.Sections
}

func (self *HttpFormInstance) SaveForm(r *http.Request) (err error) {
	parsedData := make([]formsutl.SectionData, len(self.form.Sections))

	for sidx, sec := range self.form.Sections {
		sectionData := formsutl.SectionData{
			Name:   sec.Name,
			Fields: make([]sdkforms.FieldData, len(sec.Fields)),
		}

		for fidx, fld := range sec.Fields {
			field := sdkforms.FieldData{Name: fld.GetName()}
			valstr := r.Form[sec.Name+":"+fld.GetName()]

			if fld.GetType() != sdkforms.FormFieldTypeMulti && len(valstr) == 0 {
				continue
			}

			switch fld.GetType() {
			case sdkforms.FormFieldTypeText, sdkforms.FormFieldTypeNumber, sdkforms.FormFieldTypeBoolean:
				field.Value, err = formsutl.ParseBasicValue(fld, valstr[0])
				if err != nil {
					return err
				}

			case sdkforms.FormFieldTypeList:
				field.Value, err = formsutl.ParseListFieldValue(fld, valstr)
				if err != nil {
					return err
				}

			case sdkforms.FormFieldTypeMulti:
				val, err := formsutl.ParseMultiFieldValue(sec, fld, r.Form)
				if err != nil {
					return err
				}

				field.Value = formsutl.MultiFieldData{
					Fields: val,
				}

			default:
				return errors.New("invalid field type" + fld.GetType())
			}

			sectionData.Fields[fidx] = field
		}

		parsedData[sidx] = sectionData
	}

	if err = self.writeData(parsedData); err != nil {
		self.mu.Lock()
		self.data = nil
		self.mu.Unlock()
		return
	}

	self.LoadFormData()

	return
}

func (self *HttpFormInstance) GetStringValue(section string, field string) (val string, err error) {
	v, err := self.getFieldValue(section, field)
	if err != nil {
		return val, err
	}
	str, ok := v.(string)
	if !ok {
		return val, errors.New(fmt.Sprintf("section %s, field %s is not a string slice", section, field))
	}
	return str, nil
}

func (self *HttpFormInstance) GetStringValues(section string, field string) (val []string, err error) {
	ivals, err := self.getFieldValues(section, field)
	if err != nil {
		return nil, err
	}

	val = make([]string, len(ivals))
	for i, v := range ivals {
		val[i] = v.(string)
	}

	return val, nil
}

func (self *HttpFormInstance) GetBoolValue(section string, field string) (val bool, err error) {
	v, err := self.getFieldValue(section, field)
	if err != nil {
		return
	}
	if val, ok := v.(bool); ok {
		return val, nil
	}
	return false, errors.New(fmt.Sprintf("section %s, field %s is not a boolean", section, field))
}

func (self *HttpFormInstance) GetBoolValues(section string, field string) (val []bool, err error) {
	ivals, err := self.getFieldValues(section, field)
	if err != nil {
		return
	}

	val = make([]bool, len(ivals))
	for i, v := range ivals {
		val[i] = v.(bool)
	}

	return val, nil
}

func (self *HttpFormInstance) GetFloatValue(section string, field string) (val float64, err error) {
	v, err := self.getFieldValue(section, field)
	if err != nil {
		return
	}
	if val, ok := v.(float64); ok {
		return val, nil
	}
	return val, errors.New(fmt.Sprintf("section %s, field %s is not a boolean", section, field))
}

func (self *HttpFormInstance) GetFloatValues(section string, field string) (val []float64, err error) {
	ivals, err := self.getFieldValues(section, field)
	if err != nil {
		return
	}

	val = make([]float64, len(ivals))
	for i, v := range ivals {
		val[i] = v.(float64)
	}
	return val, nil
}

func (self *HttpFormInstance) GetIntValue(section string, field string) (val int, err error) {
	v, err := self.getFieldValue(section, field)
	if err != nil {
		return
	}
	if val, ok := v.(int); ok {
		return val, nil
	}
	return val, errors.New(fmt.Sprintf("section %s, field %s is not an int", section, field))
}

func (self *HttpFormInstance) GetIntValues(section string, field string) (val []int, err error) {
	ivals, err := self.getFieldValues(section, field)
	if err != nil {
		return
	}

	val = make([]int, len(ivals))
	for i, v := range ivals {
		val[i] = v.(int)
	}

	return val, nil
}

func (self *HttpFormInstance) GetMultiField(section string, field string) (val sdkforms.IMultiField, err error) {
	fld, ok := self.getField(section, field)
	if !ok {
		return val, fmt.Errorf("multi-field with name %s does not exist", field)
	}

	mfld, ok := fld.(sdkforms.MultiField)
	if !ok {
		return val, fmt.Errorf("form field %s is not a multi-field", field)
	}

	v, err := self.getFieldValue(section, field)
	if err != nil {
		return
	}

	ivals, ok := v.(map[string]interface{})
	if !ok {
		return val, errors.New(fmt.Sprintf("section %s, field %s is not a multi-field, instead %T", section, field, v))
	}

	ifields, ok := ivals["fields"]
	if !ok {
		return val, errors.New(fmt.Sprintf("multi-field %s value has no 'fields' field", field))
	}

	irows, ok := ifields.([]interface{})
	if !ok {
		return val, fmt.Errorf("multi-field %s value is not a slice of field data, instead %T", field, ifields)
	}

	mfd := formsutl.MultiFieldData{Fields: make([][]sdkforms.FieldData, len(irows))}

	for ridx, irow := range irows {
		icols, ok := irow.([]interface{})
		if !ok {
			return val, fmt.Errorf("multi-field %s row is not a slice of field data, instead %T", field, irow)
		}

		cols := mfld.Columns()
		row := make([]sdkforms.FieldData, len(cols))

		for cidx, colfld := range cols {
			fd := sdkforms.FieldData{Name: colfld.Name}

			if cidx > (len(icols) - 1) {
				row[cidx] = fd
				continue
			}

			icol := icols[cidx]

			colmap, ok := icol.(map[string]interface{})
			if !ok {
				return val, fmt.Errorf("multi-field column %s is not a field data, instead %T", colfld.Name, icol)
			}

			v, ok := colmap["value"]
			if !ok {
				return val, fmt.Errorf("multi-field column %s does not have a value field", colfld.Name)
			}

			fd.Value = v

			row[cidx] = fd
		}

		mfd.Fields[ridx] = row
	}

	return mfd, nil
}

func (self *HttpFormInstance) GetRedirectUrl() string {
	url := self.api.HttpAPI.httpRouter.UrlForRoute(sdkhttp.PluginRouteName(self.form.CallbackRoute))
	return url
}

// start private funcs---------------------
func (self *HttpFormInstance) dataPath() string {
	return filepath.Join(sdkpaths.ConfigDir, "plugins", self.api.Pkg(), self.form.Name+".json")
}

func (self *HttpFormInstance) writeData(parsedData []formsutl.SectionData) error {
	savepath := self.dataPath()
	if err := sdkfs.EnsureDir(filepath.Dir(savepath)); err != nil {
		return err
	}
	return sdkfs.WriteJson(savepath, parsedData)
}

func (self *HttpFormInstance) getSection(section string) (sec sdkforms.FormSection, ok bool) {
	for _, s := range self.form.Sections {
		if s.Name == section {
			return s, true
		}
	}
	return
}

func (self *HttpFormInstance) getField(section string, field string) (f sdkforms.IFormField, ok bool) {
	for _, s := range self.form.Sections {
		if s.Name == section {
			for _, fld := range s.Fields {
				if fld.GetName() == field {
					return fld, true
				}
			}
		}
	}
	return
}

func (self *HttpFormInstance) getParsedSection(section string) (sec formsutl.SectionData, ok bool) {
	data := self.getFormData()
	if data == nil {
		return
	}

	for _, s := range data {
		if s.Name == section {
			return s, true
		}
	}

	return
}

func (self *HttpFormInstance) getParsedField(section string, field string) (fld sdkforms.FieldData, ok bool) {
	if s, ok := self.getParsedSection(section); ok {
		for _, f := range s.Fields {
			if f.Name == field {
				return f, true
			}
		}

	}

	return
}

func (self *HttpFormInstance) getParsedFieldValue(section string, field string) (val interface{}, ok bool) {
	if f, ok := self.getParsedField(section, field); ok {
		return f.Value, true
	}
	return
}

func (self *HttpFormInstance) getDefaultValue(secname string, field string) (val interface{}, err error) {
	if f, ok := self.getField(secname, field); ok {
		return f.GetDefaultVal(), nil
	}
	return nil, errors.New(fmt.Sprintf("section %s, field %s default value not found", secname, field))
}

func (self *HttpFormInstance) getDefaultValues(secname string, field string) (val []interface{}, err error) {
	f, ok := self.getField(secname, field)
	if !ok {
		return nil, errors.New(fmt.Sprintf("section %s, field %s default value not found", secname, field))
	}

	v := f.GetDefaultVal()
	if val, ok = v.([]interface{}); !ok {
		return nil, errors.New(fmt.Sprintf("section %s, field %s default value is not a slice", secname, field))
	}

	return val, nil
}

func (self *HttpFormInstance) getFieldValue(section string, field string) (val interface{}, err error) {
	if v, ok := self.getParsedFieldValue(section, field); ok {
		return v, nil
	}

	return self.getDefaultValue(section, field)
}

func (self *HttpFormInstance) getFieldValues(section string, field string) (val []interface{}, err error) {
	v, ok := self.getParsedFieldValue(section, field)
	if !ok {
		return self.getDefaultValues(section, field)
	}

	if val, ok = v.([]interface{}); !ok {
		return self.getDefaultValues(section, field)
	}

	return val, nil
}

func (self *HttpFormInstance) getSubmitUrl() string {
	return self.api.CoreAPI.HttpAPI.httpRouter.UrlForRoute("admin:forms:save", "pkg", self.api.Pkg(), "name", self.form.Name)
}

func (self *HttpFormInstance) getFormData() []formsutl.SectionData {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.data
}
