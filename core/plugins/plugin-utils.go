package plugins

import (
	"net/http"
	"path/filepath"

	translate "github.com/flarehotspot/core/sdk/utils/translate"
	"github.com/flarehotspot/core/utils/assets"
	"github.com/flarehotspot/core/web/response"
)

func NewPluginUtils(api *PluginApi) *PluginUtils {
	return &PluginUtils{api}
}

type PluginUtils struct {
	api *PluginApi
}

func (utl *PluginUtils) Translate(msgtype translate.MsgType, msgk string) string {
	return utl.api.trnslt(msgtype, msgk)
}

func (utl *PluginUtils) Resource(path string) string {
	return filepath.Join(utl.api.dir, "resources", path)
}

func (utl *PluginUtils) BundleAssetsWithHelper(w http.ResponseWriter, r *http.Request, entries ...assets.AssetWithData) (assets.CacheData, error) {
	entriesWithHelpers := make([]assets.AssetWithData, len(entries))
	for i, entry := range entries {
		entriesWithHelpers[i] = assets.AssetWithData{
			File: entry.File,
			Data: &response.ViewData{
				ViewData:    entry.Data,
				ViewHelpers: utl.api.HttpAPI.Helpers(),
			},
		}
	}

	return assets.BundleWithData(entriesWithHelpers...)
}
