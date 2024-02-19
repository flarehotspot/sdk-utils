package sdkstr

import (
	"regexp"
	"strings"
)

func Slugify(input string) string {
	// Convert to lowercase
	result := strings.ToLower(input)

	// Remove special characters
	re := regexp.MustCompile("[^a-z0-9]+")
	result = re.ReplaceAllString(result, "_")

	// Remove leading and trailing hyphens
	result = strings.Trim(result, "-")

	return result
}
