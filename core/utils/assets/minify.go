//go:build !dev

// package assets

// import (
// 	minifyv2 "github.com/tdewolff/minify/v2"
// 	"github.com/tdewolff/minify/v2/css"
// 	"github.com/tdewolff/minify/v2/js"
// 	"log"
// 	"os"
// 	"path/filepath"
// 	"strings"
// )

// func minifyFiles(files []string) (concat string, err error) {
// 	var sb strings.Builder

// 	mtjs := "application/javascript"
// 	mtcss := "text/css"
// 	m := minifyv2.New()
// 	m.AddFunc(mtcss, css.Minify)
// 	m.AddFunc(mtjs, js.Minify)

// 	for _, f := range files {
// 		ext := filepath.Ext(f)

// 		// css or js file
// 		var mimetype string
// 		if ext == ".css" {
// 			mimetype = mtcss
// 		}
// 		if ext == ".js" {
// 			mimetype = mtjs
// 		}
// 		r, err := os.Open(f)
// 		if err != nil {
// 			log.Println(err)
// 			return "", nil
// 		}
// 		if ext == ".js" {
// 			sb.WriteString(";")
// 		}
// 		if err = m.Minify(mimetype, &sb, r); err != nil {
// 			log.Println("Cannot minify asset file "+f+":", err)
// 		}
// 	}
// 	return sb.String(), nil
// }
