package sdkcfg

// AppCfg is the application configuration.
type AppCfg struct {
  // Examples: en, zh
	Lang     string

  // Examples: USD, PH, CNY
	Currency string

  // Application secret key
	Secret   string
}

// IApplicationCfg is used to read and write application configuration.
type IApplicationCfg interface {
	Read() (AppCfg, error)
	Write(AppCfg) error
}
