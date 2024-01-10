package views

import (
	"fmt"
	"html/template"
	"strings"
	"sync"

	"github.com/flarehotspot/core/sdk/api/http/views"
	strutil "github.com/flarehotspot/core/sdk/utils/strings"
)

var (
	cache = sync.Map{}
)

func WriteViewCache(layout *ViewInput, content ViewInput, jspath string, csspath string) (*viewCache, error) {
	hash := ViewsHash(layout, content)
	vc := &viewCache{
		jspath:  jspath,
		csspath: csspath,
	}

	views := []string{content.File}
	if layout != nil {
		views = append(views, layout.File)
		vc.layout = true
	}

	templates, err := template.New("").ParseFiles(views...)
	if err != nil {
		return nil, err
	}

	vc.templates = templates

	cache.Store(hash, vc)

	return vc, nil
}

func GetViewCache(layout *ViewInput, content ViewInput) (vc *viewCache, ok bool) {
	hash := ViewsHash(layout, content)
	sym, ok := cache.Load(hash)
	if !ok {
		return nil, false
	}

	return sym.(*viewCache), true
}

func ViewsHash(layout *ViewInput, content ViewInput) (hash string) {
	files := viewFiles(layout, content)
	return strutil.Sha1Hash(files...)
}

type viewCache struct {
	hash      string
	layout    bool
	templates *template.Template
	jspath    string
	csspath   string
}

func (vc *viewCache) RenderHTML(helpers views.IViewHelpers, data any) (html string, err error) {
	vdata := &viewData{
		helpers:     helpers,
		contentData: data,
	}

	var tpl = "content"
	if vc.layout {
		tpl = "layout"
	}

	var buff strings.Builder
	if err := vc.templates.ExecuteTemplate(&buff, tpl, vdata); err != nil {
		return "", err
	}

	html = buff.String()

	return vc.attachAssets(html), nil
}

func (vc *viewCache) attachAssets(html string) (htmlout string) {
	tags := fmt.Sprintf("<script src=\"%s\"></script>\n<link rel=\"stylesheet\" href=\"%s\" />\n", vc.jspath, vc.csspath)
	html = strings.Replace(html, "</head>", fmt.Sprintf("%s\n</head>", tags), 1)
	return html
}
