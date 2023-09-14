package config

// IConfigApi is used to access the configuration API.
type IConfigApi interface {
	// Get the plugin configuration guration api.
	Plugin() IPluginCfg

	// Get the application configuration api.
	Application() IApplicationCfg

	// Get the database configuration api.
	Database() IDatabaseCfg

	// Get the http wifi rates configuration api.
	WifiRates() ISessionRatesCfg

	// Get the sessions configuration api.
	Sessions() ISessionLimitsCfg

	// Get the bandwidth configuration api.
	Bandwidth() IBandwdCfg
}
