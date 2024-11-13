package plugins

import (
	"errors"
	sdkforms "sdk/api/forms"
)

func NewHttpFormApi(api *PluginApi) *HttpFormApi {
	return &HttpFormApi{api}
}

type HttpFormApi struct {
	api *PluginApi
}

func (self *HttpFormApi) MakeHttpForm(form sdkforms.Form) (sdkforms.IHttpForm, error) {
	if form.Name == "" {
		return nil, errors.New("config key is required")
	}

	return NewHttpForm(self.api, form), nil
}
