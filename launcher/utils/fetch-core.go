package utils

import (
	"fmt"
	"launcher/config"
	"net/url"
	"path/filepath"
)

func FetchCore(os string, os_version string, os_arch string, go_version string, go_arch string) (outpath string, err error) {
	var data struct {
		ArchBinUrl string `json:"arch_bin_url"`
	}

	url := fmt.Sprintf(config.CoreFetchUrlTmpl, url.QueryEscape(os), url.QueryEscape(os_version), url.QueryEscape(os_arch), url.QueryEscape(go_version), url.QueryEscape(go_arch))
	if err = FetchJson(url, &data); err != nil {
		return "", err
	}

	downloadFile := filepath.Join(config.TempPath, "downloads/core_arch_bins", os, os_version, os_arch, go_version, go_arch+".zip")
	if err := DownloadFile(data.ArchBinUrl, downloadFile); err != nil {
		return "", err
	}

	extractTo := config.AppPath
	if err := Unzip(downloadFile, extractTo); err != nil {
		return "", err
	}

	return "", nil
}
