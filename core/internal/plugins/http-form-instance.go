package plugins

import (
	formsutl "core/internal/utils/forms"
	formsview "core/resources/views/forms/bootstrap5"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	sdkforms "sdk/api/forms"

	"github.com/a-h/templ"
	sdkfs "github.com/flarehotspot/go-utils/fs"
	sdkpaths "github.com/flarehotspot/go-utils/paths"
)

var (
	ErrFieldMulti = errors.New("field type is multifield")
)

func NewHttpForm(api *PluginApi, form sdkforms.Form) *HttpFormInstance {
	httpForm := &HttpFormInstance{
		api:        api,
		form:       form,
		parsedData: nil,
	}
	httpForm.LoadFormData()
	return httpForm
}

type HttpFormInstance struct {
	api        *PluginApi
	form       sdkforms.Form
	parsedData []formsutl.SectionData
}

func (self *HttpFormInstance) Template(r *http.Request) templ.Component {
	csrfTag := self.api.HttpAPI.Helpers().CsrfHtmlTag(r)
	return formsview.HtmlForm(self, csrfTag, self.getSubmitUrl())
}

func (self *HttpFormInstance) LoadFormData() {
	if !sdkfs.Exists(self.dataPath()) {
		return
	}
	fmt.Println("Loading values from", self.dataPath())
	if err := sdkfs.ReadJson(self.dataPath(), &self.parsedData); err != nil {
		self.parsedData = nil
	}
}

func (self *HttpFormInstance) GetSections() []sdkforms.FormSection {
	return self.form.Sections
}

func (self *HttpFormInstance) GetFormData() []formsutl.SectionData {
	return self.parsedData
}

