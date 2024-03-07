package connmgr

import (
	"context"
	"sync"

	"github.com/flarehotspot/core/internal/db"
	"github.com/flarehotspot/core/internal/db/models"
	"github.com/flarehotspot/core/internal/utils/sse"
)

type ClientDevice struct {
	mu       sync.RWMutex
	db       *db.Database
	mdls     *models.Models
	id       int64
	mac      string
	ip       string
	hostname string
}

func NewClientDevice(dtb *db.Database, mdls *models.Models, d *models.Device) *ClientDevice {
	return &ClientDevice{
		db:       dtb,
		mdls:     mdls,
		id:       d.Id(),
		mac:      d.MacAddress(),
		ip:       d.IpAddress(),
		hostname: d.Hostname(),
	}
}

func (self *ClientDevice) Id() int64 {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.id
}

func (self *ClientDevice) Hostname() string {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.hostname
}

func (self *ClientDevice) MacAddr() string {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.mac
}

func (self *ClientDevice) IpAddr() string {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.ip
}

func (self *ClientDevice) Update(ctx context.Context, mac string, ip string, hostname string) error {
	self.mu.Lock()
	defer self.mu.Unlock()

	err := self.mdls.Device().Update(ctx, self.id, self.mac, self.ip, self.hostname)
	if err != nil {
		return err
	}

	self.hostname = hostname
	self.mac = mac
	self.ip = ip

	return nil
}

func (self *ClientDevice) Emit(t string, data interface{}) {
	sse.Emit(self.MacAddr(), t, data)
}
