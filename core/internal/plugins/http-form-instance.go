package plugins

import (
	formsview "core/resources/views/forms/bootstrap5"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	sdkforms "sdk/api/forms"
	sdkhttp "sdk/api/http"
	"strconv"

	"github.com/a-h/templ"
)

var (
	ErrFieldMulti   = errors.New("field type is multifield")
	ErrNotBasicType = fmt.Errorf("field type is not a basic type, e.g. string, integer, decimal, bool")
)

func NewHttpForm(api *PluginApi, form sdkforms.Form) *HttpFormInstance {
	return &HttpFormInstance{
		api:  api,
		form: form,
	}
}

type HttpFormInstance struct {
	api  *PluginApi
	form sdkforms.Form
	data []sdkforms.SectionData
}

func (self *HttpFormInstance) GetTemplate(r *http.Request) templ.Component {
	csrfTag := self.api.HttpAPI.Helpers().CsrfHtmlTag(r)
	submitText := "Submit"
	if self.form.SubmitLabel != "" {
		submitText = self.form.SubmitLabel
	}
	submitUrl := self.api.HttpAPI.httpRouter.UrlForRoute(sdkhttp.PluginRouteName(self.form.CallbackRoute))
	return formsview.HtmlForm(self, csrfTag, submitUrl, submitText)
}

func (self *HttpFormInstance) ParseForm(r *http.Request) (err error) {
	if err := r.ParseForm(); err != nil {
		return err
	}

	parsedData := make([]sdkforms.SectionData, len(self.form.Sections))

	for sidx, sec := range self.form.Sections {
		sectionData := sdkforms.SectionData{
			Name:   sec.Name,
			Fields: make([]sdkforms.FieldData, len(sec.Fields)),
		}

		for fidx, fld := range sec.Fields {
			field := sdkforms.FieldData{Name: fld.GetName()}
			valstr := r.Form[sec.Name+":"+fld.GetName()]

			switch fld.GetType() {

			case sdkforms.FormFieldTypeText,
				sdkforms.FormFieldTypeInteger,
				sdkforms.FormFieldTypeDecimal,
				sdkforms.FormFieldTypeBoolean:
				field.Value, err = ParseBasicValue(fld, valstr)
				if err != nil {
					field.Value = fld.GetValue()
				}

			case sdkforms.FormFieldTypeList:
				field.Value, err = ParseListFieldValue(fld, valstr)
				if err != nil {
					field.Value = fld.GetValue()
				}

			case sdkforms.FormFieldTypeMulti:
				val, err := ParseMultiFieldValue(sec, fld, r.Form)
				if err != nil {
					mfld, ok := fld.(sdkforms.MultiField)
					if !ok {
						return fmt.Errorf("section %s, field %s type is not multifield, instead %T", sec, fld.GetName(), fld)
					}

					fldvals := mfld.GetValue()
					mfldval, ok := fldvals.([][]sdkforms.FieldData)
					if !ok {
						return fmt.Errorf("section %s, field %s value is not a slice of sdkforms.FieldData, instead %T", sec, fld.GetName(), fldvals)
					}
					val = mfldval
				}
				field.Value = sdkforms.MultiFieldData{
					Fields: val,
				}

			default:
				return errors.New("invalid field type" + fld.GetType())
			}

			if field.Value == nil {
				field.Value = GetTypeDefault(fld)
			}

			sectionData.Fields[fidx] = field
		}

		parsedData[sidx] = sectionData
	}

	self.data = parsedData
	return nil
}

func (self *HttpFormInstance) GetSections() []sdkforms.FormSection {
	return self.form.Sections
}

func (self *HttpFormInstance) GetStringValue(section string, field string) (val string, err error) {
	v, err := self.getFieldValue(section, field)
	if err != nil {
		return val, err
	}
	str, ok := v.(string)
	if !ok {
		return val, errors.New(fmt.Sprintf("section %s field %s is not a string, instead %T", section, field, v))
	}
	return str, nil
}

func (self *HttpFormInstance) GetStringValues(section string, field string) (val []string, err error) {
	ivals, err := self.getFieldValues(section, field)
	if err != nil {
		return nil, err
	}

	val, ok := ivals.([]string)
	if !ok {
		return nil, errors.New(fmt.Sprintf("section %s, field %s is not a slice of strings", section, field))
	}

	return val, nil
}

func (self *HttpFormInstance) GetIntValue(section string, field string) (val int64, err error) {
	v, err := self.getFieldValue(section, field)
	if err != nil {
		return
	}
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		return int64(v.(float64)), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.(int64), nil
	}
	return val, errors.New(fmt.Sprintf("section %s, field %s is not an int", section, field))
}