func (self *HttpFormInstance) SaveForm(r *http.Request) (err error) {
	parsedData := make([]formsutl.SectionData, len(self.form.Sections))

	for sidx, sec := range self.form.Sections {
		sectionData := formsutl.SectionData{
			Name:   sec.Name,
			Fields: make([]formsutl.FieldData, len(sec.Fields)),
		}

		for fidx, fld := range sec.Fields {
			field := formsutl.FieldData{Name: fld.GetName()}
			valstr := r.Form[sec.Name+"::"+fld.GetName()]
			if len(valstr) == 0 {
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

	self.parsedData = parsedData

	if err = self.writeData(); err != nil {
		self.parsedData = nil
		return
	}

	fmt.Printf("parsedData: %+v\n", parsedData)

	self.parsedData = parsedData

	return
}

func (self *HttpFormInstance) GetStringValue(secname string, name string) (val string, err error) {
	v, err := self.getFieldValue(secname, name)
	if err != nil {
		return val, err
	}
	str, ok := v.(string)
	if !ok {
		return val, errors.New(fmt.Sprintf("section %s, field %s is not a string slice", secname, name))
	}
	return str, nil
}

func (self *HttpFormInstance) GetStringValues(secname string, name string) (val []string, err error) {
	v, err := self.getFieldValue(secname, name)
	if err != nil {
		return nil, err
	}
	str, ok := v.([]string)
	if !ok {
		return nil, errors.New(fmt.Sprintf("section %s, field %s is not a string slice", secname, name))
	}
	return str, nil
}

func (self *HttpFormInstance) GetIntValue(secname string, name string) (val int, err error) {
	v, err := self.getFieldValue(secname, name)
	if err != nil {
		return 0, err
	}
	num, ok := v.(int)
	if !ok {
		return 0, errors.New(fmt.Sprintf("section %s, field %s is not an int", secname, name))
	}
	return num, nil
}

func (self *HttpFormInstance) GetIntValues(secname string, name string) (val []int, err error) {
	v, err := self.getFieldValue(secname, name)
	if err != nil {
		return val, err
	}
	num, ok := v.([]int)
	if !ok {
		return val, errors.New(fmt.Sprintf("section %s, field %s is not an int", secname, name))
	}
	return num, nil
}

func (self *HttpFormInstance) GetBoolValue(secname string, name string) (val bool, err error) {
	v, err := self.getFieldValue(secname, name)
	if err != nil {
		return
	}
	if val, ok := v.(bool); ok {
		return val, nil
	}
	return false, errors.New(fmt.Sprintf("section %s, field %s is not a boolean", secname, name))
}

func (self *HttpFormInstance) GetBoolValues(secname string, name string) (val []bool, err error) {
	v, err := self.getFieldValue(secname, name)
	if err != nil {
		return
	}
	if val, ok := v.([]bool); ok {
		return val, nil
	}
	return val, errors.New(fmt.Sprintf("section %s, field %s is not a boolean", secname, name))
}

func (self *HttpFormInstance) GetFloatValue(secname string, name string) (val float64, err error) {
	v, err := self.getFieldValue(secname, name)
	if err != nil {
		return
	}
	if val, ok := v.(float64); ok {
		return val, nil
	}
	return val, errors.New(fmt.Sprintf("section %s, field %s is not a boolean", secname, name))
}

func (self *HttpFormInstance) GetFloatValues(secname string, name string) (val []float64, err error) {
	v, err := self.getFieldValue(secname, name)
	if err != nil {
		return
	}
	if val, ok := v.([]float64); ok {
		return val, nil
	}
	return val, errors.New(fmt.Sprintf("section %s, field %s is not a boolean", secname, name))
}

func (self *HttpFormInstance) GetMultiField(secname string, name string) (val sdkforms.IMultiField, err error) {
	v, err := self.getFieldValue(secname, name)
	if err != nil {
		return
	}

	fields, ok := v.(formsutl.MultiFieldData)
	if !ok {
		return val, errors.New(fmt.Sprintf("section %s, field %s is not a multi-field", secname, name))
	}

	return fields, nil
}

// start private funcs---------------------
func (self *HttpFormInstance) dataPath() string {
	return filepath.Join(sdkpaths.ConfigDir, "plugins", self.api.Pkg(), self.form.Name+".json")
}

func (self *HttpFormInstance) writeData() error {
	savepath := self.dataPath()
	if err := sdkfs.EnsureDir(filepath.Dir(savepath)); err != nil {
		return err
	}
	return sdkfs.WriteJson(savepath, self.parsedData)
}

func (self *HttpFormInstance) getSection(secname string) (sec sdkforms.FormSection, ok bool) {
	for _, s := range self.form.Sections {
		if s.Name == secname {
			return s, true
		}
	}
	return
}

func (self *HttpFormInstance) getField(secname string, name string) (f sdkforms.FormField, ok bool) {
	for _, s := range self.form.Sections {
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

func (self *HttpFormInstance) getParsedSection(secname string) (sec formsutl.SectionData, ok bool) {
	if self.parsedData == nil {
		return
	}

	for _, s := range self.parsedData {
		if s.Name == secname {
			return s, true
		}
	}

	return
}

func (self *HttpFormInstance) getParsedField(secname string, name string) (fld formsutl.FieldData, ok bool) {
	if s, ok := self.getParsedSection(secname); ok {
		for _, f := range s.Fields {
			if f.Name == name {
				return f, true
			}
		}

	}

	return
}

func (self *HttpFormInstance) getParsedFieldValue(secname string, name string) (val interface{}, ok bool) {
	if f, ok := self.getParsedField(secname, name); ok {
		return f.Value, true
	}
	return
}

func (self *HttpFormInstance) getDefaultValue(secname string, name string) (val interface{}, err error) {
	if f, ok := self.getField(secname, name); ok {
		return f.GetDefaultVal(), nil
	}
	return nil, errors.New(fmt.Sprintf("section %s, field %s default value not found", secname, name))
}

func (self *HttpFormInstance) getFieldValue(secname string, name string) (val interface{}, err error) {
	if v, ok := self.getParsedFieldValue(secname, name); ok {
		return v, nil
	}

	return self.getDefaultValue(secname, name)
}

func (self *HttpFormInstance) getSubmitUrl() string {
	return self.api.CoreAPI.HttpAPI.httpRouter.UrlForRoute("admin:forms:save", "pkg", self.api.Pkg(), "name", self.form.Name)
}
