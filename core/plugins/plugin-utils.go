package plugins

import (
	"net/http"

	"github.com/flarehotspot/core/utils/assets"
	"github.com/flarehotspot/core/web/response"
)

func NewPluginUtils(api *PluginApi) *PluginUtils {
	return &PluginUtils{api}
}

type PluginUtils struct {
	api *PluginApi
}

func (u *PluginUtils) BundleAssetsWithHelper(w http.ResponseWriter, r *http.Request, entries ...assets.AssetWithData) (assets.CacheData, error) {
	entriesWithHelpers := make([]assets.AssetWithData, len(entries))
	for i, entry := range entries {
		entriesWithHelpers[i] = assets.AssetWithData{
			File: entry.File,
			Data: &response.ViewData{
				ViewData:    entry.Data,
				ViewHelpers: u.api.HttpAPI.Helpers(),
			},
		}
	}

	return assets.BundleWithData(entriesWithHelpers...)
}