func (self *HttpFormInstance) GetIntValues(section string, field string) (val []int64, err error) {
	ivals, err := self.getFieldValues(section, field)
	if err != nil {
		return
	}

	t := reflect.TypeOf(ivals).Elem()
	val = []int64{}

	switch t.Kind() {
	case reflect.Int64:
		vals := ivals.([]int64)
		return vals, nil
	default:
		return nil, errors.New(fmt.Sprintf("section %s, field %s is not a slice of int64", section, field))
	}
}

func (self *HttpFormInstance) GetFloatValue(section string, field string) (val float64, err error) {
	v, err := self.getFieldValue(section, field)
	if err != nil {
		return
	}
	if val, ok := v.(float64); ok {
		return val, nil
	}
	return val, errors.New(fmt.Sprintf("section %s, field %s is not a float64", section, field))
}

func (self *HttpFormInstance) GetFloatValues(section string, field string) (val []float64, err error) {
	ivals, err := self.getFieldValues(section, field)
	if err != nil {
		return
	}

	t := reflect.TypeOf(ivals).Elem()
	switch t.Kind() {
	case reflect.Float64:
		vals := ivals.([]float64)
		return vals, nil
	default:
		return nil, errors.New(fmt.Sprintf("section %s, field %s is not a slice of float64", section, field))
	}
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

	t := reflect.TypeOf(ivals).Elem()
	switch t.Kind() {
	case reflect.Bool:
		vals := ivals.([]bool)
		return vals, nil
	default:
		return nil, errors.New(fmt.Sprintf("section %s, field %s is not a slice of boolean", section, field))
	}
}

