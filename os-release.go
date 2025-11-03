package sdkutils

import (
	"path/filepath"
)

type OsRelease struct {
	Os        string `json:"os"`
	OsVersion string `json:"os_version"`
	OsTarget  string `json:"os_target"`
	OsArch    string `json:"os_arch"`
	OsProfile string `json:"os_profile"`
	OsConfig  string `json:"os_config"`
	IsMono    bool   `json:"is_mono"`
}

func ReadOsRelease() (OsRelease, error) {
	var release OsRelease
	err := JsonRead(filepath.Join(PathAppDir, "os_release.json"), &release)
	return release, err
}

func WriteOsRelease(release OsRelease) error {
	return JsonWrite(filepath.Join(PathAppDir, "os_release.json"), &release)
}
