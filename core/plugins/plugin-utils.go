package plugins

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/flarehotspot/core/config"
	"github.com/flarehotspot/core/utils/assets"
	"github.com/flarehotspot/core/web/response"
)

func NewPluginUtils(api *PluginApi) *PluginUtils {
	return &PluginUtils{api}
}

type PluginUtils struct {
	api *PluginApi
}

func (utl *PluginUtils) Translate(msgtype string, msgk string, pairs ...string) string {
	trnsdir := utl.Resource("translations")
	appcfg, _ := config.ReadApplicationConfig()

	f := filepath.Join(trnsdir, appcfg.Lang, msgtype, msgk+".txt")
	bytes, err := os.ReadFile(f)
	if err != nil {
		return err.Error()
	}

	template := string(bytes)

	if len(pairs)%2 != 0 {
		return "Invalid number of pairs"
	}

	for i := 0; i < len(pairs); i += 2 {
		key := pairs[i]
		value := pairs[i+1]
		placeholder := fmt.Sprintf("${%s}", key)
		template = regexp.MustCompile(placeholder).ReplaceAllString(template, value)
	}

	return strings.TrimSpace(template)
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
