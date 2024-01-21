package config

const themesConfigJsonFile = "themes.json"

type ThemesConfig struct {
	Auth   string `json:"auth"`
	Portal string `json:"portal"`
	Admin  string `json:"admin"`
}

func ReadThemesConfig() (ThemesConfig, error) {
	var cfg ThemesConfig
	err := readConfigFile(&cfg, themesConfigJsonFile)
	return cfg, err
}

func WriteThemesConfig(cfg ThemesConfig) error {
	return writeConfigFile(themesConfigJsonFile, cfg)
}
