package formsutl

import (
	"errors"
	"fmt"
	"reflect"
	sdkforms "sdk/api/forms"
	"strings"
)

type MultiFieldData struct {
	Fields [][]sdkforms.FieldData `json:"fields"`
}

func (f MultiFieldData) NumRows() int {
	return len(f.Fields)
}

func (f MultiFieldData) Json() string {
	var s strings.Builder
	s.WriteString("[")

	for i, row := range f.Fields {
		if i > 0 {
			s.WriteString(", ")
		}
		s.WriteString("{ ")

		for j, field := range row {
			if j > 0 {
				s.WriteString(", ")
			}

			s.WriteString(fmt.Sprintf(`"%s": `, field.Name))

			typ := reflect.TypeOf(field.Value)

			switch typ.Kind() {
			case reflect.String:
				s.WriteString(fmt.Sprintf(`"%s" `, field.Value))

			case reflect.Float64:
				s.WriteString(fmt.Sprintf("%f", field.Value))

			case reflect.Bool:
				s.WriteString(fmt.Sprintf("%t", field.Value))

			default:
				s.WriteString("null")
			}

		}

		s.WriteString(" }")
	}

	s.WriteString("]")

	return s.String()
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
