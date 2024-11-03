package plugins

// import "strings"

// type VueRoutePath string

// func (self VueRoutePath) GetTemplate() string {
// 	return string(self)
// }

// func (self VueRoutePath) URL(pairs ...string) string {
// 	path := string(self)
// 	// get all param names in path
// 	paramKeys := []string{}
// 	for _, part := range strings.Split(path, "/") {
// 		if strings.HasPrefix(part, ":") {
// 			paramKeys = append(paramKeys, part[1:])
// 		}
// 	}

// 	for i := 0; i < len(pairs); i += 2 {
// 		path = strings.Replace(path, ":"+pairs[i], pairs[i+1], 1)
// 	}

// 	// add to query string all the params not in the path
// 	query := []string{}
// 	for i := 0; i < len(pairs); i += 2 {
// 		found := false
// 		for _, key := range paramKeys {
// 			if key == pairs[i] {
// 				found = true
// 			}
// 		}
// 		if !found {
// 			query = append(query, pairs[i]+"="+pairs[i+1])
// 		}
// 	}

// 	if len(query) > 0 {
// 		if strings.Contains(path, "?") {
// 			path += "&"
// 		} else {
// 			path += "?"
// 		}
// 		path += strings.Join(query, "&")
// 	}

// 	return path
// }
