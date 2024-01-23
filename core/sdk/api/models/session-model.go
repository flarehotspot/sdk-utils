package sdkmodels

import (
	"context"
	"database/sql"
	"time"
)

// ISessionModel is the session model.
type ISessionModel interface {

	CreateTx(tx *sql.Tx, ctx context.Context, devId int64, t uint8, timeSecs uint, dataMbytes float64, exp *uint, downMbit int, upMbit int, g bool) (ISession, error)

	FindTx(tx *sql.Tx, ctx context.Context, id int64) (ISession, error)
	UpdateTx(tx *sql.Tx, ctx context.Context, id int64, devId int64, t uint8, timeSecs uint, dataMbytes float64, timeCons uint, dataCons float64, started *time.Time, exp *uint, downMbit int, upMbit int, g bool) error
	AvlForDevTx(tx *sql.Tx, ctx context.Context, devId int64) (ISession, error)
	SessionsForDevTx(tx *sql.Tx, ctx context.Context, devId int64) ([]ISession, error)
	DevHasSessionTx(tx *sql.Tx, ctx context.Context, devId int64) (ok bool, err error)
	UpdateAllBandwidthTx(tx *sql.Tx, ctx context.Context, downMbit int, upMbit int, g bool) error

	Create(ctx context.Context, devId int64, t uint8, timeSecs uint, dataMbytes float64, exp *uint, downMbit int, upMbit int, g bool) (ISession, error)
	Find(ctx context.Context, id int64) (ISession, error)
	Update(ctx context.Context, id int64, devId int64, t uint8, timeSecs uint, dataMbytes float64, timeCons uint, dataCons float64, started *time.Time, exp *uint, downMbit int, upMbit int, g bool) error
	AvlForDev(ctx context.Context, devId int64) (ISession, error)
	SessionsForDev(ctx context.Context, devId int64) ([]ISession, error)
	DevHasSession(ctx context.Context, devId int64) (ok bool, err error)
	UpdateAllBandwidth(ctx context.Context, downMbit int, upMbit int, g bool) error
}
