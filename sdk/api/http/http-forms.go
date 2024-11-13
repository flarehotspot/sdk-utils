package sdkhttp

import (
	sdkforms "sdk/api/forms"
)

type HttpFormApi interface {
	MakeHttpForm(f sdkforms.Form) (form sdkforms.IHttpForm, err error)
	GetForm(name string) (form sdkforms.IHttpForm, ok bool)
}
