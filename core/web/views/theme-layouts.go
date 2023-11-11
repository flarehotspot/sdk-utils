package views

import (
	"path/filepath"

	"github.com/flarehotspot/core/config/themecfg"
	"github.com/flarehotspot/core/sdk/utils/paths"
)

const (
	LoginHtml        = "auth/login.html"
	BootingIndexHtml = "booting/index.html"
	PortalViewHtml   = "captive-portal/index.html"
)

const (
	portalLayoutHtml = "captive-portal/layout.html"
	adminLayoutHtml  = "web-admin/layout.html"
)

func PortalLayout() *ViewInput {
	themepkg := themecfg.Read().CaptivePortal
	extras := BundleExtras{
		ExtraJS: &[]string{
			filepath.Join(paths.CoreDir, "resources/assets/portal/js/event-source.polyfill.js"),
			filepath.Join(paths.CoreDir, "resources/assets/portal/js/events.js"),
		},
	}
	view := filepath.Join(paths.VendorDir, themepkg, "resources/views", portalLayoutHtml)
	return &ViewInput{File: view, Extras: &extras}
}

func WebAdminLayout() *ViewInput {
	extras := BundleExtras{
		ExtraJS: &[]string{filepath.Join(paths.CoreDir, "resources/assets/admin/js/events.js")},
	}
	themepkg := themecfg.Read().WebAdmin
	view := filepath.Join(paths.VendorDir, themepkg, "resources/views", adminLayoutHtml)
	return &ViewInput{File: view, Extras: &extras}
}
