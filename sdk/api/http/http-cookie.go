package sdkhttp

import (
	"net/http"
	"strings"
	"time"
)

func GetCookie(r *http.Request, name string) (string, error) {
	cookie, err := r.Cookie(name)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func SetCookie(w http.ResponseWriter, name string, value string) {
  expires := time.Now().AddDate(1, 0, 0)
  cookie := &http.Cookie{Name: name, Value: strings.TrimSpace(value), Path: "/", Expires: expires}
	http.SetCookie(w, cookie)
}
