package plugins

import (
	"github.com/flarehotspot/flarehotspot/core/connmgr"
	"github.com/flarehotspot/flarehotspot/core/db"
	"github.com/flarehotspot/flarehotspot/core/db/models"
	"github.com/flarehotspot/flarehotspot/core/network"
	paths "github.com/flarehotspot/sdk/utils/paths"
)

type CoreGlobals struct {
	Db             *db.Database
	CoreAPI        *PluginApi
	ClientRegister *connmgr.ClientRegister
	ClientMgr      *connmgr.SessionsMgr
	TrafficMgr     *network.TrafficMgr
	BootProgress   *BootProgress
	Models         *models.Models
	PluginMgr      *PluginsMgr
	PaymentsMgr    *PaymentsMgr
}

func New() *CoreGlobals {
	db, _ := db.NewDatabase()
	bp := NewBootProgress()
	mdls := models.New(db)
	clntReg := connmgr.NewClientRegister(db, mdls)
	clntMgr := connmgr.NewSessionsMgr(db, mdls)
	trfcMgr := network.NewTrafficMgr()
	pmtMgr := NewPaymentMgr()

	trfcMgr.Start()
	clntMgr.ListenTraffic(trfcMgr)

	plgnMgr := NewPluginMgr(db, mdls, pmtMgr, clntReg, clntMgr, trfcMgr)
	coreApi := NewPluginApi(paths.CoreDir, plgnMgr, trfcMgr)
	plgnMgr.InitCoreApi(coreApi)

	return &CoreGlobals{db, coreApi, clntReg, clntMgr, trfcMgr, bp, mdls, plgnMgr, pmtMgr}
}
