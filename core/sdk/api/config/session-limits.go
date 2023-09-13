package config

type ExpPauseDenom struct {
	// Matches sessions with minutes equal or above to this value.
	Minutes uint

	// Match sessions with megabytes equal or above to this value.
	Mbytes uint

	// ExpDays is the number of days the session is valid for.
	ExpDays uint

	// PauseLimit is the number of times the session can be paused.
	PauseLimit uint
}

// SessCfgData is the session settings configuration.
type SessCfgData struct {
	// Start session on boot.
	StartOnBoot bool

	// Pause the sessions when internet goes down.
	PauseInternetDown bool

	// Resume the sessions when internet comes back up.
	ResumeInterUp bool

	// Resume session when user connects to the network.
	ResumeWifiConnect bool

	// List of pause limit and expiration denominations. This is used to determine the number of pause
	// the session is allowed and its expiration time.
	PauseLimitDenoms []*ExpPauseDenom
}

// ISessionLimitsCfg is the configuration for session expiration and pause limit.
type ISessionLimitsCfg interface {
	// Reads the session limits configuration.
	Read() (*SessCfgData, error)

	// Writs the session limits configuration.
	Write(*SessCfgData) error

	// Computes the expiration days of a session based on the session's minutes and megabytes.
	ComputeExpDays(minutes uint, mbytes uint) (days uint)

  // Computes the pause limit of a session based on the session's minutes and megabytes.
	ComputePauseLimit(minutes uint, mbytes uint) (limit uint)
}
