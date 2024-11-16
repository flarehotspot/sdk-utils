package pkg

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/a-h/templ/generator"
	"github.com/a-h/templ/parser/v2"
	sdkfs "github.com/flarehotspot/go-utils/fs"
)

func BuildTemplates(pluginDir string) error {
	var templateFiles []string
	if err := sdkfs.LsFiles(filepath.Join(pluginDir, "resources/views"), &templateFiles, true); err != nil {
		return err
	}

	var (
		wg    sync.WaitGroup
		errCh = make(chan error, len(templateFiles))
		errs  []error
	)

	for _, file := range templateFiles {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()

			if filepath.Ext(file) == ".templ" {
				dir := filepath.Dir(file)
				filename := filepath.Base(file)
				out := filepath.Join(dir, strings.Replace(filename, ".templ", "_templ.go", 1))
				fmt.Println("Generating template:", file, "->", out)

				t, err := parser.Parse(file)
				if err != nil {
					fmt.Println("Error parsing template", file, err)
					errCh <- err
					return
				}

				if sdkfs.Exists(out) {
					if err = os.Remove(out); err != nil {
						errCh <- err
						return
					}
				}

				outfile, err := os.OpenFile(out, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
				if err != nil {
					fmt.Println("Error opening file:", err)
					errCh <- err
					return
				}

				defer outfile.Close() // Ensure the file is closed after writing

				_, _, err = generator.Generate(t, outfile)
				if err != nil {
					fmt.Println("Error generating template", err)
					errCh <- err
					return
				}

				errCh <- nil
			}

		}(file)
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err := range errCh {
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("errors occurred while building templates: %v", errs)
	}

	return nil
}
