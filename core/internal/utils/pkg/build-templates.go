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

	var wg sync.WaitGroup
	errCh := make(chan error)

	for _, file := range templateFiles {
		if filepath.Ext(file) == ".templ" {
			wg.Add(1)
			go func(file string) {
				defer wg.Done()

				err := func(file string) (err error) {
					dir := filepath.Dir(file)
					filename := filepath.Base(file)
					out := filepath.Join(dir, strings.Replace(filename, ".templ", "_templ.go", 1))
					fmt.Println("Generating template:", file, "->", out)

					t, err := parser.Parse(file)
					if err != nil {
						fmt.Println("Error parsing template", file, err)
						return
					}

					if sdkfs.Exists(out) {
						if err = os.Remove(out); err != nil {
							return
						}
					}

					outfile, err := os.OpenFile(out, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
					if err != nil {
						fmt.Println("Error opening file:", err)
						return err
					}

					defer outfile.Close() // Ensure the file is closed after writing

					_, _, err = generator.Generate(t, outfile)
					if err != nil {
						fmt.Println("Error generating template", err)
					}
					return

				}(file)

				errCh <- err
			}(file)
		}
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	errs := []error{}
	for err := range errCh {
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return errs[0]
	}

	return nil
}
