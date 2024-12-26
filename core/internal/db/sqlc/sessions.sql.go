// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: sessions.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createSession = `-- name: CreateSession :one
INSERT INTO sessions (
  device_id, session_type, time_secs, 
  data_mbytes, exp_days, down_mbits, 
  up_mbits, use_global
) 
VALUES 
  ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id
`

type CreateSessionParams struct {
	DeviceID    pgtype.UUID
	SessionType int16
	TimeSecs    pgtype.Int4
	DataMbytes  pgtype.Numeric
	ExpDays     pgtype.Int4
	DownMbits   int32
	UpMbits     int32
	UseGlobal   bool
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, createSession,
		arg.DeviceID,
		arg.SessionType,
		arg.TimeSecs,
		arg.DataMbytes,
		arg.ExpDays,
		arg.DownMbits,
		arg.UpMbits,
		arg.UseGlobal,
	)
	var id pgtype.UUID
	err := row.Scan(&id)
	return id, err
}

const findAvlSessionForDev = `-- name: FindAvlSessionForDev :one
SELECT 
  id, 
  device_id, 
  session_type, 
  time_secs, 
  data_mbytes, 
  consumption_secs, 
  consumption_mb, 
  started_at, 
  exp_days, 
  down_mbits, 
  up_mbits, 
  use_global, 
  created_at, 
    CASE 
        WHEN exp_days IS NOT NULL AND started_at IS NOT NULL
        THEN started_at + INTERVAL '1 day' * exp_days
        ELSE NULL
    END AS expires_at 
FROM 
  sessions 
WHERE 
  device_id = $1 
  AND (
    (
      session_type = 0 
      AND consumption_secs < time_secs
    ) 
    OR (
      session_type = 1 
      AND consumption_mb < data_mbytes
    ) 
    OR (
      session_type = 2 
      AND consumption_mb < data_mbytes 
      AND consumption_secs < time_secs
    )
  ) 
  AND (
    (
      exp_days IS NULL 
      OR started_at IS NULL
    ) 
    OR (
      exp_days IS NOT NULL 
      AND started_at IS NOT NULL 
      AND NOW() < started_at + INTERVAL '1 day' * exp_days
    )
  ) 
LIMIT 
  1
`

type FindAvlSessionForDevRow struct {
	ID              pgtype.UUID
	DeviceID        pgtype.UUID
	SessionType     int16
	TimeSecs        pgtype.Int4
	DataMbytes      pgtype.Numeric
	ConsumptionSecs pgtype.Int4
	ConsumptionMb   pgtype.Numeric
	StartedAt       pgtype.Timestamp
	ExpDays         pgtype.Int4
	DownMbits       int32
	UpMbits         int32
	UseGlobal       bool
	CreatedAt       pgtype.Timestamp
	ExpiresAt       interface{}
}

func (q *Queries) FindAvlSessionForDev(ctx context.Context, deviceID pgtype.UUID) (FindAvlSessionForDevRow, error) {
	row := q.db.QueryRow(ctx, findAvlSessionForDev, deviceID)
	var i FindAvlSessionForDevRow
	err := row.Scan(
		&i.ID,
		&i.DeviceID,
		&i.SessionType,
		&i.TimeSecs,
		&i.DataMbytes,
		&i.ConsumptionSecs,
		&i.ConsumptionMb,
		&i.StartedAt,
		&i.ExpDays,
		&i.DownMbits,
		&i.UpMbits,
		&i.UseGlobal,
		&i.CreatedAt,
		&i.ExpiresAt,
	)
	return i, err
}

const findSession = `-- name: FindSession :one
SELECT 
  id, 
  device_id, 
  session_type, 
  time_secs, 
  data_mbytes, 
  consumption_secs, 
  consumption_mb, 
  started_at, 
  exp_days, 
  down_mbits, 
  up_mbits, 
  use_global, 
  created_at, 
    CASE 
        WHEN exp_days IS NOT NULL AND started_at IS NOT NULL
        THEN started_at + INTERVAL '1 day' * exp_days
        ELSE NULL
    END AS expires_at 
FROM 
  sessions 
WHERE 
  id = $1 
LIMIT 
  1
`

type FindSessionRow struct {
	ID              pgtype.UUID
	DeviceID        pgtype.UUID
	SessionType     int16
	TimeSecs        pgtype.Int4
	DataMbytes      pgtype.Numeric
	ConsumptionSecs pgtype.Int4
	ConsumptionMb   pgtype.Numeric
	StartedAt       pgtype.Timestamp
	ExpDays         pgtype.Int4
	DownMbits       int32
	UpMbits         int32
	UseGlobal       bool
	CreatedAt       pgtype.Timestamp
	ExpiresAt       interface{}
}

func (q *Queries) FindSession(ctx context.Context, id pgtype.UUID) (FindSessionRow, error) {
	row := q.db.QueryRow(ctx, findSession, id)
	var i FindSessionRow
	err := row.Scan(
		&i.ID,
		&i.DeviceID,
		&i.SessionType,
		&i.TimeSecs,
		&i.DataMbytes,
		&i.ConsumptionSecs,
		&i.ConsumptionMb,
		&i.StartedAt,
		&i.ExpDays,
		&i.DownMbits,
		&i.UpMbits,
		&i.UseGlobal,
		&i.CreatedAt,
		&i.ExpiresAt,
	)
	return i, err
}

const findSessionsForDev = `-- name: FindSessionsForDev :many
SELECT 
  id, 
  device_id, 
  session_type, 
  time_secs, 
  data_mbytes, 
  consumption_secs, 
  consumption_mb, 
  started_at, 
  exp_days, 
  down_mbits, 
  up_mbits, 
  use_global, 
  created_at, 
    CASE 
        WHEN exp_days IS NOT NULL AND started_at IS NOT NULL
        THEN started_at + INTERVAL '1 day' * exp_days
        ELSE NULL
    END AS expires_at 
