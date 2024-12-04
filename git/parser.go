package sdkgit

import (
	"errors"
	"net/url"
	"strings"
)

// GitSource represents the parsed components of a Git source URL.
type GitSource struct {
	Owner string
	Repo  string
	Ref   string
	Token string // Field for the access token
}

// ParseGitSource parses a Git source URL into a GitSource struct.
func ParseGitSource(sourceUrl string) (source GitSource, err error) {
	// Parse the URL
	u, err := url.Parse(sourceUrl)
	if err != nil {
		return source, errors.New("invalid URL format")
	}

	// Check the URL scheme
	if u.Scheme != "https" && u.Scheme != "http" && u.Scheme != "git" {
		return source, errors.New("unsupported URL scheme: " + sourceUrl)
	}

	// Extract the path components
	pathParts := strings.Split(strings.Trim(u.Path, "/"), "/")
	if len(pathParts) < 2 {
		return source, errors.New("URL must include at least owner and repository")
	}

	// Extract owner and repo
	source.Owner = pathParts[0]
	source.Repo = strings.TrimSuffix(pathParts[1], ".git") // Remove .git if present

	// Extract ref from query parameters if present
	queryRef := u.Query().Get("ref")
	if queryRef != "" {
		source.Ref = queryRef
	} else if len(pathParts) > 2 {
		// Assume ref if a third path component exists
		source.Ref = pathParts[2]
	}

	// Extract access token if present in the User part of the URL
	if u.User != nil {
		username := u.User.Username()
		if strings.HasPrefix(username, "oauth2:") {
			source.Token = strings.TrimPrefix(username, "oauth2:")
		}

		password, ok := u.User.Password()
		if source.Token == "" && ok {
			source.Token = password
		}
	}

	return source, nil
}
