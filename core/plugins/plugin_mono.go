//go:build mono

package plugins

import (
	"html/template"
	"log"

	defaultThemes "github.com/AdoPiSoft/com.adopisoft.basic-flare-theme"
	netmgr "github.com/flarehotspot/com.flarego.basic-net-mgr"
	acct "github.com/flarehotspot/com.flarego.basic-system-account"
	wifihotspot "github.com/flarehotspot/com.flarego.basic-wifi-hotspot"
	coinslot "github.com/flarehotspot/com.flarego.wired-coinslot"
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

func (p *PluginApi) Init() error {
	log.Println("Plugin Name: ", p.Name())
	switch p.Pkg() {
	case "com.flarego.wired-coinslot":
		coinslot.Init(p)
		log.Println("Successfully loaded plugin: ", p.Name())
	case "com.adopisoft.basic-flare-theme":
		defaultThemes.Init(p)
	case "com.flarego.basic-wifi-hotspot":
		wifihotspot.Init(p)
	case "com.flarego.basic-system-account":
		acct.Init(p)
	case "com.flarego.basic-net-mgr":
		netmgr.Init(p)
	default:
		panic("Unable to load plugin: " + p.dir)
	}
	return nil
}
