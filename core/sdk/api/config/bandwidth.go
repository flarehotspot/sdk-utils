package config

// BandwdData is the bandwidth configuration for a given interface. Each interface bandwidth is configured individually.
type BandwdData struct {
	// UseGlobal is true if the global bandwidth should be used.
	UseGlobal bool

	// GlobalDownMbits is the global download bandwidth in Mbits.
	GlobalDownMbits int

	// GlobalUpMbits is the global upload bandwidth in Mbits.
	GlobalUpMbits int

	// UserDownMbits is the per user download bandwidth in Mbits.
	UserDownMbits int

	// UserUpMbits is the per user upload bandwidth in Mbits.
	UserUpMbits int
}

// IBandwdCfg is used to get and set bandwidth configuration.
type IBandwdCfg interface {
	// GetConfig returns the bandwidth configuration for a given interface.
	GetConfig(ifname string) (cfg *BandwdData, ok bool)

	// SetConfig sets the bandwidth configuration for a given interface.
  // It needs application restart for the changes to take effect.
	SetConfig(ifname string, cfg *BandwdData) error
}
