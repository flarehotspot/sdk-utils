package formsutl

import (
	"errors"
	"fmt"
	"net/url"
	sdkforms "sdk/api/forms"
	"strconv"
)

var (
	ErrNotBasicType = fmt.Errorf("field type is not a basic type, i.e. string, number, bool")
)

func ParseBasicValue(fld sdkforms.FormField, valstr string) (val interface{}, err error) {
	switch fld.GetType() {
	case sdkforms.FormFieldTypeText:
		val = valstr
	case sdkforms.FormFieldTypeNumber:
		val, err = strconv.ParseFloat(valstr, 64)
		if err != nil {
			return 0, nil
		}
	case sdkforms.FormFieldTypeBoolean:
		val, err = strconv.ParseBool(valstr)
		if err != nil {
			return false, nil
		}
	default:
		err = ErrNotBasicType
	}
	return
}

func ParseListFieldValue(fld sdkforms.FormField, valstr []string) (val interface{}, err error) {
	listField, ok := fld.(sdkforms.ListField)
	if !ok {
		err = fmt.Errorf("field %s is not a list field", fld.GetName())
		return
	}

	switch listField.Type {

	case sdkforms.FormFieldTypeText:
		vals := valstr
		val = valstr
		if !listField.Multiple && len(vals) > 0 {
			val = vals[0]
		}

	case sdkforms.FormFieldTypeNumber:
		vals := make([]float64, len(valstr))
		for i, v := range valstr {
			vals[i], err = strconv.ParseFloat(v, 64)
			if err != nil {
				return 0, nil
			}
		}
		val = vals
		if !listField.Multiple && len(vals) > 0 {
			val = vals[0]
		}

	case sdkforms.FormFieldTypeBoolean:
		vals := make([]bool, len(valstr))
		for i, v := range valstr {
			vals[i], err = strconv.ParseBool(v)
			if err != nil {
				return false, nil
			}
		}
		val = vals
		if !listField.Multiple && len(vals) > 0 {
			val = vals[0]
		}

	default:
		err = errors.New(fmt.Sprintf("%s default value %s is not supported list field", fld.GetName(), listField.Type))
	}

	return
}

func ParseMultiFieldValue(sec sdkforms.FormSection, f sdkforms.FormField, form url.Values) (val [][]sdkforms.FieldData, err error) {
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
			case sdkforms.FormFieldTypeText, sdkforms.FormFieldTypeNumber, sdkforms.FormFieldTypeBoolean:
				if ridx >= len(colarr) {
					value = getTypeDefault(colfld.GetType())
					break
				}

				valstr := colarr[ridx]
				value, err = ParseBasicValue(colfld, valstr)
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

func getTypeDefault(t string) interface{} {
	switch t {
	case sdkforms.FormFieldTypeText:
		return ""
	case sdkforms.FormFieldTypeNumber:
		return 0
	case sdkforms.FormFieldTypeBoolean:
		return false
	default:
		return nil
	}
}
