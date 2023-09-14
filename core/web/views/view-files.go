//go:build !dev

package views

func viewFiles(views ...*ViewInput) (files []string) {
	files = []string{}

	for _, v := range views {
		files = append(files, v.File)
	}

	return files
}
