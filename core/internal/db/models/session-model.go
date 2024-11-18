package models

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"core/internal/db"
	"core/internal/db/sqlc"
	"core/internal/utils/pg"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type SessionModel struct {
	db     *db.Database
	models *Models
}

func NewSessionModel(dtb *db.Database, mdls *Models) *SessionModel {
	return &SessionModel{dtb, mdls}
}

func (self *SessionModel) AvlForDevTx(tx pgx.Tx, ctx context.Context, deviceId pgtype.UUID) (*Session, error) {
	sRow, err := self.db.Queries.FindAvlSessionForDev(ctx, deviceId)
	if err != nil {
		log.Printf("error finding available session for dev %v: %w", deviceId, err)
		return nil, err
	}
	s := NewSession(self.db, self.models)

	return s, nil
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

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("could not commit transaction: %w", err)
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

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	log.Println("Successfully updated all bandwidth of valid sessions")
	return nil
}

func (self *SessionModel) Create(ctx context.Context, devId pgtype.UUID, t uint8, timeSecs uint, dataMbytes float64, exp *uint, downMbit int, upMbit int, g bool) (*Session, error) {
	sId, err := self.db.Queries.CreateSession(ctx, sqlc.CreateSessionParams{
		DeviceID:    devId,
		SessionType: int16(t),
		TimeSecs:    pgtype.Int4{Int32: int32(timeSecs)},
		DataMbytes:  pg.Float64ToNumeric(dataMbytes),
		ExpDays:     pgtype.Int4{Int32: int32(*exp)},
		DownMbits:   int32(downMbit),
		UpMbits:     int32(upMbit),
		UseGlobal:   g,
	})
	if err != nil {
		log.Println("error creating session:", err)
		return nil, err
	}

	return self.Find(ctx, sId)
}

func (self *SessionModel) Find(ctx context.Context, id pgtype.UUID) (*Session, error) {
	sRow, err := self.db.Queries.FindSession(ctx, id)
	if err != nil {
		log.Printf("error finding session %v: %w", id, err)
		return nil, err
	}

	expDays := uint(sRow.ExpDays.Int32)

	session := NewSession(self.db, self.models)
	session.id = sRow.ID
	session.deviceId = sRow.DeviceID
	session.timeSecs = uint(sRow.TimeSecs.Int32)
	session.dataMb = pg.NumericToFloat64(sRow.DataMbytes)
	session.timeCons = uint(sRow.ConsumptionSecs.Int32)
	session.dataCons = pg.NumericToFloat64(sRow.ConsumptionMb)
	session.startedAt = &sRow.StartedAt.Time
	session.expDays = &expDays
	// TODO: fix proper expiry calculation
	// session.expiresAt = sRow.ExpiresAt

	session.downMbits = int(sRow.DownMbits)
	session.upMbits = int(sRow.UpMbits)
	session.useGlobal = sRow.UseGlobal
	session.createdAt = sRow.CreatedAt.Time

	return session, nil
}

func (self *SessionModel) Update(ctx context.Context, id pgtype.UUID, devId pgtype.UUID, t uint8, timeSecs uint, dataMbytes float64, timeCons uint, dataCons float64, started *time.Time, exp *uint, downMbit int, upMbit int, g bool) error {
	err := self.db.Queries.UpdateSession(ctx, sqlc.UpdateSessionParams{
		DeviceID:        devId,
		SessionType:     int16(t),
		TimeSecs:        pgtype.Int4{Int32: int32(timeSecs)},
		DataMbytes:      pg.Float64ToNumeric(dataMbytes),
		ConsumptionSecs: pgtype.Int4{Int32: int32(timeCons)},
		ConsumptionMb:   pg.Float64ToNumeric(dataCons),
		StartedAt:       pgtype.Timestamp{Time: *started},
		ExpDays:         pgtype.Int4{Int32: int32(*exp)},
		DownMbits:       int32(downMbit),
		UpMbits:         int32(upMbit),
		UseGlobal:       g,
		ID:              id,
	})
	if err != nil {
		log.Printf("error updating session %v: %w", id, err)
		return err
	}

	log.Printf("Successfully updated device with id %d", id)
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

	return nil
}
