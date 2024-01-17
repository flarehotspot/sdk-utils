package views

import "github.com/flarehotspot/core/sdk/api/http"

type ILayoutData interface {
	Helpers() http.IHelpers
	ContentPath() string
}
