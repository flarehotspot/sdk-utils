package sdkhttp

import (
	"net/http"
	sdkforms "sdk/api/forms"

	"github.com/a-h/templ"
)

type IHttpFormApi interface {
	RegisterForms(forms ...sdkforms.Form) (err error)
	GetForm(name string) (form IHttpForm, ok bool)
}

type IHttpForm interface {
	GetTemplate(r *http.Request) templ.Component

	GetSections() []sdkforms.FormSection

	GetStringValue(section string, name string) (string, error)
	GetStringValues(section string, name string) ([]string, error)

	GetIntValue(section string, name string) (int64, error)
	GetIntValues(section string, name string) ([]int64, error)

	GetFloatValue(section string, name string) (float64, error)
	GetFloatValues(section string, name string) ([]float64, error)

	GetBoolValue(section string, name string) (bool, error)
	GetBoolValues(section string, name string) ([]bool, error)

	GetMultiField(section string, name string) (sdkforms.IMultiField, error)

	ParseForm(r *http.Request) error
}
