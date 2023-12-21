package views

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/flarehotspot/core/sdk/api/http/views"
	"github.com/flarehotspot/core/sdk/libs/jet"
	"github.com/flarehotspot/core/sdk/utils/paths"
	strutil "github.com/flarehotspot/core/sdk/utils/strings"
)

var (
	cache = sync.Map{}
)

func WriteViewCache(layout *ViewInput, content ViewInput, jspath string, csspath string) *viewCache {
	hash := ViewsHash(layout, content)
	vc := &viewCache{
		content: content.File,
		jspath:  jspath,
		csspath: csspath,
	}

	if layout != nil {
		vc.layout = &layout.File
	}

	cache.Store(hash, vc)

	return vc
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
	hash    string
	layout  *string
	content string
	jspath  string
	csspath string
}

func (vc *viewCache) RenderHTML(helpers views.IViewHelpers, data any) (html string, err error) {
	vdata := &viewData{
		helpers:     helpers,
		contentData: data,
	}

	var tmpl *jet.Template

	if vc.layout != nil {
		contpath, err := paths.RelativeFromTo(*vc.layout, vc.content)
		if err != nil {
			return "", err
		}

		vdata.contentPath = contpath

		log.Println("contpath", contpath)
		log.Println("vc.layout", *vc.layout)
		log.Println("vc.content", vc.content)

		tmpl, err = GetTemplate(paths.Strip(*vc.layout))
		if err != nil {
			return "", err
		}
	} else {
		tmpl, err = GetTemplate(paths.Strip(vc.content))
		if err != nil {
			return "", err
		}
	}

	var buff bytes.Buffer
	if err := tmpl.Execute(&buff, nil, vdata); err != nil {
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
