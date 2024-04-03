package connmgr

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/flarehotspot/core/internal/db"
	"github.com/flarehotspot/core/internal/db/models"
	jobque "github.com/flarehotspot/core/internal/utils/job-que"
	connmgr "github.com/flarehotspot/sdk/api/connmgr"
)

var regQue *jobque.JobQues = jobque.NewJobQues()

type ClientRegister struct {
	db           *db.Database
	mdls         *models.Models
	mgr          *SessionsMgr
	createdHooks []connmgr.ClientCreatedHookFn
	changedHooks []connmgr.ClientChangedHookFn
}

func NewClientRegister(dtb *db.Database, mdls *models.Models) *ClientRegister {
	return &ClientRegister{
		db:           dtb,
		mdls:         mdls,
		createdHooks: []connmgr.ClientCreatedHookFn{},
		changedHooks: []connmgr.ClientChangedHookFn{},
	}
}

func (reg *ClientRegister) ClientCreatedHook(fn connmgr.ClientCreatedHookFn) {
	reg.createdHooks = append(reg.createdHooks, fn)
}

func (reg *ClientRegister) ClientChangedHook(fn connmgr.ClientChangedHookFn) {
	reg.changedHooks = append(reg.changedHooks, fn)
}

func (reg *ClientRegister) Register(ctx context.Context, mac string, ip string, hostname string) (connmgr.ClientDevice, error) {
	var (
		clnt *ClientDevice
		dev  *models.Device
		err  error
	)

	dev, err = reg.mdls.Device().FindByMac(ctx, mac)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if errors.Is(err, sql.ErrNoRows) {
		dev, err = reg.mdls.Device().Create(ctx, mac, ip, hostname)
		if err != nil {
			return nil, err
		}
	}

	changed := ip != dev.IpAddress() || hostname != dev.Hostname()

	// Update device details if need be
	if changed {
		err := dev.Update(ctx, mac, ip, hostname)
		if err != nil {
			return nil, err
		}
	}

	clnt = NewClientDevice(reg.db, reg.mdls, dev)
	if changed && reg.mgr.IsConnected(clnt) {
		log.Println("TODO: Update connection with new ip and mac.")
		// TODO: Update connection with new mac and ip
	}

	return clnt, nil
}