func (self *HttpFormInstance) GetMultiField(section string, field string) (val sdkforms.IMultiField, err error) {
	v, err := self.getFieldValue(section, field)
	if err != nil {
		return
	}

	mfd, ok := v.(sdkforms.MultiFieldData)
	if !ok {
		return val, errors.New(fmt.Sprintf("section %s, field %s value is not sdkforms.MultiFieldData, instead %T", section, field, v))
	}

	return mfd, nil
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

func (self *HttpFormInstance) getParsedSection(section string) (sec sdkforms.SectionData, ok bool) {
	data := self.data
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

func (self *HttpFormInstance) getFieldValue(section string, field string) (val interface{}, err error) {
	if self.data == nil {
		fld, ok := self.getField(section, field)
		if !ok {
			return nil, errors.New(fmt.Sprintf("section %s, field %s value not found", section, field))
		}
		return fld.GetValue(), nil
	}

	if v, ok := self.getParsedFieldValue(section, field); ok {
		return v, nil
	}

	return nil, errors.New(fmt.Sprintf("section %s, field %s value not found", section, field))
}

func (self *HttpFormInstance) getFieldValues(section string, field string) (val interface{}, err error) {
	var ok bool
	if self.data == nil {
		fld, ok := self.getField(section, field)
		if !ok {
			return nil, errors.New(fmt.Sprintf("section %s, field %s value not found", section, field))
		}
		val = fld.GetValue()
	} else {
		val, ok = self.getParsedFieldValue(section, field)
		if !ok {
			return nil, errors.New(fmt.Sprintf("section %s, field %s values not found", section, field))
		}
	}

	if reflect.TypeOf(val).Kind() != reflect.Slice {
		return nil, errors.New(fmt.Sprintf("section %s, field %s values is not a slice", section, field))
	}

	return val, nil
}

// ----- Parser functions ----
func ParseBasicValue(fld sdkforms.IFormField, valstr []string) (val interface{}, err error) {
	switch fld.GetType() {
	case sdkforms.FormFieldTypeText:
		if len(valstr) < 1 {
			return "", nil
		}
		val = valstr[0]

	case sdkforms.FormFieldTypeInteger:
		if len(valstr) < 1 {
			return 0, nil
		}
		val, err = strconv.ParseInt(valstr[0], 10, 64)
		if err != nil {
			return 0, nil
		}
	case sdkforms.FormFieldTypeDecimal:
		if len(valstr) < 1 {
			return 0.0, nil
		}
		val, err = strconv.ParseFloat(valstr[0], 64)
		if err != nil {
			return 0, nil
		}
	case sdkforms.FormFieldTypeBoolean:
		if len(valstr) < 1 {
			return false, nil
		}
		val, err = strconv.ParseBool(valstr[0])
		if err != nil {
			return false, nil
		}
	default:
		err = ErrNotBasicType
	}
	return
}

func ParseListFieldValue(fld sdkforms.IFormField, valstr []string) (val interface{}, err error) {
	listField, ok := fld.(sdkforms.ListField)
	if !ok {
		err = fmt.Errorf("field %s is not a list field", fld.GetName())
		return
	}

	if valstr == nil {
		return GetTypeDefault(fld), nil
	}

	switch listField.Type {

	case sdkforms.FormFieldTypeText:
		vals := valstr
		val = valstr
		if !listField.Multiple {
			if len(vals) > 0 {
				val = vals[0]
				return
			}
			val = ""
		}
		return

	case sdkforms.FormFieldTypeInteger:
		vals := make([]int64, len(valstr))
		for i, v := range valstr {
			vals[i], err = strconv.ParseInt(v, 10, 64)
			if err != nil {
				return 0, nil
			}
		}
		val = vals
		if !listField.Multiple {
			if len(vals) > 0 {
				val = vals[0]
				return
			}
			val = 0
		}
		return

	case sdkforms.FormFieldTypeDecimal:
		vals := make([]float64, len(valstr))
		for i, v := range valstr {
			vals[i], err = strconv.ParseFloat(v, 64)
			if err != nil {
				return 0, nil
			}
		}
		val = vals
		if !listField.Multiple {
			if len(vals) > 0 {
				val = vals[0]
				return
			}
			val = 0.0
		}
		return

	case sdkforms.FormFieldTypeBoolean:
		vals := make([]bool, len(valstr))
		for i, v := range valstr {
			vals[i], err = strconv.ParseBool(v)
			if err != nil {
				return false, nil
			}
		}
		val = vals
		if !listField.Multiple {
			if len(vals) > 0 {
				val = vals[0]
				return
			}
			val = false
		}
		return

	default:
		err = errors.New(fmt.Sprintf("%s default value %s is not supported list field", fld.GetName(), listField.Type))
	}

	return
}

func ParseMultiFieldValue(sec sdkforms.FormSection, f sdkforms.IFormField, form url.Values) (val [][]sdkforms.FieldData, err error) {
	fld, ok := f.(sdkforms.MultiField)
	if !ok {
		err = errors.New(fmt.Sprintf("field %s in section %s is not a multi-field", f.GetName(), sec.Name))
		return
	}

	columns := fld.Columns()
	if len(columns) < 1 {
		err = errors.New(fmt.Sprintf("multi-field %s in section %s has no columns", fld.Name, sec.Name))
		return
	}

	col1 := sec.Name + ":" + fld.Name + ":" + columns[0].Name
	numRows := len(form[col1])

	vals := make([][]sdkforms.FieldData, numRows)

	for ridx := 0; ridx < numRows; ridx++ {
		row := make([]sdkforms.FieldData, len(columns))
		for cidx, colfld := range columns {
			var value interface{}

			inputName := sec.Name + ":" + fld.Name + ":" + colfld.Name
			colarr := form[inputName]

			switch colfld.GetType() {

			case sdkforms.FormFieldTypeText,
				sdkforms.FormFieldTypeInteger,
				sdkforms.FormFieldTypeDecimal,
				sdkforms.FormFieldTypeBoolean:

				if ridx >= len(colarr) {
					value = GetTypeDefault(colfld)
					break
				}

				valstr := colarr[ridx]
				value, err = ParseBasicValue(colfld, []string{valstr})
				if err != nil {
					return nil, err
				}

			default:
				err = errors.New(fmt.Sprintf("unsupported list field type %s", colfld.GetType()))
				return
			}

			row[cidx] = sdkforms.FieldData{
				Name:  colfld.GetName(),
				Value: value,
			}
		}

		vals[ridx] = row
	}

	return vals, nil

}

func GetTypeDefault(fld sdkforms.IFormField) interface{} {
	switch fld.GetType() {

	case sdkforms.FormFieldTypeText,
		sdkforms.FormFieldTypeInteger,
		sdkforms.FormFieldTypeDecimal,
		sdkforms.FormFieldTypeBoolean:
		return GetBasicTypeDefault(fld.GetType())

	case sdkforms.FormFieldTypeList:
		lsfld := fld.(sdkforms.ListField)
		if lsfld.Multiple {
			return []interface{}{}
		} else {
			return GetBasicTypeDefault(fld.GetType())
		}

	case sdkforms.FormFieldTypeMulti:
		return map[string]interface{}{}

	default:
		return nil
	}
}

func GetBasicTypeDefault(t string) interface{} {
	switch t {
	case sdkforms.FormFieldTypeText:
		return ""
	case sdkforms.FormFieldTypeInteger:
		return int64(0)
	case sdkforms.FormFieldTypeDecimal:
		return float64(0.0)
	case sdkforms.FormFieldTypeBoolean:
		return false
	default:
		return nil
	}
}
