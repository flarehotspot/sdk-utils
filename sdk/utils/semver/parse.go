package sdksemver

import (
	"log"
	"strconv"
	"strings"
)

// Parses the string versions into a Version struct
func ParseVersion(rawVersion string) (Version, error) {
	prVersion := strings.Split(rawVersion, ".")
	majorVersion, err := strconv.Atoi(prVersion[0])
	if err != nil {
		log.Println("Error parsing major version: ", err)
		return Version{}, err
	}
	minorVersion, err := strconv.Atoi(prVersion[1])
	if err != nil {
		log.Println("Error parsing minor version: ", err)
		return Version{}, err
	}
	patchVersion, err := strconv.Atoi(strings.Split(prVersion[2], "-")[0])
	if err != nil {
		log.Println("Error parsing patch version: ", err)
		return Version{}, err
	}

	return Version{
		Major: majorVersion,
		Minor: minorVersion,
		Patch: patchVersion,
	}, nil
}


