package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/flarehotspot/core/db"
)

var (
	// select all session fields sql query
	selectQuery = `SELECT id, device_id, session_type, time_secs, data_mbytes, consumption_secs, consumption_mb, started_at, exp_days, down_mbits, up_mbits, use_global, created_at, IF(exp_days IS NOT NULL AND STARTED_AT IS NOT NULL, DATE_ADD(started_at, INTERVAL exp_days DAY), NULL) AS expires_at FROM sessions`

	// valid sessions sql where query
	validWhereQuery = `(
    (session_type = 0 AND consumption_secs < time_secs) OR
    (session_type = 1 AND consumption_mb < data_mbytes) OR
    (session_type = 2 AND consumption_mb < data_mbytes AND consumption_secs < time_secs)
  )
  AND (
    (exp_days IS NULL OR started_at IS NULL) OR
    (exp_days IS NOT NULL AND started_at IS NOT NULL AND NOW() < DATE_ADD(started_at, INTERVAL exp_days DAY))
  )`
)

type SessionModel struct {
	db     *db.Database
	models *Models
}

func NewSessionModel(dtb *db.Database, mdls *Models) *SessionModel {
	return &SessionModel{dtb, mdls}
}

func (self *SessionModel) CreateTx(tx *sql.Tx, ctx context.Context, devId int64, t uint8, timeSecs uint, dataMbytes float64, expDays *uint, downMbits int, upMbits int, useGlobal bool) (*Session, error) {
	query := "INSERT INTO sessions (device_id, session_type, time_secs, data_mbytes, exp_days, down_mbits, up_mbits, use_global) VALUES(?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := tx.ExecContext(ctx, query, devId, t, timeSecs, dataMbytes, expDays, downMbits, upMbits, useGlobal)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		log.Println("SQL Lastid() Error: ", err)
		return nil, err
	}

	return self.FindTx(tx, ctx, lastId)
}

func (self *SessionModel) FindTx(tx *sql.Tx, ctx context.Context, id int64) (*Session, error) {
	s := NewSession(self.db, self.models)
	query := selectQuery + " WHERE id = ? LIMIT 1"
	err := tx.QueryRowContext(ctx, query, id).
		Scan(&s.id, &s.deviceId, &s.sessionType, &s.timeSecs, &s.dataMb, &s.timeCons, &s.dataCons, &s.startedAt, &s.expDays, &s.downMbits, &s.upMbits, &s.useGlobal, &s.createdAt, &s.expiresAt)

	return s, err
}

func (self *SessionModel) UpdateTx(tx *sql.Tx, ctx context.Context, id int64, devId int64, t uint8, timeSecs uint, dataMbytes float64, timeCons uint, dataCons float64, started *time.Time, exp *uint, downMbit int, upMbit int, g bool) error {
	query := `UPDATE sessions
  SET device_id = ?, session_type = ?, time_secs = ?, data_mbytes = ?, consumption_secs = ?, consumption_mb = ?, started_at = ?, exp_days = ?, down_mbits = ?, up_mbits = ?, use_global = ?
  WHERE id = ? LIMIT 1`

	_, err := tx.ExecContext(ctx, query, devId, t, timeSecs, dataMbytes, timeCons, dataCons, started, exp, downMbit, upMbit, g, id)
	return err
}

func (self *SessionModel) AvlForDevTx(tx *sql.Tx, ctx context.Context, deviceId int64) (*Session, error) {
	s := NewSession(self.db, self.models)
	query := selectQuery + " WHERE device_id = ? AND " + validWhereQuery + " LIMIT 1"

	err := tx.QueryRowContext(ctx, query, deviceId).
		Scan(&s.id, &s.deviceId, &s.sessionType, &s.timeSecs, &s.dataMb, &s.timeCons, &s.dataCons, &s.startedAt, &s.expDays, &s.downMbits, &s.upMbits, &s.useGlobal, &s.createdAt, &s.expiresAt)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	return s, err
}

func (self *SessionModel) SessionsForDevTx(tx *sql.Tx, ctx context.Context, devId int64) ([]*Session, error) {
	query := selectQuery + " WHERE device_id = ? AND " + validWhereQuery
	rows, err := tx.QueryContext(ctx, query, devId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sessions := []*Session{}

	for rows.Next() {
		s := NewSession(self.db, self.models)
		if err := rows.Scan(&s.id, &s.deviceId, &s.sessionType, &s.timeSecs, &s.dataMb, &s.timeCons, &s.dataCons, &s.startedAt, &s.expDays, &s.downMbits, &s.upMbits, &s.useGlobal, &s.createdAt, &s.expiresAt); err != nil {
			return sessions, err
		}
		sessions = append(sessions, s)
	}

	return sessions, rows.Err()
}

func (self *SessionModel) DevHasSessionTx(tx *sql.Tx, ctx context.Context, devId int64) (bool, error) {
	query := fmt.Sprintf("SELECT EXISTS(\n%s\n)", selectQuery+" WHERE device_id = ? AND "+validWhereQuery)

	var exists bool
	row := tx.QueryRowContext(ctx, query, devId)
	if err := row.Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

func (self *SessionModel) UpdateAllBandwidthTx(tx *sql.Tx, ctx context.Context, downMbits int, upMbits int, useGlobal bool) error {
	query := `UPDATE sessions SET down_mbits = ?, up_mbits = ?, use_global = ? WHERE ` + validWhereQuery
	_, err := tx.ExecContext(ctx, query, downMbits, upMbits, useGlobal)
	return err
}

func (self *SessionModel) Create(ctx context.Context, devId int64, t uint8, timeSecs uint, dataMbytes float64, exp *uint, downMbit int, upMbit int, g bool) (*Session, error) {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	s, err := self.CreateTx(tx, ctx, devId, t, timeSecs, dataMbytes, exp, downMbit, upMbit, g)
	if err != nil {
		return nil, err
	}

	return s, tx.Commit()
}

func (self *SessionModel) Find(ctx context.Context, id int64) (*Session, error) {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	s, err := self.FindTx(tx, ctx, id)
	if err != nil {
		return nil, err
	}

	return s, tx.Commit()
}

func (self *SessionModel) Update(ctx context.Context, id int64, devId int64, t uint8, timeSecs uint, dataMbytes float64, timeCons uint, dataCons float64, started *time.Time, exp *uint, downMbit int, upMbit int, g bool) error {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = self.UpdateTx(tx, ctx, id, devId, t, timeSecs, dataMbytes, timeCons, dataCons, started, exp, downMbit, upMbit, g)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (self *SessionModel) AvlForDev(ctx context.Context, devId int64) (*Session, error) {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	s, err := self.AvlForDevTx(tx, ctx, devId)
	if err != nil {
		return nil, err
	}

	return s, tx.Commit()
}

func (self *SessionModel) SessionsForDev(ctx context.Context, devId int64) ([]*Session, error) {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	sessions, err := self.SessionsForDevTx(tx, ctx, devId)
	if err != nil {
		return nil, err
	}

	return sessions, tx.Commit()
}

func (self *SessionModel) DevHasSession(ctx context.Context, devId int64) (bool, error) {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return false, err
	}
	defer tx.Rollback()

	ok, err := self.DevHasSessionTx(tx, ctx, devId)
	if err != nil {
		return false, err
	}

	return ok, tx.Commit()
}

func (self *SessionModel) UpdateAllBandwidth(ctx context.Context, downMbit int, upMbit int, g bool) error {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = self.UpdateAllBandwidthTx(tx, ctx, downMbit, upMbit, g)
	if err != nil {
		return err
	}

	return tx.Commit()
}
