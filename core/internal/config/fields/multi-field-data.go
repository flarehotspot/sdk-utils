package cfgfields

import (
	"errors"
	"fmt"
)

type MultiFieldData struct {
	Name   string        `json:"name"`
	Fields [][]FieldData `json:"fields"`
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

	return "", errors.New(fmt.Sprintf("field %s not found in multi-field %s", name, f.Name))
}

func (f MultiFieldData) GetStringValue(row int, name string) (val string, err error) {
	v, err := f.GetValue(row, name)
	if err != nil {
		return "", err
	}

	val, ok := v.(string)
	if !ok {
		return "", errors.New(fmt.Sprintf("field %s in row %d in multi-field %s is not a string", name, row, f.Name))
	}

	return val, nil
}

func (f MultiFieldData) GetIntValue(row int, name string) (val int, err error) {
	v, err := f.GetValue(row, name)
	if err != nil {
		return 0, err
	}

	val, ok := v.(int)
	if !ok {
		return 0, errors.New(fmt.Sprintf("field %s in row %d in multi-field %s is not an int", name, row, f.Name))
	}

	return val, nil
}
