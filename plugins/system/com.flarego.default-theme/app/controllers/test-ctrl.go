package controllers

import (
	"net/http"
	sdkplugin "sdk/api/plugin"
)

func TestCtrl(api sdkplugin.PluginApi) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("this is a test page"))
	}
}
