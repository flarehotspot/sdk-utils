package connmgr

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/flarehotspot/core/internal/db"
	"github.com/flarehotspot/core/internal/db/models"
	jobque "github.com/flarehotspot/core/internal/utils/job-que"
	connmgr "github.com/flarehotspot/sdk/api/connmgr"
)

var regQue *jobque.JobQues = jobque.NewJobQues()

func NewClientRegister(dtb *db.Database, mdls *models.Models) *ClientRegister {
	return &ClientRegister{
		db:           dtb,
		mdls:         mdls,
		finderHooks:  []connmgr.ClientFinderFn{},
		createdHooks: []connmgr.ClientCreatedHookFn{},
		changedHooks: []connmgr.ClientChangedHookFn{},
	}
}

type ClientRegister struct {
	db           *db.Database
	mdls         *models.Models
	mgr          *SessionsMgr
	finderHooks  []connmgr.ClientFinderFn
	createdHooks []connmgr.ClientCreatedHookFn
	changedHooks []connmgr.ClientChangedHookFn
}

func (reg *ClientRegister) ClientFinderHook(fn ...connmgr.ClientFinderFn) {
	reg.finderHooks = append(reg.finderHooks, fn...)
}

func (reg *ClientRegister) ClientCreatedHook(fn ...connmgr.ClientCreatedHookFn) {
	reg.createdHooks = append(reg.createdHooks, fn...)
}

func (reg *ClientRegister) ClientChangedHook(fn ...connmgr.ClientChangedHookFn) {
	reg.changedHooks = append(reg.changedHooks, fn...)
}

func (reg *ClientRegister) Register(r *http.Request, mac string, ip string, hostname string) (connmgr.ClientDevice, error) {

	if len(reg.finderHooks) > 0 {
		for _, hookFn := range reg.finderHooks {
			if clnt, ok := hookFn(r, mac, ip, hostname); ok && clnt != nil {
				return clnt, nil
			}
		}
	}

	ctx := r.Context()

    dev, err := reg.mdls.Device().FindByMac(ctx, mac)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if errors.Is(err, sql.ErrNoRows) {
		dev, err = reg.mdls.Device().Create(ctx, mac, ip, hostname)
		if err != nil {
			return nil, err
		}

        clnt := NewClientDevice(reg.db, reg.mdls, dev)

		// call createdHooks functions
		if len(reg.createdHooks) > 0 {
			for _, hookFn := range reg.createdHooks {
				if err := hookFn(r, clnt); err != nil {
					return nil, err
				}
			}
		}

		return clnt, nil
	}

    clnt := NewClientDevice(reg.db, reg.mdls, dev)
	changed := ip != dev.IpAddress() || hostname != dev.Hostname()

	// Update device details if need be
	if changed {
		connected := reg.mgr.IsConnected(clnt)
		if connected {
			err = reg.mgr.Disconnect(ctx, clnt, "Device details changed, reconnecting...")
			if err != nil {
				return nil, err
			}
		}

		old := NewClientDevice(reg.db, reg.mdls, dev.Clone())
		err := dev.Update(ctx, mac, ip, hostname)
		if err != nil {
			return nil, err
		}

		// call changedHooks functions
		if len(reg.changedHooks) > 0 {
			for _, hookFn := range reg.changedHooks {
				if err := hookFn(r, clnt, old); err != nil {
					return nil, err
				}
			}
		}

		// reconnect client device
		if connected {
			err := reg.mgr.Connect(ctx, clnt, "Device details changed, reconnected successfully")
			if err != nil {
				return nil, err
			}
		}
	}

	return clnt, nil
}
