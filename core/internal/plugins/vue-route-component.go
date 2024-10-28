package plugins

import (
	"path/filepath"

	"core/internal/utils/crypt"

	sdkstr "github.com/flarehotspot/go-utils/strings"
)

func NewVueRouteComponent(api *PluginApi, name string, path string, file string, permsReq []string, permsAny []string) *VueRouteComponent {

	compPath := filepath.Join(api.Utl.Resource("components/"), file)
	compHash, _ := crypt.SHA1Files(compPath)
	compHash = sdkstr.Sha1Hash(name, path, compPath, compHash)

	if name == "" {
		name = "empty-route-name-" + compHash
	}

	helpers := api.HttpAPI.Helpers()
	var httpCompPath string
	if file == "" {
		httpCompPath = helpers.VueComponentPath("Empty.vue")
	}
	httpCompPath = helpers.VueComponentPath(file)

	return &VueRouteComponent{
		api:                 api,
		hash:                compHash,
		Path:                path,
		Name:                name,
		File:                file,
		VueRouteName:        api.HttpAPI.vueRouter.MakeVueRouteName(name),
		VueRoutePath:        api.HttpAPI.vueRouter.MakeVueRoutePath(path),
		HttpComponentPath:   httpCompPath,
		PermissionsRequired: permsReq,
		PermissionsAnyOf:    permsAny,
	}
}

type VueRouteComponent struct {
	api                 *PluginApi
	hash                string
	Path                string
	Name                string
	File                string
	VueRoutePath        VueRoutePath `json:"vue_route_path"`
	VueRouteName        string       `json:"vue_route_name"`
	HttpComponentPath   string       `json:"http_component_path"`
	PermissionsRequired []string     `json:"permissions_required"`
	PermissionsAnyOf    []string     `json:"permissions_any_of"`
}
