package sdkhttp

import (
	sdkforms "sdk/api/forms"
)

type HttpFormApi interface {
	RegisterHttpForms(forms ...sdkforms.Form) (err error)
	GetForm(name string) (form sdkforms.IHttpForm, err error)
}
