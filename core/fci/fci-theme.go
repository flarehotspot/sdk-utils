package fci

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/flarehotspot/core/config/themecfg"
	sdkfci "github.com/flarehotspot/core/sdk/api/fci"
	paths "github.com/flarehotspot/core/sdk/utils/paths"
)

// FciComposeView returns the html form as string
func FciComposeView(cfg *FciConfig) (htm string, err error) {
	var builder strings.Builder

	for _, sec := range cfg.Sections {
		sechtml, err := FciReadFile("section.html")
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

		for _, ipt := range sec.Inputs {
			t := ipt.Type()
			f, err := FciViewFile(t)
			htmlstr, err := FciReadFile(f)
			if err != nil {
				return "", err
			}

			tplname := fmt.Sprintf("input-%s-%s", ipt.Type(), ipt.Name())
			tpl, err := template.New(tplname).Parse(htmlstr)
			if err != nil {
				return "", err
			}

			result, err := FciExecInputTemplate(tpl, ipt)
			if err != nil {
				return "", err
			}

			builder.WriteString(result)
		}
	}

	return builder.String(), nil
}

func FciViewFile(t sdkfci.IFciInputTypes) (v string, err error) {
	switch t {
	case sdkfci.FciInputField:
		return "input-field.html", nil
	case sdkfci.FciInputFieldList:
		return "input-field-list.html", nil
	}

	return "", fmt.Errorf("invalid fci input type: %s", t)
}

func FciReadFile(v string) (f string, err error) {
	themepkg := themecfg.Read().Admin
	viewdir := filepath.Join(paths.VendorDir, themepkg, "resources/views/fci")
	file := filepath.Join(viewdir, v)

	b, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func FciExecInputTemplate(tpl *template.Template, ipt sdkfci.IFciInput) (htm string, err error) {
	var result strings.Builder

	t := ipt.Type()
	switch t {
	case sdkfci.FciInputField:
		data := ipt.(*FciInputField)
		err = tpl.Execute(&result, data)
		if err != nil {
			return "", err
		}

	case sdkfci.FciInputFieldList:
		data := ipt.(*FciFieldList)
		err = tpl.Execute(&result, data)
		if err != nil {
			return "", err
		}
	}

	return result.String(), nil
}
