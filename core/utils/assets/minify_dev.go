//go:build dev

package assets

import (
	"io/ioutil"
	"log"
	"strings"
)

func concatFiles(files []string) (concat string, err error) {
	var sb strings.Builder

	for _, f := range files {
		var content string

		b, err := ioutil.ReadFile(f)
		if err != nil {
			log.Println("Asset not found: ", f)
		} else {
			content = string(b)
		}

		comment := filePathComment(f)
		_, err = sb.WriteString(comment + content)
		if err != nil {
			log.Printf("Error writing asset file \"%s\": %s", f, err.Error())
			return "", err
		}

	}

	return sb.String(), nil
}
