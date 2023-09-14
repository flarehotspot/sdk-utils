package migrate

import (
	"github.com/flarehotspot/core/sdk/utils/fs"
	"github.com/flarehotspot/core/sdk/utils/slices"
	"sort"
	"strings"
)

type MigDirection int

const (
	migration_Down MigDirection = iota
	migration_Up
)

func listFiles(dir string, d MigDirection) (files []string, err error) {
	files = []string{}
	list, err := fs.LsFiles(dir, false)
	if err != nil {
		return files, err
	}

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
