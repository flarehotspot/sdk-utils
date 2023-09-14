package views

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	stdstr "strings"

	"github.com/flarehotspot/core/utils/crypt"
	jobque "github.com/flarehotspot/core/utils/job-que"
	"github.com/flarehotspot/core/sdk/utils/slices"
	"github.com/flarehotspot/core/sdk/utils/strings"
)

var viewQue = jobque.NewJobQues()

type ViewInput struct {
	File    string
	Extras  *BundleExtras
	FuncMap template.FuncMap
}

type viewCache struct {
	tmpl *template.Template
	hash string
}

func ViewProc(fmap template.FuncMap, views ...*ViewInput) (*template.Template, error) {
	cache, err := GetViewCache(views...)
	if err != nil {
		sym, err := viewQue.Exec(func() (interface{}, error) {
			log.Println(err)
			for _, v := range views {
				log.Printf("NOT CACHED: %+v", *v)
			}

			bundles := []AssetBundle{}
			for _, v := range views {
				b, err := AssetBundles(v.File, v.Extras)
				if err != nil {
					return nil, err
				}
				bundles = append(bundles, b)
			}

			viewFiles := []string{}
			for _, v := range views {
				viewFiles = append(viewFiles, v.File)
			}

			html, err := getHtmlContents(viewFiles...)
			if err != nil {
				return nil, err
			}

			html, err = procHtml(html, bundles)
			if err != nil {
				return nil, err
			}

			tmpl, err := template.New("html").Funcs(fmap).Parse(html)
			if err != nil {
				return nil, err
			}

			err = WriteViewCache(tmpl, views...)
			if err != nil {
				return nil, err
			}

			return tmpl, nil
		})

		if err != nil {
			return nil, err
		}

		return sym.(*template.Template), nil
	}

	return cache, nil
}

func getHtmlContents(views ...string) (string, error) {
	if len(views) < 2 {
		viewBytes, err := ioutil.ReadFile(views[0])
		if err != nil {
			return "", err
		}
		return string(viewBytes), nil
	} else {
		layoutFile := views[0]
		viewFile := views[1]
		html, err := insertContent(layoutFile, viewFile)
		if err != nil {
			return "", err
		}
		return html, nil
	}
}

func insertContent(layoutFile string, viewFile string) (result string, err error) {
	layoutBytes, err := ioutil.ReadFile(layoutFile)
	if err != nil {
		return "", err
	}

	contentBytes, err := ioutil.ReadFile(viewFile)
	if err != nil {
		return "", err
	}

	html := fmt.Sprintf("%s\n<!-- END LAYOUT -->\n {{ define \"content\" }}\n%s\n{{ end }}", layoutBytes, contentBytes)
	return html, nil
}

func insertTags(html string, sources []string, t string) (result string, err error) {
	if len(sources) < 1 {
		return html, nil
	}

	var delim string

	switch t {
	case TagTypeScript:
		delim = "</body>"
	case TagTypeStyle:
		delim = "</head>"
	default:
		err := errors.New("asset type can only \"script\" or \"style\"")
		return "", err
	}

	tags := slices.MapString(sources, func(s string) string {
		if s == "" {
			return ""
		}

		switch t {
		case TagTypeScript:
			return fmt.Sprintf("<script src=\"%s\"></script>", s)
		case TagTypeStyle:
			return fmt.Sprintf("<link rel=\"stylesheet\" href=\"%s\"></link>", s)
		default:
			return ""
		}
	})

	tag := stdstr.Join(tags, "\n") + "\n" + delim

	return stdstr.Replace(html, delim, tag, 1), nil
}

func procHtml(html string, bundles []AssetBundle) (string, error) {
	scripts := []string{}
	styles := []string{}

	for _, b := range bundles {
		scripts = append(scripts, b.ScriptSrc)
		styles = append(styles, b.StyleSrc)
	}

	html, err := insertTags(html, scripts, TagTypeScript)
	if err != nil {
		return "", err
	}

	html, err = insertTags(html, styles, TagTypeStyle)
	if err != nil {
		return "", err
	}

	return html, nil
}

func filesHash(files ...string) (hash string, err error) {
	hash, err = crypt.FastHashFiles(files...)
	return hash, err
}

func viewsHash(views ...*ViewInput) (hash string, err error) {
	files := []string{}

	for _, v := range views {
		files = append(files, v.File)
	}

	return strings.Sha1Hash(files...), nil
}
