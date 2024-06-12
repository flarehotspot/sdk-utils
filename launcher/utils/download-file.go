package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// DownloadFile downloads a file from a URL and saves it to a specified path
func DownloadFile(url string, outpath string) error {
	// Create the file
	if err := EmptyDir(filepath.Dir(outpath)); err != nil {
		return err
	}

	out, err := os.Create(outpath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Custom HTTP client to follow redirects
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// Limit the number of redirects
			if len(via) >= 10 {
				return fmt.Errorf("stopped after 10 redirects")
			}
			return nil
		},
	}

	// Get the data
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
