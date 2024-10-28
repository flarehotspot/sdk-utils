package migrate

import (
	"sort"
	"strings"

	fs "github.com/flarehotspot/go-utils/fs"
	slices "github.com/flarehotspot/go-utils/slices"
)

type MigDirection int

const (
	migration_Down MigDirection = iota
	migration_Up
)

func listFiles(dir string, d MigDirection) (files []string, err error) {
	list := []string{}
	if err = fs.LsFiles(dir, &list, false); err != nil {
		return files, err
	}

	files = []string{}
	if d == migration_Down {
		for _, f := range list {
			if strings.HasSuffix(f, ".down.sql") && !strings.HasPrefix(f, ".") {
				files = append(files, f)
			}
		}
		slices.ReverseString(files)
	} else {
		for _, f := range list {
			if strings.HasSuffix(f, ".up.sql") && !strings.HasPrefix(f, ".") {
				files = append(files, f)
			}
		}
		sort.Strings(files)
	}

	return files, nil
}
