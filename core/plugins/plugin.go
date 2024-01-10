package plugins

import (
	"github.com/flarehotspot/core/connmgr"
	"github.com/flarehotspot/core/db"
	"github.com/flarehotspot/core/sdk/utils/translate"
)

type PluginApi struct {
	slug             string
	dir              string
	trnslt           translate.TranslateFn
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
