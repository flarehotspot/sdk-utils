package sdkgit

import "strings"

func NeutralizeUrl(url string) string {
	if strings.HasSuffix(url, ".git") {
		return strings.Replace(url, ".git", "", 1)
	}
	return url
}
