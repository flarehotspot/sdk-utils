package sdkforms

import (
	"net/http"

	"github.com/a-h/templ"
)

type IHttpForm interface {
	Template(r *http.Request) templ.Component

	GetSections() []FormSection

	GetStringValue(section string, name string) (string, error)
	GetStringValues(section string, name string) ([]string, error)

	GetIntValue(section string, name string) (int, error)
	GetIntValues(section string, name string) ([]int, error)

	GetFloatValue(section string, name string) (float64, error)
	GetFloatValues(section string, name string) ([]float64, error)

	GetBoolValue(section string, name string) (bool, error)
	GetBoolValues(section string, name string) ([]bool, error)

	GetMultiField(section string, name string) (IMultiField, error)
}

type Form struct {
	Name          string
	CallbackRoute string
	Sections      []FormSection
}
