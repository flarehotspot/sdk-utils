package models

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"core/internal/db"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

const (
	// select all session fields sql query
	SELECT_QUERY = `SELECT id, device_id, session_type, time_secs, data_mbytes, consumption_secs, consumption_mb, started_at, exp_days, down_mbits, up_mbits, use_global, created_at, IF(exp_days IS NOT NULL AND STARTED_AT IS NOT NULL, DATE_ADD(started_at, INTERVAL exp_days DAY), NULL) AS expires_at FROM sessions`

	// valid sessions sql where query
	VALID_WHERE_QUERY = `(
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

func (self *SessionModel) CreateTx(tx pgx.Tx, ctx context.Context, devId uuid.UUID, t uint8, timeSecs uint, dataMbytes float64, expDays *uint, downMbits int, upMbits int, useGlobal bool) (*Session, error) {
	defer func() {
		if err := tx.Rollback(ctx); err != nil && err != pgx.ErrTxClosed {
			log.Printf("Rollback failed: %v", err)
		}
	}()

	query := "INSERT INTO sessions (device_id, session_type, time_secs, data_mbytes, exp_days, down_mbits, up_mbits, use_global) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"

	var lastInsertId uuid.UUID
	err := tx.QueryRow(ctx, query, devId, t, timeSecs, dataMbytes, expDays, downMbits, upMbits, useGlobal).Scan(&lastInsertId)
	if err != nil {
		log.Printf("SQL Execution Error failed: %v", err)
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		log.Printf("SQL transaction commit failed: %v", err)
		return nil, err
	}

	return self.FindTx(tx, ctx, lastInsertId)
}

func (self *SessionModel) FindTx(tx pgx.Tx, ctx context.Context, id uuid.UUID) (*Session, error) {
	s := NewSession(self.db, self.models)
	query := SELECT_QUERY + " WHERE id = $1 LIMIT 1"

	err := tx.QueryRow(ctx, query, id).
		Scan(&s.id, &s.deviceId, &s.sessionType, &s.timeSecs, &s.dataMb, &s.timeCons, &s.dataCons, &s.startedAt, &s.expDays, &s.downMbits, &s.upMbits, &s.useGlobal, &s.createdAt, &s.expiresAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			log.Printf("No session found with id %d", id)
			return nil, nil
		}
		log.Printf("Error finding payment with id %d: %v", id, err)
		return nil, err
	}

	return s, nil
}

func (self *SessionModel) UpdateTx(tx pgx.Tx, ctx context.Context, id uuid.UUID, devId uuid.UUID, t uint8, secs uint, mb float64, timeCons uint, dataCons float64, started *time.Time, exp *uint, downMbit int, upMbit int, g bool) error {
	query := `UPDATE sessions
  SET device_id = $1, session_type = $2, time_secs = $3, data_mbytes = $4, consumption_secs = $5, consumption_mb = $6, started_at = $7, exp_days = $8, down_mbits = $9, up_mbits = $10, use_global = $11
  WHERE id = $12 LIMIT 1`

	cmdTag, err := tx.Exec(ctx, query, devId, t, secs, mb, timeCons, dataCons, started, exp, downMbit, upMbit, g, id)
	if err != nil {
		log.Printf("SQL Exec Error while updating session ID %d: %v", id, err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		log.Printf("No session found with id %d; update operation skipped", id)
		return fmt.Errorf("session with id %d not found", id)
	}

	log.Printf("Successfully updated device with id %d", id)
	return nil
}

func (self *SessionModel) AvlForDevTx(tx pgx.Tx, ctx context.Context, deviceId uuid.UUID) (*Session, error) {
	s := NewSession(self.db, self.models)
	query := SELECT_QUERY + " WHERE device_id = $1 AND " + VALID_WHERE_QUERY + " LIMIT 1"

	err := tx.QueryRow(ctx, query, deviceId).
		Scan(&s.id, &s.deviceId, &s.sessionType, &s.timeSecs, &s.dataMb, &s.timeCons, &s.dataCons, &s.startedAt, &s.expDays, &s.downMbits, &s.upMbits, &s.useGlobal, &s.createdAt, &s.expiresAt)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.New("No more available sessions for this device")
	}

	return s, err
}

func (self *SessionModel) SessionsForDevTx(tx pgx.Tx, ctx context.Context, devId uuid.UUID) ([]*Session, error) {
	sessions := []*Session{}
	query := SELECT_QUERY + " WHERE device_id = $1 AND " + VALID_WHERE_QUERY

	rows, err := tx.Query(ctx, query, devId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		s := NewSession(self.db, self.models)
		if err := rows.Scan(&s.id, &s.deviceId, &s.sessionType, &s.timeSecs, &s.dataMb, &s.timeCons, &s.dataCons, &s.startedAt, &s.expDays, &s.downMbits, &s.upMbits, &s.useGlobal, &s.createdAt, &s.expiresAt); err != nil {
			return sessions, err
		}
		sessions = append(sessions, s)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("row iteration error: %w", rows.Err())
	}

	return sessions, nil
}

func (self *SessionModel) UpdateAllBandwidthTx(tx pgx.Tx, ctx context.Context, downMbits int, upMbits int, useGlobal bool) error {
	query := `UPDATE sessions SET down_mbits = $1, up_mbits = $2, use_global = $3 WHERE ` + VALID_WHERE_QUERY

	cmdTag, err := tx.Exec(ctx, query, downMbits, upMbits, useGlobal)
	if err != nil {
		log.Printf("SQL Exec Error while updating session bandwidth: %v", err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("session not found: %w", err)
	}

	log.Println("Successfully updated all bandwidth of valid sessions")
	return nil
}

func (self *SessionModel) Create(ctx context.Context, devId uuid.UUID, t uint8, timeSecs uint, dataMbytes float64, exp *uint, downMbit int, upMbit int, g bool) (*Session, error) {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	s, err := self.CreateTx(tx, ctx, devId, t, timeSecs, dataMbytes, exp, downMbit, upMbit, g)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("could not commit transaction: %w", err)
	}

	return s, nil
}

func (self *SessionModel) Find(ctx context.Context, id uuid.UUID) (*Session, error) {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	s, err := self.FindTx(tx, ctx, id)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("could not commit transaction: %w", err)
	}

	return s, nil
}

func (self *SessionModel) Update(ctx context.Context, id uuid.UUID, devId uuid.UUID, t uint8, timeSecs uint, dataMbytes float64, timeCons uint, dataCons float64, started *time.Time, exp *uint, downMbit int, upMbit int, g bool) error {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return fmt.Errorf("could not begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	err = self.UpdateTx(tx, ctx, id, devId, t, timeSecs, dataMbytes, timeCons, dataCons, started, exp, downMbit, upMbit, g)
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	return nil
}

func (self *SessionModel) AvlForDev(ctx context.Context, devId uuid.UUID) (*Session, error) {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	s, err := self.AvlForDevTx(tx, ctx, devId)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("could not commit transaction: %w", err)
	}

	return s, nil
}

func (self *SessionModel) SessionsForDev(ctx context.Context, devId uuid.UUID) ([]*Session, error) {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	sessions, err := self.SessionsForDevTx(tx, ctx, devId)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("could not commit transaction: %w", err)
	}

	return sessions, nil
}

func (self *SessionModel) UpdateAllBandwidth(ctx context.Context, downMbit int, upMbit int, g bool) error {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return fmt.Errorf("could not begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	err = self.UpdateAllBandwidthTx(tx, ctx, downMbit, upMbit, g)
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	return nil
}
