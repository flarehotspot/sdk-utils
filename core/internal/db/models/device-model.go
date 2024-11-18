package models

import (
	"context"
	"log"

	"core/internal/db"
	"core/internal/db/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

type DeviceModel struct {
	db     *db.Database
	models *Models
}

func NewDeviceModel(database *db.Database, mdls *Models) *DeviceModel {
	return &DeviceModel{database, mdls}
}

func (self *DeviceModel) Create(ctx context.Context, mac string, ip string, hostname string) (*Device, error) {
	lastInsertId, err := self.db.Queries.CreateDevice(ctx, sqlc.CreateDeviceParams{})
	if err != nil {
		log.Println("error creating new device:", err)
		return nil, err
	}

	d, err := self.db.Queries.FindDevice(ctx, lastInsertId)
	if err != nil {
		log.Printf("error finding device %d: %w\n", lastInsertId, err)
		return nil, err
	}

	dev := &Device{
		db:        self.db,
		models:    self.models,
		id:        d.ID,
		macAddr:   d.MacAddress,
		ipAddr:    d.IpAddress,
		hostname:  d.Hostname.String,
		createdAt: d.CreatedAt.Time,
	}

	return dev, nil
}

func (self *DeviceModel) Find(ctx context.Context, id pgtype.UUID) (*Device, error) {
	device := NewDevice(self.db, self.models)

	d, err := self.db.Queries.FindDevice(ctx, id)
	if err != nil {
		log.Println("error finding device %v: %w", id, err)
		return nil, err
	}

	device.id = d.ID
	device.macAddr = d.MacAddress
	device.ipAddr = d.IpAddress
	device.hostname = d.Hostname.String
	device.createdAt = d.CreatedAt.Time

	log.Printf("Found device: %+v", device)
	return device, nil
}

func (self *DeviceModel) FindByMac(ctx context.Context, mac string) (*Device, error) {
	device := NewDevice(self.db, self.models)

	d, err := self.db.Queries.FindDeviceByMac(ctx, mac)
	if err != nil {
		log.Println("error finding device %s: %w", mac, err)
		return nil, err
	}

	device.id = d.ID
	device.macAddr = d.MacAddress
	device.ipAddr = d.IpAddress
	device.hostname = d.Hostname.String
	device.createdAt = d.CreatedAt.Time

	log.Printf("Found device: %+v", device)
	return device, nil
}

func (self *DeviceModel) Update(ctx context.Context, id pgtype.UUID, mac string, ip string, hostname string) error {
	err := self.db.Queries.UpdateDevice(ctx, sqlc.UpdateDeviceParams{
		Hostname:   pgtype.Text{String: hostname, Valid: hostname != ""},
		IpAddress:  ip,
		MacAddress: mac,
		ID:         id,
	})
	if err != nil {
		log.Println("error updating device %v: %w", id, err)
		return err
	}

	log.Printf("Successfully updated device with id %d", id)
	return nil
}
