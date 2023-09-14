package plugins

import (
	"net/http"
)

type ThemesApi struct {
	handler func(w http.ResponseWriter, r *http.Request)
}

func NewThemesApi() *ThemesApi {
	return &ThemesApi{}
}

func (t *ThemesApi) PortalIndexHandler(handler func(w http.ResponseWriter, r *http.Request)) {
	t.handler = handler
}

func (t *ThemesApi) GetPortalHandler() func(w http.ResponseWriter, r *http.Request) {
	return t.handler
}
