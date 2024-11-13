package sdkhttp

import (
	sdkforms "sdk/api/forms"
)

type HttpFormApi interface {
	MakeHttpForm(form sdkforms.Form) (sdkforms.IHttpForm, error)
}
