package formsutl

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
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
	case sdkforms.FormFieldTypeBoolean:
		val, err = strconv.ParseBool(valstr)
	default:
		err = ErrNotBasicType
	}
	return
}

func ParseListFieldValue(fld sdkforms.FormField, valstr []string) (val interface{}, err error) {
	t := reflect.TypeOf(fld.GetDefaultVal()).Kind()
	if t != reflect.Slice {
		err = fmt.Errorf("field type is not a list")
		return
	}

	listField, ok := fld.(sdkforms.ListField)
	if !ok {
		err = fmt.Errorf("field type is not a list")
		return
	}

	switch listField.Type {

	case sdkforms.FormFieldTypeText:
		val = valstr

	case sdkforms.FormFieldTypeNumber:
		vals := make([]float64, len(valstr))
		for i, v := range valstr {
			vals[i], err = strconv.ParseFloat(v, 64)
			if err != nil {
				return
			}
		}
		val = vals

	case sdkforms.FormFieldTypeBoolean:
		vals := make([]bool, len(valstr))
		for i, v := range valstr {
			vals[i], err = strconv.ParseBool(v)
			if err != nil {
				return
			}
		}
		val = vals

	default:
		err = errors.New(fmt.Sprintf("%s default value %s is not supported list field", fld.GetName(), listField.Type))
	}

	return
}

func ParseMultiFieldValue(sec sdkforms.FormSection, f sdkforms.FormField, form url.Values) (val [][]FieldData, err error) {
	fmt.Printf("parsing multifield form: %+v\n", form)

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

	fmt.Printf("multi field form: %+v\n", form)

	col1 := sec.Name + "::" + fld.Name + "::" + columns[0].Name + "[]"
	numRows := len(form[col1])

	fmt.Printf("numRows: %d\n", numRows)

	vals := make([][]FieldData, numRows)

	for ridx := 0; ridx < numRows; ridx++ {
		row := make([]FieldData, len(columns))
		for cidx, colfld := range columns {

			inputName := sec.Name + "::" + fld.Name + "::" + colfld.Name + "[]"
			colarr := form[inputName]
			valstr := colarr[ridx]

			fmt.Printf("colarr: %+v\n", colarr)
			fmt.Printf("valstr: %s\n", valstr)

			var value interface{}

			switch colfld.GetType() {
			case sdkforms.FormFieldTypeText, sdkforms.FormFieldTypeNumber, sdkforms.FormFieldTypeBoolean:
				value, err = ParseBasicValue(colfld, valstr)
				if err != nil {
					return nil, err
				}

			default:
				err = errors.New(fmt.Sprintf("unsupported list field type %s", fld.GetType()))
				return
			}

			row[cidx] = FieldData{
				Name:  colfld.GetName(),
				Value: value,
			}
		}

		fmt.Printf("row: %+v\n", row)
		vals[ridx] = row
	}

	fmt.Printf("parsed multi-field: %+v\n", fld)

	return vals, nil

}
