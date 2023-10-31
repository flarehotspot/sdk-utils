package themes

import (
	"errors"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/flarehotspot/core/config/themecfg"
	"github.com/flarehotspot/core/fci"
	sdkfci "github.com/flarehotspot/core/sdk/api/fci"
	"github.com/flarehotspot/core/sdk/utils/fs"
	"github.com/flarehotspot/core/sdk/utils/paths"
)

func FciThemeView(v string) (f string, err error) {
	themepkg := themecfg.Read().WebAdmin
	viewdir := filepath.Join(paths.VendorDir, themepkg, "resources/views/fci")
	file := filepath.Join(viewdir, v)

	if !fs.Exists(file) {
		return "", errors.New("file not found: " + file)
	}

	b, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func FciComposeView(cfg *fci.FciConfig) (htm string, err error) {
	var builder strings.Builder

	for _, sec := range cfg.Sections {
		sechtml, err := FciThemeView("section.html")
		if err != nil {
			return "", err
		}

		name := strings.TrimSpace(sec.Secname)
		tpl, err := template.New("section-" + name + ".html").Parse(sechtml)
		if err != nil {
			return "", err
		}

		var result strings.Builder
		err = tpl.Execute(&result, sec)
		if err != nil {
			return "", err
		}

		builder.WriteString(result.String())

		for _, input := range sec.Inputs {
			viewfile, err := FciViewFile(input.Type())
			inphtml, err := FciThemeView(viewfile)
			if err != nil {
				return "", err
			}

			tplname := fmt.Sprintf("input-%s-%s", input.Type(), input.Name())
			tpl, err := template.New(tplname).Parse(inphtml)
			if err != nil {
				return "", err
			}

			var result strings.Builder
			err = tpl.Execute(&result, input)
			if err != nil {
				return "", err
			}

			builder.WriteString(result.String())
		}
	}

	return builder.String(), nil
}

func FciViewFile(t sdkfci.IFciInputTypes) (v string, err error) {
	switch t {
	case sdkfci.FciInputField:
		return "input-field.html", nil
	}

	return "", fmt.Errorf("invalid fci input type: %s", t)
}