FROM 
  sessions 
WHERE 
  device_id = $1 
  AND (
    (
      session_type = 0 
      AND consumption_secs < time_secs
    ) 
    OR (
      session_type = 1 
      AND consumption_mb < data_mbytes
    ) 
    OR (
      session_type = 2 
      AND consumption_mb < data_mbytes 
      AND consumption_secs < time_secs
    )
  ) 
  AND (
    (
      exp_days IS NULL 
      OR started_at IS NULL
    ) 
    OR (
      exp_days IS NOT NULL 
      AND started_at IS NOT NULL 
      AND NOW() < started_at + INTERVAL '1 day' * exp_days
    )
  )
`

type FindSessionsForDevRow struct {
	ID              pgtype.UUID
	DeviceID        pgtype.UUID
	SessionType     int16
	TimeSecs        pgtype.Int4
	DataMbytes      pgtype.Numeric
	ConsumptionSecs pgtype.Int4
	ConsumptionMb   pgtype.Numeric
	StartedAt       pgtype.Timestamp
	ExpDays         pgtype.Int4
	DownMbits       int32
	UpMbits         int32
	UseGlobal       bool
	CreatedAt       pgtype.Timestamp
	ExpiresAt       interface{}
}

func (q *Queries) FindSessionsForDev(ctx context.Context, deviceID pgtype.UUID) ([]FindSessionsForDevRow, error) {
	rows, err := q.db.Query(ctx, findSessionsForDev, deviceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindSessionsForDevRow
	for rows.Next() {
		var i FindSessionsForDevRow
		if err := rows.Scan(
			&i.ID,
			&i.DeviceID,
			&i.SessionType,
			&i.TimeSecs,
			&i.DataMbytes,
			&i.ConsumptionSecs,
			&i.ConsumptionMb,
			&i.StartedAt,
			&i.ExpDays,
			&i.DownMbits,
			&i.UpMbits,
			&i.UseGlobal,
			&i.CreatedAt,
			&i.ExpiresAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAllBandwidth = `-- name: UpdateAllBandwidth :exec
UPDATE 
  sessions 
SET 
  down_mbits = $1, 
  up_mbits = $2, 
  use_global = $3 
WHERE 
  (
    (
      session_type = 0 
      AND consumption_secs < time_secs
    ) 
    OR (
      session_type = 1 
      AND consumption_mb < data_mbytes
    ) 
    OR (
      session_type = 2 
      AND consumption_mb < data_mbytes 
      AND consumption_secs < time_secs
    )
  ) 
  AND (
    (
      exp_days IS NULL 
      OR started_at IS NULL
    ) 
    OR (
      exp_days IS NOT NULL 
      AND started_at IS NOT NULL 
      AND NOW() < started_at + INTERVAL '1 day' * exp_days
    )
  )
`

type UpdateAllBandwidthParams struct {
	DownMbits int32
	UpMbits   int32
	UseGlobal bool
}

func (q *Queries) UpdateAllBandwidth(ctx context.Context, arg UpdateAllBandwidthParams) error {
	_, err := q.db.Exec(ctx, updateAllBandwidth, arg.DownMbits, arg.UpMbits, arg.UseGlobal)
	return err
}

const updateSession = `-- name: UpdateSession :exec
UPDATE 
  sessions 
SET 
  device_id = $1, 
  session_type = $2, 
  time_secs = $3, 
  data_mbytes = $4, 
  consumption_secs = $5, 
  consumption_mb = $6, 
  started_at = $7, 
  exp_days = $8, 
  down_mbits = $9, 
  up_mbits = $10, 
  use_global = $11 
WHERE 
  id = $12
`

type UpdateSessionParams struct {
	DeviceID        pgtype.UUID
	SessionType     int16
	TimeSecs        pgtype.Int4
	DataMbytes      pgtype.Numeric
	ConsumptionSecs pgtype.Int4
	ConsumptionMb   pgtype.Numeric
	StartedAt       pgtype.Timestamp
	ExpDays         pgtype.Int4
	DownMbits       int32
	UpMbits         int32
	UseGlobal       bool
	ID              pgtype.UUID
}

func (q *Queries) UpdateSession(ctx context.Context, arg UpdateSessionParams) error {
	_, err := q.db.Exec(ctx, updateSession,
		arg.DeviceID,
		arg.SessionType,
		arg.TimeSecs,
		arg.DataMbytes,
		arg.ConsumptionSecs,
		arg.ConsumptionMb,
		arg.StartedAt,
		arg.ExpDays,
		arg.DownMbits,
		arg.UpMbits,
		arg.UseGlobal,
		arg.ID,
	)
	return err
}
