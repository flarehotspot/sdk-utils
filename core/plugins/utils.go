package plugins

// import (
	// "log"
	// "path/filepath"

	// "github.com/flarehotspot/core/config/appcfg"
	// "github.com/flarehotspot/core/config/plugincfg"
	// coreRouter "github.com/flarehotspot/core/web/router"
	// router "github.com/flarehotspot/core/sdk/api/http/router"
	// "github.com/flarehotspot/core/sdk/libs/slug"
	// "github.com/flarehotspot/core/sdk/utils/lang"
	// "github.com/gorilla/mux"
// )

// type PluginUtils struct {
	// dir  string
	// pmgr *PluginsMgr
// }

// func (self *PluginUtils) Info() (*plugincfg.PluginInfo, error) {
	// return plugincfg.GetPluginInfo(self.dir)
// }

// func (self *PluginUtils) Slug() string {
	// info, err := self.Info()
	// if err != nil {
		// return ""
	// }
	// return slug.Make(info.Name)
// }

// func (self *PluginUtils) DirName() string {
	// return self.dir
// }

// func (self *PluginUtils) TranslationsDir() string {
	// return filepath.Join(self.dir, "resources/translations")
// }

// func (self *PluginUtils) Resource(f string) (path string) {
	// return filepath.Join(self.dir, "resources", f)
// }

// func (self *PluginUtils) Translate(msgtype string, msgk string) string {
  // cfg, err := appcfg.ReadConfig()
  // if err != nil {
    // return err.Error()
  // }
  // lang := cfg.Lang
  // translate := langutil.NewTranslator(self.TranslationsDir(), lang)
  // return translate(msgtype, msgk)
// }

// func (self *PluginUtils) PluginMgr() *PluginsMgr {
	// return self.pmgr
// }

// func (self *PluginUtils) MuxRouteName(name router.PluginRouteName) router.MuxRouteName {
	// muxname := self.Slug() + "::" + string(name)
	// return router.MuxRouteName(muxname)
// }

// func (util *PluginUtils) UrlForMuxRoute(muxname router.MuxRouteName, pairs ...string) string {
	// var url string
	// coreRouter.RootRouter().Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		// if url == "" && route.GetName() == string(muxname) {
			// result, err := route.URL(pairs...)
			// if err != nil {
				// log.Println(err)
				// return nil
			// }
			// url = result.String()
		// }
		// return nil
	// })
	// return url
// }

// func (util *PluginUtils) UrlForRoute(name router.PluginRouteName, pairs ...string) string {
	// muxname := util.MuxRouteName(name)
	// return util.UrlForMuxRoute(muxname, pairs...)
// }

// func NewPluginUtils(dir string, pmgr *PluginsMgr) *PluginUtils {
	// return &PluginUtils{dir, pmgr}
// }
