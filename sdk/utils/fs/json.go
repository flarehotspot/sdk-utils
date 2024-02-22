package sdkfs

import (
	"encoding/json"
	"os"
)

func WriteJson(f string, v any) error {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(f, b, PermFile)
}

func ReadJson(f string, v any) error {
	b, err := os.ReadFile(f)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}
