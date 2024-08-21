package plugins

import (
	"core/internal/connmgr"
	"core/internal/db"
	"core/internal/db/models"
	"core/internal/network"
	sdkpaths "sdk/utils/paths"
	"sync/atomic"
)

type AppState struct {
	NeedsRestart atomic.Bool
}

type CoreGlobals struct {
	Db             *db.Database
	State          *AppState
	CoreAPI        *PluginApi
	ClientRegister *connmgr.ClientRegister
	ClientMgr      *connmgr.SessionsMgr
	TrafficMgr     *network.TrafficMgr
	BootProgress   *BootProgress
	Models         *models.Models
	PluginMgr      *PluginsMgr
	PaymentsMgr    *PaymentsMgr
}

func NewGlobals() *CoreGlobals {
	state := &AppState{}
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
	coreApi := NewPluginApi(sdkpaths.CoreDir, plgnMgr, trfcMgr)
	plgnMgr.InitCoreApi(coreApi)

	return &CoreGlobals{
		db,
		state,
		coreApi,
		clntReg,
		clntMgr,
		trfcMgr,
		bp,
		mdls,
		plgnMgr,
		pmtMgr,
	}
}
