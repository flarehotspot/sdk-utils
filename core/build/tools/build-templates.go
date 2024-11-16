package tools

import (
	"core/internal/utils/pkg"
	"fmt"
	"sync"
)

func BuildTemplates() {
	pluginDirs := pkg.ListPluginDirs(true)

	var wg sync.WaitGroup
	var errCh = make(chan error)

	for _, p := range pluginDirs {
		wg.Add(1)
		go func(p string) {
			defer wg.Done()
			errCh <- pkg.BuildTemplates(p)
		}(p)
	}

	go func() {
		wg.Done()
		close(errCh)
	}()

	var errs []error

	for err := range errCh {
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
		panic("Failed to build templates")
	}

}
