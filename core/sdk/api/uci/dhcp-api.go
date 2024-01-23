package sdkuci

// DhcpCfg represents the DHCP configuration
type DhcpCfg struct {
	Section   string
	Ifname    string
	StartIp   string
	Limit     uint
	LeaseHour uint
}

type IDhcpApi interface {
	GetSection(ifname string) (section string, ok bool)
	GetConfig(section string) (dhcp *DhcpCfg, ok bool)
	SetConfig(ifname string, cfg *DhcpCfg) error
}
