package plugins

import "strings"

type VueRoutePath string

func (self VueRoutePath) GetTemplate() string {
	return string(self)
}

func (self VueRoutePath) URL(pairs ...string) string {
	path := string(self)
	for i := 0; i < len(pairs); i += 2 {
		path = strings.Replace(path, ":"+pairs[i], pairs[i+1], 1)
	}
	return path
}
