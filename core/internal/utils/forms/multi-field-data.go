package formsutl

import (
	"errors"
	"fmt"
	sdkforms "sdk/api/forms"
)

type MultiFieldData struct {
	Fields [][]sdkforms.FieldData `json:"fields"`
}

func (f MultiFieldData) NumRows() int {
	return len(f.Fields)
}

func (f MultiFieldData) GetValue(row int, name string) (val interface{}, err error) {
	r := f.Fields[row]
	if r == nil {
		return "", errors.New(fmt.Sprintf("row %d not found", row))
	}

	for _, field := range r {
		if field.Name == name {
			return field.Value, nil
		}
	}

	return "", errors.New(fmt.Sprintf("field %s not found in multi-field", name))
}

func (f MultiFieldData) GetStringValue(row int, name string) (val string, err error) {
	v, err := f.GetValue(row, name)
	if err != nil {
		return "", err
	}

	val, ok := v.(string)
	if !ok {
		return "", errors.New(fmt.Sprintf("field %s in row %d in multi-field is not a string", name, row))
	}

	return val, nil
}

func (f MultiFieldData) GetFloatValue(row int, name string) (val float64, err error) {
	v, err := f.GetValue(row, name)
	if err != nil {
		return 0, err
	}

	val, ok := v.(float64)
	if !ok {
		return 0, errors.New(fmt.Sprintf("field %s in row %d in multi-field is not float64", name, row))
	}

	return val, nil
}

func (f MultiFieldData) GetBoolValue(row int, name string) (val bool, err error) {
	v, err := f.GetValue(row, name)
	if err != nil {
		return
	}

	val, ok := v.(bool)
	if !ok {
		err = errors.New(fmt.Sprintf("field %s in row %d in multi-field is not a boolean", name, row))
		return
	}

	return val, nil
}
