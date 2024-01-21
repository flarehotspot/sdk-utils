package plugins

import (
	"database/sql"
	"log"
	"path/filepath"

	"github.com/flarehotspot/core/config/plugincfg"
	"github.com/flarehotspot/core/network"
	acct "github.com/flarehotspot/core/sdk/api/accounts"
	ads "github.com/flarehotspot/core/sdk/api/ads"
	config "github.com/flarehotspot/core/sdk/api/config"
	connmgr "github.com/flarehotspot/core/sdk/api/connmgr"
	http "github.com/flarehotspot/core/sdk/api/http"
	inappur "github.com/flarehotspot/core/sdk/api/inappur"
	models "github.com/flarehotspot/core/sdk/api/models"
	sdknet "github.com/flarehotspot/core/sdk/api/network"
	paymentsApi "github.com/flarehotspot/core/sdk/api/payments"
	plugin "github.com/flarehotspot/core/sdk/api/plugin"
	themes "github.com/flarehotspot/core/sdk/api/themes"
	uci "github.com/flarehotspot/core/sdk/api/uci"
	"github.com/flarehotspot/core/sdk/libs/slug"
	translate "github.com/flarehotspot/core/sdk/utils/translate"
	"github.com/flarehotspot/core/utils/migrate"
)

func (p *PluginApi) Migrate() error {
	migdir := filepath.Join(p.dir, "resources/migrations")
	err := migrate.MigrateUp(migdir, p.db.SqlDB())
	if err != nil {
		log.Println("Error in plugin migration "+p.Name(), ":", err.Error())
		return err
	}

	log.Println("Done migrating plugin:", p.Name())
	return nil
}

func (p *PluginApi) Name() string {
	return p.info.Name
}

func (p *PluginApi) Pkg() string {
	return p.info.Package
}

func (p *PluginApi) Version() string {
	return p.info.Version
}

func (p *PluginApi) Slug() string {
	return p.slug
}

func (p *PluginApi) Description() string {
	info, err := plugincfg.GetPluginInfo(p.dir)
	if err != nil {
		return ""
	}
	return info.Description
}

func (p *PluginApi) Dir() string {
	return p.dir
}

func (p *PluginApi) Translate(msgtype translate.MsgType, msgk string) string {
	return p.trnslt(msgtype, msgk)
}

func (p *PluginApi) Resource(path string) string {
	return filepath.Join(p.dir, "resources", path)
}

func (p *PluginApi) DbApi() *sql.DB {
	return p.db.SqlDB()
}

func (p *PluginApi) ModelsApi() models.IModelsApi {
	return p.models
}

func (p *PluginApi) AcctApi() acct.IAcctApi {
	return p.AcctAPI
}

func (p *PluginApi) HttpApi() http.IHttpApi {
	return p.HttpAPI
}

func (p *PluginApi) ConfigApi() config.IConfigApi {
	return p.ConfigAPI
}

func (p *PluginApi) PaymentsApi() paymentsApi.IPaymentsApi {
	return p.PaymentsAPI
}

func (p *PluginApi) AdsApi() ads.IAdsApi {
	return p.AdsAPI
}

func (p *PluginApi) PluginMgr() plugin.IPluginMgr {
	return p.PluginsMgr
}

func (p *PluginApi) InAppPurchaseApi() inappur.InAppPurchaseApi {
	return p.InAppPurchaseAPI
}

func (p *PluginApi) NetworkApi() sdknet.INetworkApi {
	return p.NetworkAPI
}

func (p *PluginApi) ClientReg() connmgr.IClientRegister {
	return p.ClntReg
}

func (p *PluginApi) ClientMgr() connmgr.IClientMgr {
	return p.ClntMgr
}

func (p *PluginApi) UciApi() uci.IUciApi {
	return p.UciAPI
}

func (p *PluginApi) ThemesApi() themes.IThemesApi {
	return p.ThemesAPI
}

func NewPluginApi(dir string, pmgr *PluginsMgr, trfkMgr *network.TrafficMgr) *PluginApi {
	info, err := plugincfg.GetPluginInfo(dir)
	if err != nil {
		log.Println("Error getting plugin info: ", err.Error())
	}

	pluginApi := &PluginApi{
		info:       info,
		slug:       slug.Make(info.Package),
		dir:        dir,
		db:         pmgr.db,
		PluginsMgr: pmgr,
		ClntReg:    pmgr.clntReg,
		ClntMgr:    pmgr.clntMgr,
	}

	translateFn := translate.NewTranslator(dir)
	mdls := NewPluginModels(pmgr.models)
	acctApi := NewAcctApi(pluginApi)
	httpApi := NewHttpApi(pluginApi, pmgr.db, pmgr.clntReg, pmgr.models, pmgr.clntReg, pmgr.paymgr)
	configApi := NewConfigApi(pluginApi)
	paymentsApi := NewPaymentsApi(pluginApi, pmgr.paymgr)
	themesApi := NewThemesApi(pluginApi)
	networkApi := NewNetworkApi(trfkMgr)
	adsApi := NewAdsApi(pluginApi)
	inappur := NewInAppPurchaseApi(pluginApi)
	uciApi := NewUciApi()
	utl := NewPluginUtils(pluginApi)

	pluginApi.trnslt = translateFn
	pluginApi.models = mdls
	pluginApi.AcctAPI = acctApi
	pluginApi.HttpAPI = httpApi
	pluginApi.ConfigAPI = configApi
	pluginApi.PaymentsAPI = paymentsApi
	pluginApi.ThemesAPI = themesApi
	pluginApi.NetworkAPI = networkApi
	pluginApi.AdsAPI = adsApi
	pluginApi.InAppPurchaseAPI = inappur
	pluginApi.UciAPI = uciApi
	pluginApi.Utl = utl

	log.Println("NewPluginApi: ", dir, " - ", info.Package, " - ", info.Name, " - ", info.Version, " - ", info.Description)

	return pluginApi
}
