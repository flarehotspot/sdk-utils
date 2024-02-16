package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	sdkpaths "github.com/flarehotspot/flarehotspot/core/sdk/utils/paths"
)

func readConfigFile(out any, f string) error {
	location := filepath.Join(sdkpaths.ConfigDir, f)
	bytes, err := os.ReadFile(location)
	if err != nil {
		// read from defaults
		location = filepath.Join(sdkpaths.ConfigDir, ".defaults", f)
		bytes, err = os.ReadFile(location)
		if err != nil {
			return err
		}
	}

	return json.Unmarshal(bytes, out)
}

func writeConfigFile(f string, config any) error {
	bytes, err := json.Marshal(config)
	if err != nil {
		return err
	}

	location := filepath.Join(sdkpaths.ConfigDir, f)
	return os.WriteFile(location, bytes, 0644)
}
