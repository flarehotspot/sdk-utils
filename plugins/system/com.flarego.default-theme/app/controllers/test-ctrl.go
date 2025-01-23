package controllers

import (
	"net/http"
	sdkplugin "sdk/api"
)

func TestCtrl(api sdkplugin.IPluginApi) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("this is a test page"))
	}
}
