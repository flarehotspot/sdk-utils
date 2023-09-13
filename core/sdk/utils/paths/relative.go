package paths

import (
	"strings"
)

/*
returns the relative path from file "from" to file "to".

	Example:

	from := "/path/to/dir1/file1.jpg"
	to   := "/path/to/dir2/file2.jpg"

  result := path.RelativeFromTo(from, to)
  log.Println(result)

  -> "../dir2/file2.jpg"

*/
func RelativeFromTo(from string, to string) string {

	var csb strings.Builder
	var rpsb strings.Builder
	farr := strings.Split(from, "/")
	tarr := strings.Split(to, "/")
	p := 0

	for i, fs := range farr {
		ts := tarr[i]
		if fs == ts {
			if i > 0 {
				csb.WriteString("/" + fs)
			} else {
				csb.WriteString(fs)
			}
		} else {
			p = len(farr) - (i + 1)
			break
		}
	}

	for i := 0; i < p; i++ {
		if i > 0 {
			rpsb.WriteString("/..")
		} else {
			rpsb.WriteString("..")
		}
	}

	return strings.Replace(to, csb.String(), rpsb.String(), 1)
}
