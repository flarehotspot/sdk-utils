package plugins

import (
	"errors"
	sdkforms "sdk/api/forms"
	"sync"
)

func NewHttpFormApi(api *PluginApi) *HttpFormApi {
	return &HttpFormApi{
		api:   api,
		forms: sync.Map{},
	}
}

type HttpFormApi struct {
	api   *PluginApi
	forms sync.Map
}

func (self *HttpFormApi) MakeHttpForm(form sdkforms.Form) (sdkforms.IHttpForm, error) {
	if form.Name == "" {
		return nil, errors.New("config key is required")
	}

	f := NewHttpForm(self.api, form)
	self.forms.Store(form.Name, f)

	return f, nil
}

func (self *HttpFormApi) GetForm(name string) (form sdkforms.IHttpForm, ok bool) {
	f, ok := self.forms.Load(name)
	if !ok {
		return
	}

	form, ok = f.(sdkforms.IHttpForm)
	return
}
