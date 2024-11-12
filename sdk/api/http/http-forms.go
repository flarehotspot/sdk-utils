package sdkhttp

import (
	"net/http"

	"github.com/a-h/templ"
)

type ColField struct {
	Name        string
	Validations string
	DefaultVal  interface{}
}

type FieldOption struct {
	Label string
	Value interface{}
}

type Field struct {
	Name        string
	Label       string
	Columns     []ColField    // for multi-field inputs
	Options     []FieldOption // for list/select inputs
	Validations string
	DefaultVal  interface{}
}

type Section struct {
	Name        string
	Title       string
	Description string
	Fields      []Field
}

type IMultiField interface {
	GetStringValue(row int, name string) (string, error)
	GetStringValues(row int, name string) ([]string, error)

	GetIntValue(row int, name string) (int, error)
	GetIntValues(row int, name string) ([]int, error)

	GetFloatValue(row int, name string) (float64, error)
	GetFloatValues(row int, name string) ([]float64, error)

	GetBoolValue(row int, name string) (bool, error)
	GetBoolValues(row int, name string) ([]bool, error)
}

type HttpForm interface {
	Template(r *http.Request) templ.Component

	GetStringValue(section string, name string) (string, error)
	GetStringValues(section string, name string) ([]string, error)

	GetIntValue(section string, name string) (int, error)
	GetIntValues(section string, name string) ([]int, error)

	GetFloatValue(section string, name string) (float64, error)
	GetFloatValues(section string, name string) ([]float64, error)

	GetBoolValue(section string, name string) (bool, error)
	GetBoolValues(section string, name string) ([]bool, error)
}

type HttpFormApi interface {
	NewHttpForm(key string, sections []Section) (HttpForm, error)
	SaveForm(r *http.Request) (HttpForm, error)
	GetForm(key string) (HttpForm, error)
}
