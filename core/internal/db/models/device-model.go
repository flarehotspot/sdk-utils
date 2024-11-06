package models

import (
	"context"
	"fmt"
	"log"

	"core/internal/db"

	"github.com/jackc/pgx/v5"
)

type DeviceModel struct {
	db     *db.Database
	models *Models
}

func NewDeviceModel(database *db.Database, mdls *Models) *DeviceModel {
	return &DeviceModel{database, mdls}
}

func (self *DeviceModel) CreateTx(tx pgx.Tx, ctx context.Context, mac string, ip string, hostname string) (*Device, error) {
	defer func() {
		if err := tx.Rollback(ctx); err != nil && err != pgx.ErrTxClosed {
			log.Printf("Rollback failed: %v", err)
		}
	}()

	query := "INSERT INTO devices (mac_address, ip_address, hostname) VALUES($1, $2, UPPER($3)) RETURNING id"
	var lastInsertId int
	err := tx.QueryRow(ctx, query, mac, ip, hostname).Scan(&lastInsertId)
	if err != nil {
		log.Printf("SQL Execution failed: %v", err)
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		log.Printf("SQL transaction commit failed: %v", err)
		return nil, err
	}

	return self.FindTx(tx, ctx, int64(lastInsertId))
}

func (self *DeviceModel) FindTx(tx pgx.Tx, ctx context.Context, id int64) (*Device, error) {
	device := NewDevice(self.db, self.models)

	err := tx.QueryRow(ctx, "SELECT id, mac_address, ip_address, hostname, created_at FROM devices WHERE id = $1 LIMIT 1", id).
		Scan(&device.id, &device.macAddr, &device.ipAddr, &device.hostname, &device.createdAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			log.Printf("No device found with id %d", id)
			return nil, nil
		}
		log.Printf("Error finding device with id %d: %v", id, err)
		return nil, err
	}

	log.Printf("Found device: %+v", device)
	return device, nil
}

func (self *DeviceModel) FindByMacTx(tx pgx.Tx, ctx context.Context, mac string) (*Device, error) {
	device := NewDevice(self.db, self.models)

	query := "SELECT id, hostname, ip_address, mac_address, created_at FROM devices WHERE UPPER(mac_address) = UPPER($1) LIMIT 1"
	err := tx.QueryRow(ctx, query, mac).
		Scan(&device.id, &device.hostname, &device.ipAddr, &device.macAddr, &device.createdAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			log.Printf("No device found with mac %s", mac)
			return nil, nil
		}
		log.Printf("Error finding device with id %s: %v", mac, err)
		return nil, err
	}

	log.Println("Found device: ", device)
	return device, nil
}

func (self *DeviceModel) UpdateTx(tx pgx.Tx, ctx context.Context, id int64, mac string, ip string, hostname string) error {
	query := "UPDATE devices SET hostname = $1, ip_address = $2, mac_address = $3 WHERE id = $4"

	cmdTag, err := tx.Exec(ctx, query, hostname, ip, mac, id)
	if err != nil {
		log.Printf("SQL Exec Error while updating device ID %d: %v", id, err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		log.Printf("No device found with id %d; update operation skipped", id)
		return fmt.Errorf("device with id %d not found", id)
	}

	log.Printf("Successfully updated device with id %d", id)
	return nil
}

func (self *DeviceModel) Create(ctx context.Context, mac string, ip string, hostname string) (*Device, error) {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	dev, err := self.CreateTx(tx, ctx, mac, ip, hostname)
	if err != nil {
		return nil, fmt.Errorf("failed to create device: %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("could not commit transaction: %w", err)
	}

	return dev, nil
}

func (self *DeviceModel) Find(ctx context.Context, id int64) (*Device, error) {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	device, err := self.FindTx(tx, ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find device: %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("could not commit transaction: %w", err)
	}

	return device, nil
}

func (self *DeviceModel) FindByMac(ctx context.Context, mac string) (*Device, error) {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	dev, err := self.FindByMacTx(tx, ctx, mac)
	if err != nil {
		return nil, err
	}

	return dev, tx.Commit(ctx)
}

func (self *DeviceModel) Update(ctx context.Context, id int64, mac string, ip string, hostname string) error {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return fmt.Errorf("could not begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	err = self.UpdateTx(tx, ctx, id, mac, ip, hostname)
	if err != nil {
		return fmt.Errorf("could not update device: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	return nil
}
