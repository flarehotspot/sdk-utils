package models

import (
	"context"
	"database/sql"
	"log"

	"github.com/flarehotspot/core/internal/db"
)

type DeviceModel struct {
	db     *db.Database
	models *Models
}

func NewDeviceModel(database *db.Database, mdls *Models) *DeviceModel {
	return &DeviceModel{database, mdls}
}

func (self *DeviceModel) CreateTx(tx *sql.Tx, ctx context.Context, mac string, ip string, hostname string) (*Device, error) {
	query := "INSERT INTO devices (mac_address, ip_address, hostname) VALUES(?, ?, UPPER(?))"
	result, err := tx.ExecContext(ctx, query, mac, ip, hostname)
	if err != nil {
		log.Println("SQL Exec Error: ", err)
		return nil, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		log.Println("SQL Exec Error: ", err)
		return nil, err
	}

	return self.FindTx(tx, ctx, lastId)
}

func (self *DeviceModel) FindTx(tx *sql.Tx, ctx context.Context, id int64) (*Device, error) {
	device := NewDevice(self.db, self.models)
	err := tx.QueryRowContext(ctx, "SELECT id, mac_address, ip_address, hostname, created_at FROM devices WHERE id = ? LIMIT 1", id).
		Scan(&device.id, &device.macAddr, &device.ipAddr, &device.hostname, &device.createdAt)

	if err != nil {
		log.Println("Error finding device with id "+string(rune(id)), err.Error())
		return nil, err
	}

	log.Println("Found device: ", device)
	return device, nil
}

func (self *DeviceModel) FindByMacTx(tx *sql.Tx, ctx context.Context, mac string) (*Device, error) {
	device := NewDevice(self.db, self.models)
	query := "SELECT id, hostname, ip_address, mac_address, created_at FROM devices WHERE UPPER(mac_address) = UPPER(?) LIMIT 1"
	err := tx.QueryRowContext(ctx, query, mac).
		Scan(&device.id, &device.hostname, &device.ipAddr, &device.macAddr, &device.createdAt)

	if err != nil {
		log.Println("Error finding device with mac "+mac, err.Error())
		return nil, err
	}

	log.Println("Found device: ", device)
	return device, nil
}

func (self *DeviceModel) UpdateTx(tx *sql.Tx, ctx context.Context, id int64, mac string, ip string, hostname string) error {
	query := "UPDATE devices SET hostname = ?, ip_address = ?, mac_address = ? WHERE id = ? LIMIT 1"
	_, err := tx.ExecContext(ctx, query, hostname, ip, mac, id)
	if err != nil {
		log.Println("SQL Exec Error: ", err)
		return err
	}
	return nil
}

func (self *DeviceModel) Create(ctx context.Context, mac string, ip string, hostname string) (*Device, error) {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	dev, err := self.CreateTx(tx, ctx, mac, ip, hostname)
	if err != nil {
		return nil, err
	}

	return dev, tx.Commit()
}

func (self *DeviceModel) Find(ctx context.Context, id int64) (*Device, error) {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	device, err := self.FindTx(tx, ctx, id)
	if err != nil {
		return nil, err
	}

	return device, tx.Commit()
}

func (self *DeviceModel) FindByMac(ctx context.Context, mac string) (*Device, error) {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	dev, err := self.FindByMacTx(tx, ctx, mac)
	if err != nil {
		return nil, err
	}

	return dev, tx.Commit()
}

func (self *DeviceModel) Update(ctx context.Context, id int64, mac string, ip string, hostname string) error {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Commit()

	err = self.UpdateTx(tx, ctx, id, hostname, mac, ip)
	if err != nil {
		return err
	}

	return tx.Commit()
}
