package sdkgit

import (
	"fmt"
	"io"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	sdkfs "github.com/flarehotspot/go-utils/fs"
	sdkpaths "github.com/flarehotspot/go-utils/paths"
	sdkstr "github.com/flarehotspot/go-utils/strings"
)

var (
	cacheDir = filepath.Join(sdkpaths.TmpDir, "git-cache")
)

type RepoSource struct {
	URL string
	Ref string // Can be branch, tag, commit, or empty
}

func MakeCachePath(repo RepoSource) string {
	return filepath.Join(sdkpaths.TmpDir, "git-cache", sdkstr.Slugify(repo.URL, "_"), repo.Ref)
}

func IsCached(repo RepoSource) bool {
	cachePath := MakeCachePath(repo)
	return repo.Ref != "" && sdkfs.Exists(cachePath)
}

func Cache(repo RepoSource, clonePath string) error {
	cachePath := MakeCachePath(repo)
	if err := sdkfs.EmptyDir(cachePath); err != nil {
		return err
	}
	// Copy the cloned repository to the cache directory
	if err := sdkfs.CopyDir(clonePath, cachePath, nil); err != nil {
		return err
	}
	log.Printf("Repository cached to %s", cachePath)
	return nil
}

func Clone(w io.Writer, repo RepoSource, clonePath string) error {
	// Ensure the parent directory of clonePath exists
	parentDir := filepath.Dir(clonePath)
	if err := sdkfs.EmptyDir(parentDir); err != nil {
		return err
	}

	if IsCached(repo) {
		cachePath := MakeCachePath(repo)
		if err := sdkfs.CopyDir(cachePath, clonePath, nil); err != nil {
			return err
		}
	} else {
		// Clone the repository using the "git clone" command with the provided URL
		var stderr strings.Builder

		cmd := exec.Command("git", "clone", repo.URL, clonePath)
		cmd.Stderr = &stderr

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("Error: %s\nStderr: %s", err.Error(), stderr.String())
		}

		log.Printf("Repository cloned to %s", clonePath)

		// If a specific ref (branch, tag, commit) is provided, checkout that ref
		if repo.Ref != "" {
			// Prepare the checkout command
			checkoutCmd := exec.Command("git", "checkout", repo.Ref)
			checkoutCmd.Stdout = w
			checkoutCmd.Stderr = w
			checkoutCmd.Dir = clonePath // Set the working directory for the command
			if err := checkoutCmd.Run(); err != nil {
				return err
			}

			if err := Cache(repo, clonePath); err != nil {
				return err
			}

			log.Printf("Checked out ref %s", repo.Ref)
		}
	}

	return nil
}
