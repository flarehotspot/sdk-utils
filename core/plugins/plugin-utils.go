package plugins

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/flarehotspot/flarehotspot/core/config"
	"github.com/flarehotspot/flarehotspot/core/utils/assets"
	"github.com/flarehotspot/flarehotspot/core/utils/flaretmpl"
	"github.com/flarehotspot/flarehotspot/core/web/response"
)

func NewPluginUtils(api *PluginApi) *PluginUtils {
	return &PluginUtils{api}
}

type PluginUtils struct {
	api *PluginApi
}

func (utl *PluginUtils) Translate(msgtype string, msgk string, pairs ...any) string {
	if len(pairs)%2 != 0 {
		return "Invalid number of translation params."
	}

	trnsdir := utl.Resource("translations")
	appcfg, _ := config.ReadApplicationConfig()

	f := filepath.Join(trnsdir, appcfg.Lang, msgtype, msgk+".txt")
	tmpl, err := flaretmpl.GetTextTemplate(f)
	if err != nil {
		log.Println("Warning: Translation file not found: ", f)
		return msgk
	}

	vdata := map[any]any{}
	for i := 0; i < len(pairs); i += 2 {
		key := pairs[i]
		value := pairs[i+1]
		vdata[key] = value
	}

	var output strings.Builder
	err = tmpl.Execute(&output, vdata)
	if err != nil {
		log.Println("Error executing translation template "+f, err)
		return msgk
	}
	return output.String()
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
