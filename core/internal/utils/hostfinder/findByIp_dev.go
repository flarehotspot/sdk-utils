//go:build dev

package hostfinder

func FindByIp(ip string) (*HostData, error) {
	return &HostData{
		Hostname: "localhost",
		IpAddr:   "10.0.0.2",
		MacAddr:  "00:00:00:00:00:00",
	}, nil
}
