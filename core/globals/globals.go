package globals

import (
	"github.com/flarehotspot/core/connmgr"
	"github.com/flarehotspot/core/db"
	"github.com/flarehotspot/core/db/models"
	"github.com/flarehotspot/core/network"
	"github.com/flarehotspot/core/payments"
	"github.com/flarehotspot/core/plugins"
	"github.com/flarehotspot/core/sdk/utils/paths"
)

type CoreGlobals struct {
	Db             *db.Database
	CoreApi        *plugins.PluginApi
	ClientRegister *connmgr.ClientRegister
	ClientMgr      *connmgr.ClientMgr
	TrafficMgr     *network.TrafficMgr
	BootProgress   *BootProgress
	Models         *models.Models
	PluginMgr      *plugins.PluginsMgr
	PaymentsMgr    *payments.PaymentsMgr
}

func New() *CoreGlobals {
	db, _ := db.NewDatabase()
	bp := NewBootProgress()
	mdls := models.New(db)
	clntReg := connmgr.NewClientRegister(db, mdls)
	clntMgr := connmgr.NewClientMgr(db, mdls)
	trfcMgr := network.NewTrafficMgr()
	pmtMgr := payments.NewPaymentMgr()

	trfcMgr.Start()
	clntMgr.ListenTraffic(trfcMgr)

	plgnMgr := plugins.NewPluginMgr(db, mdls, pmtMgr, clntReg, clntMgr, trfcMgr)
	coreApi := plugins.NewPluginApi(paths.CoreDir, plgnMgr, trfcMgr)

	plgnMgr.RegisterPlugin(coreApi)

	return &CoreGlobals{db, coreApi, clntReg, clntMgr, trfcMgr, bp, mdls, plgnMgr, pmtMgr}
}
