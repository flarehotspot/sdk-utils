package plugins

import (
	"html/template"

	"github.com/flarehotspot/core/connmgr"
	"github.com/flarehotspot/core/db"
	"github.com/flarehotspot/core/sdk/utils/translate"
)

type PluginApi struct {
	slug             string
	dir              string
	trnslt           translate.TranslateFn
	vfmap            template.FuncMap
	db               *db.Database
	models           *PluginModels
	AcctAPI          *AccountsApi
	HttpAPI          *HttpApi
	NavAPI           *NavApi
	ConfigAPI        *ConfigApi
	PaymentsAPI      *PaymentsApi
	ThemesAPI        *ThemesApi
	NetworkAPI       *NetworkApi
	AdsAPI           *AdsApi
	InAppPurchaseAPI *InAppPurchaseApi
	PluginsMgr       *PluginsMgr
	ClntReg          *connmgr.ClientRegister
	ClntMgr          *connmgr.ClientMgr
	UciAPI           *UciApi
}

// func (api *PluginApi) Init() error {
// 	pluginLib := filepath.Join(api.dir, "plugin.so")
// 	p, err := plugin.Open(pluginLib)
// 	if err != nil {
// 		return err
// 	}
//
// 	initSym, err := p.Lookup("Init")
// 	if err != nil {
// 		return err
// 	}
//
// 	initFn := initSym.(func(sdk.IPluginApi))
// 	initFn(api)
//
// 	return nil
// }
