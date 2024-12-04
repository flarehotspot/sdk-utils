package sdkgit

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadGitHubTarball(repoUrl string, outputFile string) error {
	repo, err := ParseGitSource(repoUrl)
	if err != nil {
		return err
	}

	// Construct the GitHub API URL
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/tarball/%s", repo.Owner, repo.Repo, repo.Ref)

	// Create an HTTP client and a request
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Add headers
	if repo.Token != "" {
		req.Header.Set("Accept", "application/vnd.github+json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", repo.Token))
		req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	}

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Check for a successful response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Create the output file
	out, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer out.Close()

	// Copy the response body to the output file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write to output file: %w", err)
	}

	return nil
}
