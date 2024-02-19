package sdkcfg

// AppCfg is the application configuration.
type AppCfg struct {
	// Examples: en, zh
	Lang string `json:"lang"`

	// Examples: USD, PH, CNY
	Currency string `json:"currency"`

	// Application secret key
	Secret string `json:"secret"`
}

// ApplicationCfg is used to read and write application configuration.
type ApplicationCfg interface {
	Read() (AppCfg, error)
	Write(AppCfg) error
}
