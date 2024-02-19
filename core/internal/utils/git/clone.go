package git

import (
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type RepoSource struct {
	URL string
	Ref string // Can be branch, tag, commit, or empty
}

func Clone(w io.Writer, repo RepoSource, clonePath string) error {
	// Ensure the parent directory of clonePath exists
	parentDir := filepath.Dir(clonePath)
	if err := os.MkdirAll(parentDir, os.ModePerm); err != nil {
		return err
	}

	// Clone the repository using the "git clone" command with the provided URL
	cloneCmd := exec.Command("git", "clone", repo.URL, clonePath)
	cloneCmd.Stdout = w
	cloneCmd.Stderr = w
	if err := cloneCmd.Run(); err != nil {
		return err
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

		log.Printf("Checked out ref %s", repo.Ref)
	}

	return nil
}
