package views

import (
	"github.com/flarehotspot/core/sdk/api/http/views"
	"github.com/flarehotspot/core/sdk/libs/jet"
	"github.com/flarehotspot/core/sdk/utils/paths"
	"github.com/flarehotspot/core/utils/assets"
	jobque "github.com/flarehotspot/core/utils/job-que"
)

var viewQue = jobque.NewJobQues()
var loader = NewJetLoader()
var viewSet = jet.NewSet(loader)

type ViewInput struct {
	File   string
	Extras *BundleExtras
}

func ViewProc(layout *ViewInput, content ViewInput, helpers views.IViewHelpers, data any) (html string, err error) {
	cache, ok := GetViewCache(layout, content)
	if !ok {
		sym, err := viewQue.Exec(func() (interface{}, error) {
			views := []*ViewInput{&content}
			if layout != nil {
				views = []*ViewInput{layout, &content}
			}

			// gather assets
			scripts := []string{}
			styles := []string{}
			for _, v := range views {
				va := ViewAssets(v.File)
				scripts = append(scripts, va.Scripts...)
				styles = append(styles, va.Styles...)

				if v.Extras != nil {
					if v.Extras.ExtraJS != nil {
						scripts = append(scripts, *v.Extras.ExtraJS...)
					}
					if v.Extras.ExtraCSS != nil {
						styles = append(styles, *v.Extras.ExtraCSS...)
					}
				}

				err = CopyDirsToPublic(v.File)
				if err != nil {
					return "", err
				}
			}

			jsbundle, err := assets.Bundle(scripts...)
			if err != nil {
				return "", err
			}

			cssbundle, err := assets.Bundle(styles...)
			if err != nil {
				return "", err
			}

			vc := WriteViewCache(layout, content, paths.Strip(jsbundle), paths.Strip(cssbundle))
			return vc, nil
		})

		if err != nil {
			return "", err
		}

		cache = sym.(*viewCache)
	}

	return cache.RenderHTML(helpers, data)
}
