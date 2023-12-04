//go:build dev

package views

import "github.com/flarehotspot/core/sdk/utils/fs"

func viewFiles(views ...*ViewInput) (files []string) {
	files = []string{}

	for _, v := range views {
		files = append(files, v.File)
		if v.Extras != nil {
			extras := v.Extras
			if extras.ExtraJS != nil {
				files = append(files, *extras.ExtraJS...)
			}

			if extras.ExtraCSS != nil {
				files = append(files, *extras.ExtraCSS...)
			}

			if extras.ExtraDirs != nil {
				for _, extraDir := range *extras.ExtraDirs {
					dirfiles, err := fs.LsFiles(extraDir.Src, true)
					if err != nil {
						continue
					}
					files = append(files, dirfiles...)
				}
			}
		}

		files = append(files, viewAssets(v.File)...)
	}
	return files
}
