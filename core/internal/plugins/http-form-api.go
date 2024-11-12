package plugins

import (
	"errors"
	"net/http"
	"path/filepath"
	sdkhttp "sdk/api/http"

	sdkpaths "github.com/flarehotspot/go-utils/paths"
)

func NewHttpFormApi(api *PluginApi) *HttpFormApi {
	return &HttpFormApi{api}
}

type HttpFormApi struct {
	api *PluginApi
}

func (self *HttpFormApi) NewHttpForm(key string, sections []sdkhttp.Section) (sdkhttp.HttpForm, error) {
	if key == "" {
		return nil, errors.New("config key is required")
	}

	configDir := filepath.Join(sdkpaths.ConfigDir, "plugins", self.api.Pkg(), "config", key)

	return NewHttpForm(self.api, configDir, sections)
}

func (self *HttpFormApi) SaveForm(r *http.Request) (sdkhttp.HttpForm, error) {
	key := r.FormValue("config::key")
	configDir := filepath.Join(sdkpaths.ConfigDir, "plugins", self.api.Pkg(), "config", key)
	return LoadHttpForm(self.api, configDir)
}

func (self *HttpFormApi) GetForm(key string) (sdkhttp.HttpForm, error) {
	configDir := filepath.Join(sdkpaths.ConfigDir, "plugins", self.api.Pkg(), "config", key)
	return LoadHttpForm(self.api, configDir)
}
