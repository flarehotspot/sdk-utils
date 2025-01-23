package plugins

import (
	"errors"
	sdkapi "sdk/api"
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

func (self *HttpFormApi) RegisterForms(forms ...sdkapi.HttpForm) error {
	for _, form := range forms {
		if form.Name == "" {
			return errors.New("form name key is required")
		}

		f := NewHttpForm(self.api, form)
		self.forms.Store(form.Name, f)
	}
	return nil
}

func (self *HttpFormApi) GetForm(name string) (form sdkapi.IHttpForm, ok bool) {
	f, ok := self.forms.Load(name)
	if !ok {
		return form, false
	}

	form, ok = f.(sdkapi.IHttpForm)
	if !ok {
		return form, false
	}

	return
}
